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

// StopDataSet invokes the airec.StopDataSet API synchronously
func (client *Client) StopDataSet(request *StopDataSetRequest) (response *StopDataSetResponse, err error) {
	response = CreateStopDataSetResponse()
	err = client.DoAction(request, response)
	return
}

// StopDataSetWithChan invokes the airec.StopDataSet API asynchronously
func (client *Client) StopDataSetWithChan(request *StopDataSetRequest) (<-chan *StopDataSetResponse, <-chan error) {
	responseChan := make(chan *StopDataSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopDataSet(request)
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

// StopDataSetWithCallback invokes the airec.StopDataSet API asynchronously
func (client *Client) StopDataSetWithCallback(request *StopDataSetRequest, callback func(response *StopDataSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopDataSetResponse
		var err error
		defer close(result)
		response, err = client.StopDataSet(request)
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

// StopDataSetRequest is the request struct for api StopDataSet
type StopDataSetRequest struct {
	*requests.RoaRequest
	VersionId  string `position:"Path" name:"versionId"`
	InstanceId string `position:"Path" name:"instanceId"`
}

// StopDataSetResponse is the response struct for api StopDataSet
type StopDataSetResponse struct {
	*responses.BaseResponse
	Code      string `json:"code" xml:"code"`
	Message   string `json:"message" xml:"message"`
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateStopDataSetRequest creates a request to invoke StopDataSet API
func CreateStopDataSetRequest() (request *StopDataSetRequest) {
	request = &StopDataSetRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "StopDataSet", "/v2/openapi/instances/[instanceId]/dataSets/[versionId]/actions/stop", "airec", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopDataSetResponse creates a response to parse from StopDataSet response
func CreateStopDataSetResponse() (response *StopDataSetResponse) {
	response = &StopDataSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}