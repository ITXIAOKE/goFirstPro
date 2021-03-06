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

// EurekaInstance is a nested struct in mse response
type EurekaInstance struct {
	Status                string                 `json:"Status" xml:"Status"`
	LastDirtyTimestamp    int64                  `json:"LastDirtyTimestamp" xml:"LastDirtyTimestamp"`
	IpAddr                string                 `json:"IpAddr" xml:"IpAddr"`
	HomePageUrl           string                 `json:"HomePageUrl" xml:"HomePageUrl"`
	HostName              string                 `json:"HostName" xml:"HostName"`
	InstanceId            string                 `json:"InstanceId" xml:"InstanceId"`
	Port                  int                    `json:"Port" xml:"Port"`
	SecurePort            int                    `json:"SecurePort" xml:"SecurePort"`
	App                   string                 `json:"App" xml:"App"`
	DurationInSecs        int                    `json:"DurationInSecs" xml:"DurationInSecs"`
	LastUpdatedTimestamp  int64                  `json:"LastUpdatedTimestamp" xml:"LastUpdatedTimestamp"`
	RenewalIntervalInSecs int                    `json:"RenewalIntervalInSecs" xml:"RenewalIntervalInSecs"`
	VipAddress            string                 `json:"VipAddress" xml:"VipAddress"`
	Metadata              map[string]interface{} `json:"Metadata" xml:"Metadata"`
}
