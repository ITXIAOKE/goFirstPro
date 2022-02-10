package imm

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

// CreateSet invokes the imm.CreateSet API synchronously
func (client *Client) CreateSet(request *CreateSetRequest) (response *CreateSetResponse, err error) {
	response = CreateCreateSetResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSetWithChan invokes the imm.CreateSet API asynchronously
func (client *Client) CreateSetWithChan(request *CreateSetRequest) (<-chan *CreateSetResponse, <-chan error) {
	responseChan := make(chan *CreateSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSet(request)
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

// CreateSetWithCallback invokes the imm.CreateSet API asynchronously
func (client *Client) CreateSetWithCallback(request *CreateSetRequest, callback func(response *CreateSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSetResponse
		var err error
		defer close(result)
		response, err = client.CreateSet(request)
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

// CreateSetRequest is the request struct for api CreateSet
type CreateSetRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SetName string `position:"Query" name:"SetName"`
	SetId   string `position:"Query" name:"SetId"`
}

// CreateSetResponse is the response struct for api CreateSet
type CreateSetResponse struct {
	*responses.BaseResponse
	ModifyTime  string `json:"ModifyTime" xml:"ModifyTime"`
	VideoCount  int    `json:"VideoCount" xml:"VideoCount"`
	ImageCount  int    `json:"ImageCount" xml:"ImageCount"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	CreateTime  string `json:"CreateTime" xml:"CreateTime"`
	SetName     string `json:"SetName" xml:"SetName"`
	SetId       string `json:"SetId" xml:"SetId"`
	VideoLength int    `json:"VideoLength" xml:"VideoLength"`
	FaceCount   int    `json:"FaceCount" xml:"FaceCount"`
}

// CreateCreateSetRequest creates a request to invoke CreateSet API
func CreateCreateSetRequest() (request *CreateSetRequest) {
	request = &CreateSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateSet", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSetResponse creates a response to parse from CreateSet response
func CreateCreateSetResponse() (response *CreateSetResponse) {
	response = &CreateSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
