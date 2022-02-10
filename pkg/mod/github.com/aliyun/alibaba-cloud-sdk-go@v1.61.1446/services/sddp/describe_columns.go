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

// DescribeColumns invokes the sddp.DescribeColumns API synchronously
func (client *Client) DescribeColumns(request *DescribeColumnsRequest) (response *DescribeColumnsResponse, err error) {
	response = CreateDescribeColumnsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeColumnsWithChan invokes the sddp.DescribeColumns API asynchronously
func (client *Client) DescribeColumnsWithChan(request *DescribeColumnsRequest) (<-chan *DescribeColumnsResponse, <-chan error) {
	responseChan := make(chan *DescribeColumnsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeColumns(request)
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

// DescribeColumnsWithCallback invokes the sddp.DescribeColumns API asynchronously
func (client *Client) DescribeColumnsWithCallback(request *DescribeColumnsRequest, callback func(response *DescribeColumnsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeColumnsResponse
		var err error
		defer close(result)
		response, err = client.DescribeColumns(request)
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

// DescribeColumnsRequest is the request struct for api DescribeColumns
type DescribeColumnsRequest struct {
	*requests.RpcRequest
	ProductCode   string           `position:"Query" name:"ProductCode"`
	RiskLevels    string           `position:"Query" name:"RiskLevels"`
	RuleName      string           `position:"Query" name:"RuleName"`
	QueryName     string           `position:"Query" name:"QueryName"`
	RiskLevelId   requests.Integer `position:"Query" name:"RiskLevelId"`
	SensLevelName string           `position:"Query" name:"SensLevelName"`
	SourceIp      string           `position:"Query" name:"SourceIp"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	TableId       requests.Integer `position:"Query" name:"TableId"`
	Lang          string           `position:"Query" name:"Lang"`
	TableName     string           `position:"Query" name:"TableName"`
	FeatureType   requests.Integer `position:"Query" name:"FeatureType"`
	OrderBy       string           `position:"Query" name:"OrderBy"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	InstanceId    requests.Integer `position:"Query" name:"InstanceId"`
	InstanceName  string           `position:"Query" name:"InstanceName"`
	Name          string           `position:"Query" name:"Name"`
	RuleId        requests.Integer `position:"Query" name:"RuleId"`
}

// DescribeColumnsResponse is the response struct for api DescribeColumns
type DescribeColumnsResponse struct {
	*responses.BaseResponse
	CurrentPage int      `json:"CurrentPage" xml:"CurrentPage"`
	RequestId   string   `json:"RequestId" xml:"RequestId"`
	PageSize    int      `json:"PageSize" xml:"PageSize"`
	TotalCount  int      `json:"TotalCount" xml:"TotalCount"`
	Items       []Column `json:"Items" xml:"Items"`
}

// CreateDescribeColumnsRequest creates a request to invoke DescribeColumns API
func CreateDescribeColumnsRequest() (request *DescribeColumnsRequest) {
	request = &DescribeColumnsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeColumns", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeColumnsResponse creates a response to parse from DescribeColumns response
func CreateDescribeColumnsResponse() (response *DescribeColumnsResponse) {
	response = &DescribeColumnsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
