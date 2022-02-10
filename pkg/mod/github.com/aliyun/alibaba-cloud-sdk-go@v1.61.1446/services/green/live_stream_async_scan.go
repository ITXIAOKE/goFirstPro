package green

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

// LiveStreamAsyncScan invokes the green.LiveStreamAsyncScan API synchronously
func (client *Client) LiveStreamAsyncScan(request *LiveStreamAsyncScanRequest) (response *LiveStreamAsyncScanResponse, err error) {
	response = CreateLiveStreamAsyncScanResponse()
	err = client.DoAction(request, response)
	return
}

// LiveStreamAsyncScanWithChan invokes the green.LiveStreamAsyncScan API asynchronously
func (client *Client) LiveStreamAsyncScanWithChan(request *LiveStreamAsyncScanRequest) (<-chan *LiveStreamAsyncScanResponse, <-chan error) {
	responseChan := make(chan *LiveStreamAsyncScanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LiveStreamAsyncScan(request)
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

// LiveStreamAsyncScanWithCallback invokes the green.LiveStreamAsyncScan API asynchronously
func (client *Client) LiveStreamAsyncScanWithCallback(request *LiveStreamAsyncScanRequest, callback func(response *LiveStreamAsyncScanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LiveStreamAsyncScanResponse
		var err error
		defer close(result)
		response, err = client.LiveStreamAsyncScan(request)
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

// LiveStreamAsyncScanRequest is the request struct for api LiveStreamAsyncScan
type LiveStreamAsyncScanRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// LiveStreamAsyncScanResponse is the response struct for api LiveStreamAsyncScan
type LiveStreamAsyncScanResponse struct {
	*responses.BaseResponse
}

// CreateLiveStreamAsyncScanRequest creates a request to invoke LiveStreamAsyncScan API
func CreateLiveStreamAsyncScanRequest() (request *LiveStreamAsyncScanRequest) {
	request = &LiveStreamAsyncScanRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "LiveStreamAsyncScan", "/green/livestream/asyncscan", "", "")
	request.Method = requests.POST
	return
}

// CreateLiveStreamAsyncScanResponse creates a response to parse from LiveStreamAsyncScan response
func CreateLiveStreamAsyncScanResponse() (response *LiveStreamAsyncScanResponse) {
	response = &LiveStreamAsyncScanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
