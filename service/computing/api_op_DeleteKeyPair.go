// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DeleteKeyPairInput struct {
	_ struct{} `type:"structure"`

	KeyName *string `locationName:"KeyName" type:"string"`
}

// String returns the string representation
func (s DeleteKeyPairInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteKeyPairOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DeleteKeyPairOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteKeyPair = "DeleteKeyPair"

// DeleteKeyPairRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DeleteKeyPairRequest.
//    req := client.DeleteKeyPairRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DeleteKeyPair
func (c *Client) DeleteKeyPairRequest(input *DeleteKeyPairInput) DeleteKeyPairRequest {
	op := &aws.Operation{
		Name:       opDeleteKeyPair,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DeleteKeyPairInput{}
	}

	req := c.newRequest(op, input, &DeleteKeyPairOutput{})
	return DeleteKeyPairRequest{Request: req, Input: input, Copy: c.DeleteKeyPairRequest}
}

// DeleteKeyPairRequest is the request type for the
// DeleteKeyPair API operation.
type DeleteKeyPairRequest struct {
	*aws.Request
	Input *DeleteKeyPairInput
	Copy  func(*DeleteKeyPairInput) DeleteKeyPairRequest
}

// Send marshals and sends the DeleteKeyPair API request.
func (r DeleteKeyPairRequest) Send(ctx context.Context) (*DeleteKeyPairResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteKeyPairResponse{
		DeleteKeyPairOutput: r.Request.Data.(*DeleteKeyPairOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteKeyPairResponse is the response type for the
// DeleteKeyPair API operation.
type DeleteKeyPairResponse struct {
	*DeleteKeyPairOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteKeyPair request.
func (r *DeleteKeyPairResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
