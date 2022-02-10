package codeup

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

// EnableRepositoryDeployKey invokes the codeup.EnableRepositoryDeployKey API synchronously
func (client *Client) EnableRepositoryDeployKey(request *EnableRepositoryDeployKeyRequest) (response *EnableRepositoryDeployKeyResponse, err error) {
	response = CreateEnableRepositoryDeployKeyResponse()
	err = client.DoAction(request, response)
	return
}

// EnableRepositoryDeployKeyWithChan invokes the codeup.EnableRepositoryDeployKey API asynchronously
func (client *Client) EnableRepositoryDeployKeyWithChan(request *EnableRepositoryDeployKeyRequest) (<-chan *EnableRepositoryDeployKeyResponse, <-chan error) {
	responseChan := make(chan *EnableRepositoryDeployKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableRepositoryDeployKey(request)
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

// EnableRepositoryDeployKeyWithCallback invokes the codeup.EnableRepositoryDeployKey API asynchronously
func (client *Client) EnableRepositoryDeployKeyWithCallback(request *EnableRepositoryDeployKeyRequest, callback func(response *EnableRepositoryDeployKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableRepositoryDeployKeyResponse
		var err error
		defer close(result)
		response, err = client.EnableRepositoryDeployKey(request)
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

// EnableRepositoryDeployKeyRequest is the request struct for api EnableRepositoryDeployKey
type EnableRepositoryDeployKeyRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	SubUserId      string           `position:"Query" name:"SubUserId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	KeyId          requests.Integer `position:"Path" name:"KeyId"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
}

// EnableRepositoryDeployKeyResponse is the response struct for api EnableRepositoryDeployKey
type EnableRepositoryDeployKeyResponse struct {
	*responses.BaseResponse
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateEnableRepositoryDeployKeyRequest creates a request to invoke EnableRepositoryDeployKey API
func CreateEnableRepositoryDeployKeyRequest() (request *EnableRepositoryDeployKeyRequest) {
	request = &EnableRepositoryDeployKeyRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "EnableRepositoryDeployKey", "/api/v3/projects/[ProjectId]/keys/[KeyId]/enable", "", "")
	request.Method = requests.POST
	return
}

// CreateEnableRepositoryDeployKeyResponse creates a response to parse from EnableRepositoryDeployKey response
func CreateEnableRepositoryDeployKeyResponse() (response *EnableRepositoryDeployKeyResponse) {
	response = &EnableRepositoryDeployKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
