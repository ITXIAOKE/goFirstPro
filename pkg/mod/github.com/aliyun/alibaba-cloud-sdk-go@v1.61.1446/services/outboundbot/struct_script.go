package outboundbot

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

// Script is a nested struct in outboundbot response
type Script struct {
	IsDrafted                 bool   `json:"IsDrafted" xml:"IsDrafted"`
	ScriptName                string `json:"ScriptName" xml:"ScriptName"`
	UpdateTime                int64  `json:"UpdateTime" xml:"UpdateTime"`
	MiniPlaybackConfigEnabled bool   `json:"MiniPlaybackConfigEnabled" xml:"MiniPlaybackConfigEnabled"`
	Industry                  string `json:"Industry" xml:"Industry"`
	RejectReason              string `json:"RejectReason" xml:"RejectReason"`
	FailReason                string `json:"FailReason" xml:"FailReason"`
	MiniPlaybackEnabled       bool   `json:"MiniPlaybackEnabled" xml:"MiniPlaybackEnabled"`
	DebugStatus               string `json:"DebugStatus" xml:"DebugStatus"`
	ScriptId                  string `json:"ScriptId" xml:"ScriptId"`
	Scene                     string `json:"Scene" xml:"Scene"`
	ChatbotId                 string `json:"ChatbotId" xml:"ChatbotId"`
	TtsConfig                 string `json:"TtsConfig" xml:"TtsConfig"`
	Status                    string `json:"Status" xml:"Status"`
	ScriptDescription         string `json:"ScriptDescription" xml:"ScriptDescription"`
	IsDebugDrafted            bool   `json:"IsDebugDrafted" xml:"IsDebugDrafted"`
	AsrConfig                 string `json:"AsrConfig" xml:"AsrConfig"`
	AppliedVersion            string `json:"AppliedVersion" xml:"AppliedVersion"`
	NewBargeInEnable          bool   `json:"NewBargeInEnable" xml:"NewBargeInEnable"`
	DebugVersion              string `json:"DebugVersion" xml:"DebugVersion"`
}
