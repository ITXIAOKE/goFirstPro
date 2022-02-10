package cloudauth

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

// LivenessFaceVerify invokes the cloudauth.LivenessFaceVerify API synchronously
func (client *Client) LivenessFaceVerify(request *LivenessFaceVerifyRequest) (response *LivenessFaceVerifyResponse, err error) {
	response = CreateLivenessFaceVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// LivenessFaceVerifyWithChan invokes the cloudauth.LivenessFaceVerify API asynchronously
func (client *Client) LivenessFaceVerifyWithChan(request *LivenessFaceVerifyRequest) (<-chan *LivenessFaceVerifyResponse, <-chan error) {
	responseChan := make(chan *LivenessFaceVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LivenessFaceVerify(request)
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

// LivenessFaceVerifyWithCallback invokes the cloudauth.LivenessFaceVerify API asynchronously
func (client *Client) LivenessFaceVerifyWithCallback(request *LivenessFaceVerifyRequest, callback func(response *LivenessFaceVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LivenessFaceVerifyResponse
		var err error
		defer close(result)
		response, err = client.LivenessFaceVerify(request)
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

// LivenessFaceVerifyRequest is the request struct for api LivenessFaceVerify
type LivenessFaceVerifyRequest struct {
	*requests.RpcRequest
	ProductCode            string           `position:"Body" name:"ProductCode"`
	OssObjectName          string           `position:"Body" name:"OssObjectName"`
	FaceContrastPicture    string           `position:"Body" name:"FaceContrastPicture"`
	Ip                     string           `position:"Body" name:"Ip"`
	Mobile                 string           `position:"Body" name:"Mobile"`
	DeviceToken            string           `position:"Body" name:"DeviceToken"`
	UserId                 string           `position:"Body" name:"UserId"`
	CertifyId              string           `position:"Body" name:"CertifyId"`
	OuterOrderNo           string           `position:"Body" name:"OuterOrderNo"`
	FaceContrastPictureUrl string           `position:"Body" name:"FaceContrastPictureUrl"`
	SceneId                requests.Integer `position:"Body" name:"SceneId"`
	OssBucketName          string           `position:"Body" name:"OssBucketName"`
	Model                  string           `position:"Query" name:"Model"`
	Crop                   string           `position:"Body" name:"Crop"`
}

// LivenessFaceVerifyResponse is the response struct for api LivenessFaceVerify
type LivenessFaceVerifyResponse struct {
	*responses.BaseResponse
	Code         string       `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateLivenessFaceVerifyRequest creates a request to invoke LivenessFaceVerify API
func CreateLivenessFaceVerifyRequest() (request *LivenessFaceVerifyRequest) {
	request = &LivenessFaceVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "LivenessFaceVerify", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateLivenessFaceVerifyResponse creates a response to parse from LivenessFaceVerify response
func CreateLivenessFaceVerifyResponse() (response *LivenessFaceVerifyResponse) {
	response = &LivenessFaceVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
