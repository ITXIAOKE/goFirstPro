package airec

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

// DecribeRankingModel invokes the airec.DecribeRankingModel API synchronously
func (client *Client) DecribeRankingModel(request *DecribeRankingModelRequest) (response *DecribeRankingModelResponse, err error) {
	response = CreateDecribeRankingModelResponse()
	err = client.DoAction(request, response)
	return
}

// DecribeRankingModelWithChan invokes the airec.DecribeRankingModel API asynchronously
func (client *Client) DecribeRankingModelWithChan(request *DecribeRankingModelRequest) (<-chan *DecribeRankingModelResponse, <-chan error) {
	responseChan := make(chan *DecribeRankingModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DecribeRankingModel(request)
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

// DecribeRankingModelWithCallback invokes the airec.DecribeRankingModel API asynchronously
func (client *Client) DecribeRankingModelWithCallback(request *DecribeRankingModelRequest, callback func(response *DecribeRankingModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DecribeRankingModelResponse
		var err error
		defer close(result)
		response, err = client.DecribeRankingModel(request)
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

// DecribeRankingModelRequest is the request struct for api DecribeRankingModel
type DecribeRankingModelRequest struct {
	*requests.RoaRequest
	InstanceId     string `position:"Path" name:"instanceId"`
	RankingModelId string `position:"Path" name:"rankingModelId"`
}

// DecribeRankingModelResponse is the response struct for api DecribeRankingModel
type DecribeRankingModelResponse struct {
	*responses.BaseResponse
	Code      string                      `json:"code" xml:"code"`
	Message   string                      `json:"message" xml:"message"`
	RequestId string                      `json:"requestId" xml:"requestId"`
	Result    ResultInDecribeRankingModel `json:"result" xml:"result"`
}

// CreateDecribeRankingModelRequest creates a request to invoke DecribeRankingModel API
func CreateDecribeRankingModelRequest() (request *DecribeRankingModelRequest) {
	request = &DecribeRankingModelRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "DecribeRankingModel", "/v2/openapi/instances/[instanceId]/ranking-models/[rankingModelId]", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDecribeRankingModelResponse creates a response to parse from DecribeRankingModel response
func CreateDecribeRankingModelResponse() (response *DecribeRankingModelResponse) {
	response = &DecribeRankingModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
