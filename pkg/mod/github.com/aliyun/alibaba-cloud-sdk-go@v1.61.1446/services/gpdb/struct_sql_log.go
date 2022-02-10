package gpdb

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

// SQLLog is a nested struct in gpdb response
type SQLLog struct {
	OperationClass       string   `json:"OperationClass" xml:"OperationClass"`
	ExecuteState         string   `json:"ExecuteState" xml:"ExecuteState"`
	ExecuteCost          float64  `json:"ExecuteCost" xml:"ExecuteCost"`
	SQLText              string   `json:"SQLText" xml:"SQLText"`
	SourcePort           int      `json:"SourcePort" xml:"SourcePort"`
	DBRole               string   `json:"DBRole" xml:"DBRole"`
	OperationType        string   `json:"OperationType" xml:"OperationType"`
	SourceIP             string   `json:"SourceIP" xml:"SourceIP"`
	SQLPlan              string   `json:"SQLPlan" xml:"SQLPlan"`
	ReturnRowCounts      int64    `json:"ReturnRowCounts" xml:"ReturnRowCounts"`
	DBName               string   `json:"DBName" xml:"DBName"`
	OperationExecuteTime string   `json:"OperationExecuteTime" xml:"OperationExecuteTime"`
	ScanRowCounts        int64    `json:"ScanRowCounts" xml:"ScanRowCounts"`
	AccountName          string   `json:"AccountName" xml:"AccountName"`
	QueryId              string   `json:"QueryId" xml:"QueryId"`
	SliceIds             []string `json:"SliceIds" xml:"SliceIds"`
}
