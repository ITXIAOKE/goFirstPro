package mts

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

// JobsItem is a nested struct in mts response
type JobsItem struct {
	Status    string `json:"Status" xml:"Status"`
	Result    string `json:"Result" xml:"Result"`
	JobId     string `json:"JobId" xml:"JobId"`
	UserId    int64  `json:"UserId" xml:"UserId"`
	Message   string `json:"Message" xml:"Message"`
	JobTime   int64  `json:"JobTime" xml:"JobTime"`
	JobParams string `json:"JobParams" xml:"JobParams"`
}
