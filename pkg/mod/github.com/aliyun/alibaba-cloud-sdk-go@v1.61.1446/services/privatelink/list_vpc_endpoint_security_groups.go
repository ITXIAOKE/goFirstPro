package privatelink

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

// ListVpcEndpointSecurityGroups invokes the privatelink.ListVpcEndpointSecurityGroups API synchronously
func (client *Client) ListVpcEndpointSecurityGroups(request *ListVpcEndpointSecurityGroupsRequest) (response *ListVpcEndpointSecurityGroupsResponse, err error) {
	response = CreateListVpcEndpointSecurityGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListVpcEndpointSecurityGroupsWithChan invokes the privatelink.ListVpcEndpointSecurityGroups API asynchronously
func (client *Client) ListVpcEndpointSecurityGroupsWithChan(request *ListVpcEndpointSecurityGroupsRequest) (<-chan *ListVpcEndpointSecurityGroupsResponse, <-chan error) {
	responseChan := make(chan *ListVpcEndpointSecurityGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVpcEndpointSecurityGroups(request)
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

// ListVpcEndpointSecurityGroupsWithCallback invokes the privatelink.ListVpcEndpointSecurityGroups API asynchronously
func (client *Client) ListVpcEndpointSecurityGroupsWithCallback(request *ListVpcEndpointSecurityGroupsRequest, callback func(response *ListVpcEndpointSecurityGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVpcEndpointSecurityGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListVpcEndpointSecurityGroups(request)
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

// ListVpcEndpointSecurityGroupsRequest is the request struct for api ListVpcEndpointSecurityGroups
type ListVpcEndpointSecurityGroupsRequest struct {
	*requests.RpcRequest
	EndpointId string           `position:"Query" name:"EndpointId"`
	NextToken  string           `position:"Query" name:"NextToken"`
	MaxResults requests.Integer `position:"Query" name:"MaxResults"`
}

// ListVpcEndpointSecurityGroupsResponse is the response struct for api ListVpcEndpointSecurityGroups
type ListVpcEndpointSecurityGroupsResponse struct {
	*responses.BaseResponse
	NextToken      string          `json:"NextToken" xml:"NextToken"`
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	MaxResults     string          `json:"MaxResults" xml:"MaxResults"`
	SecurityGroups []SecurityGroup `json:"SecurityGroups" xml:"SecurityGroups"`
}

// CreateListVpcEndpointSecurityGroupsRequest creates a request to invoke ListVpcEndpointSecurityGroups API
func CreateListVpcEndpointSecurityGroupsRequest() (request *ListVpcEndpointSecurityGroupsRequest) {
	request = &ListVpcEndpointSecurityGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Privatelink", "2020-04-15", "ListVpcEndpointSecurityGroups", "privatelink", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListVpcEndpointSecurityGroupsResponse creates a response to parse from ListVpcEndpointSecurityGroups response
func CreateListVpcEndpointSecurityGroupsResponse() (response *ListVpcEndpointSecurityGroupsResponse) {
	response = &ListVpcEndpointSecurityGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}