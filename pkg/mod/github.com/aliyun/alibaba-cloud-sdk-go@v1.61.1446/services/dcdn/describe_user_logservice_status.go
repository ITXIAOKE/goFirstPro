package dcdn

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

// DescribeUserLogserviceStatus invokes the dcdn.DescribeUserLogserviceStatus API synchronously
func (client *Client) DescribeUserLogserviceStatus(request *DescribeUserLogserviceStatusRequest) (response *DescribeUserLogserviceStatusResponse, err error) {
	response = CreateDescribeUserLogserviceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserLogserviceStatusWithChan invokes the dcdn.DescribeUserLogserviceStatus API asynchronously
func (client *Client) DescribeUserLogserviceStatusWithChan(request *DescribeUserLogserviceStatusRequest) (<-chan *DescribeUserLogserviceStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeUserLogserviceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserLogserviceStatus(request)
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

// DescribeUserLogserviceStatusWithCallback invokes the dcdn.DescribeUserLogserviceStatus API asynchronously
func (client *Client) DescribeUserLogserviceStatusWithCallback(request *DescribeUserLogserviceStatusRequest, callback func(response *DescribeUserLogserviceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserLogserviceStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserLogserviceStatus(request)
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

// DescribeUserLogserviceStatusRequest is the request struct for api DescribeUserLogserviceStatus
type DescribeUserLogserviceStatusRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeUserLogserviceStatusResponse is the response struct for api DescribeUserLogserviceStatus
type DescribeUserLogserviceStatusResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Enabled       bool   `json:"Enabled" xml:"Enabled"`
	OnService     bool   `json:"OnService" xml:"OnService"`
	InDebt        bool   `json:"InDebt" xml:"InDebt"`
	InDebtOverdue bool   `json:"InDebtOverdue" xml:"InDebtOverdue"`
}

// CreateDescribeUserLogserviceStatusRequest creates a request to invoke DescribeUserLogserviceStatus API
func CreateDescribeUserLogserviceStatusRequest() (request *DescribeUserLogserviceStatusRequest) {
	request = &DescribeUserLogserviceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeUserLogserviceStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeUserLogserviceStatusResponse creates a response to parse from DescribeUserLogserviceStatus response
func CreateDescribeUserLogserviceStatusResponse() (response *DescribeUserLogserviceStatusResponse) {
	response = &DescribeUserLogserviceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
