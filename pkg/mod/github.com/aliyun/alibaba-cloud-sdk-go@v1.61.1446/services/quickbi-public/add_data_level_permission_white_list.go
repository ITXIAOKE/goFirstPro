package quickbi_public

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

// AddDataLevelPermissionWhiteList invokes the quickbi_public.AddDataLevelPermissionWhiteList API synchronously
func (client *Client) AddDataLevelPermissionWhiteList(request *AddDataLevelPermissionWhiteListRequest) (response *AddDataLevelPermissionWhiteListResponse, err error) {
	response = CreateAddDataLevelPermissionWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// AddDataLevelPermissionWhiteListWithChan invokes the quickbi_public.AddDataLevelPermissionWhiteList API asynchronously
func (client *Client) AddDataLevelPermissionWhiteListWithChan(request *AddDataLevelPermissionWhiteListRequest) (<-chan *AddDataLevelPermissionWhiteListResponse, <-chan error) {
	responseChan := make(chan *AddDataLevelPermissionWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDataLevelPermissionWhiteList(request)
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

// AddDataLevelPermissionWhiteListWithCallback invokes the quickbi_public.AddDataLevelPermissionWhiteList API asynchronously
func (client *Client) AddDataLevelPermissionWhiteListWithCallback(request *AddDataLevelPermissionWhiteListRequest, callback func(response *AddDataLevelPermissionWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDataLevelPermissionWhiteListResponse
		var err error
		defer close(result)
		response, err = client.AddDataLevelPermissionWhiteList(request)
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

// AddDataLevelPermissionWhiteListRequest is the request struct for api AddDataLevelPermissionWhiteList
type AddDataLevelPermissionWhiteListRequest struct {
	*requests.RpcRequest
	TargetType  string `position:"Query" name:"TargetType"`
	TargetIds   string `position:"Query" name:"TargetIds"`
	ClientToken string `position:"Query" name:"ClientToken"`
	RuleType    string `position:"Query" name:"RuleType"`
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	OperateType string `position:"Query" name:"OperateType"`
	CubeId      string `position:"Query" name:"CubeId"`
}

// AddDataLevelPermissionWhiteListResponse is the response struct for api AddDataLevelPermissionWhiteList
type AddDataLevelPermissionWhiteListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateAddDataLevelPermissionWhiteListRequest creates a request to invoke AddDataLevelPermissionWhiteList API
func CreateAddDataLevelPermissionWhiteListRequest() (request *AddDataLevelPermissionWhiteListRequest) {
	request = &AddDataLevelPermissionWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2020-11-11", "AddDataLevelPermissionWhiteList", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddDataLevelPermissionWhiteListResponse creates a response to parse from AddDataLevelPermissionWhiteList response
func CreateAddDataLevelPermissionWhiteListResponse() (response *AddDataLevelPermissionWhiteListResponse) {
	response = &AddDataLevelPermissionWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
