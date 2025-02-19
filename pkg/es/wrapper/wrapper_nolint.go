// Copyright (c) 2019 The Jaeger Authors.
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

package eswrapper

import "github.com/kjschnei001/jaeger/pkg/es"

// Some of the functions of elastic.BulkIndexRequest violate golint rules,
// e.g. Id() should be ID() and BodyJson() should be BodyJSON().

// Id calls this function to internal service.
func (i IndexServiceWrapper) Id(id string) es.IndexService {
	return WrapESIndexService(i.bulkIndexReq.Id(id), i.bulkService, i.esVersion)
}

// BodyJson calls this function to internal service.
func (i IndexServiceWrapper) BodyJson(body interface{}) es.IndexService {
	return WrapESIndexService(i.bulkIndexReq.Doc(body), i.bulkService, i.esVersion)
}
