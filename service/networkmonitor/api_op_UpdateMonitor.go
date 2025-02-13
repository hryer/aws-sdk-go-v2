// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmonitor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the aggregationPeriod for a monitor. Monitors support an
// aggregationPeriod of either 30 or 60 seconds.
func (c *Client) UpdateMonitor(ctx context.Context, params *UpdateMonitorInput, optFns ...func(*Options)) (*UpdateMonitorOutput, error) {
	if params == nil {
		params = &UpdateMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMonitor", params, optFns, c.addOperationUpdateMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMonitorInput struct {

	// The aggregation time, in seconds, to change to. This must be either 30 or 60 .
	//
	// This member is required.
	AggregationPeriod *int64

	// The name of the monitor to update. Run ListMonitors to get a list of monitor
	// names.
	//
	// This member is required.
	MonitorName *string

	noSmithyDocumentSerde
}

type UpdateMonitorOutput struct {

	// The ARN of the monitor that was updated.
	//
	// This member is required.
	MonitorArn *string

	// The name of the monitor that was updated.
	//
	// This member is required.
	MonitorName *string

	// The state of the updated monitor.
	//
	// This member is required.
	State types.MonitorState

	// The changed aggregation period.
	AggregationPeriod *int64

	// The list of key-value pairs associated with the monitor.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateMonitor"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMonitor(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateMonitor",
	}
}
