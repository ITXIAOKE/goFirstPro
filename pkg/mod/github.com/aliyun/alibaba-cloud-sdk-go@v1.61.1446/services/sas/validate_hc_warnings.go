package sas

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

// ValidateHcWarnings invokes the sas.ValidateHcWarnings API synchronously
func (client *Client) ValidateHcWarnings(request *ValidateHcWarningsRequest) (response *ValidateHcWarningsResponse, err error) {
	response = CreateValidateHcWarningsResponse()
	err = client.DoAction(request, response)
	return
}

// ValidateHcWarningsWithChan invokes the sas.ValidateHcWarnings API asynchronously
func (client *Client) ValidateHcWarningsWithChan(request *ValidateHcWarningsRequest) (<-chan *ValidateHcWarningsResponse, <-chan error) {
	responseChan := make(chan *ValidateHcWarningsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ValidateHcWarnings(request)
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

// ValidateHcWarningsWithCallback invokes the sas.ValidateHcWarnings API asynchronously
func (client *Client) ValidateHcWarningsWithCallback(request *ValidateHcWarningsRequest, callback func(response *ValidateHcWarningsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ValidateHcWarningsResponse
		var err error
		defer close(result)
		response, err = client.ValidateHcWarnings(request)
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

// ValidateHcWarningsRequest is the request struct for api ValidateHcWarnings
type ValidateHcWarningsRequest struct {
	*requests.RpcRequest
	RiskIds  string `position:"Query" name:"RiskIds"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Uuids    string `position:"Query" name:"Uuids"`
}

// ValidateHcWarningsResponse is the response struct for api ValidateHcWarnings
type ValidateHcWarningsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateValidateHcWarningsRequest creates a request to invoke ValidateHcWarnings API
func CreateValidateHcWarningsRequest() (request *ValidateHcWarningsRequest) {
	request = &ValidateHcWarningsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ValidateHcWarnings", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateValidateHcWarningsResponse creates a response to parse from ValidateHcWarnings response
func CreateValidateHcWarningsResponse() (response *ValidateHcWarningsResponse) {
	response = &ValidateHcWarningsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}