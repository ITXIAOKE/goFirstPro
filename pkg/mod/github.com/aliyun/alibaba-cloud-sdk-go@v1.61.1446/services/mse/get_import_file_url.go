package mse

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

// GetImportFileUrl invokes the mse.GetImportFileUrl API synchronously
func (client *Client) GetImportFileUrl(request *GetImportFileUrlRequest) (response *GetImportFileUrlResponse, err error) {
	response = CreateGetImportFileUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetImportFileUrlWithChan invokes the mse.GetImportFileUrl API asynchronously
func (client *Client) GetImportFileUrlWithChan(request *GetImportFileUrlRequest) (<-chan *GetImportFileUrlResponse, <-chan error) {
	responseChan := make(chan *GetImportFileUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetImportFileUrl(request)
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

// GetImportFileUrlWithCallback invokes the mse.GetImportFileUrl API asynchronously
func (client *Client) GetImportFileUrlWithCallback(request *GetImportFileUrlRequest, callback func(response *GetImportFileUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetImportFileUrlResponse
		var err error
		defer close(result)
		response, err = client.GetImportFileUrl(request)
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

// GetImportFileUrlRequest is the request struct for api GetImportFileUrl
type GetImportFileUrlRequest struct {
	*requests.RpcRequest
	ContentType string `position:"Query" name:"ContentType"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	NamespaceId string `position:"Query" name:"NamespaceId"`
}

// GetImportFileUrlResponse is the response struct for api GetImportFileUrl
type GetImportFileUrlResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Code           int    `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetImportFileUrlRequest creates a request to invoke GetImportFileUrl API
func CreateGetImportFileUrlRequest() (request *GetImportFileUrlRequest) {
	request = &GetImportFileUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetImportFileUrl", "", "")
	request.Method = requests.POST
	return
}

// CreateGetImportFileUrlResponse creates a response to parse from GetImportFileUrl response
func CreateGetImportFileUrlResponse() (response *GetImportFileUrlResponse) {
	response = &GetImportFileUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
