package live

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

// DescribeLiveStagingIp invokes the live.DescribeLiveStagingIp API synchronously
func (client *Client) DescribeLiveStagingIp(request *DescribeLiveStagingIpRequest) (response *DescribeLiveStagingIpResponse, err error) {
	response = CreateDescribeLiveStagingIpResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStagingIpWithChan invokes the live.DescribeLiveStagingIp API asynchronously
func (client *Client) DescribeLiveStagingIpWithChan(request *DescribeLiveStagingIpRequest) (<-chan *DescribeLiveStagingIpResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStagingIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStagingIp(request)
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

// DescribeLiveStagingIpWithCallback invokes the live.DescribeLiveStagingIp API asynchronously
func (client *Client) DescribeLiveStagingIpWithCallback(request *DescribeLiveStagingIpRequest, callback func(response *DescribeLiveStagingIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStagingIpResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStagingIp(request)
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

// DescribeLiveStagingIpRequest is the request struct for api DescribeLiveStagingIp
type DescribeLiveStagingIpRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStagingIpResponse is the response struct for api DescribeLiveStagingIp
type DescribeLiveStagingIpResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	IPV4s     IPV4s  `json:"IPV4s" xml:"IPV4s"`
}

// CreateDescribeLiveStagingIpRequest creates a request to invoke DescribeLiveStagingIp API
func CreateDescribeLiveStagingIpRequest() (request *DescribeLiveStagingIpRequest) {
	request = &DescribeLiveStagingIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStagingIp", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveStagingIpResponse creates a response to parse from DescribeLiveStagingIp response
func CreateDescribeLiveStagingIpResponse() (response *DescribeLiveStagingIpResponse) {
	response = &DescribeLiveStagingIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
