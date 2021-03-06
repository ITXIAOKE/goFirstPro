package smartag

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

// CloudConnectNetwork is a nested struct in smartag response
type CloudConnectNetwork struct {
	IsDefault               bool   `json:"IsDefault" xml:"IsDefault"`
	CreateTime              int64  `json:"CreateTime" xml:"CreateTime"`
	AssociatedCenOwnerId    string `json:"AssociatedCenOwnerId" xml:"AssociatedCenOwnerId"`
	AssociatedCloudBoxCount string `json:"AssociatedCloudBoxCount" xml:"AssociatedCloudBoxCount"`
	InterworkingStatus      string `json:"InterworkingStatus" xml:"InterworkingStatus"`
	CcnId                   string `json:"CcnId" xml:"CcnId"`
	AvailableCloudBoxCount  string `json:"AvailableCloudBoxCount" xml:"AvailableCloudBoxCount"`
	CidrBlock               string `json:"CidrBlock" xml:"CidrBlock"`
	Description             string `json:"Description" xml:"Description"`
	SnatCidrBlock           string `json:"SnatCidrBlock" xml:"SnatCidrBlock"`
	AssociatedCenId         string `json:"AssociatedCenId" xml:"AssociatedCenId"`
	Name                    string `json:"Name" xml:"Name"`
	ResourceGroupId         string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Tags                    Tags   `json:"Tags" xml:"Tags"`
}
