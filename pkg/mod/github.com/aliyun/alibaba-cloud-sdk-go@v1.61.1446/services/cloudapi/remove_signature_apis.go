package cloudapi

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

// RemoveSignatureApis invokes the cloudapi.RemoveSignatureApis API synchronously
// api document: https://help.aliyun.com/api/cloudapi/removesignatureapis.html
func (client *Client) RemoveSignatureApis(request *RemoveSignatureApisRequest) (response *RemoveSignatureApisResponse, err error) {
	response = CreateRemoveSignatureApisResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveSignatureApisWithChan invokes the cloudapi.RemoveSignatureApis API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/removesignatureapis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveSignatureApisWithChan(request *RemoveSignatureApisRequest) (<-chan *RemoveSignatureApisResponse, <-chan error) {
	responseChan := make(chan *RemoveSignatureApisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveSignatureApis(request)
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

// RemoveSignatureApisWithCallback invokes the cloudapi.RemoveSignatureApis API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/removesignatureapis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveSignatureApisWithCallback(request *RemoveSignatureApisRequest, callback func(response *RemoveSignatureApisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveSignatureApisResponse
		var err error
		defer close(result)
		response, err = client.RemoveSignatureApis(request)
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

// RemoveSignatureApisRequest is the request struct for api RemoveSignatureApis
type RemoveSignatureApisRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	GroupId       string `position:"Query" name:"GroupId"`
	SignatureId   string `position:"Query" name:"SignatureId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ApiIds        string `position:"Query" name:"ApiIds"`
}

// RemoveSignatureApisResponse is the response struct for api RemoveSignatureApis
type RemoveSignatureApisResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveSignatureApisRequest creates a request to invoke RemoveSignatureApis API
func CreateRemoveSignatureApisRequest() (request *RemoveSignatureApisRequest) {
	request = &RemoveSignatureApisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveSignatureApis", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveSignatureApisResponse creates a response to parse from RemoveSignatureApis response
func CreateRemoveSignatureApisResponse() (response *RemoveSignatureApisResponse) {
	response = &RemoveSignatureApisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
