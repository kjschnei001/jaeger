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

package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/kjschnei001/jaeger/model"
	"github.com/kjschnei001/jaeger/storage/dependencystore"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// GetDependencies provides a mock function with given fields: endTs, lookback
func (_m *Reader) GetDependencies(ctx context.Context, endTs time.Time, lookback time.Duration) ([]model.DependencyLink, error) {
	ret := _m.Called(endTs, lookback)

	var r0 []model.DependencyLink
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Duration) []model.DependencyLink); ok {
		r0 = rf(ctx, endTs, lookback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DependencyLink)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Duration) error); ok {
		r1 = rf(ctx, endTs, lookback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ dependencystore.Reader = (*Reader)(nil)
