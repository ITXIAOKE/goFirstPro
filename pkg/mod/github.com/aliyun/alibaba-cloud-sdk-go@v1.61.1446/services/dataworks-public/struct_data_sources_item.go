package dataworks_public

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

// DataSourcesItem is a nested struct in dataworks_public response
type DataSourcesItem struct {
	Name                string `json:"Name" xml:"Name"`
	BindingCalcEngineId int    `json:"BindingCalcEngineId" xml:"BindingCalcEngineId"`
	DefaultEngine       bool   `json:"DefaultEngine" xml:"DefaultEngine"`
	Id                  int    `json:"Id" xml:"Id"`
	Status              int    `json:"Status" xml:"Status"`
	SubType             string `json:"SubType" xml:"SubType"`
	DataSourceType      string `json:"DataSourceType" xml:"DataSourceType"`
	GmtModified         string `json:"GmtModified" xml:"GmtModified"`
	TenantId            int64  `json:"TenantId" xml:"TenantId"`
	ConnectStatus       int    `json:"ConnectStatus" xml:"ConnectStatus"`
	Operator            string `json:"Operator" xml:"Operator"`
	Sequence            int    `json:"Sequence" xml:"Sequence"`
	GmtCreate           string `json:"GmtCreate" xml:"GmtCreate"`
	Content             string `json:"Content" xml:"Content"`
	EnvType             int    `json:"EnvType" xml:"EnvType"`
	ProjectId           int    `json:"ProjectId" xml:"ProjectId"`
	Shared              bool   `json:"Shared" xml:"Shared"`
	Description         string `json:"Description" xml:"Description"`
}
