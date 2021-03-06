package retailcloud

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

// AppInstanceDetail is a nested struct in retailcloud response
type AppInstanceDetail struct {
	AppInstanceId string `json:"AppInstanceId" xml:"AppInstanceId"`
	CreateTime    string `json:"CreateTime" xml:"CreateTime"`
	Spec          string `json:"Spec" xml:"Spec"`
	RestartCount  int    `json:"RestartCount" xml:"RestartCount"`
	HostIp        string `json:"HostIp" xml:"HostIp"`
	PodIp         string `json:"PodIp" xml:"PodIp"`
	Health        string `json:"Health" xml:"Health"`
	Requests      string `json:"Requests" xml:"Requests"`
	Limits        string `json:"Limits" xml:"Limits"`
	Version       string `json:"Version" xml:"Version"`
	Status        string `json:"Status" xml:"Status"`
}
