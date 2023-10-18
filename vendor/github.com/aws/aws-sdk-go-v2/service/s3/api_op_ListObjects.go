// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns some or all (up to 1,000) of the objects in a bucket. You can use the
// request parameters as selection criteria to return a subset of the objects in a
// bucket. A 200 OK response can contain valid or invalid XML. Be sure to design
// your application to parse the contents of the response and handle it
// appropriately. This action has been revised. We recommend that you use the newer
// version, ListObjectsV2 (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectsV2.html)
// , when developing applications. For backward compatibility, Amazon S3 continues
// to support ListObjects . The following operations are related to ListObjects :
//   - ListObjectsV2 (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectsV2.html)
//   - GetObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
//   - PutObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//   - CreateBucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
//   - ListBuckets (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html)
func (c *Client) ListObjects(ctx context.Context, params *ListObjectsInput, optFns ...func(*Options)) (*ListObjectsOutput, error) {
	if params == nil {
		params = &ListObjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListObjects", params, optFns, c.addOperationListObjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectsInput struct {

	// The name of the bucket containing the objects. When using this action with an
	// access point, you must direct requests to the access point hostname. The access
	// point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see Using access points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide. When you use this action with Amazon S3 on
	// Outposts, you must direct requests to the S3 on Outposts hostname. The S3 on
	// Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com . When you
	// use this action with S3 on Outposts through the Amazon Web Services SDKs, you
	// provide the Outposts access point ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see What is S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Bucket *string

	// A delimiter is a character you use to group keys.
	Delimiter *string

	// Requests Amazon S3 to encode the object keys in the response and specifies the
	// encoding method to use. An object key may contain any Unicode character;
	// however, XML 1.0 parser cannot parse some characters, such as characters with an
	// ASCII value from 0 to 10. For characters that are not supported in XML 1.0, you
	// can add this parameter to request that Amazon S3 encode the keys in the
	// response.
	EncodingType types.EncodingType

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	// Marker is where you want Amazon S3 to start listing from. Amazon S3 starts
	// listing after this specified key. Marker can be any key in the bucket.
	Marker *string

	// Sets the maximum number of keys returned in the response. By default the action
	// returns up to 1,000 key names. The response might contain fewer keys but will
	// never contain more.
	MaxKeys int32

	// Limits the response to keys that begin with the specified prefix.
	Prefix *string

	// Confirms that the requester knows that she or he will be charged for the list
	// objects request. Bucket owners need not specify this parameter in their
	// requests.
	RequestPayer types.RequestPayer

	noSmithyDocumentSerde
}

type ListObjectsOutput struct {

	// All of the keys (up to 1,000) rolled up in a common prefix count as a single
	// return when calculating the number of returns. A response can contain
	// CommonPrefixes only if you specify a delimiter. CommonPrefixes contains all (if
	// there are any) keys between Prefix and the next occurrence of the string
	// specified by the delimiter. CommonPrefixes lists keys that act like
	// subdirectories in the directory specified by Prefix. For example, if the prefix
	// is notes/ and the delimiter is a slash (/) as in notes/summer/july, the common
	// prefix is notes/summer/. All of the keys that roll up into a common prefix count
	// as a single return when calculating the number of returns.
	CommonPrefixes []types.CommonPrefix

	// Metadata about each object returned.
	Contents []types.Object

	// Causes keys that contain the same string between the prefix and the first
	// occurrence of the delimiter to be rolled up into a single result element in the
	// CommonPrefixes collection. These rolled-up keys are not returned elsewhere in
	// the response. Each rolled-up result counts as only one return against the
	// MaxKeys value.
	Delimiter *string

	// Encoding type used by Amazon S3 to encode object keys in the response.
	EncodingType types.EncodingType

	// A flag that indicates whether Amazon S3 returned all of the results that
	// satisfied the search criteria.
	IsTruncated bool

	// Indicates where in the bucket listing begins. Marker is included in the
	// response if it was sent with the request.
	Marker *string

	// The maximum number of keys returned in the response body.
	MaxKeys int32

	// The bucket name.
	Name *string

	// When response is truncated (the IsTruncated element value in the response is
	// true), you can use the key name in this field as marker in the subsequent
	// request to get next set of objects. Amazon S3 lists objects in alphabetical
	// order Note: This element is returned only if you have delimiter request
	// parameter specified. If response does not include the NextMarker and it is
	// truncated, you can use the value of the last Key in the response as the marker
	// in the subsequent request to get the next set of object keys.
	NextMarker *string

	// Keys that begin with the indicated prefix.
	Prefix *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListObjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListObjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListObjects{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListObjectsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListObjects(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addListObjectsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListObjects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListObjects",
	}
}

// getListObjectsBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getListObjectsBucketMember(input interface{}) (*string, bool) {
	in := input.(*ListObjectsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addListObjectsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getListObjectsBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
