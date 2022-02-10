package foas

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

// ExpandCluster invokes the foas.ExpandCluster API synchronously
func (client *Client) ExpandCluster(request *ExpandClusterRequest) (response *ExpandClusterResponse, err error) {
	response = CreateExpandClusterResponse()
	err = client.DoAction(request, response)
	return
}

// ExpandClusterWithChan invokes the foas.ExpandCluster API asynchronously
func (client *Client) ExpandClusterWithChan(request *ExpandClusterRequest) (<-chan *ExpandClusterResponse, <-chan error) {
	responseChan := make(chan *ExpandClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExpandCluster(request)
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

// ExpandClusterWithCallback invokes the foas.ExpandCluster API asynchronously
func (client *Client) ExpandClusterWithCallback(request *ExpandClusterRequest, callback func(response *ExpandClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExpandClusterResponse
		var err error
		defer close(result)
		response, err = client.ExpandCluster(request)
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

// ExpandClusterRequest is the request struct for api ExpandCluster
type ExpandClusterRequest struct {
	*requests.RoaRequest
	Count       requests.Integer `position:"Body" name:"count"`
	Model       string           `position:"Body" name:"model"`
	UserVSwitch string           `position:"Body" name:"userVSwitch"`
	ClusterId   string           `position:"Path" name:"clusterId"`
}

// ExpandClusterResponse is the response struct for api ExpandCluster
type ExpandClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateExpandClusterRequest creates a request to invoke ExpandCluster API
func CreateExpandClusterRequest() (request *ExpandClusterRequest) {
	request = &ExpandClusterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "ExpandCluster", "/api/v2/clusters/[clusterId]/expand", "foas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateExpandClusterResponse creates a response to parse from ExpandCluster response
func CreateExpandClusterResponse() (response *ExpandClusterResponse) {
	response = &ExpandClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}