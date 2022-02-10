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

// MetricInfo is a nested struct in dataworks_public response
type MetricInfo struct {
	Message             string `json:"Message" xml:"Message"`
	SumReaderRecords    int64  `json:"SumReaderRecords" xml:"SumReaderRecords"`
	SumWriterRecords    int64  `json:"SumWriterRecords" xml:"SumWriterRecords"`
	LastTaskDelay       int64  `json:"LastTaskDelay" xml:"LastTaskDelay"`
	InsertReaderRecords int64  `json:"InsertReaderRecords" xml:"InsertReaderRecords"`
	UpdateReaderRecords int64  `json:"UpdateReaderRecords" xml:"UpdateReaderRecords"`
	DeleteReaderRecords int64  `json:"DeleteReaderRecords" xml:"DeleteReaderRecords"`
	InsertWriterRecords int64  `json:"InsertWriterRecords" xml:"InsertWriterRecords"`
	UpdateWriterRecords int64  `json:"UpdateWriterRecords" xml:"UpdateWriterRecords"`
	DeleteWriterRecords int64  `json:"DeleteWriterRecords" xml:"DeleteWriterRecords"`
}
