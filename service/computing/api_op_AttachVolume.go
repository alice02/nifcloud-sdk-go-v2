// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"
	"time"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type AttachVolumeInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `locationName:"InstanceId" type:"string"`

	VolumeId *string `locationName:"VolumeId" type:"string"`
}

// String returns the string representation
func (s AttachVolumeInput) String() string {
	return nifcloudutil.Prettify(s)
}

type AttachVolumeOutput struct {
	_ struct{} `type:"structure"`

	AttachTime *time.Time `locationName:"attachTime" type:"timestamp"`

	Device *string `locationName:"device" type:"string"`

	InstanceId *string `locationName:"instanceId" type:"string"`

	InstanceUniqueId *string `locationName:"instanceUniqueId" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`

	Status *string `locationName:"status" type:"string"`

	VolumeId *string `locationName:"volumeId" type:"string"`
}

// String returns the string representation
func (s AttachVolumeOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opAttachVolume = "AttachVolume"

// AttachVolumeRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using AttachVolumeRequest.
//    req := client.AttachVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/AttachVolume
func (c *Client) AttachVolumeRequest(input *AttachVolumeInput) AttachVolumeRequest {
	op := &aws.Operation{
		Name:       opAttachVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &AttachVolumeInput{}
	}

	req := c.newRequest(op, input, &AttachVolumeOutput{})
	return AttachVolumeRequest{Request: req, Input: input, Copy: c.AttachVolumeRequest}
}

// AttachVolumeRequest is the request type for the
// AttachVolume API operation.
type AttachVolumeRequest struct {
	*aws.Request
	Input *AttachVolumeInput
	Copy  func(*AttachVolumeInput) AttachVolumeRequest
}

// Send marshals and sends the AttachVolume API request.
func (r AttachVolumeRequest) Send(ctx context.Context) (*AttachVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachVolumeResponse{
		AttachVolumeOutput: r.Request.Data.(*AttachVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachVolumeResponse is the response type for the
// AttachVolume API operation.
type AttachVolumeResponse struct {
	*AttachVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachVolume request.
func (r *AttachVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
