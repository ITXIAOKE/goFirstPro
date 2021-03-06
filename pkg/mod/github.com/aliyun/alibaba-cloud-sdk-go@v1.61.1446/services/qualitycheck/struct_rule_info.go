package qualitycheck

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

// RuleInfo is a nested struct in qualitycheck response
type RuleInfo struct {
	Comments                 string                            `json:"Comments" xml:"Comments"`
	Name                     string                            `json:"Name" xml:"Name"`
	Status                   int                               `json:"Status" xml:"Status"`
	AutoReview               int                               `json:"AutoReview" xml:"AutoReview"`
	CreateTime               string                            `json:"CreateTime" xml:"CreateTime"`
	CreateEmpid              string                            `json:"CreateEmpid" xml:"CreateEmpid"`
	LastUpdateTime           string                            `json:"LastUpdateTime" xml:"LastUpdateTime"`
	IsOnline                 int                               `json:"IsOnline" xml:"IsOnline"`
	Rid                      string                            `json:"Rid" xml:"Rid"`
	ScoreSubName             string                            `json:"ScoreSubName" xml:"ScoreSubName"`
	StartTime                string                            `json:"StartTime" xml:"StartTime"`
	ScoreName                string                            `json:"ScoreName" xml:"ScoreName"`
	EndTime                  string                            `json:"EndTime" xml:"EndTime"`
	ScoreId                  int                               `json:"ScoreId" xml:"ScoreId"`
	RuleLambda               string                            `json:"RuleLambda" xml:"RuleLambda"`
	Weight                   string                            `json:"Weight" xml:"Weight"`
	ScoreSubId               int                               `json:"ScoreSubId" xml:"ScoreSubId"`
	LastUpdateEmpid          string                            `json:"LastUpdateEmpid" xml:"LastUpdateEmpid"`
	RuleScoreType            int                               `json:"RuleScoreType" xml:"RuleScoreType"`
	Type                     int                               `json:"Type" xml:"Type"`
	IsDelete                 int                               `json:"IsDelete" xml:"IsDelete"`
	BusinessCategoryNameList BusinessCategoryNameListInGetRule `json:"BusinessCategoryNameList" xml:"BusinessCategoryNameList"`
	Conditions               ConditionsInConfigDataSet         `json:"Conditions" xml:"Conditions"`
	Rules                    RulesInConfigDataSet              `json:"Rules" xml:"Rules"`
}
