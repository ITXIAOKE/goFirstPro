package scsp

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

// QueryHotlineSession invokes the scsp.QueryHotlineSession API synchronously
func (client *Client) QueryHotlineSession(request *QueryHotlineSessionRequest) (response *QueryHotlineSessionResponse, err error) {
	response = CreateQueryHotlineSessionResponse()
	err = client.DoAction(request, response)
	return
}

// QueryHotlineSessionWithChan invokes the scsp.QueryHotlineSession API asynchronously
func (client *Client) QueryHotlineSessionWithChan(request *QueryHotlineSessionRequest) (<-chan *QueryHotlineSessionResponse, <-chan error) {
	responseChan := make(chan *QueryHotlineSessionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryHotlineSession(request)
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

// QueryHotlineSessionWithCallback invokes the scsp.QueryHotlineSession API asynchronously
func (client *Client) QueryHotlineSessionWithCallback(request *QueryHotlineSessionRequest, callback func(response *QueryHotlineSessionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryHotlineSessionResponse
		var err error
		defer close(result)
		response, err = client.QueryHotlineSession(request)
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

// QueryHotlineSessionRequest is the request struct for api QueryHotlineSession
type QueryHotlineSessionRequest struct {
	*requests.RpcRequest
	GroupId        requests.Integer `position:"Query"`
	ServicerId     string           `position:"Query"`
	Params         string           `position:"Query"`
	GroupName      string           `position:"Query"`
	Acid           string           `position:"Query"`
	CallingNumber  string           `position:"Query"`
	QueryEndTime   requests.Integer `position:"Query"`
	InstanceId     string           `position:"Query"`
	CalledNumber   string           `position:"Query"`
	RequestId      string           `position:"Query"`
	PageNo         requests.Integer `position:"Query"`
	QueryStartTime requests.Integer `position:"Query"`
	ServicerName   string           `position:"Query"`
	PageSize       requests.Integer `position:"Query"`
	CallResult     string           `position:"Query"`
	CallType       requests.Integer `position:"Query"`
	MemberName     string           `position:"Query"`
	MemberId       string           `position:"Query"`
}

// QueryHotlineSessionResponse is the response struct for api QueryHotlineSession
type QueryHotlineSessionResponse struct {
	*responses.BaseResponse
	Message   string                    `json:"Message" xml:"Message"`
	RequestId string                    `json:"RequestId" xml:"RequestId"`
	Code      string                    `json:"Code" xml:"Code"`
	Success   bool                      `json:"Success" xml:"Success"`
	Data      DataInQueryHotlineSession `json:"Data" xml:"Data"`
}

// CreateQueryHotlineSessionRequest creates a request to invoke QueryHotlineSession API
func CreateQueryHotlineSessionRequest() (request *QueryHotlineSessionRequest) {
	request = &QueryHotlineSessionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "QueryHotlineSession", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryHotlineSessionResponse creates a response to parse from QueryHotlineSession response
func CreateQueryHotlineSessionResponse() (response *QueryHotlineSessionResponse) {
	response = &QueryHotlineSessionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
