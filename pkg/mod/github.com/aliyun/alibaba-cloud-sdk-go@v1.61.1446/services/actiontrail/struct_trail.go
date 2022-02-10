package actiontrail

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

// Trail is a nested struct in actiontrail response
type Trail struct {
	Status              string `json:"Status" xml:"Status"`
	HomeRegion          string `json:"HomeRegion" xml:"HomeRegion"`
	StartLoggingTime    string `json:"StartLoggingTime" xml:"StartLoggingTime"`
	CreateTime          string `json:"CreateTime" xml:"CreateTime"`
	StopLoggingTime     string `json:"StopLoggingTime" xml:"StopLoggingTime"`
	OrganizationId      string `json:"OrganizationId" xml:"OrganizationId"`
	SlsWriteRoleArn     string `json:"SlsWriteRoleArn" xml:"SlsWriteRoleArn"`
	OssBucketLocation   string `json:"OssBucketLocation" xml:"OssBucketLocation"`
	TrailRegion         string `json:"TrailRegion" xml:"TrailRegion"`
	Name                string `json:"Name" xml:"Name"`
	IsOrganizationTrail bool   `json:"IsOrganizationTrail" xml:"IsOrganizationTrail"`
	SlsProjectArn       string `json:"SlsProjectArn" xml:"SlsProjectArn"`
	EventRW             string `json:"EventRW" xml:"EventRW"`
	OssKeyPrefix        string `json:"OssKeyPrefix" xml:"OssKeyPrefix"`
	UpdateTime          string `json:"UpdateTime" xml:"UpdateTime"`
	Region              string `json:"Region" xml:"Region"`
	OssBucketName       string `json:"OssBucketName" xml:"OssBucketName"`
	OssWriteRoleArn     string `json:"OssWriteRoleArn" xml:"OssWriteRoleArn"`
	IsShadowTrail       int64  `json:"IsShadowTrail" xml:"IsShadowTrail"`
}
