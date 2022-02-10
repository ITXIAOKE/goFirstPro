package ga

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

// GetAcl invokes the ga.GetAcl API synchronously
func (client *Client) GetAcl(request *GetAclRequest) (response *GetAclResponse, err error) {
	response = CreateGetAclResponse()
	err = client.DoAction(request, response)
	return
}

// GetAclWithChan invokes the ga.GetAcl API asynchronously
func (client *Client) GetAclWithChan(request *GetAclRequest) (<-chan *GetAclResponse, <-chan error) {
	responseChan := make(chan *GetAclResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAcl(request)
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

// GetAclWithCallback invokes the ga.GetAcl API asynchronously
func (client *Client) GetAclWithCallback(request *GetAclRequest, callback func(response *GetAclResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAclResponse
		var err error
		defer close(result)
		response, err = client.GetAcl(request)
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

// GetAclRequest is the request struct for api GetAcl
type GetAclRequest struct {
	*requests.RpcRequest
	AclId string `position:"Query" name:"AclId"`
}

// GetAclResponse is the response struct for api GetAcl
type GetAclResponse struct {
	*responses.BaseResponse
	AclStatus        string                 `json:"AclStatus" xml:"AclStatus"`
	RequestId        string                 `json:"RequestId" xml:"RequestId"`
	AddressIPVersion string                 `json:"AddressIPVersion" xml:"AddressIPVersion"`
	AclId            string                 `json:"AclId" xml:"AclId"`
	AclName          string                 `json:"AclName" xml:"AclName"`
	AclEntries       []AclEntriesItem       `json:"AclEntries" xml:"AclEntries"`
	RelatedListeners []RelatedListenersItem `json:"RelatedListeners" xml:"RelatedListeners"`
}

// CreateGetAclRequest creates a request to invoke GetAcl API
func CreateGetAclRequest() (request *GetAclRequest) {
	request = &GetAclRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "GetAcl", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAclResponse creates a response to parse from GetAcl response
func CreateGetAclResponse() (response *GetAclResponse) {
	response = &GetAclResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
