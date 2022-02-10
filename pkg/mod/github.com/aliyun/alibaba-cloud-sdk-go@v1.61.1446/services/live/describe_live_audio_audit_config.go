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

// DescribeLiveAudioAuditConfig invokes the live.DescribeLiveAudioAuditConfig API synchronously
func (client *Client) DescribeLiveAudioAuditConfig(request *DescribeLiveAudioAuditConfigRequest) (response *DescribeLiveAudioAuditConfigResponse, err error) {
	response = CreateDescribeLiveAudioAuditConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveAudioAuditConfigWithChan invokes the live.DescribeLiveAudioAuditConfig API asynchronously
func (client *Client) DescribeLiveAudioAuditConfigWithChan(request *DescribeLiveAudioAuditConfigRequest) (<-chan *DescribeLiveAudioAuditConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveAudioAuditConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveAudioAuditConfig(request)
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

// DescribeLiveAudioAuditConfigWithCallback invokes the live.DescribeLiveAudioAuditConfig API asynchronously
func (client *Client) DescribeLiveAudioAuditConfigWithCallback(request *DescribeLiveAudioAuditConfigRequest, callback func(response *DescribeLiveAudioAuditConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveAudioAuditConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveAudioAuditConfig(request)
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

// DescribeLiveAudioAuditConfigRequest is the request struct for api DescribeLiveAudioAuditConfig
type DescribeLiveAudioAuditConfigRequest struct {
	*requests.RpcRequest
	AppName    string           `position:"Query" name:"AppName"`
	StreamName string           `position:"Query" name:"StreamName"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveAudioAuditConfigResponse is the response struct for api DescribeLiveAudioAuditConfig
type DescribeLiveAudioAuditConfigResponse struct {
	*responses.BaseResponse
	RequestId                string                   `json:"RequestId" xml:"RequestId"`
	LiveAudioAuditConfigList LiveAudioAuditConfigList `json:"LiveAudioAuditConfigList" xml:"LiveAudioAuditConfigList"`
}

// CreateDescribeLiveAudioAuditConfigRequest creates a request to invoke DescribeLiveAudioAuditConfig API
func CreateDescribeLiveAudioAuditConfigRequest() (request *DescribeLiveAudioAuditConfigRequest) {
	request = &DescribeLiveAudioAuditConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveAudioAuditConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveAudioAuditConfigResponse creates a response to parse from DescribeLiveAudioAuditConfig response
func CreateDescribeLiveAudioAuditConfigResponse() (response *DescribeLiveAudioAuditConfigResponse) {
	response = &DescribeLiveAudioAuditConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}