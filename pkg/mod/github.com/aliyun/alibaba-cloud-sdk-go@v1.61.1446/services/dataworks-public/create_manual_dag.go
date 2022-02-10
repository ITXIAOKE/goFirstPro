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

// CreateManualDag invokes the dataworks_public.CreateManualDag API synchronously
func (client *Client) CreateManualDag(request *CreateManualDagRequest) (response *CreateManualDagResponse, err error) {
	response = CreateCreateManualDagResponse()
	err = client.DoAction(request, response)
	return
}

// CreateManualDagWithChan invokes the dataworks_public.CreateManualDag API asynchronously
func (client *Client) CreateManualDagWithChan(request *CreateManualDagRequest) (<-chan *CreateManualDagResponse, <-chan error) {
	responseChan := make(chan *CreateManualDagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateManualDag(request)
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

// CreateManualDagWithCallback invokes the dataworks_public.CreateManualDag API asynchronously
func (client *Client) CreateManualDagWithCallback(request *CreateManualDagRequest, callback func(response *CreateManualDagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateManualDagResponse
		var err error
		defer close(result)
		response, err = client.CreateManualDag(request)
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

// CreateManualDagRequest is the request struct for api CreateManualDag
type CreateManualDagRequest struct {
	*requests.RpcRequest
	ProjectEnv     string `position:"Body" name:"ProjectEnv"`
	ProjectName    string `position:"Body" name:"ProjectName"`
	DagParameters  string `position:"Body" name:"DagParameters"`
	IncludeNodeIds string `position:"Body" name:"IncludeNodeIds"`
	BizDate        string `position:"Body" name:"BizDate"`
	ExcludeNodeIds string `position:"Body" name:"ExcludeNodeIds"`
	FlowName       string `position:"Body" name:"FlowName"`
	NodeParameters string `position:"Body" name:"NodeParameters"`
}

// CreateManualDagResponse is the response struct for api CreateManualDag
type CreateManualDagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	DagId     int64  `json:"DagId" xml:"DagId"`
}

// CreateCreateManualDagRequest creates a request to invoke CreateManualDag API
func CreateCreateManualDagRequest() (request *CreateManualDagRequest) {
	request = &CreateManualDagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateManualDag", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateManualDagResponse creates a response to parse from CreateManualDag response
func CreateCreateManualDagResponse() (response *CreateManualDagResponse) {
	response = &CreateManualDagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
