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

// DescribeRuleCategory invokes the sddp.DescribeRuleCategory API synchronously
func (client *Client) DescribeRuleCategory(request *DescribeRuleCategoryRequest) (response *DescribeRuleCategoryResponse, err error) {
	response = CreateDescribeRuleCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRuleCategoryWithChan invokes the sddp.DescribeRuleCategory API asynchronously
func (client *Client) DescribeRuleCategoryWithChan(request *DescribeRuleCategoryRequest) (<-chan *DescribeRuleCategoryResponse, <-chan error) {
	responseChan := make(chan *DescribeRuleCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRuleCategory(request)
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

// DescribeRuleCategoryWithCallback invokes the sddp.DescribeRuleCategory API asynchronously
func (client *Client) DescribeRuleCategoryWithCallback(request *DescribeRuleCategoryRequest, callback func(response *DescribeRuleCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRuleCategoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeRuleCategory(request)
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

// DescribeRuleCategoryRequest is the request struct for api DescribeRuleCategory
type DescribeRuleCategoryRequest struct {
	*requests.RpcRequest
	ProductId requests.Integer `position:"Query" name:"ProductId"`
	SourceIp  string           `position:"Query" name:"SourceIp"`
	Lang      string           `position:"Query" name:"Lang"`
}

// DescribeRuleCategoryResponse is the response struct for api DescribeRuleCategory
type DescribeRuleCategoryResponse struct {
	*responses.BaseResponse
	RequestId   string    `json:"RequestId" xml:"RequestId"`
	RuleList    []Rule    `json:"RuleList" xml:"RuleList"`
	ProductList []Product `json:"ProductList" xml:"ProductList"`
}

// CreateDescribeRuleCategoryRequest creates a request to invoke DescribeRuleCategory API
func CreateDescribeRuleCategoryRequest() (request *DescribeRuleCategoryRequest) {
	request = &DescribeRuleCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeRuleCategory", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRuleCategoryResponse creates a response to parse from DescribeRuleCategory response
func CreateDescribeRuleCategoryResponse() (response *DescribeRuleCategoryResponse) {
	response = &DescribeRuleCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
