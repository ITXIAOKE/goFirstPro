package vs

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

// BatchStartStreams invokes the vs.BatchStartStreams API synchronously
func (client *Client) BatchStartStreams(request *BatchStartStreamsRequest) (response *BatchStartStreamsResponse, err error) {
	response = CreateBatchStartStreamsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchStartStreamsWithChan invokes the vs.BatchStartStreams API asynchronously
func (client *Client) BatchStartStreamsWithChan(request *BatchStartStreamsRequest) (<-chan *BatchStartStreamsResponse, <-chan error) {
	responseChan := make(chan *BatchStartStreamsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchStartStreams(request)
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

// BatchStartStreamsWithCallback invokes the vs.BatchStartStreams API asynchronously
func (client *Client) BatchStartStreamsWithCallback(request *BatchStartStreamsRequest, callback func(response *BatchStartStreamsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchStartStreamsResponse
		var err error
		defer close(result)
		response, err = client.BatchStartStreams(request)
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

// BatchStartStreamsRequest is the request struct for api BatchStartStreams
type BatchStartStreamsRequest struct {
	*requests.RpcRequest
	Id      string           `position:"Query" name:"Id"`
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// BatchStartStreamsResponse is the response struct for api BatchStartStreams
type BatchStartStreamsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Results   []Result `json:"Results" xml:"Results"`
}

// CreateBatchStartStreamsRequest creates a request to invoke BatchStartStreams API
func CreateBatchStartStreamsRequest() (request *BatchStartStreamsRequest) {
	request = &BatchStartStreamsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "BatchStartStreams", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchStartStreamsResponse creates a response to parse from BatchStartStreams response
func CreateBatchStartStreamsResponse() (response *BatchStartStreamsResponse) {
	response = &BatchStartStreamsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}