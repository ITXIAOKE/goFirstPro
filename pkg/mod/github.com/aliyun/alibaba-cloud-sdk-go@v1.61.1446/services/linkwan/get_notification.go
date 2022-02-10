package linkwan

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

// GetNotification invokes the linkwan.GetNotification API synchronously
func (client *Client) GetNotification(request *GetNotificationRequest) (response *GetNotificationResponse, err error) {
	response = CreateGetNotificationResponse()
	err = client.DoAction(request, response)
	return
}

// GetNotificationWithChan invokes the linkwan.GetNotification API asynchronously
func (client *Client) GetNotificationWithChan(request *GetNotificationRequest) (<-chan *GetNotificationResponse, <-chan error) {
	responseChan := make(chan *GetNotificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNotification(request)
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

// GetNotificationWithCallback invokes the linkwan.GetNotification API asynchronously
func (client *Client) GetNotificationWithCallback(request *GetNotificationRequest, callback func(response *GetNotificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNotificationResponse
		var err error
		defer close(result)
		response, err = client.GetNotification(request)
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

// GetNotificationRequest is the request struct for api GetNotification
type GetNotificationRequest struct {
	*requests.RpcRequest
	ApiProduct     string `position:"Body" name:"ApiProduct"`
	ApiRevision    string `position:"Body" name:"ApiRevision"`
	NotificationId string `position:"Query" name:"NotificationId"`
}

// GetNotificationResponse is the response struct for api GetNotification
type GetNotificationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetNotificationRequest creates a request to invoke GetNotification API
func CreateGetNotificationRequest() (request *GetNotificationRequest) {
	request = &GetNotificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "GetNotification", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetNotificationResponse creates a response to parse from GetNotification response
func CreateGetNotificationResponse() (response *GetNotificationResponse) {
	response = &GetNotificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
