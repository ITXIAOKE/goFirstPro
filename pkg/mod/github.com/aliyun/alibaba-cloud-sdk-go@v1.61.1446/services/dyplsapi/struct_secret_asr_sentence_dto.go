package dyplsapi

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

// SecretAsrSentenceDTO is a nested struct in dyplsapi response
type SecretAsrSentenceDTO struct {
	EndTime         int64  `json:"EndTime" xml:"EndTime"`
	SilenceDuration int64  `json:"SilenceDuration" xml:"SilenceDuration"`
	BeginTime       int64  `json:"BeginTime" xml:"BeginTime"`
	Text            string `json:"Text" xml:"Text"`
	ChannelId       int    `json:"ChannelId" xml:"ChannelId"`
	SpeechRate      int    `json:"SpeechRate" xml:"SpeechRate"`
	EmotionValue    string `json:"EmotionValue" xml:"EmotionValue"`
}
