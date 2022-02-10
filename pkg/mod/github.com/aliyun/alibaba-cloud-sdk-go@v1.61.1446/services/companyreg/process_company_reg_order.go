package companyreg

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

// ProcessCompanyRegOrder invokes the companyreg.ProcessCompanyRegOrder API synchronously
func (client *Client) ProcessCompanyRegOrder(request *ProcessCompanyRegOrderRequest) (response *ProcessCompanyRegOrderResponse, err error) {
	response = CreateProcessCompanyRegOrderResponse()
	err = client.DoAction(request, response)
	return
}

// ProcessCompanyRegOrderWithChan invokes the companyreg.ProcessCompanyRegOrder API asynchronously
func (client *Client) ProcessCompanyRegOrderWithChan(request *ProcessCompanyRegOrderRequest) (<-chan *ProcessCompanyRegOrderResponse, <-chan error) {
	responseChan := make(chan *ProcessCompanyRegOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ProcessCompanyRegOrder(request)
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

// ProcessCompanyRegOrderWithCallback invokes the companyreg.ProcessCompanyRegOrder API asynchronously
func (client *Client) ProcessCompanyRegOrderWithCallback(request *ProcessCompanyRegOrderRequest, callback func(response *ProcessCompanyRegOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ProcessCompanyRegOrderResponse
		var err error
		defer close(result)
		response, err = client.ProcessCompanyRegOrder(request)
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

// ProcessCompanyRegOrderRequest is the request struct for api ProcessCompanyRegOrder
type ProcessCompanyRegOrderRequest struct {
	*requests.RpcRequest
	ActionType      string           `position:"Query" name:"ActionType"`
	ActionRequestId string           `position:"Query" name:"ActionRequestId"`
	OperatorType    requests.Integer `position:"Query" name:"OperatorType"`
	ActionInfo      string           `position:"Query" name:"ActionInfo"`
	BizCode         string           `position:"Query" name:"BizCode"`
	BizId           string           `position:"Query" name:"BizId"`
	BizSubCode      string           `position:"Query" name:"BizSubCode"`
}

// ProcessCompanyRegOrderResponse is the response struct for api ProcessCompanyRegOrder
type ProcessCompanyRegOrderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateProcessCompanyRegOrderRequest creates a request to invoke ProcessCompanyRegOrder API
func CreateProcessCompanyRegOrderRequest() (request *ProcessCompanyRegOrderRequest) {
	request = &ProcessCompanyRegOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "ProcessCompanyRegOrder", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateProcessCompanyRegOrderResponse creates a response to parse from ProcessCompanyRegOrder response
func CreateProcessCompanyRegOrderResponse() (response *ProcessCompanyRegOrderResponse) {
	response = &ProcessCompanyRegOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}