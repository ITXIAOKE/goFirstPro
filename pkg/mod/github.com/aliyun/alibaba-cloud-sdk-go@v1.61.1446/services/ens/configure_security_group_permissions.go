package ens

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

// ConfigureSecurityGroupPermissions invokes the ens.ConfigureSecurityGroupPermissions API synchronously
func (client *Client) ConfigureSecurityGroupPermissions(request *ConfigureSecurityGroupPermissionsRequest) (response *ConfigureSecurityGroupPermissionsResponse, err error) {
	response = CreateConfigureSecurityGroupPermissionsResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigureSecurityGroupPermissionsWithChan invokes the ens.ConfigureSecurityGroupPermissions API asynchronously
func (client *Client) ConfigureSecurityGroupPermissionsWithChan(request *ConfigureSecurityGroupPermissionsRequest) (<-chan *ConfigureSecurityGroupPermissionsResponse, <-chan error) {
	responseChan := make(chan *ConfigureSecurityGroupPermissionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigureSecurityGroupPermissions(request)
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

// ConfigureSecurityGroupPermissionsWithCallback invokes the ens.ConfigureSecurityGroupPermissions API asynchronously
func (client *Client) ConfigureSecurityGroupPermissionsWithCallback(request *ConfigureSecurityGroupPermissionsRequest, callback func(response *ConfigureSecurityGroupPermissionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigureSecurityGroupPermissionsResponse
		var err error
		defer close(result)
		response, err = client.ConfigureSecurityGroupPermissions(request)
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

// ConfigureSecurityGroupPermissionsRequest is the request struct for api ConfigureSecurityGroupPermissions
type ConfigureSecurityGroupPermissionsRequest struct {
	*requests.RpcRequest
	SecurityGroupId      string                                                   `position:"Query" name:"SecurityGroupId"`
	RevokePermissions    *[]ConfigureSecurityGroupPermissionsRevokePermissions    `position:"Query" name:"RevokePermissions"  type:"Json"`
	AuthorizePermissions *[]ConfigureSecurityGroupPermissionsAuthorizePermissions `position:"Query" name:"AuthorizePermissions"  type:"Json"`
}

// ConfigureSecurityGroupPermissionsRevokePermissions is a repeated param struct in ConfigureSecurityGroupPermissionsRequest
type ConfigureSecurityGroupPermissionsRevokePermissions struct {
	SourcePortRange string `name:"SourcePortRange"`
	PortRange       string `name:"PortRange"`
	IpProtocol      string `name:"IpProtocol"`
	SourceCidrIp    string `name:"SourceCidrIp"`
	Priority        string `name:"Priority"`
	DestCidrIp      string `name:"DestCidrIp"`
	Direction       string `name:"Direction"`
	Policy          string `name:"Policy"`
}

// ConfigureSecurityGroupPermissionsAuthorizePermissions is a repeated param struct in ConfigureSecurityGroupPermissionsRequest
type ConfigureSecurityGroupPermissionsAuthorizePermissions struct {
	SourcePortRange string `name:"SourcePortRange"`
	PortRange       string `name:"PortRange"`
	IpProtocol      string `name:"IpProtocol"`
	SourceCidrIp    string `name:"SourceCidrIp"`
	Description     string `name:"Description"`
	Priority        string `name:"Priority"`
	DestCidrIp      string `name:"DestCidrIp"`
	Direction       string `name:"Direction"`
	Policy          string `name:"Policy"`
}

// ConfigureSecurityGroupPermissionsResponse is the response struct for api ConfigureSecurityGroupPermissions
type ConfigureSecurityGroupPermissionsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigureSecurityGroupPermissionsRequest creates a request to invoke ConfigureSecurityGroupPermissions API
func CreateConfigureSecurityGroupPermissionsRequest() (request *ConfigureSecurityGroupPermissionsRequest) {
	request = &ConfigureSecurityGroupPermissionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ConfigureSecurityGroupPermissions", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigureSecurityGroupPermissionsResponse creates a response to parse from ConfigureSecurityGroupPermissions response
func CreateConfigureSecurityGroupPermissionsResponse() (response *ConfigureSecurityGroupPermissionsResponse) {
	response = &ConfigureSecurityGroupPermissionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
