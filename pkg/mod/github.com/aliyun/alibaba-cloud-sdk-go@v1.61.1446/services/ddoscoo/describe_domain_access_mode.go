package ddoscoo

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

// DescribeDomainAccessMode invokes the ddoscoo.DescribeDomainAccessMode API synchronously
func (client *Client) DescribeDomainAccessMode(request *DescribeDomainAccessModeRequest) (response *DescribeDomainAccessModeResponse, err error) {
	response = CreateDescribeDomainAccessModeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainAccessModeWithChan invokes the ddoscoo.DescribeDomainAccessMode API asynchronously
func (client *Client) DescribeDomainAccessModeWithChan(request *DescribeDomainAccessModeRequest) (<-chan *DescribeDomainAccessModeResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainAccessModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainAccessMode(request)
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

// DescribeDomainAccessModeWithCallback invokes the ddoscoo.DescribeDomainAccessMode API asynchronously
func (client *Client) DescribeDomainAccessModeWithCallback(request *DescribeDomainAccessModeRequest, callback func(response *DescribeDomainAccessModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainAccessModeResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainAccessMode(request)
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

// DescribeDomainAccessModeRequest is the request struct for api DescribeDomainAccessMode
type DescribeDomainAccessModeRequest struct {
	*requests.RpcRequest
	DomainList *[]string `position:"Query" name:"DomainList"  type:"Repeated"`
	SourceIp   string    `position:"Query" name:"SourceIp"`
}

// DescribeDomainAccessModeResponse is the response struct for api DescribeDomainAccessMode
type DescribeDomainAccessModeResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	DomainModeList []DomainMode `json:"DomainModeList" xml:"DomainModeList"`
}

// CreateDescribeDomainAccessModeRequest creates a request to invoke DescribeDomainAccessMode API
func CreateDescribeDomainAccessModeRequest() (request *DescribeDomainAccessModeRequest) {
	request = &DescribeDomainAccessModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "DescribeDomainAccessMode", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainAccessModeResponse creates a response to parse from DescribeDomainAccessMode response
func CreateDescribeDomainAccessModeResponse() (response *DescribeDomainAccessModeResponse) {
	response = &DescribeDomainAccessModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
