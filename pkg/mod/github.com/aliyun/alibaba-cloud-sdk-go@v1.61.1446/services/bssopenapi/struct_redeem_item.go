package bssopenapi

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

// RedeemItem is a nested struct in bssopenapi response
type RedeemItem struct {
	ExpiryTime         string `json:"ExpiryTime" xml:"ExpiryTime"`
	GrantedTime        string `json:"GrantedTime" xml:"GrantedTime"`
	Status             string `json:"Status" xml:"Status"`
	Specification      string `json:"Specification" xml:"Specification"`
	NominalValue       string `json:"NominalValue" xml:"NominalValue"`
	EffectiveTime      string `json:"EffectiveTime" xml:"EffectiveTime"`
	RedeemId           string `json:"RedeemId" xml:"RedeemId"`
	RedeemNo           string `json:"RedeemNo" xml:"RedeemNo"`
	ApplicableProducts string `json:"ApplicableProducts" xml:"ApplicableProducts"`
	Balance            string `json:"Balance" xml:"Balance"`
}
