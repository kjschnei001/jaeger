// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package processors

import (
	"context"
	"fmt"
	"sync"

	"github.com/apache/thrift/lib/go/thrift"
	"go.uber.org/zap"

	"github.com/kjschnei001/jaeger/cmd/agent/app/customtransport"
	"github.com/kjschnei001/jaeger/cmd/agent/app/servers"
	"github.com/kjschnei001/jaeger/pkg/metrics"
)

// ThriftProcessor is a server that processes spans using a TBuffered Server
type ThriftProcessor struct {
	server        servers.Server
	handler       AgentProcessor
	protocolPool  *sync.Pool
	numProcessors int
	processing    sync.WaitGroup
	logger        *zap.Logger
	metrics       struct {
		// Amount of time taken for processor to close
		ProcessorCloseTimer metrics.Timer `metric:"thrift.udp.t-processor.close-time"`

		// Number of failed buffer process operations
		HandlerProcessError metrics.Counter `metric:"thrift.udp.t-processor.handler-errors"`
	}
}

// AgentProcessor handler used by the processor to process thrift and call the reporter
// with the deserialized struct. This interface is implemented directly by Thrift generated
// code, e.g. jaegerThrift.NewAgentProcessor(handler), where handler implements the Agent
// Thrift service interface, which is invoked with the deserialized struct.
type AgentProcessor interface {
	Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException)
}

// NewThriftProcessor creates a TBufferedServer backed ThriftProcessor
func NewThriftProcessor(
	server servers.Server,
	numProcessors int,
	mFactory metrics.Factory,
	factory thrift.TProtocolFactory,
	handler AgentProcessor,
	logger *zap.Logger,
) (*ThriftProcessor, error) {
	if numProcessors <= 0 {
		return nil, fmt.Errorf(
			"number of processors must be greater than 0, called with %d", numProcessors)
	}
	protocolPool := &sync.Pool{
		New: func() interface{} {
			trans := &customtransport.TBufferedReadTransport{}
			return factory.GetProtocol(trans)
		},
	}

	res := &ThriftProcessor{
		server:        server,
		handler:       handler,
		protocolPool:  protocolPool,
		logger:        logger,
		numProcessors: numProcessors,
	}
	metrics.Init(&res.metrics, mFactory, nil)
	res.processing.Add(res.numProcessors)
	for i := 0; i < res.numProcessors; i++ {
		go func() {
			res.processBuffer()
			res.processing.Done()
		}()
	}
	return res, nil
}

// Serve starts serving traffic
func (s *ThriftProcessor) Serve() {
	s.server.Serve()
}

// IsServing indicates whether the server is currently serving traffic
func (s *ThriftProcessor) IsServing() bool {
	return s.server.IsServing()
}

// Stop stops the serving of traffic and waits until the queue is
// emptied by the readers
func (s *ThriftProcessor) Stop() {
	stopwatch := metrics.StartStopwatch(s.metrics.ProcessorCloseTimer)
	s.server.Stop()
	s.processing.Wait()
	stopwatch.Stop()
}

// processBuffer reads data off the channel and puts it into a custom transport for
// the processor to process
func (s *ThriftProcessor) processBuffer() {
	for readBuf := range s.server.DataChan() {
		protocol := s.protocolPool.Get().(thrift.TProtocol)
		payload := readBuf.GetBytes()
		protocol.Transport().Write(payload)
		s.logger.Debug("Span(s) received by the agent", zap.Int("bytes-received", len(payload)))

		// NB: oddly, thrift-gen/agent/agent.go:L156 does this: `return true, thrift.WrapTException(err2)`
		// So we check for both OK and error.
		if ok, err := s.handler.Process(context.Background(), protocol, protocol); !ok || err != nil {
			s.logger.Error("Processor failed", zap.Error(err))
			s.metrics.HandlerProcessError.Inc(1)
		}
		s.protocolPool.Put(protocol)
		s.server.DataRecd(readBuf) // acknowledge receipt and release the buffer
	}
}
