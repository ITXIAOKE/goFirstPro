package clickhouse

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

// LoghubInfoInDescribeLoghubDetail is a nested struct in clickhouse response
type LoghubInfoInDescribeLoghubDetail struct {
	AccessKey       string                             `json:"AccessKey" xml:"AccessKey"`
	TableName       string                             `json:"TableName" xml:"TableName"`
	AccessSecret    string                             `json:"AccessSecret" xml:"AccessSecret"`
	ProjectName     string                             `json:"ProjectName" xml:"ProjectName"`
	SchemaName      string                             `json:"SchemaName" xml:"SchemaName"`
	DBType          string                             `json:"DBType" xml:"DBType"`
	DeliverName     string                             `json:"DeliverName" xml:"DeliverName"`
	RegionId        string                             `json:"RegionId" xml:"RegionId"`
	Password        string                             `json:"Password" xml:"Password"`
	DBClusterId     string                             `json:"DBClusterId" xml:"DBClusterId"`
	Description     string                             `json:"Description" xml:"Description"`
	FilterDirtyData bool                               `json:"FilterDirtyData" xml:"FilterDirtyData"`
	ZoneId          string                             `json:"ZoneId" xml:"ZoneId"`
	LogStoreName    string                             `json:"LogStoreName" xml:"LogStoreName"`
	UserName        string                             `json:"UserName" xml:"UserName"`
	DomainUrl       string                             `json:"DomainUrl" xml:"DomainUrl"`
	DeliverTime     string                             `json:"DeliverTime" xml:"DeliverTime"`
	LogHubStores    LogHubStoresInDescribeLoghubDetail `json:"LogHubStores" xml:"LogHubStores"`
}
