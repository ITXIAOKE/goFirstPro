package dbs

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

// GetDBListFromAgent invokes the dbs.GetDBListFromAgent API synchronously
func (client *Client) GetDBListFromAgent(request *GetDBListFromAgentRequest) (response *GetDBListFromAgentResponse, err error) {
	response = CreateGetDBListFromAgentResponse()
	err = client.DoAction(request, response)
	return
}

// GetDBListFromAgentWithChan invokes the dbs.GetDBListFromAgent API asynchronously
func (client *Client) GetDBListFromAgentWithChan(request *GetDBListFromAgentRequest) (<-chan *GetDBListFromAgentResponse, <-chan error) {
	responseChan := make(chan *GetDBListFromAgentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDBListFromAgent(request)
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

// GetDBListFromAgentWithCallback invokes the dbs.GetDBListFromAgent API asynchronously
func (client *Client) GetDBListFromAgentWithCallback(request *GetDBListFromAgentRequest, callback func(response *GetDBListFromAgentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDBListFromAgentResponse
		var err error
		defer close(result)
		response, err = client.GetDBListFromAgent(request)
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

// GetDBListFromAgentRequest is the request struct for api GetDBListFromAgent
type GetDBListFromAgentRequest struct {
	*requests.RpcRequest
	SourceEndpointRegion string           `position:"Query" name:"SourceEndpointRegion"`
	BackupGatewayId      requests.Integer `position:"Query" name:"BackupGatewayId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	TaskId               requests.Integer `position:"Query" name:"TaskId"`
}

// GetDBListFromAgentResponse is the response struct for api GetDBListFromAgent
type GetDBListFromAgentResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DbList         DbList `json:"DbList" xml:"DbList"`
}

// CreateGetDBListFromAgentRequest creates a request to invoke GetDBListFromAgent API
func CreateGetDBListFromAgentRequest() (request *GetDBListFromAgentRequest) {
	request = &GetDBListFromAgentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "GetDBListFromAgent", "cbs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDBListFromAgentResponse creates a response to parse from GetDBListFromAgent response
func CreateGetDBListFromAgentResponse() (response *GetDBListFromAgentResponse) {
	response = &GetDBListFromAgentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
