package dataworks_public

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

// DeletedFile is a nested struct in dataworks_public response
type DeletedFile struct {
	Owner          string `json:"Owner" xml:"Owner"`
	FileType       int64  `json:"FileType" xml:"FileType"`
	CurrentVersion int64  `json:"CurrentVersion" xml:"CurrentVersion"`
	BusinessId     int64  `json:"BusinessId" xml:"BusinessId"`
	FileName       string `json:"FileName" xml:"FileName"`
	DataSourceName string `json:"DataSourceName" xml:"DataSourceName"`
	UseType        string `json:"UseType" xml:"UseType"`
	FolderId       string `json:"FolderId" xml:"FolderId"`
	ParentFileId   int64  `json:"ParentFileId" xml:"ParentFileId"`
	Content        string `json:"Content" xml:"Content"`
	NodeId         int64  `json:"NodeId" xml:"NodeId"`
	FileId         int64  `json:"FileId" xml:"FileId"`
}
