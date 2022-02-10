package oam

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

// GetAccessKeyByBucEmplId invokes the oam.GetAccessKeyByBucEmplId API synchronously
// api document: https://help.aliyun.com/api/oam/getaccesskeybybucemplid.html
func (client *Client) GetAccessKeyByBucEmplId(request *GetAccessKeyByBucEmplIdRequest) (response *GetAccessKeyByBucEmplIdResponse, err error) {
	response = CreateGetAccessKeyByBucEmplIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetAccessKeyByBucEmplIdWithChan invokes the oam.GetAccessKeyByBucEmplId API asynchronously
// api document: https://help.aliyun.com/api/oam/getaccesskeybybucemplid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAccessKeyByBucEmplIdWithChan(request *GetAccessKeyByBucEmplIdRequest) (<-chan *GetAccessKeyByBucEmplIdResponse, <-chan error) {
	responseChan := make(chan *GetAccessKeyByBucEmplIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccessKeyByBucEmplId(request)
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

// GetAccessKeyByBucEmplIdWithCallback invokes the oam.GetAccessKeyByBucEmplId API asynchronously
// api document: https://help.aliyun.com/api/oam/getaccesskeybybucemplid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAccessKeyByBucEmplIdWithCallback(request *GetAccessKeyByBucEmplIdRequest, callback func(response *GetAccessKeyByBucEmplIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAccessKeyByBucEmplIdResponse
		var err error
		defer close(result)
		response, err = client.GetAccessKeyByBucEmplId(request)
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

// GetAccessKeyByBucEmplIdRequest is the request struct for api GetAccessKeyByBucEmplId
type GetAccessKeyByBucEmplIdRequest struct {
	*requests.RpcRequest
	ExpectedExpireTime requests.Integer `position:"Query" name:"ExpectedExpireTime"`
	BucEmplId          string           `position:"Query" name:"BucEmplId"`
}

// GetAccessKeyByBucEmplIdResponse is the response struct for api GetAccessKeyByBucEmplId
type GetAccessKeyByBucEmplIdResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
	Data    Data   `json:"Data" xml:"Data"`
}

// CreateGetAccessKeyByBucEmplIdRequest creates a request to invoke GetAccessKeyByBucEmplId API
func CreateGetAccessKeyByBucEmplIdRequest() (request *GetAccessKeyByBucEmplIdRequest) {
	request = &GetAccessKeyByBucEmplIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "GetAccessKeyByBucEmplId", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAccessKeyByBucEmplIdResponse creates a response to parse from GetAccessKeyByBucEmplId response
func CreateGetAccessKeyByBucEmplIdResponse() (response *GetAccessKeyByBucEmplIdResponse) {
	response = &GetAccessKeyByBucEmplIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
