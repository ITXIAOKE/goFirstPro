package teambition_aliyun

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

// ListScenarioFieldConfig invokes the teambition_aliyun.ListScenarioFieldConfig API synchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/listscenariofieldconfig.html
func (client *Client) ListScenarioFieldConfig(request *ListScenarioFieldConfigRequest) (response *ListScenarioFieldConfigResponse, err error) {
	response = CreateListScenarioFieldConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ListScenarioFieldConfigWithChan invokes the teambition_aliyun.ListScenarioFieldConfig API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/listscenariofieldconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListScenarioFieldConfigWithChan(request *ListScenarioFieldConfigRequest) (<-chan *ListScenarioFieldConfigResponse, <-chan error) {
	responseChan := make(chan *ListScenarioFieldConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListScenarioFieldConfig(request)
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

// ListScenarioFieldConfigWithCallback invokes the teambition_aliyun.ListScenarioFieldConfig API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/listscenariofieldconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListScenarioFieldConfigWithCallback(request *ListScenarioFieldConfigRequest, callback func(response *ListScenarioFieldConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListScenarioFieldConfigResponse
		var err error
		defer close(result)
		response, err = client.ListScenarioFieldConfig(request)
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

// ListScenarioFieldConfigRequest is the request struct for api ListScenarioFieldConfig
type ListScenarioFieldConfigRequest struct {
	*requests.RpcRequest
	ProjectId string `position:"Body" name:"ProjectId"`
	OrgId     string `position:"Body" name:"OrgId"`
}

// ListScenarioFieldConfigResponse is the response struct for api ListScenarioFieldConfig
type ListScenarioFieldConfigResponse struct {
	*responses.BaseResponse
	Successful bool                  `json:"Successful" xml:"Successful"`
	ErrorCode  string                `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string                `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string                `json:"RequestId" xml:"RequestId"`
	Object     []ScenarioFieldConfig `json:"Object" xml:"Object"`
}

// CreateListScenarioFieldConfigRequest creates a request to invoke ListScenarioFieldConfig API
func CreateListScenarioFieldConfigRequest() (request *ListScenarioFieldConfigRequest) {
	request = &ListScenarioFieldConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("teambition-aliyun", "2020-02-26", "ListScenarioFieldConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateListScenarioFieldConfigResponse creates a response to parse from ListScenarioFieldConfig response
func CreateListScenarioFieldConfigResponse() (response *ListScenarioFieldConfigResponse) {
	response = &ListScenarioFieldConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
