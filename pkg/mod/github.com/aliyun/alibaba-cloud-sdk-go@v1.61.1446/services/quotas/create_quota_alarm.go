package quotas

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateQuotaAlarm invokes the quotas.CreateQuotaAlarm API synchronously
func (client *Client) CreateQuotaAlarm(request *CreateQuotaAlarmRequest) (response *CreateQuotaAlarmResponse, err error) {
	response = CreateCreateQuotaAlarmResponse()
	err = client.DoAction(request, response)
	return
}

// CreateQuotaAlarmWithChan invokes the quotas.CreateQuotaAlarm API asynchronously
func (client *Client) CreateQuotaAlarmWithChan(request *CreateQuotaAlarmRequest) (<-chan *CreateQuotaAlarmResponse, <-chan error) {
	responseChan := make(chan *CreateQuotaAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateQuotaAlarm(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateQuotaAlarmWithCallback invokes the quotas.CreateQuotaAlarm API asynchronously
func (client *Client) CreateQuotaAlarmWithCallback(request *CreateQuotaAlarmRequest, callback func(response *CreateQuotaAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateQuotaAlarmResponse
		var err error
		defer close(result)
		response, err = client.CreateQuotaAlarm(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateQuotaAlarmRequest is the request struct for api CreateQuotaAlarm
type CreateQuotaAlarmRequest struct {
	*requests.RpcRequest
	ProductCode      string                             `position:"Body" name:"ProductCode"`
	WebHook          string                             `position:"Body" name:"WebHook"`
	Threshold        requests.Float                     `position:"Body" name:"Threshold"`
	QuotaActionCode  string                             `position:"Body" name:"QuotaActionCode"`
	ThresholdType    string                             `position:"Body" name:"ThresholdType"`
	QuotaDimensions  *[]CreateQuotaAlarmQuotaDimensions `position:"Body" name:"QuotaDimensions"  type:"Repeated"`
	ThresholdPercent requests.Float                     `position:"Body" name:"ThresholdPercent"`
	AlarmName        string                             `position:"Body" name:"AlarmName"`
}

// CreateQuotaAlarmQuotaDimensions is a repeated param struct in CreateQuotaAlarmRequest
type CreateQuotaAlarmQuotaDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateQuotaAlarmResponse is the response struct for api CreateQuotaAlarm
type CreateQuotaAlarmResponse struct {
	*responses.BaseResponse
	AlarmId   string `json:"AlarmId" xml:"AlarmId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateQuotaAlarmRequest creates a request to invoke CreateQuotaAlarm API
func CreateCreateQuotaAlarmRequest() (request *CreateQuotaAlarmRequest) {
	request = &CreateQuotaAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quotas", "2020-05-10", "CreateQuotaAlarm", "quotas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateQuotaAlarmResponse creates a response to parse from CreateQuotaAlarm response
func CreateCreateQuotaAlarmResponse() (response *CreateQuotaAlarmResponse) {
	response = &CreateQuotaAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}