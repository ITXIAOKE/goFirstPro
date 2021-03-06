package gpdb

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

// DBInstanceAttributeInDescribeDBInstanceOnECSAttribute is a nested struct in gpdb response
type DBInstanceAttributeInDescribeDBInstanceOnECSAttribute struct {
	CreationTime          string                                 `json:"CreationTime" xml:"CreationTime"`
	VpcId                 string                                 `json:"VpcId" xml:"VpcId"`
	EncryptionType        string                                 `json:"EncryptionType" xml:"EncryptionType"`
	InstanceDeployType    string                                 `json:"InstanceDeployType" xml:"InstanceDeployType"`
	PayType               string                                 `json:"PayType" xml:"PayType"`
	StorageType           string                                 `json:"StorageType" xml:"StorageType"`
	ConnectionMode        string                                 `json:"ConnectionMode" xml:"ConnectionMode"`
	Port                  string                                 `json:"Port" xml:"Port"`
	LockMode              string                                 `json:"LockMode" xml:"LockMode"`
	EngineVersion         string                                 `json:"EngineVersion" xml:"EngineVersion"`
	MemorySize            int                                    `json:"MemorySize" xml:"MemorySize"`
	SegNodeNum            int                                    `json:"SegNodeNum" xml:"SegNodeNum"`
	ConnectionString      string                                 `json:"ConnectionString" xml:"ConnectionString"`
	InstanceNetworkType   string                                 `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	EncryptionKey         string                                 `json:"EncryptionKey" xml:"EncryptionKey"`
	DBInstanceDescription string                                 `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
	CpuCores              int                                    `json:"CpuCores" xml:"CpuCores"`
	ExpireTime            string                                 `json:"ExpireTime" xml:"ExpireTime"`
	DBInstanceStatus      string                                 `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
	StorageSize           int                                    `json:"StorageSize" xml:"StorageSize"`
	RegionId              string                                 `json:"RegionId" xml:"RegionId"`
	VSwitchId             string                                 `json:"VSwitchId" xml:"VSwitchId"`
	ZoneId                string                                 `json:"ZoneId" xml:"ZoneId"`
	DBInstanceId          string                                 `json:"DBInstanceId" xml:"DBInstanceId"`
	Engine                string                                 `json:"Engine" xml:"Engine"`
	DBInstanceClass       string                                 `json:"DBInstanceClass" xml:"DBInstanceClass"`
	SupportRestore        bool                                   `json:"SupportRestore" xml:"SupportRestore"`
	MinorVersion          string                                 `json:"MinorVersion" xml:"MinorVersion"`
	MasterNodeNum         int                                    `json:"MasterNodeNum" xml:"MasterNodeNum"`
	DBInstanceCategory    string                                 `json:"DBInstanceCategory" xml:"DBInstanceCategory"`
	Tags                  TagsInDescribeDBInstanceOnECSAttribute `json:"Tags" xml:"Tags"`
}
