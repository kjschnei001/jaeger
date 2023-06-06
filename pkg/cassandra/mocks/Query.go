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

import cassandra "github.com/kjschnei001/jaeger/pkg/cassandra"
import mock "github.com/stretchr/testify/mock"

// Query is an autogenerated mock type for the Query type
type Query struct {
	mock.Mock
}

// Bind provides a mock function with given fields: v
func (_m *Query) Bind(v ...interface{}) cassandra.Query {
	ret := _m.Called(v)

	var r0 cassandra.Query
	if rf, ok := ret.Get(0).(func(...interface{}) cassandra.Query); ok {
		r0 = rf(v...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cassandra.Query)
		}
	}

	return r0
}

// Consistency provides a mock function with given fields: level
func (_m *Query) Consistency(level cassandra.Consistency) cassandra.Query {
	ret := _m.Called(level)

	var r0 cassandra.Query
	if rf, ok := ret.Get(0).(func(cassandra.Consistency) cassandra.Query); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cassandra.Query)
		}
	}

	return r0
}

// PageSize provides a mock function with given fields: n
func (_m *Query) PageSize(n int) cassandra.Query {
	ret := _m.Called(n)

	var r0 cassandra.Query
	if rf, ok := ret.Get(0).(func(int) cassandra.Query); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cassandra.Query)
		}
	}

	return r0
}

// Exec provides a mock function with given fields:
func (_m *Query) Exec() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Iter provides a mock function with given fields:
func (_m *Query) Iter() cassandra.Iterator {
	ret := _m.Called()

	var r0 cassandra.Iterator
	if rf, ok := ret.Get(0).(func() cassandra.Iterator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cassandra.Iterator)
		}
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Query) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ScanCAS provides a mock function with given fields: dest
func (_m *Query) ScanCAS(dest ...interface{}) (bool, error) {
	ret := _m.Called(dest)

	var r0 bool
	if rf, ok := ret.Get(0).(func(...interface{}) bool); ok {
		r0 = rf(dest...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(dest...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ cassandra.Query = (*Query)(nil)
