package mse

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

// NacosAnsInstance is a nested struct in mse response
type NacosAnsInstance struct {
	DefaultKey                string                 `json:"DefaultKey" xml:"DefaultKey"`
	Ephemeral                 bool                   `json:"Ephemeral" xml:"Ephemeral"`
	Marked                    bool                   `json:"Marked" xml:"Marked"`
	Ip                        string                 `json:"Ip" xml:"Ip"`
	InstanceId                string                 `json:"InstanceId" xml:"InstanceId"`
	Port                      int                    `json:"Port" xml:"Port"`
	LastBeat                  int64                  `json:"LastBeat" xml:"LastBeat"`
	OkCount                   int                    `json:"OkCount" xml:"OkCount"`
	Weight                    int                    `json:"Weight" xml:"Weight"`
	InstanceHeartBeatInterval int                    `json:"InstanceHeartBeatInterval" xml:"InstanceHeartBeatInterval"`
	IpDeleteTimeout           int                    `json:"IpDeleteTimeout" xml:"IpDeleteTimeout"`
	App                       string                 `json:"App" xml:"App"`
	FailCount                 int                    `json:"FailCount" xml:"FailCount"`
	Healthy                   bool                   `json:"Healthy" xml:"Healthy"`
	Enabled                   bool                   `json:"Enabled" xml:"Enabled"`
	DatumKey                  string                 `json:"DatumKey" xml:"DatumKey"`
	ClusterName               string                 `json:"ClusterName" xml:"ClusterName"`
	InstanceHeartBeatTimeOut  int                    `json:"InstanceHeartBeatTimeOut" xml:"InstanceHeartBeatTimeOut"`
	ServiceName               string                 `json:"ServiceName" xml:"ServiceName"`
	Metadata                  map[string]interface{} `json:"Metadata" xml:"Metadata"`
}
