package sgw

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

// FileShare is a nested struct in sgw response
type FileShare struct {
	SerialNumber           string `json:"SerialNumber" xml:"SerialNumber"`
	ClientSideCmk          string `json:"ClientSideCmk" xml:"ClientSideCmk"`
	PartialSyncPaths       string `json:"PartialSyncPaths" xml:"PartialSyncPaths"`
	PathPrefix             string `json:"PathPrefix" xml:"PathPrefix"`
	BucketsStub            bool   `json:"BucketsStub" xml:"BucketsStub"`
	UploadQueue            int64  `json:"UploadQueue" xml:"UploadQueue"`
	DiskId                 string `json:"DiskId" xml:"DiskId"`
	OssUsed                int64  `json:"OssUsed" xml:"OssUsed"`
	DownloadQueue          int64  `json:"DownloadQueue" xml:"DownloadQueue"`
	ActiveMessages         int64  `json:"ActiveMessages" xml:"ActiveMessages"`
	IndexId                string `json:"IndexId" xml:"IndexId"`
	MnsHealth              string `json:"MnsHealth" xml:"MnsHealth"`
	AccessBasedEnumeration bool   `json:"AccessBasedEnumeration" xml:"AccessBasedEnumeration"`
	DownloadRate           int64  `json:"DownloadRate" xml:"DownloadRate"`
	FeLimit                int    `json:"FeLimit" xml:"FeLimit"`
	SupportArchive         bool   `json:"SupportArchive" xml:"SupportArchive"`
	OssEndpoint            string `json:"OssEndpoint" xml:"OssEndpoint"`
	OssHealth              string `json:"OssHealth" xml:"OssHealth"`
	OutRate                int64  `json:"OutRate" xml:"OutRate"`
	ServerSideCmk          string `json:"ServerSideCmk" xml:"ServerSideCmk"`
	RwClientList           string `json:"RwClientList" xml:"RwClientList"`
	State                  string `json:"State" xml:"State"`
	Protocol               string `json:"Protocol" xml:"Protocol"`
	OssBucketSsl           bool   `json:"OssBucketSsl" xml:"OssBucketSsl"`
	DownloadLimit          int    `json:"DownloadLimit" xml:"DownloadLimit"`
	InPlace                bool   `json:"InPlace" xml:"InPlace"`
	RemoteSync             bool   `json:"RemoteSync" xml:"RemoteSync"`
	FileNumLimit           int64  `json:"FileNumLimit" xml:"FileNumLimit"`
	Squash                 string `json:"Squash" xml:"Squash"`
	NfsFullPath            string `json:"NfsFullPath" xml:"NfsFullPath"`
	RemainingMetaSpace     int64  `json:"RemainingMetaSpace" xml:"RemainingMetaSpace"`
	TransferAcceleration   bool   `json:"TransferAcceleration" xml:"TransferAcceleration"`
	Size                   int64  `json:"Size" xml:"Size"`
	ServerSideEncryption   bool   `json:"ServerSideEncryption" xml:"ServerSideEncryption"`
	WindowsAcl             bool   `json:"WindowsAcl" xml:"WindowsAcl"`
	RemoteSyncDownload     bool   `json:"RemoteSyncDownload" xml:"RemoteSyncDownload"`
	ClientSideEncryption   bool   `json:"ClientSideEncryption" xml:"ClientSideEncryption"`
	BucketInfos            string `json:"BucketInfos" xml:"BucketInfos"`
	NfsV4Optimization      bool   `json:"NfsV4Optimization" xml:"NfsV4Optimization"`
	TotalUpload            int64  `json:"TotalUpload" xml:"TotalUpload"`
	DiskType               string `json:"DiskType" xml:"DiskType"`
	Used                   int64  `json:"Used" xml:"Used"`
	IgnoreDelete           bool   `json:"IgnoreDelete" xml:"IgnoreDelete"`
	RoUserList             string `json:"RoUserList" xml:"RoUserList"`
	FsSizeLimit            int64  `json:"FsSizeLimit" xml:"FsSizeLimit"`
	TotalDownload          int64  `json:"TotalDownload" xml:"TotalDownload"`
	Enabled                bool   `json:"Enabled" xml:"Enabled"`
	HighWatermark          int    `json:"HighWatermark" xml:"HighWatermark"`
	KmsRotatePeriod        string `json:"KmsRotatePeriod" xml:"KmsRotatePeriod"`
	Address                string `json:"Address" xml:"Address"`
	PollingInterval        int    `json:"PollingInterval" xml:"PollingInterval"`
	Name                   string `json:"Name" xml:"Name"`
	OssBucketName          string `json:"OssBucketName" xml:"OssBucketName"`
	ExpressSyncId          string `json:"ExpressSyncId" xml:"ExpressSyncId"`
	LagPeriod              int64  `json:"LagPeriod" xml:"LagPeriod"`
	DirectIO               bool   `json:"DirectIO" xml:"DirectIO"`
	CacheMode              string `json:"CacheMode" xml:"CacheMode"`
	InRate                 int64  `json:"InRate" xml:"InRate"`
	LowWatermark           int    `json:"LowWatermark" xml:"LowWatermark"`
	SyncProgress           int    `json:"SyncProgress" xml:"SyncProgress"`
	ServerSideAlgorithm    string `json:"ServerSideAlgorithm" xml:"ServerSideAlgorithm"`
	ObsoleteBuckets        string `json:"ObsoleteBuckets" xml:"ObsoleteBuckets"`
	BeLimit                int    `json:"BeLimit" xml:"BeLimit"`
	LocalPath              string `json:"LocalPath" xml:"LocalPath"`
	RoClientList           string `json:"RoClientList" xml:"RoClientList"`
	RwUserList             string `json:"RwUserList" xml:"RwUserList"`
	FastReclaim            bool   `json:"FastReclaim" xml:"FastReclaim"`
	Browsable              bool   `json:"Browsable" xml:"Browsable"`
	Throttling             bool   `json:"Throttling" xml:"Throttling"`
	BypassCacheRead        bool   `json:"BypassCacheRead" xml:"BypassCacheRead"`
}
