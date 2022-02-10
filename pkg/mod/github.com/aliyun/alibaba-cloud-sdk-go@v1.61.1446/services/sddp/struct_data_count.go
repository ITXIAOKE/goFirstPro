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

// DataCount is a nested struct in sddp response
type DataCount struct {
	InstanceName string   `json:"InstanceName" xml:"InstanceName"`
	UserName     string   `json:"UserName" xml:"UserName"`
	RemoteIp     string   `json:"RemoteIp" xml:"RemoteIp"`
	Id           int64    `json:"Id" xml:"Id"`
	EventName    string   `json:"EventName" xml:"EventName"`
	ProductCode  string   `json:"ProductCode" xml:"ProductCode"`
	Table        Table    `json:"Table" xml:"Table"`
	Package      Package  `json:"Package" xml:"Package"`
	Column       Column   `json:"Column" xml:"Column"`
	Oss          Oss      `json:"Oss" xml:"Oss"`
	Instance     Instance `json:"Instance" xml:"Instance"`
}
