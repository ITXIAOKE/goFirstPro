package cloudauth

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

// IdCardInfo is a nested struct in cloudauth response
type IdCardInfo struct {
	Name          string `json:"Name" xml:"Name"`
	Sex           string `json:"Sex" xml:"Sex"`
	EndDate       string `json:"EndDate" xml:"EndDate"`
	Nationality   string `json:"Nationality" xml:"Nationality"`
	BackImageUrl  string `json:"BackImageUrl" xml:"BackImageUrl"`
	Authority     string `json:"Authority" xml:"Authority"`
	Birth         string `json:"Birth" xml:"Birth"`
	Address       string `json:"Address" xml:"Address"`
	StartDate     string `json:"StartDate" xml:"StartDate"`
	Number        string `json:"Number" xml:"Number"`
	FrontImageUrl string `json:"FrontImageUrl" xml:"FrontImageUrl"`
}
