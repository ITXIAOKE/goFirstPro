package dts

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

// DescribeSynchronizationObjectModifyStatus invokes the dts.DescribeSynchronizationObjectModifyStatus API synchronously
func (client *Client) DescribeSynchronizationObjectModifyStatus(request *DescribeSynchronizationObjectModifyStatusRequest) (response *DescribeSynchronizationObjectModifyStatusResponse, err error) {
	response = CreateDescribeSynchronizationObjectModifyStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSynchronizationObjectModifyStatusWithChan invokes the dts.DescribeSynchronizationObjectModifyStatus API asynchronously
func (client *Client) DescribeSynchronizationObjectModifyStatusWithChan(request *DescribeSynchronizationObjectModifyStatusRequest) (<-chan *DescribeSynchronizationObjectModifyStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeSynchronizationObjectModifyStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSynchronizationObjectModifyStatus(request)
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

// DescribeSynchronizationObjectModifyStatusWithCallback invokes the dts.DescribeSynchronizationObjectModifyStatus API asynchronously
func (client *Client) DescribeSynchronizationObjectModifyStatusWithCallback(request *DescribeSynchronizationObjectModifyStatusRequest, callback func(response *DescribeSynchronizationObjectModifyStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSynchronizationObjectModifyStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeSynchronizationObjectModifyStatus(request)
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

// DescribeSynchronizationObjectModifyStatusRequest is the request struct for api DescribeSynchronizationObjectModifyStatus
type DescribeSynchronizationObjectModifyStatusRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	OwnerId     string `position:"Query" name:"OwnerId"`
	AccountId   string `position:"Query" name:"AccountId"`
	TaskId      string `position:"Query" name:"TaskId"`
}

// DescribeSynchronizationObjectModifyStatusResponse is the response struct for api DescribeSynchronizationObjectModifyStatus
type DescribeSynchronizationObjectModifyStatusResponse struct {
	*responses.BaseResponse
	Status                        string                                                    `json:"Status" xml:"Status"`
	ErrorMessage                  string                                                    `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId                     string                                                    `json:"RequestId" xml:"RequestId"`
	ErrCode                       string                                                    `json:"ErrCode" xml:"ErrCode"`
	Success                       string                                                    `json:"Success" xml:"Success"`
	ErrMessage                    string                                                    `json:"ErrMessage" xml:"ErrMessage"`
	DataInitializationStatus      DataInitializationStatus                                  `json:"DataInitializationStatus" xml:"DataInitializationStatus"`
	DataSynchronizationStatus     DataSynchronizationStatus                                 `json:"DataSynchronizationStatus" xml:"DataSynchronizationStatus"`
	PrecheckStatus                PrecheckStatusInDescribeSynchronizationObjectModifyStatus `json:"PrecheckStatus" xml:"PrecheckStatus"`
	StructureInitializationStatus StructureInitializationStatus                             `json:"StructureInitializationStatus" xml:"StructureInitializationStatus"`
}

// CreateDescribeSynchronizationObjectModifyStatusRequest creates a request to invoke DescribeSynchronizationObjectModifyStatus API
func CreateDescribeSynchronizationObjectModifyStatusRequest() (request *DescribeSynchronizationObjectModifyStatusRequest) {
	request = &DescribeSynchronizationObjectModifyStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeSynchronizationObjectModifyStatus", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSynchronizationObjectModifyStatusResponse creates a response to parse from DescribeSynchronizationObjectModifyStatus response
func CreateDescribeSynchronizationObjectModifyStatusResponse() (response *DescribeSynchronizationObjectModifyStatusResponse) {
	response = &DescribeSynchronizationObjectModifyStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
