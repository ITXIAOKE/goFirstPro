package retailcloud

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

// ListClusterNode invokes the retailcloud.ListClusterNode API synchronously
func (client *Client) ListClusterNode(request *ListClusterNodeRequest) (response *ListClusterNodeResponse, err error) {
	response = CreateListClusterNodeResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterNodeWithChan invokes the retailcloud.ListClusterNode API asynchronously
func (client *Client) ListClusterNodeWithChan(request *ListClusterNodeRequest) (<-chan *ListClusterNodeResponse, <-chan error) {
	responseChan := make(chan *ListClusterNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterNode(request)
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

// ListClusterNodeWithCallback invokes the retailcloud.ListClusterNode API asynchronously
func (client *Client) ListClusterNodeWithCallback(request *ListClusterNodeRequest, callback func(response *ListClusterNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterNodeResponse
		var err error
		defer close(result)
		response, err = client.ListClusterNode(request)
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

// ListClusterNodeRequest is the request struct for api ListClusterNode
type ListClusterNodeRequest struct {
	*requests.RpcRequest
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	PageNum           requests.Integer `position:"Query" name:"PageNum"`
	ClusterInstanceId string           `position:"Query" name:"ClusterInstanceId"`
}

// ListClusterNodeResponse is the response struct for api ListClusterNode
type ListClusterNodeResponse struct {
	*responses.BaseResponse
	Code       int               `json:"Code" xml:"Code"`
	ErrorMsg   string            `json:"ErrorMsg" xml:"ErrorMsg"`
	PageNumber int               `json:"PageNumber" xml:"PageNumber"`
	PageSize   int               `json:"PageSize" xml:"PageSize"`
	RequestId  string            `json:"RequestId" xml:"RequestId"`
	TotalCount int64             `json:"TotalCount" xml:"TotalCount"`
	Data       []ClusterNodeInfo `json:"Data" xml:"Data"`
}

// CreateListClusterNodeRequest creates a request to invoke ListClusterNode API
func CreateListClusterNodeRequest() (request *ListClusterNodeRequest) {
	request = &ListClusterNodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ListClusterNode", "", "")
	request.Method = requests.GET
	return
}

// CreateListClusterNodeResponse creates a response to parse from ListClusterNode response
func CreateListClusterNodeResponse() (response *ListClusterNodeResponse) {
	response = &ListClusterNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
