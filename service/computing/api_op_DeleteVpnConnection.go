// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DeleteVpnConnectionInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	VpnConnectionId *string `locationName:"VpnConnectionId" type:"string"`
}

// String returns the string representation
func (s DeleteVpnConnectionInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteVpnConnectionOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DeleteVpnConnectionOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteVpnConnection = "DeleteVpnConnection"

// DeleteVpnConnectionRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DeleteVpnConnectionRequest.
//    req := client.DeleteVpnConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DeleteVpnConnection
func (c *Client) DeleteVpnConnectionRequest(input *DeleteVpnConnectionInput) DeleteVpnConnectionRequest {
	op := &aws.Operation{
		Name:       opDeleteVpnConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DeleteVpnConnectionInput{}
	}

	req := c.newRequest(op, input, &DeleteVpnConnectionOutput{})
	return DeleteVpnConnectionRequest{Request: req, Input: input, Copy: c.DeleteVpnConnectionRequest}
}

// DeleteVpnConnectionRequest is the request type for the
// DeleteVpnConnection API operation.
type DeleteVpnConnectionRequest struct {
	*aws.Request
	Input *DeleteVpnConnectionInput
	Copy  func(*DeleteVpnConnectionInput) DeleteVpnConnectionRequest
}

// Send marshals and sends the DeleteVpnConnection API request.
func (r DeleteVpnConnectionRequest) Send(ctx context.Context) (*DeleteVpnConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVpnConnectionResponse{
		DeleteVpnConnectionOutput: r.Request.Data.(*DeleteVpnConnectionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVpnConnectionResponse is the response type for the
// DeleteVpnConnection API operation.
type DeleteVpnConnectionResponse struct {
	*DeleteVpnConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVpnConnection request.
func (r *DeleteVpnConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
