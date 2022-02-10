package dms_enterprise

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

// CreateUploadOSSFileJob invokes the dms_enterprise.CreateUploadOSSFileJob API synchronously
func (client *Client) CreateUploadOSSFileJob(request *CreateUploadOSSFileJobRequest) (response *CreateUploadOSSFileJobResponse, err error) {
	response = CreateCreateUploadOSSFileJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUploadOSSFileJobWithChan invokes the dms_enterprise.CreateUploadOSSFileJob API asynchronously
func (client *Client) CreateUploadOSSFileJobWithChan(request *CreateUploadOSSFileJobRequest) (<-chan *CreateUploadOSSFileJobResponse, <-chan error) {
	responseChan := make(chan *CreateUploadOSSFileJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUploadOSSFileJob(request)
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

// CreateUploadOSSFileJobWithCallback invokes the dms_enterprise.CreateUploadOSSFileJob API asynchronously
func (client *Client) CreateUploadOSSFileJobWithCallback(request *CreateUploadOSSFileJobRequest, callback func(response *CreateUploadOSSFileJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUploadOSSFileJobResponse
		var err error
		defer close(result)
		response, err = client.CreateUploadOSSFileJob(request)
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

// CreateUploadOSSFileJobRequest is the request struct for api CreateUploadOSSFileJob
type CreateUploadOSSFileJobRequest struct {
	*requests.RpcRequest
	UploadType   string                             `position:"Query" name:"UploadType"`
	FileSource   string                             `position:"Query" name:"FileSource"`
	Tid          requests.Integer                   `position:"Query" name:"Tid"`
	FileName     string                             `position:"Query" name:"FileName"`
	UploadTarget CreateUploadOSSFileJobUploadTarget `position:"Query" name:"UploadTarget"  type:"Struct"`
}

// CreateUploadOSSFileJobUploadTarget is a repeated param struct in CreateUploadOSSFileJobRequest
type CreateUploadOSSFileJobUploadTarget struct {
	Endpoint   string `name:"Endpoint"`
	BucketName string `name:"BucketName"`
	ObjectName string `name:"ObjectName"`
}

// CreateUploadOSSFileJobResponse is the response struct for api CreateUploadOSSFileJob
type CreateUploadOSSFileJobResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	JobKey       string `json:"JobKey" xml:"JobKey"`
}

// CreateCreateUploadOSSFileJobRequest creates a request to invoke CreateUploadOSSFileJob API
func CreateCreateUploadOSSFileJobRequest() (request *CreateUploadOSSFileJobRequest) {
	request = &CreateUploadOSSFileJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "CreateUploadOSSFileJob", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateUploadOSSFileJobResponse creates a response to parse from CreateUploadOSSFileJob response
func CreateCreateUploadOSSFileJobResponse() (response *CreateUploadOSSFileJobResponse) {
	response = &CreateUploadOSSFileJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}