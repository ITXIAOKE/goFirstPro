package cdrs

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

// PaginateDevice invokes the cdrs.PaginateDevice API synchronously
func (client *Client) PaginateDevice(request *PaginateDeviceRequest) (response *PaginateDeviceResponse, err error) {
	response = CreatePaginateDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// PaginateDeviceWithChan invokes the cdrs.PaginateDevice API asynchronously
func (client *Client) PaginateDeviceWithChan(request *PaginateDeviceRequest) (<-chan *PaginateDeviceResponse, <-chan error) {
	responseChan := make(chan *PaginateDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PaginateDevice(request)
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

// PaginateDeviceWithCallback invokes the cdrs.PaginateDevice API asynchronously
func (client *Client) PaginateDeviceWithCallback(request *PaginateDeviceRequest, callback func(response *PaginateDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PaginateDeviceResponse
		var err error
		defer close(result)
		response, err = client.PaginateDevice(request)
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

// PaginateDeviceRequest is the request struct for api PaginateDevice
type PaginateDeviceRequest struct {
	*requests.RpcRequest
	CorpId        string           `position:"Body" name:"CorpId"`
	PageNumber    requests.Integer `position:"Body" name:"PageNumber"`
	CountTotalNum requests.Boolean `position:"Body" name:"CountTotalNum"`
	AppName       string           `position:"Body" name:"AppName"`
	NameSpace     string           `position:"Body" name:"NameSpace"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
}

// PaginateDeviceResponse is the response struct for api PaginateDevice
type PaginateDeviceResponse struct {
	*responses.BaseResponse
	RequestId string               `json:"RequestId" xml:"RequestId"`
	Message   string               `json:"Message" xml:"Message"`
	Code      string               `json:"Code" xml:"Code"`
	Data      DataInPaginateDevice `json:"Data" xml:"Data"`
}

// CreatePaginateDeviceRequest creates a request to invoke PaginateDevice API
func CreatePaginateDeviceRequest() (request *PaginateDeviceRequest) {
	request = &PaginateDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "PaginateDevice", "", "")
	request.Method = requests.POST
	return
}

// CreatePaginateDeviceResponse creates a response to parse from PaginateDevice response
func CreatePaginateDeviceResponse() (response *PaginateDeviceResponse) {
	response = &PaginateDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
