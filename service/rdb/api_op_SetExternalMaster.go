// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type SetExternalMasterInput struct {
	_ struct{} `type:"structure"`

	BinlogFileName *string `locationName:"BinlogFileName" type:"string"`

	BinlogPosition *int64 `locationName:"BinlogPosition" type:"integer"`

	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string"`

	MasterHost *string `locationName:"MasterHost" type:"string"`

	MasterPort *int64 `locationName:"MasterPort" type:"integer"`

	ReplicationUserName *string `locationName:"ReplicationUserName" type:"string"`

	ReplicationUserPassword *string `locationName:"ReplicationUserPassword" type:"string"`
}

// String returns the string representation
func (s SetExternalMasterInput) String() string {
	return nifcloudutil.Prettify(s)
}

type SetExternalMasterOutput struct {
	_ struct{} `type:"structure"`

	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s SetExternalMasterOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opSetExternalMaster = "SetExternalMaster"

// SetExternalMasterRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using SetExternalMasterRequest.
//    req := client.SetExternalMasterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/SetExternalMaster
func (c *Client) SetExternalMasterRequest(input *SetExternalMasterInput) SetExternalMasterRequest {
	op := &aws.Operation{
		Name:       opSetExternalMaster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetExternalMasterInput{}
	}

	req := c.newRequest(op, input, &SetExternalMasterOutput{})
	return SetExternalMasterRequest{Request: req, Input: input, Copy: c.SetExternalMasterRequest}
}

// SetExternalMasterRequest is the request type for the
// SetExternalMaster API operation.
type SetExternalMasterRequest struct {
	*aws.Request
	Input *SetExternalMasterInput
	Copy  func(*SetExternalMasterInput) SetExternalMasterRequest
}

// Send marshals and sends the SetExternalMaster API request.
func (r SetExternalMasterRequest) Send(ctx context.Context) (*SetExternalMasterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetExternalMasterResponse{
		SetExternalMasterOutput: r.Request.Data.(*SetExternalMasterOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetExternalMasterResponse is the response type for the
// SetExternalMaster API operation.
type SetExternalMasterResponse struct {
	*SetExternalMasterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetExternalMaster request.
func (r *SetExternalMasterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
