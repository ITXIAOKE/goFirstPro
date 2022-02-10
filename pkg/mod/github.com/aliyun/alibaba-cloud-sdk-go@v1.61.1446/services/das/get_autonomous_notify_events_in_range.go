package das

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

// GetAutonomousNotifyEventsInRange invokes the das.GetAutonomousNotifyEventsInRange API synchronously
func (client *Client) GetAutonomousNotifyEventsInRange(request *GetAutonomousNotifyEventsInRangeRequest) (response *GetAutonomousNotifyEventsInRangeResponse, err error) {
	response = CreateGetAutonomousNotifyEventsInRangeResponse()
	err = client.DoAction(request, response)
	return
}

// GetAutonomousNotifyEventsInRangeWithChan invokes the das.GetAutonomousNotifyEventsInRange API asynchronously
func (client *Client) GetAutonomousNotifyEventsInRangeWithChan(request *GetAutonomousNotifyEventsInRangeRequest) (<-chan *GetAutonomousNotifyEventsInRangeResponse, <-chan error) {
	responseChan := make(chan *GetAutonomousNotifyEventsInRangeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAutonomousNotifyEventsInRange(request)
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

// GetAutonomousNotifyEventsInRangeWithCallback invokes the das.GetAutonomousNotifyEventsInRange API asynchronously
func (client *Client) GetAutonomousNotifyEventsInRangeWithCallback(request *GetAutonomousNotifyEventsInRangeRequest, callback func(response *GetAutonomousNotifyEventsInRangeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAutonomousNotifyEventsInRangeResponse
		var err error
		defer close(result)
		response, err = client.GetAutonomousNotifyEventsInRange(request)
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

// GetAutonomousNotifyEventsInRangeRequest is the request struct for api GetAutonomousNotifyEventsInRange
type GetAutonomousNotifyEventsInRangeRequest struct {
	*requests.RpcRequest
	Context      string `position:"Query" name:"__context"`
	Level        string `position:"Query" name:"Level"`
	EndTime      string `position:"Query" name:"EndTime"`
	StartTime    string `position:"Query" name:"StartTime"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	EventContext string `position:"Query" name:"EventContext"`
	MinLevel     string `position:"Query" name:"MinLevel"`
	PageOffset   string `position:"Query" name:"PageOffset"`
	PageSize     string `position:"Query" name:"PageSize"`
	NodeId       string `position:"Query" name:"NodeId"`
}

// GetAutonomousNotifyEventsInRangeResponse is the response struct for api GetAutonomousNotifyEventsInRange
type GetAutonomousNotifyEventsInRangeResponse struct {
	*responses.BaseResponse
	Message   string                                 `json:"Message" xml:"Message"`
	RequestId string                                 `json:"RequestId" xml:"RequestId"`
	Code      string                                 `json:"Code" xml:"Code"`
	Success   string                                 `json:"Success" xml:"Success"`
	Data      DataInGetAutonomousNotifyEventsInRange `json:"Data" xml:"Data"`
}

// CreateGetAutonomousNotifyEventsInRangeRequest creates a request to invoke GetAutonomousNotifyEventsInRange API
func CreateGetAutonomousNotifyEventsInRangeRequest() (request *GetAutonomousNotifyEventsInRangeRequest) {
	request = &GetAutonomousNotifyEventsInRangeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetAutonomousNotifyEventsInRange", "das", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAutonomousNotifyEventsInRangeResponse creates a response to parse from GetAutonomousNotifyEventsInRange response
func CreateGetAutonomousNotifyEventsInRangeResponse() (response *GetAutonomousNotifyEventsInRangeResponse) {
	response = &GetAutonomousNotifyEventsInRangeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}