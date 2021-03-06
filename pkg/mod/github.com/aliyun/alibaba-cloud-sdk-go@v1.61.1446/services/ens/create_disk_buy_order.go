package ens

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

// CreateDiskBuyOrder invokes the ens.CreateDiskBuyOrder API synchronously
func (client *Client) CreateDiskBuyOrder(request *CreateDiskBuyOrderRequest) (response *CreateDiskBuyOrderResponse, err error) {
	response = CreateCreateDiskBuyOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDiskBuyOrderWithChan invokes the ens.CreateDiskBuyOrder API asynchronously
func (client *Client) CreateDiskBuyOrderWithChan(request *CreateDiskBuyOrderRequest) (<-chan *CreateDiskBuyOrderResponse, <-chan error) {
	responseChan := make(chan *CreateDiskBuyOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDiskBuyOrder(request)
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

// CreateDiskBuyOrderWithCallback invokes the ens.CreateDiskBuyOrder API asynchronously
func (client *Client) CreateDiskBuyOrderWithCallback(request *CreateDiskBuyOrderRequest, callback func(response *CreateDiskBuyOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDiskBuyOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateDiskBuyOrder(request)
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

// CreateDiskBuyOrderRequest is the request struct for api CreateDiskBuyOrder
type CreateDiskBuyOrderRequest struct {
	*requests.RpcRequest
	OrderDetails string `position:"Query" name:"OrderDetails"`
}

// CreateDiskBuyOrderResponse is the response struct for api CreateDiskBuyOrder
type CreateDiskBuyOrderResponse struct {
	*responses.BaseResponse
	OrderId   string `json:"OrderId" xml:"OrderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDiskBuyOrderRequest creates a request to invoke CreateDiskBuyOrder API
func CreateCreateDiskBuyOrderRequest() (request *CreateDiskBuyOrderRequest) {
	request = &CreateDiskBuyOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CreateDiskBuyOrder", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDiskBuyOrderResponse creates a response to parse from CreateDiskBuyOrder response
func CreateCreateDiskBuyOrderResponse() (response *CreateDiskBuyOrderResponse) {
	response = &CreateDiskBuyOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
