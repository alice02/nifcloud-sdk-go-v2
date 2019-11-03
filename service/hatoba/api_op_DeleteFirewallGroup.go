// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteFirewallGroupInput struct {
	_ struct{} `type:"structure"`

	// FirewallGroupName is a required field
	FirewallGroupName *string `location:"uri" locationName:"FirewallGroupName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFirewallGroupInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFirewallGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFirewallGroupInput"}

	if s.FirewallGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FirewallGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFirewallGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FirewallGroupName != nil {
		v := *s.FirewallGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FirewallGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteFirewallGroupOutput struct {
	_ struct{} `type:"structure"`

	FirewallGroup *FirewallGroupResponse `locationName:"firewallGroup" type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s DeleteFirewallGroupOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFirewallGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FirewallGroup != nil {
		v := s.FirewallGroup

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "firewallGroup", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteFirewallGroup = "DeleteFirewallGroup"

// DeleteFirewallGroupRequest returns a request value for making API operation for
// NIFCLOUD Hatoba beta.
//
//    // Example sending a request using DeleteFirewallGroupRequest.
//    req := client.DeleteFirewallGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/hatoba-2019-03-27/DeleteFirewallGroup
func (c *Client) DeleteFirewallGroupRequest(input *DeleteFirewallGroupInput) DeleteFirewallGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteFirewallGroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/firewallGroups/{FirewallGroupName}",
	}

	if input == nil {
		input = &DeleteFirewallGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteFirewallGroupOutput{})
	return DeleteFirewallGroupRequest{Request: req, Input: input, Copy: c.DeleteFirewallGroupRequest}
}

// DeleteFirewallGroupRequest is the request type for the
// DeleteFirewallGroup API operation.
type DeleteFirewallGroupRequest struct {
	*aws.Request
	Input *DeleteFirewallGroupInput
	Copy  func(*DeleteFirewallGroupInput) DeleteFirewallGroupRequest
}

// Send marshals and sends the DeleteFirewallGroup API request.
func (r DeleteFirewallGroupRequest) Send(ctx context.Context) (*DeleteFirewallGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFirewallGroupResponse{
		DeleteFirewallGroupOutput: r.Request.Data.(*DeleteFirewallGroupOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFirewallGroupResponse is the response type for the
// DeleteFirewallGroup API operation.
type DeleteFirewallGroupResponse struct {
	*DeleteFirewallGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFirewallGroup request.
func (r *DeleteFirewallGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
