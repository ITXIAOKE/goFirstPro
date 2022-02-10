package imm

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

// DetectImageQRCodes invokes the imm.DetectImageQRCodes API synchronously
func (client *Client) DetectImageQRCodes(request *DetectImageQRCodesRequest) (response *DetectImageQRCodesResponse, err error) {
	response = CreateDetectImageQRCodesResponse()
	err = client.DoAction(request, response)
	return
}

// DetectImageQRCodesWithChan invokes the imm.DetectImageQRCodes API asynchronously
func (client *Client) DetectImageQRCodesWithChan(request *DetectImageQRCodesRequest) (<-chan *DetectImageQRCodesResponse, <-chan error) {
	responseChan := make(chan *DetectImageQRCodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectImageQRCodes(request)
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

// DetectImageQRCodesWithCallback invokes the imm.DetectImageQRCodes API asynchronously
func (client *Client) DetectImageQRCodesWithCallback(request *DetectImageQRCodesRequest, callback func(response *DetectImageQRCodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectImageQRCodesResponse
		var err error
		defer close(result)
		response, err = client.DetectImageQRCodes(request)
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

// DetectImageQRCodesRequest is the request struct for api DetectImageQRCodes
type DetectImageQRCodesRequest struct {
	*requests.RpcRequest
	Project  string `position:"Query" name:"Project"`
	ImageUri string `position:"Query" name:"ImageUri"`
}

// DetectImageQRCodesResponse is the response struct for api DetectImageQRCodes
type DetectImageQRCodesResponse struct {
	*responses.BaseResponse
	RequestId string        `json:"RequestId" xml:"RequestId"`
	ImageUri  string        `json:"ImageUri" xml:"ImageUri"`
	QRCodes   []QRCodesItem `json:"QRCodes" xml:"QRCodes"`
}

// CreateDetectImageQRCodesRequest creates a request to invoke DetectImageQRCodes API
func CreateDetectImageQRCodesRequest() (request *DetectImageQRCodesRequest) {
	request = &DetectImageQRCodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DetectImageQRCodes", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetectImageQRCodesResponse creates a response to parse from DetectImageQRCodes response
func CreateDetectImageQRCodesResponse() (response *DetectImageQRCodesResponse) {
	response = &DetectImageQRCodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
