package mse

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

// RestartCluster invokes the mse.RestartCluster API synchronously
func (client *Client) RestartCluster(request *RestartClusterRequest) (response *RestartClusterResponse, err error) {
	response = CreateRestartClusterResponse()
	err = client.DoAction(request, response)
	return
}

// RestartClusterWithChan invokes the mse.RestartCluster API asynchronously
func (client *Client) RestartClusterWithChan(request *RestartClusterRequest) (<-chan *RestartClusterResponse, <-chan error) {
	responseChan := make(chan *RestartClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestartCluster(request)
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

// RestartClusterWithCallback invokes the mse.RestartCluster API asynchronously
func (client *Client) RestartClusterWithCallback(request *RestartClusterRequest, callback func(response *RestartClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestartClusterResponse
		var err error
		defer close(result)
		response, err = client.RestartCluster(request)
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

// RestartClusterRequest is the request struct for api RestartCluster
type RestartClusterRequest struct {
	*requests.RpcRequest
	ClusterId   string `position:"Query" name:"ClusterId"`
	PodNameList string `position:"Query" name:"PodNameList"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	RequestPars string `position:"Query" name:"RequestPars"`
}

// RestartClusterResponse is the response struct for api RestartCluster
type RestartClusterResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateRestartClusterRequest creates a request to invoke RestartCluster API
func CreateRestartClusterRequest() (request *RestartClusterRequest) {
	request = &RestartClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "RestartCluster", "", "")
	request.Method = requests.POST
	return
}

// CreateRestartClusterResponse creates a response to parse from RestartCluster response
func CreateRestartClusterResponse() (response *RestartClusterResponse) {
	response = &RestartClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}