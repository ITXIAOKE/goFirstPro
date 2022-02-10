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

// UpdateCluster invokes the mse.UpdateCluster API synchronously
func (client *Client) UpdateCluster(request *UpdateClusterRequest) (response *UpdateClusterResponse, err error) {
	response = CreateUpdateClusterResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateClusterWithChan invokes the mse.UpdateCluster API asynchronously
func (client *Client) UpdateClusterWithChan(request *UpdateClusterRequest) (<-chan *UpdateClusterResponse, <-chan error) {
	responseChan := make(chan *UpdateClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCluster(request)
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

// UpdateClusterWithCallback invokes the mse.UpdateCluster API asynchronously
func (client *Client) UpdateClusterWithCallback(request *UpdateClusterRequest, callback func(response *UpdateClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateClusterResponse
		var err error
		defer close(result)
		response, err = client.UpdateCluster(request)
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

// UpdateClusterRequest is the request struct for api UpdateCluster
type UpdateClusterRequest struct {
	*requests.RpcRequest
	ClusterAliasName string `position:"Query" name:"ClusterAliasName"`
	ClusterId        string `position:"Query" name:"ClusterId"`
	InstanceId       string `position:"Query" name:"InstanceId"`
	RequestPars      string `position:"Query" name:"RequestPars"`
}

// UpdateClusterResponse is the response struct for api UpdateCluster
type UpdateClusterResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateClusterRequest creates a request to invoke UpdateCluster API
func CreateUpdateClusterRequest() (request *UpdateClusterRequest) {
	request = &UpdateClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateCluster", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateClusterResponse creates a response to parse from UpdateCluster response
func CreateUpdateClusterResponse() (response *UpdateClusterResponse) {
	response = &UpdateClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
