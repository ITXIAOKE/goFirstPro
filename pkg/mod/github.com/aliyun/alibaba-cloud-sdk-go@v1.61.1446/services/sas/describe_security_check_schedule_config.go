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

// DescribeSecurityCheckScheduleConfig invokes the sas.DescribeSecurityCheckScheduleConfig API synchronously
func (client *Client) DescribeSecurityCheckScheduleConfig(request *DescribeSecurityCheckScheduleConfigRequest) (response *DescribeSecurityCheckScheduleConfigResponse, err error) {
	response = CreateDescribeSecurityCheckScheduleConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSecurityCheckScheduleConfigWithChan invokes the sas.DescribeSecurityCheckScheduleConfig API asynchronously
func (client *Client) DescribeSecurityCheckScheduleConfigWithChan(request *DescribeSecurityCheckScheduleConfigRequest) (<-chan *DescribeSecurityCheckScheduleConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeSecurityCheckScheduleConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecurityCheckScheduleConfig(request)
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

// DescribeSecurityCheckScheduleConfigWithCallback invokes the sas.DescribeSecurityCheckScheduleConfig API asynchronously
func (client *Client) DescribeSecurityCheckScheduleConfigWithCallback(request *DescribeSecurityCheckScheduleConfigRequest, callback func(response *DescribeSecurityCheckScheduleConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecurityCheckScheduleConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecurityCheckScheduleConfig(request)
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

// DescribeSecurityCheckScheduleConfigRequest is the request struct for api DescribeSecurityCheckScheduleConfig
type DescribeSecurityCheckScheduleConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Lang            string           `position:"Query" name:"Lang"`
}

// DescribeSecurityCheckScheduleConfigResponse is the response struct for api DescribeSecurityCheckScheduleConfig
type DescribeSecurityCheckScheduleConfigResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	RiskCheckJobConfig RiskCheckJobConfig `json:"RiskCheckJobConfig" xml:"RiskCheckJobConfig"`
}

// CreateDescribeSecurityCheckScheduleConfigRequest creates a request to invoke DescribeSecurityCheckScheduleConfig API
func CreateDescribeSecurityCheckScheduleConfigRequest() (request *DescribeSecurityCheckScheduleConfigRequest) {
	request = &DescribeSecurityCheckScheduleConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeSecurityCheckScheduleConfig", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSecurityCheckScheduleConfigResponse creates a response to parse from DescribeSecurityCheckScheduleConfig response
func CreateDescribeSecurityCheckScheduleConfigResponse() (response *DescribeSecurityCheckScheduleConfigResponse) {
	response = &DescribeSecurityCheckScheduleConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}