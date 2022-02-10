package clickhouse

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

// CreateBackupPolicy invokes the clickhouse.CreateBackupPolicy API synchronously
func (client *Client) CreateBackupPolicy(request *CreateBackupPolicyRequest) (response *CreateBackupPolicyResponse, err error) {
	response = CreateCreateBackupPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBackupPolicyWithChan invokes the clickhouse.CreateBackupPolicy API asynchronously
func (client *Client) CreateBackupPolicyWithChan(request *CreateBackupPolicyRequest) (<-chan *CreateBackupPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateBackupPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBackupPolicy(request)
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

// CreateBackupPolicyWithCallback invokes the clickhouse.CreateBackupPolicy API asynchronously
func (client *Client) CreateBackupPolicyWithCallback(request *CreateBackupPolicyRequest, callback func(response *CreateBackupPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBackupPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateBackupPolicy(request)
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

// CreateBackupPolicyRequest is the request struct for api CreateBackupPolicy
type CreateBackupPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PreferredBackupPeriod string           `position:"Query" name:"PreferredBackupPeriod"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId           string           `position:"Query" name:"DBClusterId"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	PreferredBackupTime   string           `position:"Query" name:"PreferredBackupTime"`
	BackupRetentionPeriod string           `position:"Query" name:"BackupRetentionPeriod"`
}

// CreateBackupPolicyResponse is the response struct for api CreateBackupPolicy
type CreateBackupPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateBackupPolicyRequest creates a request to invoke CreateBackupPolicy API
func CreateCreateBackupPolicyRequest() (request *CreateBackupPolicyRequest) {
	request = &CreateBackupPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "CreateBackupPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateBackupPolicyResponse creates a response to parse from CreateBackupPolicy response
func CreateCreateBackupPolicyResponse() (response *CreateBackupPolicyResponse) {
	response = &CreateBackupPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
