package dataworks_public

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

// RunManualDagNodes invokes the dataworks_public.RunManualDagNodes API synchronously
func (client *Client) RunManualDagNodes(request *RunManualDagNodesRequest) (response *RunManualDagNodesResponse, err error) {
	response = CreateRunManualDagNodesResponse()
	err = client.DoAction(request, response)
	return
}

// RunManualDagNodesWithChan invokes the dataworks_public.RunManualDagNodes API asynchronously
func (client *Client) RunManualDagNodesWithChan(request *RunManualDagNodesRequest) (<-chan *RunManualDagNodesResponse, <-chan error) {
	responseChan := make(chan *RunManualDagNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunManualDagNodes(request)
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

// RunManualDagNodesWithCallback invokes the dataworks_public.RunManualDagNodes API asynchronously
func (client *Client) RunManualDagNodesWithCallback(request *RunManualDagNodesRequest, callback func(response *RunManualDagNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunManualDagNodesResponse
		var err error
		defer close(result)
		response, err = client.RunManualDagNodes(request)
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

// RunManualDagNodesRequest is the request struct for api RunManualDagNodes
type RunManualDagNodesRequest struct {
	*requests.RpcRequest
	ProjectEnv     string           `position:"Body" name:"ProjectEnv"`
	ProjectName    string           `position:"Body" name:"ProjectName"`
	DagParameters  string           `position:"Body" name:"DagParameters"`
	IncludeNodeIds string           `position:"Body" name:"IncludeNodeIds"`
	BizDate        string           `position:"Body" name:"BizDate"`
	ExcludeNodeIds string           `position:"Body" name:"ExcludeNodeIds"`
	FlowName       string           `position:"Body" name:"FlowName"`
	ProjectId      requests.Integer `position:"Body" name:"ProjectId"`
	NodeParameters string           `position:"Body" name:"NodeParameters"`
}

// RunManualDagNodesResponse is the response struct for api RunManualDagNodes
type RunManualDagNodesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	DagId     int64  `json:"DagId" xml:"DagId"`
}

// CreateRunManualDagNodesRequest creates a request to invoke RunManualDagNodes API
func CreateRunManualDagNodesRequest() (request *RunManualDagNodesRequest) {
	request = &RunManualDagNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "RunManualDagNodes", "", "")
	request.Method = requests.POST
	return
}

// CreateRunManualDagNodesResponse creates a response to parse from RunManualDagNodes response
func CreateRunManualDagNodesResponse() (response *RunManualDagNodesResponse) {
	response = &RunManualDagNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}