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

package sanitizer

import (
	"github.com/kjschnei001/jaeger/model"
)

// SanitizeSpan sanitizes/normalizes spans. Any business logic that needs to be applied to normalize the contents of a
// span should implement this interface.
type SanitizeSpan func(span *model.Span) *model.Span

// NewStandardSanitizers are automatically applied by SpanProcessor.
func NewStandardSanitizers() []SanitizeSpan {
	return []SanitizeSpan{
		NewEmptyServiceNameSanitizer(),
	}
}

// NewChainedSanitizer creates a Sanitizer from the variadic list of passed Sanitizers.
// If the list only has one element, it is returned directly to minimize indirection.
func NewChainedSanitizer(sanitizers ...SanitizeSpan) SanitizeSpan {
	if len(sanitizers) == 1 {
		return sanitizers[0]
	}
	return func(span *model.Span) *model.Span {
		for _, s := range sanitizers {
			span = s(span)
		}
		return span
	}
}
