package edas

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

// SpringCloudMockItem is a nested struct in edas response
type SpringCloudMockItem struct {
	Value              string             `json:"Value" xml:"Value"`
	Path               string             `json:"Path" xml:"Path"`
	ExceptionClassName string             `json:"ExceptionClassName" xml:"ExceptionClassName"`
	Timeout            int64              `json:"Timeout" xml:"Timeout"`
	ExecuteCondition   string             `json:"ExecuteCondition" xml:"ExecuteCondition"`
	Condition          string             `json:"Condition" xml:"Condition"`
	Method             string             `json:"Method" xml:"Method"`
	Oper               string             `json:"Oper" xml:"Oper"`
	ServiceName        string             `json:"ServiceName" xml:"ServiceName"`
	ArgumentMockItems  []ArgumentMockItem `json:"ArgumentMockItems" xml:"ArgumentMockItems"`
}
