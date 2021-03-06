package alb

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

// BackendServer is a nested struct in alb response
type BackendServer struct {
	Description     string `json:"Description" xml:"Description"`
	Port            int    `json:"Port" xml:"Port"`
	ServerId        string `json:"ServerId" xml:"ServerId"`
	ServerIp        string `json:"ServerIp" xml:"ServerIp"`
	ServerType      string `json:"ServerType" xml:"ServerType"`
	Status          string `json:"Status" xml:"Status"`
	Weight          int    `json:"Weight" xml:"Weight"`
	ServerGroupId   string `json:"ServerGroupId" xml:"ServerGroupId"`
	RemoteIpEnabled bool   `json:"RemoteIpEnabled" xml:"RemoteIpEnabled"`
}
