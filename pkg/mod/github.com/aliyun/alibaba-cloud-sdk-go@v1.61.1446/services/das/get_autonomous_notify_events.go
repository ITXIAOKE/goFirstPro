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

// GetAutonomousNotifyEvents invokes the das.GetAutonomousNotifyEvents API synchronously
func (client *Client) GetAutonomousNotifyEvents(request *GetAutonomousNotifyEventsRequest) (response *GetAutonomousNotifyEventsResponse, err error) {
	response = CreateGetAutonomousNotifyEventsResponse()
	err = client.DoAction(request, response)
	return
}

// GetAutonomousNotifyEventsWithChan invokes the das.GetAutonomousNotifyEvents API asynchronously
func (client *Client) GetAutonomousNotifyEventsWithChan(request *GetAutonomousNotifyEventsRequest) (<-chan *GetAutonomousNotifyEventsResponse, <-chan error) {
	responseChan := make(chan *GetAutonomousNotifyEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAutonomousNotifyEvents(request)
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

// GetAutonomousNotifyEventsWithCallback invokes the das.GetAutonomousNotifyEvents API asynchronously
func (client *Client) GetAutonomousNotifyEventsWithCallback(request *GetAutonomousNotifyEventsRequest, callback func(response *GetAutonomousNotifyEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAutonomousNotifyEventsResponse
		var err error
		defer close(result)
		response, err = client.GetAutonomousNotifyEvents(request)
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

// GetAutonomousNotifyEventsRequest is the request struct for api GetAutonomousNotifyEvents
type GetAutonomousNotifyEventsRequest struct {
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

// GetAutonomousNotifyEventsResponse is the response struct for api GetAutonomousNotifyEvents
type GetAutonomousNotifyEventsResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateGetAutonomousNotifyEventsRequest creates a request to invoke GetAutonomousNotifyEvents API
func CreateGetAutonomousNotifyEventsRequest() (request *GetAutonomousNotifyEventsRequest) {
	request = &GetAutonomousNotifyEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetAutonomousNotifyEvents", "das", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAutonomousNotifyEventsResponse creates a response to parse from GetAutonomousNotifyEvents response
func CreateGetAutonomousNotifyEventsResponse() (response *GetAutonomousNotifyEventsResponse) {
	response = &GetAutonomousNotifyEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
