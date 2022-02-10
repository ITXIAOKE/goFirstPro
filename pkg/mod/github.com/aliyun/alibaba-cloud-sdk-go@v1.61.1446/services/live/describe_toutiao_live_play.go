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

// DescribeToutiaoLivePlay invokes the live.DescribeToutiaoLivePlay API synchronously
func (client *Client) DescribeToutiaoLivePlay(request *DescribeToutiaoLivePlayRequest) (response *DescribeToutiaoLivePlayResponse, err error) {
	response = CreateDescribeToutiaoLivePlayResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeToutiaoLivePlayWithChan invokes the live.DescribeToutiaoLivePlay API asynchronously
func (client *Client) DescribeToutiaoLivePlayWithChan(request *DescribeToutiaoLivePlayRequest) (<-chan *DescribeToutiaoLivePlayResponse, <-chan error) {
	responseChan := make(chan *DescribeToutiaoLivePlayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeToutiaoLivePlay(request)
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

// DescribeToutiaoLivePlayWithCallback invokes the live.DescribeToutiaoLivePlay API asynchronously
func (client *Client) DescribeToutiaoLivePlayWithCallback(request *DescribeToutiaoLivePlayRequest, callback func(response *DescribeToutiaoLivePlayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeToutiaoLivePlayResponse
		var err error
		defer close(result)
		response, err = client.DescribeToutiaoLivePlay(request)
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

// DescribeToutiaoLivePlayRequest is the request struct for api DescribeToutiaoLivePlay
type DescribeToutiaoLivePlayRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	Stream    string           `position:"Query" name:"Stream"`
	App       string           `position:"Query" name:"App"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	Domain    string           `position:"Query" name:"Domain"`
}

// DescribeToutiaoLivePlayResponse is the response struct for api DescribeToutiaoLivePlay
type DescribeToutiaoLivePlayResponse struct {
	*responses.BaseResponse
	RequestId   string        `json:"RequestId" xml:"RequestId"`
	Description string        `json:"Description" xml:"Description"`
	Content     []ContentItem `json:"Content" xml:"Content"`
}

// CreateDescribeToutiaoLivePlayRequest creates a request to invoke DescribeToutiaoLivePlay API
func CreateDescribeToutiaoLivePlayRequest() (request *DescribeToutiaoLivePlayRequest) {
	request = &DescribeToutiaoLivePlayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeToutiaoLivePlay", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeToutiaoLivePlayResponse creates a response to parse from DescribeToutiaoLivePlay response
func CreateDescribeToutiaoLivePlayResponse() (response *DescribeToutiaoLivePlayResponse) {
	response = &DescribeToutiaoLivePlayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
