package ccc

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

// DeleteSurvey invokes the ccc.DeleteSurvey API synchronously
func (client *Client) DeleteSurvey(request *DeleteSurveyRequest) (response *DeleteSurveyResponse, err error) {
	response = CreateDeleteSurveyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSurveyWithChan invokes the ccc.DeleteSurvey API asynchronously
func (client *Client) DeleteSurveyWithChan(request *DeleteSurveyRequest) (<-chan *DeleteSurveyResponse, <-chan error) {
	responseChan := make(chan *DeleteSurveyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSurvey(request)
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

// DeleteSurveyWithCallback invokes the ccc.DeleteSurvey API asynchronously
func (client *Client) DeleteSurveyWithCallback(request *DeleteSurveyRequest, callback func(response *DeleteSurveyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSurveyResponse
		var err error
		defer close(result)
		response, err = client.DeleteSurvey(request)
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

// DeleteSurveyRequest is the request struct for api DeleteSurvey
type DeleteSurveyRequest struct {
	*requests.RpcRequest
	SurveyId   string `position:"Query" name:"SurveyId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	ScenarioId string `position:"Query" name:"ScenarioId"`
}

// DeleteSurveyResponse is the response struct for api DeleteSurvey
type DeleteSurveyResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateDeleteSurveyRequest creates a request to invoke DeleteSurvey API
func CreateDeleteSurveyRequest() (request *DeleteSurveyRequest) {
	request = &DeleteSurveyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "DeleteSurvey", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteSurveyResponse creates a response to parse from DeleteSurvey response
func CreateDeleteSurveyResponse() (response *DeleteSurveyResponse) {
	response = &DeleteSurveyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
