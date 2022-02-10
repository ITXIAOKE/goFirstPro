package ccc

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

// ListScenarios invokes the ccc.ListScenarios API synchronously
func (client *Client) ListScenarios(request *ListScenariosRequest) (response *ListScenariosResponse, err error) {
	response = CreateListScenariosResponse()
	err = client.DoAction(request, response)
	return
}

// ListScenariosWithChan invokes the ccc.ListScenarios API asynchronously
func (client *Client) ListScenariosWithChan(request *ListScenariosRequest) (<-chan *ListScenariosResponse, <-chan error) {
	responseChan := make(chan *ListScenariosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListScenarios(request)
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

// ListScenariosWithCallback invokes the ccc.ListScenarios API asynchronously
func (client *Client) ListScenariosWithCallback(request *ListScenariosRequest, callback func(response *ListScenariosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListScenariosResponse
		var err error
		defer close(result)
		response, err = client.ListScenarios(request)
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

// ListScenariosRequest is the request struct for api ListScenarios
type ListScenariosRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListScenariosResponse is the response struct for api ListScenarios
type ListScenariosResponse struct {
	*responses.BaseResponse
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	Code           string     `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Scenarios      []Scenario `json:"Scenarios" xml:"Scenarios"`
}

// CreateListScenariosRequest creates a request to invoke ListScenarios API
func CreateListScenariosRequest() (request *ListScenariosRequest) {
	request = &ListScenariosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListScenarios", "", "")
	request.Method = requests.POST
	return
}

// CreateListScenariosResponse creates a response to parse from ListScenarios response
func CreateListScenariosResponse() (response *ListScenariosResponse) {
	response = &ListScenariosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
