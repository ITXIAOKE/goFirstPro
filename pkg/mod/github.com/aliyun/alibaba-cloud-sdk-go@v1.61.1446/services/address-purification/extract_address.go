package address_purification

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

// ExtractAddress invokes the address_purification.ExtractAddress API synchronously
func (client *Client) ExtractAddress(request *ExtractAddressRequest) (response *ExtractAddressResponse, err error) {
	response = CreateExtractAddressResponse()
	err = client.DoAction(request, response)
	return
}

// ExtractAddressWithChan invokes the address_purification.ExtractAddress API asynchronously
func (client *Client) ExtractAddressWithChan(request *ExtractAddressRequest) (<-chan *ExtractAddressResponse, <-chan error) {
	responseChan := make(chan *ExtractAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExtractAddress(request)
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

// ExtractAddressWithCallback invokes the address_purification.ExtractAddress API asynchronously
func (client *Client) ExtractAddressWithCallback(request *ExtractAddressRequest, callback func(response *ExtractAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExtractAddressResponse
		var err error
		defer close(result)
		response, err = client.ExtractAddress(request)
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

// ExtractAddressRequest is the request struct for api ExtractAddress
type ExtractAddressRequest struct {
	*requests.RpcRequest
	DefaultProvince string `position:"Body" name:"DefaultProvince"`
	ServiceCode     string `position:"Body" name:"ServiceCode"`
	DefaultCity     string `position:"Body" name:"DefaultCity"`
	DefaultDistrict string `position:"Body" name:"DefaultDistrict"`
	AppKey          string `position:"Body" name:"AppKey"`
	Text            string `position:"Body" name:"Text"`
}

// ExtractAddressResponse is the response struct for api ExtractAddress
type ExtractAddressResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateExtractAddressRequest creates a request to invoke ExtractAddress API
func CreateExtractAddressRequest() (request *ExtractAddressRequest) {
	request = &ExtractAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("address-purification", "2019-11-18", "ExtractAddress", "addrp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExtractAddressResponse creates a response to parse from ExtractAddress response
func CreateExtractAddressResponse() (response *ExtractAddressResponse) {
	response = &ExtractAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
