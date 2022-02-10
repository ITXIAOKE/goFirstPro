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

// SubType is a nested struct in sddp response
type SubType struct {
	Id                int64  `json:"Id" xml:"Id"`
	Name              string `json:"Name" xml:"Name"`
	Code              string `json:"Code" xml:"Code"`
	Description       string `json:"Description" xml:"Description"`
	Status            int    `json:"Status" xml:"Status"`
	EventHitCount     int    `json:"EventHitCount" xml:"EventHitCount"`
	AdaptedProduct    string `json:"AdaptedProduct" xml:"AdaptedProduct"`
	ConfigCode        string `json:"ConfigCode" xml:"ConfigCode"`
	ConfigContentType int    `json:"ConfigContentType" xml:"ConfigContentType"`
	ConfigDescription string `json:"ConfigDescription" xml:"ConfigDescription"`
	ConfigValue       string `json:"ConfigValue" xml:"ConfigValue"`
}
