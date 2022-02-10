package cloudcallcenter

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

// UploadMediaCallback invokes the cloudcallcenter.UploadMediaCallback API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/uploadmediacallback.html
func (client *Client) UploadMediaCallback(request *UploadMediaCallbackRequest) (response *UploadMediaCallbackResponse, err error) {
	response = CreateUploadMediaCallbackResponse()
	err = client.DoAction(request, response)
	return
}

// UploadMediaCallbackWithChan invokes the cloudcallcenter.UploadMediaCallback API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/uploadmediacallback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadMediaCallbackWithChan(request *UploadMediaCallbackRequest) (<-chan *UploadMediaCallbackResponse, <-chan error) {
	responseChan := make(chan *UploadMediaCallbackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadMediaCallback(request)
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

// UploadMediaCallbackWithCallback invokes the cloudcallcenter.UploadMediaCallback API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/uploadmediacallback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadMediaCallbackWithCallback(request *UploadMediaCallbackRequest, callback func(response *UploadMediaCallbackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadMediaCallbackResponse
		var err error
		defer close(result)
		response, err = client.UploadMediaCallback(request)
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

// UploadMediaCallbackRequest is the request struct for api UploadMediaCallback
type UploadMediaCallbackRequest struct {
	*requests.RpcRequest
	CallbackBody string `position:"Query" name:"CallbackBody"`
	Instance     string `position:"Query" name:"Instance"`
}

// UploadMediaCallbackResponse is the response struct for api UploadMediaCallback
type UploadMediaCallbackResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateUploadMediaCallbackRequest creates a request to invoke UploadMediaCallback API
func CreateUploadMediaCallbackRequest() (request *UploadMediaCallbackRequest) {
	request = &UploadMediaCallbackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "UploadMediaCallback", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadMediaCallbackResponse creates a response to parse from UploadMediaCallback response
func CreateUploadMediaCallbackResponse() (response *UploadMediaCallbackResponse) {
	response = &UploadMediaCallbackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
