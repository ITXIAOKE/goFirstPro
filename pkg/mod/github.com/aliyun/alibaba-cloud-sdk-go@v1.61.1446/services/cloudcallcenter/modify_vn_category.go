package cloudcallcenter

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

// ModifyVnCategory invokes the cloudcallcenter.ModifyVnCategory API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvncategory.html
func (client *Client) ModifyVnCategory(request *ModifyVnCategoryRequest) (response *ModifyVnCategoryResponse, err error) {
	response = CreateModifyVnCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVnCategoryWithChan invokes the cloudcallcenter.ModifyVnCategory API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvncategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnCategoryWithChan(request *ModifyVnCategoryRequest) (<-chan *ModifyVnCategoryResponse, <-chan error) {
	responseChan := make(chan *ModifyVnCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVnCategory(request)
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

// ModifyVnCategoryWithCallback invokes the cloudcallcenter.ModifyVnCategory API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvncategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnCategoryWithCallback(request *ModifyVnCategoryRequest, callback func(response *ModifyVnCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVnCategoryResponse
		var err error
		defer close(result)
		response, err = client.ModifyVnCategory(request)
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

// ModifyVnCategoryRequest is the request struct for api ModifyVnCategory
type ModifyVnCategoryRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
	CategoryId string `position:"Query" name:"CategoryId"`
}

// ModifyVnCategoryResponse is the response struct for api ModifyVnCategory
type ModifyVnCategoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVnCategoryRequest creates a request to invoke ModifyVnCategory API
func CreateModifyVnCategoryRequest() (request *ModifyVnCategoryRequest) {
	request = &ModifyVnCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyVnCategory", "", "")
	request.Method = requests.GET
	return
}

// CreateModifyVnCategoryResponse creates a response to parse from ModifyVnCategory response
func CreateModifyVnCategoryResponse() (response *ModifyVnCategoryResponse) {
	response = &ModifyVnCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
