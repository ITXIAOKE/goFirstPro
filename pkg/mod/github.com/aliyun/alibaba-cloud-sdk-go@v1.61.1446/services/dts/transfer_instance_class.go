package dts

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

// TransferInstanceClass invokes the dts.TransferInstanceClass API synchronously
func (client *Client) TransferInstanceClass(request *TransferInstanceClassRequest) (response *TransferInstanceClassResponse, err error) {
	response = CreateTransferInstanceClassResponse()
	err = client.DoAction(request, response)
	return
}

// TransferInstanceClassWithChan invokes the dts.TransferInstanceClass API asynchronously
func (client *Client) TransferInstanceClassWithChan(request *TransferInstanceClassRequest) (<-chan *TransferInstanceClassResponse, <-chan error) {
	responseChan := make(chan *TransferInstanceClassResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TransferInstanceClass(request)
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

// TransferInstanceClassWithCallback invokes the dts.TransferInstanceClass API asynchronously
func (client *Client) TransferInstanceClassWithCallback(request *TransferInstanceClassRequest, callback func(response *TransferInstanceClassResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TransferInstanceClassResponse
		var err error
		defer close(result)
		response, err = client.TransferInstanceClass(request)
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

// TransferInstanceClassRequest is the request struct for api TransferInstanceClass
type TransferInstanceClassRequest struct {
	*requests.RpcRequest
	InstanceClass string `position:"Query" name:"InstanceClass"`
	DtsJobId      string `position:"Query" name:"DtsJobId"`
	OrderType     string `position:"Query" name:"OrderType"`
}

// TransferInstanceClassResponse is the response struct for api TransferInstanceClass
type TransferInstanceClassResponse struct {
	*responses.BaseResponse
	EndTime        string `json:"EndTime" xml:"EndTime"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DtsJobId       string `json:"DtsJobId" xml:"DtsJobId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	ChargeType     string `json:"ChargeType" xml:"ChargeType"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	Code           string `json:"Code" xml:"Code"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	InstanceId     string `json:"InstanceId" xml:"InstanceId"`
}

// CreateTransferInstanceClassRequest creates a request to invoke TransferInstanceClass API
func CreateTransferInstanceClassRequest() (request *TransferInstanceClassRequest) {
	request = &TransferInstanceClassRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "TransferInstanceClass", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTransferInstanceClassResponse creates a response to parse from TransferInstanceClass response
func CreateTransferInstanceClassResponse() (response *TransferInstanceClassResponse) {
	response = &TransferInstanceClassResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
