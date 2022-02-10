package rds

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

// DescribeHASwitchConfig invokes the rds.DescribeHASwitchConfig API synchronously
func (client *Client) DescribeHASwitchConfig(request *DescribeHASwitchConfigRequest) (response *DescribeHASwitchConfigResponse, err error) {
	response = CreateDescribeHASwitchConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHASwitchConfigWithChan invokes the rds.DescribeHASwitchConfig API asynchronously
func (client *Client) DescribeHASwitchConfigWithChan(request *DescribeHASwitchConfigRequest) (<-chan *DescribeHASwitchConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeHASwitchConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHASwitchConfig(request)
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

// DescribeHASwitchConfigWithCallback invokes the rds.DescribeHASwitchConfig API asynchronously
func (client *Client) DescribeHASwitchConfigWithCallback(request *DescribeHASwitchConfigRequest, callback func(response *DescribeHASwitchConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHASwitchConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeHASwitchConfig(request)
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

// DescribeHASwitchConfigRequest is the request struct for api DescribeHASwitchConfig
type DescribeHASwitchConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeHASwitchConfigResponse is the response struct for api DescribeHASwitchConfig
type DescribeHASwitchConfigResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	HAConfig     string `json:"HAConfig" xml:"HAConfig"`
	ManualHATime string `json:"ManualHATime" xml:"ManualHATime"`
}

// CreateDescribeHASwitchConfigRequest creates a request to invoke DescribeHASwitchConfig API
func CreateDescribeHASwitchConfigRequest() (request *DescribeHASwitchConfigRequest) {
	request = &DescribeHASwitchConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeHASwitchConfig", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeHASwitchConfigResponse creates a response to parse from DescribeHASwitchConfig response
func CreateDescribeHASwitchConfigResponse() (response *DescribeHASwitchConfigResponse) {
	response = &DescribeHASwitchConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
