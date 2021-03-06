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

// EnableLayer7CC invokes the ddoscoo.EnableLayer7CC API synchronously
func (client *Client) EnableLayer7CC(request *EnableLayer7CCRequest) (response *EnableLayer7CCResponse, err error) {
	response = CreateEnableLayer7CCResponse()
	err = client.DoAction(request, response)
	return
}

// EnableLayer7CCWithChan invokes the ddoscoo.EnableLayer7CC API asynchronously
func (client *Client) EnableLayer7CCWithChan(request *EnableLayer7CCRequest) (<-chan *EnableLayer7CCResponse, <-chan error) {
	responseChan := make(chan *EnableLayer7CCResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableLayer7CC(request)
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

// EnableLayer7CCWithCallback invokes the ddoscoo.EnableLayer7CC API asynchronously
func (client *Client) EnableLayer7CCWithCallback(request *EnableLayer7CCRequest, callback func(response *EnableLayer7CCResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableLayer7CCResponse
		var err error
		defer close(result)
		response, err = client.EnableLayer7CC(request)
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

// EnableLayer7CCRequest is the request struct for api EnableLayer7CC
type EnableLayer7CCRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
}

// EnableLayer7CCResponse is the response struct for api EnableLayer7CC
type EnableLayer7CCResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableLayer7CCRequest creates a request to invoke EnableLayer7CC API
func CreateEnableLayer7CCRequest() (request *EnableLayer7CCRequest) {
	request = &EnableLayer7CCRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "EnableLayer7CC", "", "")
	request.Method = requests.POST
	return
}

// CreateEnableLayer7CCResponse creates a response to parse from EnableLayer7CC response
func CreateEnableLayer7CCResponse() (response *EnableLayer7CCResponse) {
	response = &EnableLayer7CCResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
