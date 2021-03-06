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

// ConfigureSynchronizationJob invokes the dts.ConfigureSynchronizationJob API synchronously
func (client *Client) ConfigureSynchronizationJob(request *ConfigureSynchronizationJobRequest) (response *ConfigureSynchronizationJobResponse, err error) {
	response = CreateConfigureSynchronizationJobResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigureSynchronizationJobWithChan invokes the dts.ConfigureSynchronizationJob API asynchronously
func (client *Client) ConfigureSynchronizationJobWithChan(request *ConfigureSynchronizationJobRequest) (<-chan *ConfigureSynchronizationJobResponse, <-chan error) {
	responseChan := make(chan *ConfigureSynchronizationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigureSynchronizationJob(request)
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

// ConfigureSynchronizationJobWithCallback invokes the dts.ConfigureSynchronizationJob API asynchronously
func (client *Client) ConfigureSynchronizationJobWithCallback(request *ConfigureSynchronizationJobRequest, callback func(response *ConfigureSynchronizationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigureSynchronizationJobResponse
		var err error
		defer close(result)
		response, err = client.ConfigureSynchronizationJob(request)
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

// ConfigureSynchronizationJobRequest is the request struct for api ConfigureSynchronizationJob
type ConfigureSynchronizationJobRequest struct {
	*requests.RpcRequest
	SourceEndpointInstanceId        string           `position:"Query" name:"SourceEndpoint.InstanceId"`
	Checkpoint                      string           `position:"Query" name:"Checkpoint"`
	DestinationEndpointInstanceId   string           `position:"Query" name:"DestinationEndpoint.InstanceId"`
	SourceEndpointIP                string           `position:"Query" name:"SourceEndpoint.IP"`
	SynchronizationObjects          string           `position:"Body" name:"SynchronizationObjects"`
	DestinationEndpointPassword     string           `position:"Query" name:"DestinationEndpoint.Password"`
	DataInitialization              requests.Boolean `position:"Query" name:"DataInitialization"`
	StructureInitialization         requests.Boolean `position:"Query" name:"StructureInitialization"`
	PartitionKeyModifyTimeMinute    requests.Boolean `position:"Query" name:"PartitionKey.ModifyTime_Minute"`
	PartitionKeyModifyTimeDay       requests.Boolean `position:"Query" name:"PartitionKey.ModifyTime_Day"`
	SourceEndpointInstanceType      string           `position:"Query" name:"SourceEndpoint.InstanceType"`
	SynchronizationJobId            string           `position:"Query" name:"SynchronizationJobId"`
	SynchronizationJobName          string           `position:"Query" name:"SynchronizationJobName"`
	AccountId                       string           `position:"Query" name:"AccountId"`
	SourceEndpointUserName          string           `position:"Query" name:"SourceEndpoint.UserName"`
	SourceEndpointDatabaseName      string           `position:"Query" name:"SourceEndpoint.DatabaseName"`
	PartitionKeyModifyTimeMonth     requests.Boolean `position:"Query" name:"PartitionKey.ModifyTime_Month"`
	SourceEndpointPort              string           `position:"Query" name:"SourceEndpoint.Port"`
	SourceEndpointOwnerID           string           `position:"Query" name:"SourceEndpoint.OwnerID"`
	DestinationEndpointUserName     string           `position:"Query" name:"DestinationEndpoint.UserName"`
	DestinationEndpointPort         string           `position:"Query" name:"DestinationEndpoint.Port"`
	PartitionKeyModifyTimeYear      requests.Boolean `position:"Query" name:"PartitionKey.ModifyTime_Year"`
	SourceEndpointRole              string           `position:"Query" name:"SourceEndpoint.Role"`
	OwnerId                         string           `position:"Query" name:"OwnerId"`
	PartitionKeyModifyTimeHour      requests.Boolean `position:"Query" name:"PartitionKey.ModifyTime_Hour"`
	DestinationEndpointDataBaseName string           `position:"Query" name:"DestinationEndpoint.DataBaseName"`
	SourceEndpointPassword          string           `position:"Query" name:"SourceEndpoint.Password"`
	MigrationReserved               string           `position:"Query" name:"MigrationReserved"`
	DestinationEndpointIP           string           `position:"Query" name:"DestinationEndpoint.IP"`
	DestinationEndpointInstanceType string           `position:"Query" name:"DestinationEndpoint.InstanceType"`
	SynchronizationDirection        string           `position:"Query" name:"SynchronizationDirection"`
}

// ConfigureSynchronizationJobResponse is the response struct for api ConfigureSynchronizationJob
type ConfigureSynchronizationJobResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	Success    string `json:"Success" xml:"Success"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateConfigureSynchronizationJobRequest creates a request to invoke ConfigureSynchronizationJob API
func CreateConfigureSynchronizationJobRequest() (request *ConfigureSynchronizationJobRequest) {
	request = &ConfigureSynchronizationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "ConfigureSynchronizationJob", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigureSynchronizationJobResponse creates a response to parse from ConfigureSynchronizationJob response
func CreateConfigureSynchronizationJobResponse() (response *ConfigureSynchronizationJobResponse) {
	response = &ConfigureSynchronizationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
