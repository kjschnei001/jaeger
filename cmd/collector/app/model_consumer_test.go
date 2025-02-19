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

package app

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kjschnei001/jaeger/model"
)

func TestChainedProcessSpan(t *testing.T) {
	happened1 := false
	happened2 := false
	func1 := func(span *model.Span, tenant string) { happened1 = true }
	func2 := func(span *model.Span, tenant string) { happened2 = true }
	chained := ChainedProcessSpan(func1, func2)
	chained(&model.Span{}, "")
	assert.True(t, happened1)
	assert.True(t, happened2)
}
