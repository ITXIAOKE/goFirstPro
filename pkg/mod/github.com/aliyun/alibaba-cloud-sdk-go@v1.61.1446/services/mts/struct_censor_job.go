package mts

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

// CensorJob is a nested struct in mts response
type CensorJob struct {
	CreationTime          string                `json:"CreationTime" xml:"CreationTime"`
	State                 string                `json:"State" xml:"State"`
	TitleCensorResult     string                `json:"TitleCensorResult" xml:"TitleCensorResult"`
	Message               string                `json:"Message" xml:"Message"`
	BarrageCensorResult   string                `json:"BarrageCensorResult" xml:"BarrageCensorResult"`
	DescCensorResult      string                `json:"DescCensorResult" xml:"DescCensorResult"`
	ResultSaveObject      string                `json:"ResultSaveObject" xml:"ResultSaveObject"`
	UserData              string                `json:"UserData" xml:"UserData"`
	Code                  string                `json:"Code" xml:"Code"`
	PipelineId            string                `json:"PipelineId" xml:"PipelineId"`
	Id                    string                `json:"Id" xml:"Id"`
	CensorTerrorismResult CensorTerrorismResult `json:"CensorTerrorismResult" xml:"CensorTerrorismResult"`
	Input                 Input                 `json:"Input" xml:"Input"`
	CensorConfig          CensorConfig          `json:"CensorConfig" xml:"CensorConfig"`
	CensorPornResult      CensorPornResult      `json:"CensorPornResult" xml:"CensorPornResult"`
	ImageCensorResults    ImageCensorResults    `json:"ImageCensorResults" xml:"ImageCensorResults"`
}
