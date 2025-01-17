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

package clientcfghttp

import (
	"encoding/json"
	"errors"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kjschnei001/jaeger/internal/metricstest"
	p2json "github.com/kjschnei001/jaeger/model/converter/json"
	tSampling092 "github.com/kjschnei001/jaeger/pkg/clientcfg/clientcfghttp/thrift-0.9.2"
	"github.com/kjschnei001/jaeger/proto-gen/api_v2"
	"github.com/kjschnei001/jaeger/thrift-gen/baggage"
)

type testServer struct {
	metricsFactory *metricstest.Factory
	samplingStore  *mockSamplingStore
	bgMgr          *mockBaggageMgr
	server         *httptest.Server
	handler        *HTTPHandler
}

func withServer(
	basePath string,
	mockSamplingResponse *api_v2.SamplingStrategyResponse,
	mockBaggageResponse []*baggage.BaggageRestriction,
	testFn func(server *testServer),
) {
	metricsFactory := metricstest.NewFactory(0)
	samplingStore := &mockSamplingStore{samplingResponse: mockSamplingResponse}
	bgMgr := &mockBaggageMgr{baggageResponse: mockBaggageResponse}
	cfgMgr := &ConfigManager{
		SamplingStrategyStore: samplingStore,
		BaggageManager:        bgMgr,
	}
	handler := NewHTTPHandler(HTTPHandlerParams{
		ConfigManager:          cfgMgr,
		MetricsFactory:         metricsFactory,
		BasePath:               basePath,
		LegacySamplingEndpoint: true,
	})
	r := mux.NewRouter()
	handler.RegisterRoutes(r)
	server := httptest.NewServer(r)
	defer server.Close()
	testFn(&testServer{
		metricsFactory: metricsFactory,
		samplingStore:  samplingStore,
		bgMgr:          bgMgr,
		server:         server,
		handler:        handler,
	})
}

func TestHTTPHandler(t *testing.T) {
	testHTTPHandler(t, "")
}

func TestHTTPHandlerWithBasePath(t *testing.T) {
	testHTTPHandler(t, "/foo")
}

func testHTTPHandler(t *testing.T, basePath string) {
	withServer(basePath, rateLimiting(42), restrictions("luggage", 10), func(ts *testServer) {
		tests := []struct {
			endpoint  string
			expOutput string
		}{
			{
				endpoint:  "/",
				expOutput: `{"strategyType":1,"rateLimitingSampling":{"maxTracesPerSecond":42}}`,
			},
			{
				endpoint:  "/sampling",
				expOutput: `{"strategyType":"RATE_LIMITING","rateLimitingSampling":{"maxTracesPerSecond":42}}`,
			},
		}
		for _, test := range tests {
			t.Run("endpoint="+test.endpoint, func(t *testing.T) {
				resp, err := http.Get(ts.server.URL + basePath + test.endpoint + "?service=Y")
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, resp.StatusCode)
				body, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				err = resp.Body.Close()
				require.NoError(t, err)
				assert.Equal(t, test.expOutput, string(body))
				if test.endpoint == "/" {
					objResp := &tSampling092.SamplingStrategyResponse{}
					require.NoError(t, json.Unmarshal(body, objResp))
					assert.EqualValues(t,
						ts.samplingStore.samplingResponse.GetStrategyType(),
						objResp.GetStrategyType())
					assert.EqualValues(t,
						ts.samplingStore.samplingResponse.GetRateLimitingSampling().GetMaxTracesPerSecond(),
						objResp.GetRateLimitingSampling().GetMaxTracesPerSecond())
				} else {
					objResp, err := p2json.SamplingStrategyResponseFromJSON(body)
					require.NoError(t, err)
					assert.EqualValues(t, ts.samplingStore.samplingResponse, objResp)
				}
			})
		}

		t.Run("request against endpoint /baggageRestrictions", func(t *testing.T) {
			resp, err := http.Get(ts.server.URL + basePath + "/baggageRestrictions?service=Y")
			require.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			require.NoError(t, err)
			var objResp []*baggage.BaggageRestriction
			require.NoError(t, json.Unmarshal(body, &objResp))
			assert.EqualValues(t, ts.bgMgr.baggageResponse, objResp)
		})

		// handler must emit metrics
		ts.metricsFactory.AssertCounterMetrics(t, []metricstest.ExpectedMetric{
			{Name: "http-server.requests", Tags: map[string]string{"type": "sampling"}, Value: 1},
			{Name: "http-server.requests", Tags: map[string]string{"type": "sampling-legacy"}, Value: 1},
			{Name: "http-server.requests", Tags: map[string]string{"type": "baggage"}, Value: 1},
		}...)
	})
}

func TestHTTPHandlerErrors(t *testing.T) {
	testCases := []struct {
		description          string
		mockSamplingResponse *api_v2.SamplingStrategyResponse
		mockBaggageResponse  []*baggage.BaggageRestriction
		url                  string
		statusCode           int
		body                 string
		metrics              []metricstest.ExpectedMetric
	}{
		{
			description: "no service name",
			url:         "",
			statusCode:  http.StatusBadRequest,
			body:        "'service' parameter must be provided once\n",
			metrics: []metricstest.ExpectedMetric{
				{Name: "http-server.errors", Tags: map[string]string{"source": "all", "status": "4xx"}, Value: 1},
			},
		},
		{
			description: "sampling endpoint too many service names",
			url:         "?service=Y&service=Y",
			statusCode:  http.StatusBadRequest,
			body:        "'service' parameter must be provided once\n",
			metrics: []metricstest.ExpectedMetric{
				{Name: "http-server.errors", Tags: map[string]string{"source": "all", "status": "4xx"}, Value: 1},
			},
		},
		{
			description: "baggage endpoint too many service names",
			url:         "/baggageRestrictions?service=Y&service=Y",
			statusCode:  http.StatusBadRequest,
			body:        "'service' parameter must be provided once\n",
			metrics: []metricstest.ExpectedMetric{
				{Name: "http-server.errors", Tags: map[string]string{"source": "all", "status": "4xx"}, Value: 1},
			},
		},
		{
			description: "sampler collector error",
			url:         "?service=Y",
			statusCode:  http.StatusInternalServerError,
			body:        "collector error: no mock response provided\n",
			metrics: []metricstest.ExpectedMetric{
				{Name: "http-server.errors", Tags: map[string]string{"source": "collector-proxy", "status": "5xx"}, Value: 1},
			},
		},
		{
			description: "baggage collector error",
			url:         "/baggageRestrictions?service=Y",
			statusCode:  http.StatusInternalServerError,
			body:        "collector error: no mock response provided\n",
			metrics: []metricstest.ExpectedMetric{
				{Name: "http-server.errors", Tags: map[string]string{"source": "collector-proxy", "status": "5xx"}, Value: 1},
			},
		},
		{
			description:          "sampler marshalling error",
			mockSamplingResponse: probabilistic(math.NaN()),
			url:                  "?service=Y",
			statusCode:           http.StatusInternalServerError,
			body:                 "cannot marshall to JSON\n",
			metrics: []metricstest.ExpectedMetric{
				{Name: "http-server.errors", Tags: map[string]string{"source": "thrift", "status": "5xx"}, Value: 1},
			},
		},
	}
	for _, tc := range testCases {
		testCase := tc // capture loop var
		t.Run(testCase.description, func(t *testing.T) {
			withServer("", testCase.mockSamplingResponse, testCase.mockBaggageResponse, func(ts *testServer) {
				resp, err := http.Get(ts.server.URL + testCase.url)
				require.NoError(t, err)
				assert.Equal(t, testCase.statusCode, resp.StatusCode)
				if testCase.body != "" {
					body, err := io.ReadAll(resp.Body)
					assert.NoError(t, err)
					assert.Equal(t, testCase.body, string(body))
				}

				if len(testCase.metrics) > 0 {
					ts.metricsFactory.AssertCounterMetrics(t, testCase.metrics...)
				}
			})
		})
	}

	t.Run("failure to write a response", func(t *testing.T) {
		withServer("", probabilistic(0.001), restrictions("luggage", 10), func(ts *testServer) {
			handler := ts.handler

			req := httptest.NewRequest("GET", "http://localhost:80/?service=X", nil)
			w := &mockWriter{header: make(http.Header)}
			handler.serveSamplingHTTP(w, req, handler.encodeThriftLegacy)

			ts.metricsFactory.AssertCounterMetrics(t,
				metricstest.ExpectedMetric{Name: "http-server.errors", Tags: map[string]string{"source": "write", "status": "5xx"}, Value: 1})

			req = httptest.NewRequest("GET", "http://localhost:80/baggageRestrictions?service=X", nil)
			handler.serveBaggageHTTP(w, req)

			ts.metricsFactory.AssertCounterMetrics(t,
				metricstest.ExpectedMetric{Name: "http-server.errors", Tags: map[string]string{"source": "write", "status": "5xx"}, Value: 2})
		})
	})
}

func TestEncodeErrors(t *testing.T) {
	withServer("", nil, nil, func(server *testServer) {
		_, err := server.handler.encodeThriftLegacy(&api_v2.SamplingStrategyResponse{
			StrategyType: -1,
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "ConvertSamplingResponseFromDomain failed")
		server.metricsFactory.AssertCounterMetrics(t, []metricstest.ExpectedMetric{
			{Name: "http-server.errors", Tags: map[string]string{"source": "thrift", "status": "5xx"}, Value: 1},
		}...)

		_, err = server.handler.encodeProto(nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "SamplingStrategyResponseToJSON failed")
		server.metricsFactory.AssertCounterMetrics(t, []metricstest.ExpectedMetric{
			{Name: "http-server.errors", Tags: map[string]string{"source": "proto", "status": "5xx"}, Value: 1},
		}...)
	})
}

func rateLimiting(rate int32) *api_v2.SamplingStrategyResponse {
	return &api_v2.SamplingStrategyResponse{
		StrategyType: api_v2.SamplingStrategyType_RATE_LIMITING,
		RateLimitingSampling: &api_v2.RateLimitingSamplingStrategy{
			MaxTracesPerSecond: rate,
		},
	}
}

func probabilistic(probability float64) *api_v2.SamplingStrategyResponse {
	return &api_v2.SamplingStrategyResponse{
		StrategyType: api_v2.SamplingStrategyType_PROBABILISTIC,
		ProbabilisticSampling: &api_v2.ProbabilisticSamplingStrategy{
			SamplingRate: probability,
		},
	}
}

func restrictions(key string, size int32) []*baggage.BaggageRestriction {
	return []*baggage.BaggageRestriction{
		{BaggageKey: key, MaxValueLength: size},
	}
}

type mockWriter struct {
	header http.Header
}

func (w *mockWriter) Header() http.Header {
	return w.header
}

func (w *mockWriter) Write([]byte) (int, error) {
	return 0, errors.New("write error")
}

func (w *mockWriter) WriteHeader(int) {}
