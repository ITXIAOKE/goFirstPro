package sddp

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

// DataLimit is a nested struct in sddp response
type DataLimit struct {
	Connector           string               `json:"Connector" xml:"Connector"`
	Enable              int                  `json:"Enable" xml:"Enable"`
	ErrorMessage        string               `json:"ErrorMessage" xml:"ErrorMessage"`
	TenantName          string               `json:"TenantName" xml:"TenantName"`
	LocalName           string               `json:"LocalName" xml:"LocalName"`
	SupportAgentInstall bool                 `json:"SupportAgentInstall" xml:"SupportAgentInstall"`
	InstanceDescription string               `json:"InstanceDescription" xml:"InstanceDescription"`
	SupportAudit        bool                 `json:"SupportAudit" xml:"SupportAudit"`
	CheckStatusName     string               `json:"CheckStatusName" xml:"CheckStatusName"`
	DbVersion           string               `json:"DbVersion" xml:"DbVersion"`
	SupportEvent        bool                 `json:"SupportEvent" xml:"SupportEvent"`
	RegionId            string               `json:"RegionId" xml:"RegionId"`
	TotalCount          int                  `json:"TotalCount" xml:"TotalCount"`
	InstanceId          string               `json:"InstanceId" xml:"InstanceId"`
	OcrStatus           int                  `json:"OcrStatus" xml:"OcrStatus"`
	AuditStatus         int                  `json:"AuditStatus" xml:"AuditStatus"`
	LogStoreDay         int                  `json:"LogStoreDay" xml:"LogStoreDay"`
	EngineType          string               `json:"EngineType" xml:"EngineType"`
	ErrorCode           string               `json:"ErrorCode" xml:"ErrorCode"`
	ProcessTotalCount   int                  `json:"ProcessTotalCount" xml:"ProcessTotalCount"`
	CheckStatus         int                  `json:"CheckStatus" xml:"CheckStatus"`
	ResourceTypeCode    string               `json:"ResourceTypeCode" xml:"ResourceTypeCode"`
	Port                int                  `json:"Port" xml:"Port"`
	SupportDatamask     bool                 `json:"SupportDatamask" xml:"SupportDatamask"`
	SamplingSize        int                  `json:"SamplingSize" xml:"SamplingSize"`
	ParentId            string               `json:"ParentId" xml:"ParentId"`
	EventStatus         int                  `json:"EventStatus" xml:"EventStatus"`
	ProcessStatus       int                  `json:"ProcessStatus" xml:"ProcessStatus"`
	SupportScan         bool                 `json:"SupportScan" xml:"SupportScan"`
	Id                  int64                `json:"Id" xml:"Id"`
	NextStartTime       int64                `json:"NextStartTime" xml:"NextStartTime"`
	GmtCreate           int64                `json:"GmtCreate" xml:"GmtCreate"`
	ResourceType        int64                `json:"ResourceType" xml:"ResourceType"`
	AgentId             string               `json:"AgentId" xml:"AgentId"`
	AutoScan            int                  `json:"AutoScan" xml:"AutoScan"`
	UserName            string               `json:"UserName" xml:"UserName"`
	LastFinishedTime    int64                `json:"LastFinishedTime" xml:"LastFinishedTime"`
	DatamaskStatus      int                  `json:"DatamaskStatus" xml:"DatamaskStatus"`
	SupportOcr          bool                 `json:"SupportOcr" xml:"SupportOcr"`
	AgentState          int                  `json:"AgentState" xml:"AgentState"`
	DataLimitList       []DataLimitListInner `json:"DataLimitList" xml:"DataLimitList"`
}
