package companyreg

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

// DeletePayroll invokes the companyreg.DeletePayroll API synchronously
func (client *Client) DeletePayroll(request *DeletePayrollRequest) (response *DeletePayrollResponse, err error) {
	response = CreateDeletePayrollResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePayrollWithChan invokes the companyreg.DeletePayroll API asynchronously
func (client *Client) DeletePayrollWithChan(request *DeletePayrollRequest) (<-chan *DeletePayrollResponse, <-chan error) {
	responseChan := make(chan *DeletePayrollResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePayroll(request)
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

// DeletePayrollWithCallback invokes the companyreg.DeletePayroll API asynchronously
func (client *Client) DeletePayrollWithCallback(request *DeletePayrollRequest, callback func(response *DeletePayrollResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePayrollResponse
		var err error
		defer close(result)
		response, err = client.DeletePayroll(request)
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

// DeletePayrollRequest is the request struct for api DeletePayroll
type DeletePayrollRequest struct {
	*requests.RpcRequest
	BizId string           `position:"Query" name:"BizId"`
	Id    requests.Integer `position:"Query" name:"Id"`
}

// DeletePayrollResponse is the response struct for api DeletePayroll
type DeletePayrollResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateDeletePayrollRequest creates a request to invoke DeletePayroll API
func CreateDeletePayrollRequest() (request *DeletePayrollRequest) {
	request = &DeletePayrollRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "DeletePayroll", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeletePayrollResponse creates a response to parse from DeletePayroll response
func CreateDeletePayrollResponse() (response *DeletePayrollResponse) {
	response = &DeletePayrollResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
