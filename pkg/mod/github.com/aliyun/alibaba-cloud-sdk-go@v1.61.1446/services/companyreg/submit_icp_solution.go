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

// SubmitIcpSolution invokes the companyreg.SubmitIcpSolution API synchronously
func (client *Client) SubmitIcpSolution(request *SubmitIcpSolutionRequest) (response *SubmitIcpSolutionResponse, err error) {
	response = CreateSubmitIcpSolutionResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitIcpSolutionWithChan invokes the companyreg.SubmitIcpSolution API asynchronously
func (client *Client) SubmitIcpSolutionWithChan(request *SubmitIcpSolutionRequest) (<-chan *SubmitIcpSolutionResponse, <-chan error) {
	responseChan := make(chan *SubmitIcpSolutionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitIcpSolution(request)
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

// SubmitIcpSolutionWithCallback invokes the companyreg.SubmitIcpSolution API asynchronously
func (client *Client) SubmitIcpSolutionWithCallback(request *SubmitIcpSolutionRequest, callback func(response *SubmitIcpSolutionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitIcpSolutionResponse
		var err error
		defer close(result)
		response, err = client.SubmitIcpSolution(request)
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

// SubmitIcpSolutionRequest is the request struct for api SubmitIcpSolution
type SubmitIcpSolutionRequest struct {
	*requests.RpcRequest
	Area           string           `position:"Body" name:"Area"`
	ActionType     string           `position:"Body" name:"ActionType"`
	IntentionBizId string           `position:"Body" name:"IntentionBizId"`
	Source         string           `position:"Body" name:"Source"`
	UserId         string           `position:"Body" name:"UserId"`
	IcpType        requests.Integer `position:"Body" name:"IcpType"`
	CompanyAddress string           `position:"Body" name:"CompanyAddress"`
	CompanyName    string           `position:"Body" name:"CompanyName"`
	BizId          string           `position:"Body" name:"BizId"`
}

// SubmitIcpSolutionResponse is the response struct for api SubmitIcpSolution
type SubmitIcpSolutionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	BizId     string `json:"BizId" xml:"BizId"`
}

// CreateSubmitIcpSolutionRequest creates a request to invoke SubmitIcpSolution API
func CreateSubmitIcpSolutionRequest() (request *SubmitIcpSolutionRequest) {
	request = &SubmitIcpSolutionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "SubmitIcpSolution", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitIcpSolutionResponse creates a response to parse from SubmitIcpSolution response
func CreateSubmitIcpSolutionResponse() (response *SubmitIcpSolutionResponse) {
	response = &SubmitIcpSolutionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
