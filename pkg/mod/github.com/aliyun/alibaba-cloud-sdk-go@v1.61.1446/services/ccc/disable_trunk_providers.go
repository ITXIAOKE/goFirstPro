package ccc

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

// DisableTrunkProviders invokes the ccc.DisableTrunkProviders API synchronously
func (client *Client) DisableTrunkProviders(request *DisableTrunkProvidersRequest) (response *DisableTrunkProvidersResponse, err error) {
	response = CreateDisableTrunkProvidersResponse()
	err = client.DoAction(request, response)
	return
}

// DisableTrunkProvidersWithChan invokes the ccc.DisableTrunkProviders API asynchronously
func (client *Client) DisableTrunkProvidersWithChan(request *DisableTrunkProvidersRequest) (<-chan *DisableTrunkProvidersResponse, <-chan error) {
	responseChan := make(chan *DisableTrunkProvidersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableTrunkProviders(request)
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

// DisableTrunkProvidersWithCallback invokes the ccc.DisableTrunkProviders API asynchronously
func (client *Client) DisableTrunkProvidersWithCallback(request *DisableTrunkProvidersRequest, callback func(response *DisableTrunkProvidersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableTrunkProvidersResponse
		var err error
		defer close(result)
		response, err = client.DisableTrunkProviders(request)
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

// DisableTrunkProvidersRequest is the request struct for api DisableTrunkProviders
type DisableTrunkProvidersRequest struct {
	*requests.RpcRequest
	ProviderName *[]string `position:"Query" name:"ProviderName"  type:"Repeated"`
}

// DisableTrunkProvidersResponse is the response struct for api DisableTrunkProviders
type DisableTrunkProvidersResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateDisableTrunkProvidersRequest creates a request to invoke DisableTrunkProviders API
func CreateDisableTrunkProvidersRequest() (request *DisableTrunkProvidersRequest) {
	request = &DisableTrunkProvidersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "DisableTrunkProviders", "", "")
	request.Method = requests.POST
	return
}

// CreateDisableTrunkProvidersResponse creates a response to parse from DisableTrunkProviders response
func CreateDisableTrunkProvidersResponse() (response *DisableTrunkProvidersResponse) {
	response = &DisableTrunkProvidersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
