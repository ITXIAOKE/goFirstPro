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

// BasicInfo is a nested struct in retailcloud response
type BasicInfo struct {
	BusinessCode            string   `json:"BusinessCode" xml:"BusinessCode"`
	ClusterId               int64    `json:"ClusterId" xml:"ClusterId"`
	InstallLogInProcess     bool     `json:"InstallLogInProcess" xml:"InstallLogInProcess"`
	MainUserId              string   `json:"MainUserId" xml:"MainUserId"`
	RegionName              string   `json:"RegionName" xml:"RegionName"`
	HasInstallArmsPilot     bool     `json:"HasInstallArmsPilot" xml:"HasInstallArmsPilot"`
	HasInstallLogController bool     `json:"HasInstallLogController" xml:"HasInstallLogController"`
	InstallArmsInProcess    bool     `json:"InstallArmsInProcess" xml:"InstallArmsInProcess"`
	RegionId                string   `json:"RegionId" xml:"RegionId"`
	EnvType                 string   `json:"EnvType" xml:"EnvType"`
	UserNick                string   `json:"UserNick" xml:"UserNick"`
	ClusterName             string   `json:"ClusterName" xml:"ClusterName"`
	ClusterInstanceId       string   `json:"ClusterInstanceId" xml:"ClusterInstanceId"`
	VpcId                   string   `json:"VpcId" xml:"VpcId"`
	Vswitchs                []string `json:"Vswitchs" xml:"Vswitchs"`
	EcsIds                  []string `json:"EcsIds" xml:"EcsIds"`
}