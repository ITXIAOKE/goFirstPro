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

// DeleteInvoice invokes the companyreg.DeleteInvoice API synchronously
func (client *Client) DeleteInvoice(request *DeleteInvoiceRequest) (response *DeleteInvoiceResponse, err error) {
	response = CreateDeleteInvoiceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteInvoiceWithChan invokes the companyreg.DeleteInvoice API asynchronously
func (client *Client) DeleteInvoiceWithChan(request *DeleteInvoiceRequest) (<-chan *DeleteInvoiceResponse, <-chan error) {
	responseChan := make(chan *DeleteInvoiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteInvoice(request)
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

// DeleteInvoiceWithCallback invokes the companyreg.DeleteInvoice API asynchronously
func (client *Client) DeleteInvoiceWithCallback(request *DeleteInvoiceRequest, callback func(response *DeleteInvoiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteInvoiceResponse
		var err error
		defer close(result)
		response, err = client.DeleteInvoice(request)
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

// DeleteInvoiceRequest is the request struct for api DeleteInvoice
type DeleteInvoiceRequest struct {
	*requests.RpcRequest
	BizId string           `position:"Query" name:"BizId"`
	Id    requests.Integer `position:"Query" name:"Id"`
}

// DeleteInvoiceResponse is the response struct for api DeleteInvoice
type DeleteInvoiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateDeleteInvoiceRequest creates a request to invoke DeleteInvoice API
func CreateDeleteInvoiceRequest() (request *DeleteInvoiceRequest) {
	request = &DeleteInvoiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "DeleteInvoice", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteInvoiceResponse creates a response to parse from DeleteInvoice response
func CreateDeleteInvoiceResponse() (response *DeleteInvoiceResponse) {
	response = &DeleteInvoiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}