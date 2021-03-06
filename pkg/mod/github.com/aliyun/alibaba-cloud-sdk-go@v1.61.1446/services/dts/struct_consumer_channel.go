package dts

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

// ConsumerChannel is a nested struct in dts response
type ConsumerChannel struct {
	ConsumerGroupUserName string `json:"ConsumerGroupUserName" xml:"ConsumerGroupUserName"`
	ConsumerGroupId       string `json:"ConsumerGroupId" xml:"ConsumerGroupId"`
	MessageDelay          int64  `json:"MessageDelay" xml:"MessageDelay"`
	ConsumerGroupName     string `json:"ConsumerGroupName" xml:"ConsumerGroupName"`
	ConsumptionCheckpoint string `json:"ConsumptionCheckpoint" xml:"ConsumptionCheckpoint"`
	UnconsumedData        int64  `json:"UnconsumedData" xml:"UnconsumedData"`
}
