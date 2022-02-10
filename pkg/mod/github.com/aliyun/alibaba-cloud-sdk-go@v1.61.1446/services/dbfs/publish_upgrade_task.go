package dbfs

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

// PublishUpgradeTask invokes the dbfs.PublishUpgradeTask API synchronously
func (client *Client) PublishUpgradeTask(request *PublishUpgradeTaskRequest) (response *PublishUpgradeTaskResponse, err error) {
	response = CreatePublishUpgradeTaskResponse()
	err = client.DoAction(request, response)
	return
}

// PublishUpgradeTaskWithChan invokes the dbfs.PublishUpgradeTask API asynchronously
func (client *Client) PublishUpgradeTaskWithChan(request *PublishUpgradeTaskRequest) (<-chan *PublishUpgradeTaskResponse, <-chan error) {
	responseChan := make(chan *PublishUpgradeTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublishUpgradeTask(request)
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

// PublishUpgradeTaskWithCallback invokes the dbfs.PublishUpgradeTask API asynchronously
func (client *Client) PublishUpgradeTaskWithCallback(request *PublishUpgradeTaskRequest, callback func(response *PublishUpgradeTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublishUpgradeTaskResponse
		var err error
		defer close(result)
		response, err = client.PublishUpgradeTask(request)
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

// PublishUpgradeTaskRequest is the request struct for api PublishUpgradeTask
type PublishUpgradeTaskRequest struct {
	*requests.RpcRequest
	PageNumber        requests.Integer `position:"Query" name:"PageNumber"`
	BatchStrategyList string           `position:"Query" name:"BatchStrategyList"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
}

// PublishUpgradeTaskResponse is the response struct for api PublishUpgradeTask
type PublishUpgradeTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePublishUpgradeTaskRequest creates a request to invoke PublishUpgradeTask API
func CreatePublishUpgradeTaskRequest() (request *PublishUpgradeTaskRequest) {
	request = &PublishUpgradeTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "PublishUpgradeTask", "", "")
	request.Method = requests.POST
	return
}

// CreatePublishUpgradeTaskResponse creates a response to parse from PublishUpgradeTask response
func CreatePublishUpgradeTaskResponse() (response *PublishUpgradeTaskResponse) {
	response = &PublishUpgradeTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
