package outboundbot

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

// Progress is a nested struct in outboundbot response
type Progress struct {
	FailErrorCode    string         `json:"FailErrorCode" xml:"FailErrorCode"`
	TotalJobs        int            `json:"TotalJobs" xml:"TotalJobs"`
	PausedNum        int            `json:"PausedNum" xml:"PausedNum"`
	FeedbackUrl      string         `json:"FeedbackUrl" xml:"FeedbackUrl"`
	Scheduling       int            `json:"Scheduling" xml:"Scheduling"`
	Cancelled        int            `json:"Cancelled" xml:"Cancelled"`
	Executing        int            `json:"Executing" xml:"Executing"`
	HandledJobCount  int            `json:"HandledJobCount" xml:"HandledJobCount"`
	FailReason       string         `json:"FailReason" xml:"FailReason"`
	TotalNotAnswered int            `json:"TotalNotAnswered" xml:"TotalNotAnswered"`
	CancelledNum     int            `json:"CancelledNum" xml:"CancelledNum"`
	StartTime        int64          `json:"StartTime" xml:"StartTime"`
	Failed           int            `json:"Failed" xml:"Failed"`
	ExecutingNum     int            `json:"ExecutingNum" xml:"ExecutingNum"`
	Duration         int            `json:"Duration" xml:"Duration"`
	FailedNum        int            `json:"FailedNum" xml:"FailedNum"`
	Status           string         `json:"Status" xml:"Status"`
	Paused           int            `json:"Paused" xml:"Paused"`
	TotalCompleted   int            `json:"TotalCompleted" xml:"TotalCompleted"`
	TotalJobCount    int            `json:"TotalJobCount" xml:"TotalJobCount"`
	Categories       []KeyValuePair `json:"Categories" xml:"Categories"`
	Briefs           []KeyValuePair `json:"Briefs" xml:"Briefs"`
}
