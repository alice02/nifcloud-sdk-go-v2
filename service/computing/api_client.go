// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"github.com/aokumasan/nifcloud-sdk-go-v2/nifcloud"
	"github.com/aokumasan/nifcloud-sdk-go-v2/nifcloud/signer/v2"
	"github.com/aokumasan/nifcloud-sdk-go-v2/private/protocol/computing"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
)

// Client provides the API operation methods for making requests to
// computing. See this package's package overview docs
// for details on the service.
//
// The client's methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Client struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*Client)

// Used for custom request initialization logic
var initRequest func(*Client, *aws.Request)

const (
	ServiceName = "computing" // Service's name
	ServiceID   = "Computing" // Service's identifier
	EndpointsID = "computing" // Service's Endpoint identifier
)

// New creates a new instance of the client from the provided Config.
//
// Example:
//     // Create a client from just a config.
//     svc := computing.New(myConfig)
func New(config nifcloud.Config) *Client {
	svc := &Client{
		Client: aws.NewClient(
			config.AWSConfig(),
			aws.Metadata{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				EndpointsID:   EndpointsID,
				SigningName:   "computing",
				SigningRegion: config.Region,
				APIVersion:    "2016-11-15",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v2.SignRequestHandler)
	svc.Handlers.Sign.PushBackNamed(defaults.BuildContentLengthHandler)
	svc.Handlers.Build.PushBackNamed(computing.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(computing.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(computing.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(computing.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a client operation and runs any
// custom request initialization.
func (c *Client) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
