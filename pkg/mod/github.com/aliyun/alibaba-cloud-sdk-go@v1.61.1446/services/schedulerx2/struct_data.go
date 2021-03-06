package schedulerx2

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

// Data is a nested struct in schedulerx2 response
type Data struct {
	JobId              int64                    `json:"JobId" xml:"JobId"`
	JobInstanceId      int64                    `json:"JobInstanceId" xml:"JobInstanceId"`
	WfInstanceId       int64                    `json:"WfInstanceId" xml:"WfInstanceId"`
	AppGroupId         int64                    `json:"AppGroupId" xml:"AppGroupId"`
	JobInstanceDetail  JobInstanceDetail        `json:"JobInstanceDetail" xml:"JobInstanceDetail"`
	JobConfigInfo      JobConfigInfo            `json:"JobConfigInfo" xml:"JobConfigInfo"`
	WorkFlowInfo       WorkFlowInfo             `json:"WorkFlowInfo" xml:"WorkFlowInfo"`
	WorkFlowNodeInfo   WorkFlowNodeInfo         `json:"WorkFlowNodeInfo" xml:"WorkFlowNodeInfo"`
	WorkerInfos        []WorkerInfo             `json:"WorkerInfos" xml:"WorkerInfos"`
	Jobs               []Job                    `json:"Jobs" xml:"Jobs"`
	Namespaces         []Namespace              `json:"Namespaces" xml:"Namespaces"`
	JobInstanceDetails []JobInstanceDetailsItem `json:"JobInstanceDetails" xml:"JobInstanceDetails"`
	AppGroups          []AppGroup               `json:"AppGroups" xml:"AppGroups"`
}
