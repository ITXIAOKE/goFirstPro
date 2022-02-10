package sddp

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

// DescribeDataLimits invokes the sddp.DescribeDataLimits API synchronously
func (client *Client) DescribeDataLimits(request *DescribeDataLimitsRequest) (response *DescribeDataLimitsResponse, err error) {
	response = CreateDescribeDataLimitsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataLimitsWithChan invokes the sddp.DescribeDataLimits API asynchronously
func (client *Client) DescribeDataLimitsWithChan(request *DescribeDataLimitsRequest) (<-chan *DescribeDataLimitsResponse, <-chan error) {
	responseChan := make(chan *DescribeDataLimitsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataLimits(request)
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

// DescribeDataLimitsWithCallback invokes the sddp.DescribeDataLimits API asynchronously
func (client *Client) DescribeDataLimitsWithCallback(request *DescribeDataLimitsRequest, callback func(response *DescribeDataLimitsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataLimitsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataLimits(request)
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

// DescribeDataLimitsRequest is the request struct for api DescribeDataLimits
type DescribeDataLimitsRequest struct {
	*requests.RpcRequest
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	ParentId        string           `position:"Query" name:"ParentId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Enable          requests.Integer `position:"Query" name:"Enable"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	CheckStatus     requests.Integer `position:"Query" name:"CheckStatus"`
	DatamaskStatus  requests.Integer `position:"Query" name:"DatamaskStatus"`
	Lang            string           `position:"Query" name:"Lang"`
	ServiceRegionId string           `position:"Query" name:"ServiceRegionId"`
	EngineType      string           `position:"Query" name:"EngineType"`
	AuditStatus     requests.Integer `position:"Query" name:"AuditStatus"`
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	ResourceType    requests.Integer `position:"Query" name:"ResourceType"`
}

// DescribeDataLimitsResponse is the response struct for api DescribeDataLimits
type DescribeDataLimitsResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	CurrentPage int         `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	Items       []DataLimit `json:"Items" xml:"Items"`
}

// CreateDescribeDataLimitsRequest creates a request to invoke DescribeDataLimits API
func CreateDescribeDataLimitsRequest() (request *DescribeDataLimitsRequest) {
	request = &DescribeDataLimitsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeDataLimits", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDataLimitsResponse creates a response to parse from DescribeDataLimits response
func CreateDescribeDataLimitsResponse() (response *DescribeDataLimitsResponse) {
	response = &DescribeDataLimitsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
