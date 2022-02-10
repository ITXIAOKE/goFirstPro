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

// DescribeOssObjects invokes the sddp.DescribeOssObjects API synchronously
func (client *Client) DescribeOssObjects(request *DescribeOssObjectsRequest) (response *DescribeOssObjectsResponse, err error) {
	response = CreateDescribeOssObjectsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOssObjectsWithChan invokes the sddp.DescribeOssObjects API asynchronously
func (client *Client) DescribeOssObjectsWithChan(request *DescribeOssObjectsRequest) (<-chan *DescribeOssObjectsResponse, <-chan error) {
	responseChan := make(chan *DescribeOssObjectsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOssObjects(request)
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

// DescribeOssObjectsWithCallback invokes the sddp.DescribeOssObjects API asynchronously
func (client *Client) DescribeOssObjectsWithCallback(request *DescribeOssObjectsRequest, callback func(response *DescribeOssObjectsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOssObjectsResponse
		var err error
		defer close(result)
		response, err = client.DescribeOssObjects(request)
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

// DescribeOssObjectsRequest is the request struct for api DescribeOssObjects
type DescribeOssObjectsRequest struct {
	*requests.RpcRequest
	RiskLevels        string           `position:"Query" name:"RiskLevels"`
	NeedRiskCount     requests.Boolean `position:"Query" name:"NeedRiskCount"`
	QueryName         string           `position:"Query" name:"QueryName"`
	RiskLevelId       requests.Integer `position:"Query" name:"RiskLevelId"`
	LastScanTimeEnd   requests.Integer `position:"Query" name:"LastScanTimeEnd"`
	LastScanTimeStart requests.Integer `position:"Query" name:"LastScanTimeStart"`
	SourceIp          string           `position:"Query" name:"SourceIp"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	Lang              string           `position:"Query" name:"Lang"`
	ServiceRegionId   string           `position:"Query" name:"ServiceRegionId"`
	FeatureType       requests.Integer `position:"Query" name:"FeatureType"`
	OrderBy           string           `position:"Query" name:"OrderBy"`
	CurrentPage       requests.Integer `position:"Query" name:"CurrentPage"`
	RuleIds           string           `position:"Query" name:"RuleIds"`
	InstanceId        string           `position:"Query" name:"InstanceId"`
	Name              string           `position:"Query" name:"Name"`
	RuleId            requests.Integer `position:"Query" name:"RuleId"`
	Category          requests.Integer `position:"Query" name:"Category"`
}

// DescribeOssObjectsResponse is the response struct for api DescribeOssObjects
type DescribeOssObjectsResponse struct {
	*responses.BaseResponse
	RequestId   string                       `json:"RequestId" xml:"RequestId"`
	PageSize    int                          `json:"PageSize" xml:"PageSize"`
	CurrentPage int                          `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int                          `json:"TotalCount" xml:"TotalCount"`
	Items       []ColumnInDescribeOssObjects `json:"Items" xml:"Items"`
}

// CreateDescribeOssObjectsRequest creates a request to invoke DescribeOssObjects API
func CreateDescribeOssObjectsRequest() (request *DescribeOssObjectsRequest) {
	request = &DescribeOssObjectsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeOssObjects", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeOssObjectsResponse creates a response to parse from DescribeOssObjects response
func CreateDescribeOssObjectsResponse() (response *DescribeOssObjectsResponse) {
	response = &DescribeOssObjectsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
