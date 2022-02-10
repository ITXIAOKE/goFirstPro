package arms

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

// AlertRuleEntity is a nested struct in arms response
type AlertRuleEntity struct {
	Status             string       `json:"Status" xml:"Status"`
	UpdateTime         int64        `json:"UpdateTime" xml:"UpdateTime"`
	ContactGroupIdList string       `json:"ContactGroupIdList" xml:"ContactGroupIdList"`
	CreateTime         int64        `json:"CreateTime" xml:"CreateTime"`
	AlertTitle         string       `json:"AlertTitle" xml:"AlertTitle"`
	UserId             string       `json:"UserId" xml:"UserId"`
	AlertVersion       int          `json:"AlertVersion" xml:"AlertVersion"`
	HostByAlertManager bool         `json:"HostByAlertManager" xml:"HostByAlertManager"`
	AlertType          int          `json:"AlertType" xml:"AlertType"`
	ContactGroupIds    string       `json:"ContactGroupIds" xml:"ContactGroupIds"`
	Config             string       `json:"Config" xml:"Config"`
	RegionId           string       `json:"RegionId" xml:"RegionId"`
	AlertLevel         string       `json:"AlertLevel" xml:"AlertLevel"`
	TaskStatus         string       `json:"TaskStatus" xml:"TaskStatus"`
	Title              string       `json:"Title" xml:"Title"`
	TaskId             int64        `json:"TaskId" xml:"TaskId"`
	Id                 int64        `json:"Id" xml:"Id"`
	AlertWays          []string     `json:"AlertWays" xml:"AlertWays"`
	AlertWay           []string     `json:"AlertWay" xml:"AlertWay"`
	AlarmContext       AlarmContext `json:"AlarmContext" xml:"AlarmContext"`
	AlertRule          AlertRule    `json:"AlertRule" xml:"AlertRule"`
	MetricParam        MetricParam  `json:"MetricParam" xml:"MetricParam"`
	Notice             Notice       `json:"Notice" xml:"Notice"`
}