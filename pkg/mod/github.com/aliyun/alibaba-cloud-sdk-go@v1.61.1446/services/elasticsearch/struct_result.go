package elasticsearch

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

// Result is a nested struct in elasticsearch response
type Result struct {
	Unit                      string                     `json:"unit" xml:"unit"`
	QueueType                 string                     `json:"queueType" xml:"queueType"`
	Endpoints                 string                     `json:"endpoints" xml:"endpoints"`
	IlmPolicy                 string                     `json:"ilmPolicy" xml:"ilmPolicy"`
	Priority                  int                        `json:"priority" xml:"priority"`
	EndpointName              string                     `json:"endpointName" xml:"endpointName"`
	Tags                      map[string]interface{}     `json:"tags" xml:"tags"`
	GmtUpdateTime             string                     `json:"gmtUpdateTime" xml:"gmtUpdateTime"`
	BatchSize                 int                        `json:"batchSize" xml:"batchSize"`
	PipelineId                string                     `json:"pipelineId" xml:"pipelineId"`
	RegionId                  string                     `json:"regionId" xml:"regionId"`
	EndpointDomain            string                     `json:"endpointDomain" xml:"endpointDomain"`
	Env                       string                     `json:"env" xml:"env"`
	VpcId                     string                     `json:"vpcId" xml:"vpcId"`
	Connectable               bool                       `json:"connectable" xml:"connectable"`
	Dps                       map[string]interface{}     `json:"dps" xml:"dps"`
	Phases                    map[string]interface{}     `json:"phases" xml:"phases"`
	Status                    string                     `json:"Status" xml:"Status"`
	PaymentType               string                     `json:"PaymentType" xml:"PaymentType"`
	Enable                    bool                       `json:"enable" xml:"enable"`
	State                     string                     `json:"state" xml:"state"`
	Workers                   int                        `json:"workers" xml:"workers"`
	CreateUrl                 string                     `json:"createUrl" xml:"createUrl"`
	PipelineManagementType    string                     `json:"pipelineManagementType" xml:"pipelineManagementType"`
	InstanceId                string                     `json:"InstanceId" xml:"InstanceId"`
	OutputES                  string                     `json:"OutputES" xml:"OutputES"`
	VswitchId                 string                     `json:"VswitchId" xml:"VswitchId"`
	ReportId                  string                     `json:"reportId" xml:"reportId"`
	VsArea                    string                     `json:"VsArea" xml:"VsArea"`
	CreatedAt                 string                     `json:"CreatedAt" xml:"CreatedAt"`
	Metric                    string                     `json:"metric" xml:"metric"`
	Summary                   float64                    `json:"summary" xml:"summary"`
	Replica                   int64                      `json:"Replica" xml:"Replica"`
	Type                      string                     `json:"type" xml:"type"`
	Version                   string                     `json:"Version" xml:"Version"`
	PipelineStatus            string                     `json:"pipelineStatus" xml:"pipelineStatus"`
	DataStream                bool                       `json:"dataStream" xml:"dataStream"`
	ServiceId                 string                     `json:"serviceId" xml:"serviceId"`
	Description               string                     `json:"description" xml:"description"`
	InstanceCategory          string                     `json:"InstanceCategory" xml:"InstanceCategory"`
	Trigger                   string                     `json:"trigger" xml:"trigger"`
	OwnerId                   string                     `json:"OwnerId" xml:"OwnerId"`
	OutputEsDescription       string                     `json:"OutputEsDescription" xml:"OutputEsDescription"`
	EsInstanceId              string                     `json:"esInstanceId" xml:"esInstanceId"`
	Scene                     string                     `json:"scene" xml:"scene"`
	OutputESUserName          string                     `json:"OutputESUserName" xml:"OutputESUserName"`
	GmtCreatedTime            string                     `json:"gmtCreatedTime" xml:"gmtCreatedTime"`
	ValidateType              string                     `json:"validateType" xml:"validateType"`
	Health                    string                     `json:"health" xml:"health"`
	UserName                  string                     `json:"userName" xml:"userName"`
	ClusterType               string                     `json:"clusterType" xml:"clusterType"`
	NodeAmount                int64                      `json:"NodeAmount" xml:"NodeAmount"`
	ResId                     string                     `json:"resId" xml:"resId"`
	EndTime                   int64                      `json:"EndTime" xml:"EndTime"`
	QueueCheckPointWrites     int                        `json:"queueCheckPointWrites" xml:"queueCheckPointWrites"`
	Name                      string                     `json:"name" xml:"name"`
	MessageWatermark          int64                      `json:"messageWatermark" xml:"messageWatermark"`
	DeployedReplica           int64                      `json:"DeployedReplica" xml:"DeployedReplica"`
	QuartzRegex               string                     `json:"QuartzRegex" xml:"QuartzRegex"`
	EndpointId                string                     `json:"endpointId" xml:"endpointId"`
	Integrity                 float64                    `json:"integrity" xml:"integrity"`
	EnablePublic              bool                       `json:"enablePublic" xml:"enablePublic"`
	OversizedCluster          bool                       `json:"OversizedCluster" xml:"OversizedCluster"`
	ResourceSpec              string                     `json:"ResourceSpec" xml:"ResourceSpec"`
	Region                    string                     `json:"Region" xml:"Region"`
	ClusterId                 string                     `json:"clusterId" xml:"clusterId"`
	FileSize                  int64                      `json:"fileSize" xml:"fileSize"`
	Config                    string                     `json:"config" xml:"config"`
	Value                     int64                      `json:"value" xml:"value"`
	UpdateTime                int64                      `json:"updateTime" xml:"updateTime"`
	IndexTemplate             string                     `json:"indexTemplate" xml:"indexTemplate"`
	LastDayUsage              int64                      `json:"lastDayUsage" xml:"lastDayUsage"`
	QueueMaxBytes             int                        `json:"queueMaxBytes" xml:"queueMaxBytes"`
	BatchDelay                int                        `json:"batchDelay" xml:"batchDelay"`
	CreateTime                int64                      `json:"createTime" xml:"createTime"`
	CurrentUsage              int64                      `json:"currentUsage" xml:"currentUsage"`
	MasterSpec                []string                   `json:"masterSpec" xml:"masterSpec"`
	PrivateNetworkIpWhiteList []string                   `json:"privateNetworkIpWhiteList" xml:"privateNetworkIpWhiteList"`
	EsVersions                []string                   `json:"esVersions" xml:"esVersions"`
	EsIPWhitelist             []string                   `json:"esIPWhitelist" xml:"esIPWhitelist"`
	PublicIpWhitelist         []string                   `json:"publicIpWhitelist" xml:"publicIpWhitelist"`
	KibanaPrivateIPWhitelist  []string                   `json:"kibanaPrivateIPWhitelist" xml:"kibanaPrivateIPWhitelist"`
	Zones                     []string                   `json:"zones" xml:"zones"`
	KibanaIPWhitelist         []string                   `json:"kibanaIPWhitelist" xml:"kibanaIPWhitelist"`
	EsIPBlacklist             []string                   `json:"esIPBlacklist" xml:"esIPBlacklist"`
	IndexPatterns             []string                   `json:"indexPatterns" xml:"indexPatterns"`
	ClientNodeSpec            []string                   `json:"clientNodeSpec" xml:"clientNodeSpec"`
	InstanceSupportNodes      []string                   `json:"instanceSupportNodes" xml:"instanceSupportNodes"`
	PipelineIds               []string                   `json:"pipelineIds" xml:"pipelineIds"`
	Node                      Node                       `json:"node" xml:"node"`
	ElasticExpansionTask      ElasticExpansionTask       `json:"elasticExpansionTask" xml:"elasticExpansionTask"`
	JvmConfine                JvmConfine                 `json:"jvmConfine" xml:"jvmConfine"`
	MetaInfo                  MetaInfo                   `json:"metaInfo" xml:"metaInfo"`
	ElasticShrinkTask         ElasticShrinkTask          `json:"elasticShrinkTask" xml:"elasticShrinkTask"`
	KibanaNodeProperties      KibanaNodeProperties       `json:"kibanaNodeProperties" xml:"kibanaNodeProperties"`
	ClientNodeAmountRange     ClientNodeAmountRange      `json:"clientNodeAmountRange" xml:"clientNodeAmountRange"`
	Template                  Template                   `json:"template" xml:"template"`
	ValidateResult            ValidateResult             `json:"validateResult" xml:"validateResult"`
	OssObject                 OssObject                  `json:"ossObject" xml:"ossObject"`
	ElasticNodeProperties     ElasticNodeProperties      `json:"elasticNodeProperties" xml:"elasticNodeProperties"`
	WarmNodeProperties        WarmNodeProperties         `json:"warmNodeProperties" xml:"warmNodeProperties"`
	MasterDiskList            []Disk                     `json:"masterDiskList" xml:"masterDiskList"`
	ClientNodeDiskList        []Disk                     `json:"clientNodeDiskList" xml:"clientNodeDiskList"`
	NodeConfigurations        []NodeConfigurationsItem   `json:"NodeConfigurations" xml:"NodeConfigurations"`
	ExtendConfigs             []ExtendConfigsItem        `json:"ExtendConfigs" xml:"ExtendConfigs"`
	SupportVersions           []CategoryEntity           `json:"supportVersions" xml:"supportVersions"`
	EsVersionsLatestList      []EsVersionsLatestListItem `json:"esVersionsLatestList" xml:"esVersionsLatestList"`
	NodeSpecList              []NodeSpecListItem         `json:"nodeSpecList" xml:"nodeSpecList"`
	DataDiskList              []DataDiskListItem         `json:"dataDiskList" xml:"dataDiskList"`
	DiagnoseItems             []DiagnoseItemsItem        `json:"diagnoseItems" xml:"diagnoseItems"`
}
