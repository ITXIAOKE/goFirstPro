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

// DescribeImageSharePermission invokes the ens.DescribeImageSharePermission API synchronously
func (client *Client) DescribeImageSharePermission(request *DescribeImageSharePermissionRequest) (response *DescribeImageSharePermissionResponse, err error) {
	response = CreateDescribeImageSharePermissionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImageSharePermissionWithChan invokes the ens.DescribeImageSharePermission API asynchronously
func (client *Client) DescribeImageSharePermissionWithChan(request *DescribeImageSharePermissionRequest) (<-chan *DescribeImageSharePermissionResponse, <-chan error) {
	responseChan := make(chan *DescribeImageSharePermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImageSharePermission(request)
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

// DescribeImageSharePermissionWithCallback invokes the ens.DescribeImageSharePermission API asynchronously
func (client *Client) DescribeImageSharePermissionWithCallback(request *DescribeImageSharePermissionRequest, callback func(response *DescribeImageSharePermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImageSharePermissionResponse
		var err error
		defer close(result)
		response, err = client.DescribeImageSharePermission(request)
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

// DescribeImageSharePermissionRequest is the request struct for api DescribeImageSharePermission
type DescribeImageSharePermissionRequest struct {
	*requests.RpcRequest
	ImageId    string           `position:"Query" name:"ImageId"`
	PageNumber string           `position:"Query" name:"PageNumber"`
	PageSize   string           `position:"Query" name:"PageSize"`
	AliyunId   requests.Integer `position:"Query" name:"AliyunId"`
}

// DescribeImageSharePermissionResponse is the response struct for api DescribeImageSharePermission
type DescribeImageSharePermissionResponse struct {
	*responses.BaseResponse
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	PageSize   int      `json:"PageSize" xml:"PageSize"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	PageNumber int      `json:"PageNumber" xml:"PageNumber"`
	ImageId    string   `json:"ImageId" xml:"ImageId"`
	Accounts   Accounts `json:"Accounts" xml:"Accounts"`
}

// CreateDescribeImageSharePermissionRequest creates a request to invoke DescribeImageSharePermission API
func CreateDescribeImageSharePermissionRequest() (request *DescribeImageSharePermissionRequest) {
	request = &DescribeImageSharePermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeImageSharePermission", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeImageSharePermissionResponse creates a response to parse from DescribeImageSharePermission response
func CreateDescribeImageSharePermissionResponse() (response *DescribeImageSharePermissionResponse) {
	response = &DescribeImageSharePermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}