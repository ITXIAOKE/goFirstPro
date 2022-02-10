package xtrace

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Span is a nested struct in xtrace response
type Span struct {
	TraceID       string                 `json:"TraceID" xml:"TraceID"`
	OperationName string                 `json:"OperationName" xml:"OperationName"`
	Duration      int64                  `json:"Duration" xml:"Duration"`
	ServiceName   string                 `json:"ServiceName" xml:"ServiceName"`
	ServiceIp     string                 `json:"ServiceIp" xml:"ServiceIp"`
	Timestamp     int64                  `json:"Timestamp" xml:"Timestamp"`
	RpcId         string                 `json:"RpcId" xml:"RpcId"`
	ResultCode    string                 `json:"ResultCode" xml:"ResultCode"`
	HaveStack     bool                   `json:"HaveStack" xml:"HaveStack"`
	SpanId        string                 `json:"SpanId" xml:"SpanId"`
	ParentSpanId  string                 `json:"ParentSpanId" xml:"ParentSpanId"`
	TagEntryList  TagEntryListInGetTrace `json:"TagEntryList" xml:"TagEntryList"`
	LogEventList  LogEventList           `json:"LogEventList" xml:"LogEventList"`
}