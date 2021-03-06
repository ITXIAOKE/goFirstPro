package cas

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

// DeleteSSLCertificate invokes the cas.DeleteSSLCertificate API synchronously
func (client *Client) DeleteSSLCertificate(request *DeleteSSLCertificateRequest) (response *DeleteSSLCertificateResponse, err error) {
	response = CreateDeleteSSLCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSSLCertificateWithChan invokes the cas.DeleteSSLCertificate API asynchronously
func (client *Client) DeleteSSLCertificateWithChan(request *DeleteSSLCertificateRequest) (<-chan *DeleteSSLCertificateResponse, <-chan error) {
	responseChan := make(chan *DeleteSSLCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSSLCertificate(request)
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

// DeleteSSLCertificateWithCallback invokes the cas.DeleteSSLCertificate API asynchronously
func (client *Client) DeleteSSLCertificateWithCallback(request *DeleteSSLCertificateRequest, callback func(response *DeleteSSLCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSSLCertificateResponse
		var err error
		defer close(result)
		response, err = client.DeleteSSLCertificate(request)
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

// DeleteSSLCertificateRequest is the request struct for api DeleteSSLCertificate
type DeleteSSLCertificateRequest struct {
	*requests.RpcRequest
	SourceIp       string `position:"Query" name:"SourceIp"`
	CertIdentifier string `position:"Query" name:"CertIdentifier"`
}

// DeleteSSLCertificateResponse is the response struct for api DeleteSSLCertificate
type DeleteSSLCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteSSLCertificateRequest creates a request to invoke DeleteSSLCertificate API
func CreateDeleteSSLCertificateRequest() (request *DeleteSSLCertificateRequest) {
	request = &DeleteSSLCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cas", "2020-06-19", "DeleteSSLCertificate", "cas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteSSLCertificateResponse creates a response to parse from DeleteSSLCertificate response
func CreateDeleteSSLCertificateResponse() (response *DeleteSSLCertificateResponse) {
	response = &DeleteSSLCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
