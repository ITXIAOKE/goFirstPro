package dms_enterprise

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

// SQLExecAuditLog is a nested struct in dms_enterprise response
type SQLExecAuditLog struct {
	OpTime       string `json:"OpTime" xml:"OpTime"`
	UserName     string `json:"UserName" xml:"UserName"`
	UserId       int64  `json:"UserId" xml:"UserId"`
	InstanceName string `json:"InstanceName" xml:"InstanceName"`
	InstanceId   int64  `json:"InstanceId" xml:"InstanceId"`
	SchemaName   string `json:"SchemaName" xml:"SchemaName"`
	DbId         int64  `json:"DbId" xml:"DbId"`
	Logic        bool   `json:"Logic" xml:"Logic"`
	SQLType      string `json:"SQLType" xml:"SQLType"`
	SQL          string `json:"SQL" xml:"SQL"`
	ExecState    string `json:"ExecState" xml:"ExecState"`
	AffectRows   int64  `json:"AffectRows" xml:"AffectRows"`
	ElapsedTime  int64  `json:"ElapsedTime" xml:"ElapsedTime"`
	Remark       string `json:"Remark" xml:"Remark"`
}
