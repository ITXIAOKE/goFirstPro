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

// PutProject invokes the imm.PutProject API synchronously
func (client *Client) PutProject(request *PutProjectRequest) (response *PutProjectResponse, err error) {
	response = CreatePutProjectResponse()
	err = client.DoAction(request, response)
	return
}

// PutProjectWithChan invokes the imm.PutProject API asynchronously
func (client *Client) PutProjectWithChan(request *PutProjectRequest) (<-chan *PutProjectResponse, <-chan error) {
	responseChan := make(chan *PutProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutProject(request)
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

// PutProjectWithCallback invokes the imm.PutProject API asynchronously
func (client *Client) PutProjectWithCallback(request *PutProjectRequest, callback func(response *PutProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutProjectResponse
		var err error
		defer close(result)
		response, err = client.PutProject(request)
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

// PutProjectRequest is the request struct for api PutProject
type PutProjectRequest struct {
	*requests.RpcRequest
	Project       string `position:"Query" name:"Project"`
	BCTaskVersion string `position:"Query" name:"BCTaskVersion"`
	ServiceRole   string `position:"Query" name:"ServiceRole"`
}

// PutProjectResponse is the response struct for api PutProject
type PutProjectResponse struct {
	*responses.BaseResponse
	Project     string `json:"Project" xml:"Project"`
	ModifyTime  string `json:"ModifyTime" xml:"ModifyTime"`
	Type        string `json:"Type" xml:"Type"`
	CU          int    `json:"CU" xml:"CU"`
	ServiceRole string `json:"ServiceRole" xml:"ServiceRole"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Endpoint    string `json:"Endpoint" xml:"Endpoint"`
	CreateTime  string `json:"CreateTime" xml:"CreateTime"`
	RegionId    string `json:"RegionId" xml:"RegionId"`
	BillingType string `json:"BillingType" xml:"BillingType"`
}

// CreatePutProjectRequest creates a request to invoke PutProject API
func CreatePutProjectRequest() (request *PutProjectRequest) {
	request = &PutProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "PutProject", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePutProjectResponse creates a response to parse from PutProject response
func CreatePutProjectResponse() (response *PutProjectResponse) {
	response = &PutProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
