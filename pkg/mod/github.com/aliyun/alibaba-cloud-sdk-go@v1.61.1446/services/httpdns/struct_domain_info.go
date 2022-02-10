package httpdns

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

// DomainInfo is a nested struct in httpdns response
type DomainInfo struct {
	DomainName     string `json:"DomainName" xml:"DomainName"`
	Resolved       int64  `json:"Resolved" xml:"Resolved"`
	ResolvedHttps  int64  `json:"ResolvedHttps" xml:"ResolvedHttps"`
	Resolved6      int64  `json:"Resolved6" xml:"Resolved6"`
	ResolvedHttps6 int64  `json:"ResolvedHttps6" xml:"ResolvedHttps6"`
}