package eais

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

// DeleteEai invokes the eais.DeleteEai API synchronously
func (client *Client) DeleteEai(request *DeleteEaiRequest) (response *DeleteEaiResponse, err error) {
	response = CreateDeleteEaiResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEaiWithChan invokes the eais.DeleteEai API asynchronously
func (client *Client) DeleteEaiWithChan(request *DeleteEaiRequest) (<-chan *DeleteEaiResponse, <-chan error) {
	responseChan := make(chan *DeleteEaiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEai(request)
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

// DeleteEaiWithCallback invokes the eais.DeleteEai API asynchronously
func (client *Client) DeleteEaiWithCallback(request *DeleteEaiRequest, callback func(response *DeleteEaiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEaiResponse
		var err error
		defer close(result)
		response, err = client.DeleteEai(request)
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

// DeleteEaiRequest is the request struct for api DeleteEai
type DeleteEaiRequest struct {
	*requests.RpcRequest
	ElasticAcceleratedInstanceId string           `position:"Query" name:"ElasticAcceleratedInstanceId"`
	Force                        requests.Boolean `position:"Query" name:"Force"`
}

// DeleteEaiResponse is the response struct for api DeleteEai
type DeleteEaiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEaiRequest creates a request to invoke DeleteEai API
func CreateDeleteEaiRequest() (request *DeleteEaiRequest) {
	request = &DeleteEaiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eais", "2019-06-24", "DeleteEai", "eais", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEaiResponse creates a response to parse from DeleteEai response
func CreateDeleteEaiResponse() (response *DeleteEaiResponse) {
	response = &DeleteEaiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}