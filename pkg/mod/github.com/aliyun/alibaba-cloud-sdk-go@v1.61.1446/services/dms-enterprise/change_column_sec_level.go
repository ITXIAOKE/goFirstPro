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

// ChangeColumnSecLevel invokes the dms_enterprise.ChangeColumnSecLevel API synchronously
func (client *Client) ChangeColumnSecLevel(request *ChangeColumnSecLevelRequest) (response *ChangeColumnSecLevelResponse, err error) {
	response = CreateChangeColumnSecLevelResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeColumnSecLevelWithChan invokes the dms_enterprise.ChangeColumnSecLevel API asynchronously
func (client *Client) ChangeColumnSecLevelWithChan(request *ChangeColumnSecLevelRequest) (<-chan *ChangeColumnSecLevelResponse, <-chan error) {
	responseChan := make(chan *ChangeColumnSecLevelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeColumnSecLevel(request)
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

// ChangeColumnSecLevelWithCallback invokes the dms_enterprise.ChangeColumnSecLevel API asynchronously
func (client *Client) ChangeColumnSecLevelWithCallback(request *ChangeColumnSecLevelRequest, callback func(response *ChangeColumnSecLevelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeColumnSecLevelResponse
		var err error
		defer close(result)
		response, err = client.ChangeColumnSecLevel(request)
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

// ChangeColumnSecLevelRequest is the request struct for api ChangeColumnSecLevel
type ChangeColumnSecLevelRequest struct {
	*requests.RpcRequest
	SchemaName string           `position:"Query" name:"SchemaName"`
	IsLogic    requests.Boolean `position:"Query" name:"IsLogic"`
	NewLevel   string           `position:"Query" name:"NewLevel"`
	ColumnName string           `position:"Query" name:"ColumnName"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	DbId       requests.Integer `position:"Query" name:"DbId"`
	TableName  string           `position:"Query" name:"TableName"`
}

// ChangeColumnSecLevelResponse is the response struct for api ChangeColumnSecLevel
type ChangeColumnSecLevelResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateChangeColumnSecLevelRequest creates a request to invoke ChangeColumnSecLevel API
func CreateChangeColumnSecLevelRequest() (request *ChangeColumnSecLevelRequest) {
	request = &ChangeColumnSecLevelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ChangeColumnSecLevel", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateChangeColumnSecLevelResponse creates a response to parse from ChangeColumnSecLevel response
func CreateChangeColumnSecLevelResponse() (response *ChangeColumnSecLevelResponse) {
	response = &ChangeColumnSecLevelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}