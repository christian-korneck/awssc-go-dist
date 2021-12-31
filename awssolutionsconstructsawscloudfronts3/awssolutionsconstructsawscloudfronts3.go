// CDK Constructs for AWS Cloudfront to AWS S3 integration.
package awssolutionsconstructsawscloudfronts3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscloudfronts3/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscloudfronts3/v2/internal"
)

type CloudFrontToS3 interface {
	constructs.Construct
	CloudFrontFunction() awscloudfront.Function
	CloudFrontLoggingBucket() awss3.Bucket
	CloudFrontWebDistribution() awscloudfront.Distribution
	Node() constructs.Node
	S3Bucket() awss3.Bucket
	S3BucketInterface() awss3.IBucket
	S3LoggingBucket() awss3.Bucket
	ToString() *string
}

// The jsii proxy struct for CloudFrontToS3
type jsiiProxy_CloudFrontToS3 struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CloudFrontToS3) CloudFrontFunction() awscloudfront.Function {
	var returns awscloudfront.Function
	_jsii_.Get(
		j,
		"cloudFrontFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToS3) CloudFrontLoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"cloudFrontLoggingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToS3) CloudFrontWebDistribution() awscloudfront.Distribution {
	var returns awscloudfront.Distribution
	_jsii_.Get(
		j,
		"cloudFrontWebDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToS3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToS3) S3Bucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToS3) S3BucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3BucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToS3) S3LoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3LoggingBucket",
		&returns,
	)
	return returns
}


func NewCloudFrontToS3(scope constructs.Construct, id *string, props *CloudFrontToS3Props) CloudFrontToS3 {
	_init_.Initialize()

	j := jsiiProxy_CloudFrontToS3{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-cloudfront-s3.CloudFrontToS3",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCloudFrontToS3_Override(c CloudFrontToS3, scope constructs.Construct, id *string, props *CloudFrontToS3Props) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-cloudfront-s3.CloudFrontToS3",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CloudFrontToS3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-cloudfront-s3.CloudFrontToS3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CloudFrontToS3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudFrontToS3Props struct {
	// Optional user provided props to override the default props for the S3 Bucket.
	BucketProps *awss3.BucketProps `json:"bucketProps"`
	// Optional user provided props to override the default props.
	CloudFrontDistributionProps interface{} `json:"cloudFrontDistributionProps"`
	// Optional user provided props to override the default props for the CloudFront Logging Bucket.
	CloudFrontLoggingBucketProps *awss3.BucketProps `json:"cloudFrontLoggingBucketProps"`
	// Existing instance of S3 Bucket object, providing both this and `bucketProps` will cause an error.
	ExistingBucketObj awss3.IBucket `json:"existingBucketObj"`
	// Optional user provided props to turn on/off the automatic injection of best practice HTTP security headers in all responses from cloudfront.
	InsertHttpSecurityHeaders *bool `json:"insertHttpSecurityHeaders"`
	// Optional user provided props to override the default props for the S3 Logging Bucket.
	LoggingBucketProps *awss3.BucketProps `json:"loggingBucketProps"`
	// Whether to turn on Access Logs for the S3 bucket with the associated storage costs.
	//
	// Enabling Access Logging is a best practice.
	LogS3AccessLogs *bool `json:"logS3AccessLogs"`
	// Optional user provided props to provide an originPath that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// The string should start with a `/`, for example `/production`.
	OriginPath *string `json:"originPath"`
}

