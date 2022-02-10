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

// CreateSynchronizationJob invokes the dts.CreateSynchronizationJob API synchronously
func (client *Client) CreateSynchronizationJob(request *CreateSynchronizationJobRequest) (response *CreateSynchronizationJobResponse, err error) {
	response = CreateCreateSynchronizationJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSynchronizationJobWithChan invokes the dts.CreateSynchronizationJob API asynchronously
func (client *Client) CreateSynchronizationJobWithChan(request *CreateSynchronizationJobRequest) (<-chan *CreateSynchronizationJobResponse, <-chan error) {
	responseChan := make(chan *CreateSynchronizationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSynchronizationJob(request)
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

// CreateSynchronizationJobWithCallback invokes the dts.CreateSynchronizationJob API asynchronously
func (client *Client) CreateSynchronizationJobWithCallback(request *CreateSynchronizationJobRequest, callback func(response *CreateSynchronizationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSynchronizationJobResponse
		var err error
		defer close(result)
		response, err = client.CreateSynchronizationJob(request)
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

// CreateSynchronizationJobRequest is the request struct for api CreateSynchronizationJob
type CreateSynchronizationJobRequest struct {
	*requests.RpcRequest
	ClientToken                     string           `position:"Query" name:"ClientToken"`
	NetworkType                     string           `position:"Query" name:"networkType"`
	SourceEndpointInstanceType      string           `position:"Query" name:"SourceEndpoint.InstanceType"`
	AccountId                       string           `position:"Query" name:"AccountId"`
	SynchronizationJobClass         string           `position:"Query" name:"SynchronizationJobClass"`
	Period                          string           `position:"Query" name:"Period"`
	DestRegion                      string           `position:"Query" name:"DestRegion"`
	Topology                        string           `position:"Query" name:"Topology"`
	OwnerId                         string           `position:"Query" name:"OwnerId"`
	UsedTime                        requests.Integer `position:"Query" name:"UsedTime"`
	DBInstanceCount                 requests.Integer `position:"Query" name:"DBInstanceCount"`
	SourceRegion                    string           `position:"Query" name:"SourceRegion"`
	PayType                         string           `position:"Query" name:"PayType"`
	DestinationEndpointInstanceType string           `position:"Query" name:"DestinationEndpoint.InstanceType"`
}

// CreateSynchronizationJobResponse is the response struct for api CreateSynchronizationJob
type CreateSynchronizationJobResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	ErrCode              string `json:"ErrCode" xml:"ErrCode"`
	Success              string `json:"Success" xml:"Success"`
	SynchronizationJobId string `json:"SynchronizationJobId" xml:"SynchronizationJobId"`
	ErrMessage           string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateCreateSynchronizationJobRequest creates a request to invoke CreateSynchronizationJob API
func CreateCreateSynchronizationJobRequest() (request *CreateSynchronizationJobRequest) {
	request = &CreateSynchronizationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "CreateSynchronizationJob", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSynchronizationJobResponse creates a response to parse from CreateSynchronizationJob response
func CreateCreateSynchronizationJobResponse() (response *CreateSynchronizationJobResponse) {
	response = &CreateSynchronizationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}