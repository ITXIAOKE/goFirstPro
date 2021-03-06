package edas

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

// GetGrayApps invokes the edas.GetGrayApps API synchronously
func (client *Client) GetGrayApps(request *GetGrayAppsRequest) (response *GetGrayAppsResponse, err error) {
	response = CreateGetGrayAppsResponse()
	err = client.DoAction(request, response)
	return
}

// GetGrayAppsWithChan invokes the edas.GetGrayApps API asynchronously
func (client *Client) GetGrayAppsWithChan(request *GetGrayAppsRequest) (<-chan *GetGrayAppsResponse, <-chan error) {
	responseChan := make(chan *GetGrayAppsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGrayApps(request)
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

// GetGrayAppsWithCallback invokes the edas.GetGrayApps API asynchronously
func (client *Client) GetGrayAppsWithCallback(request *GetGrayAppsRequest, callback func(response *GetGrayAppsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGrayAppsResponse
		var err error
		defer close(result)
		response, err = client.GetGrayApps(request)
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

// GetGrayAppsRequest is the request struct for api GetGrayApps
type GetGrayAppsRequest struct {
	*requests.RoaRequest
	ClusterTypes     requests.Integer `position:"Query" name:"ClusterTypes"`
	Appname          requests.Integer `position:"Query" name:"Appname"`
	PhysicalRegionId string           `position:"Query" name:"PhysicalRegionId"`
}

// GetGrayAppsResponse is the response struct for api GetGrayApps
type GetGrayAppsResponse struct {
	*responses.BaseResponse
	RequestId string            `json:"RequestId" xml:"RequestId"`
	Code      int               `json:"Code" xml:"Code"`
	Message   string            `json:"Message" xml:"Message"`
	Data      DataInGetGrayApps `json:"Data" xml:"Data"`
}

// CreateGetGrayAppsRequest creates a request to invoke GetGrayApps API
func CreateGetGrayAppsRequest() (request *GetGrayAppsRequest) {
	request = &GetGrayAppsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetGrayApps", "/pop/v5/gray/app_list", "edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetGrayAppsResponse creates a response to parse from GetGrayApps response
func CreateGetGrayAppsResponse() (response *GetGrayAppsResponse) {
	response = &GetGrayAppsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
