// CDK Constructs for AWS IoT to AWS Kinesis Firehose to AWS S3 integration.
package awssolutionsconstructsawsiotkinesisfirehoses3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsiotkinesisfirehoses3/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsiotkinesisfirehoses3/v2/internal"
)

type IotToKinesisFirehoseToS3 interface {
	constructs.Construct
	IotActionsRole() awsiam.Role
	IotTopicRule() awsiot.CfnTopicRule
	KinesisFirehose() awskinesisfirehose.CfnDeliveryStream
	KinesisFirehoseLogGroup() awslogs.LogGroup
	KinesisFirehoseRole() awsiam.Role
	Node() constructs.Node
	S3Bucket() awss3.Bucket
	S3BucketInterface() awss3.IBucket
	S3LoggingBucket() awss3.Bucket
	ToString() *string
}

// The jsii proxy struct for IotToKinesisFirehoseToS3
type jsiiProxy_IotToKinesisFirehoseToS3 struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IotToKinesisFirehoseToS3) IotActionsRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"iotActionsRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisFirehoseToS3) IotTopicRule() awsiot.CfnTopicRule {
	var returns awsiot.CfnTopicRule
	_jsii_.Get(
		j,
		"iotTopicRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisFirehoseToS3) KinesisFirehose() awskinesisfirehose.CfnDeliveryStream {
	var returns awskinesisfirehose.CfnDeliveryStream
	_jsii_.Get(
		j,
		"kinesisFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisFirehoseToS3) KinesisFirehoseLogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"kinesisFirehoseLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisFirehoseToS3) KinesisFirehoseRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"kinesisFirehoseRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisFirehoseToS3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisFirehoseToS3) S3Bucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisFirehoseToS3) S3BucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3BucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisFirehoseToS3) S3LoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3LoggingBucket",
		&returns,
	)
	return returns
}


func NewIotToKinesisFirehoseToS3(scope constructs.Construct, id *string, props *IotToKinesisFirehoseToS3Props) IotToKinesisFirehoseToS3 {
	_init_.Initialize()

	j := jsiiProxy_IotToKinesisFirehoseToS3{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-iot-kinesisfirehose-s3.IotToKinesisFirehoseToS3",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIotToKinesisFirehoseToS3_Override(i IotToKinesisFirehoseToS3, scope constructs.Construct, id *string, props *IotToKinesisFirehoseToS3Props) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-iot-kinesisfirehose-s3.IotToKinesisFirehoseToS3",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func IotToKinesisFirehoseToS3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-iot-kinesisfirehose-s3.IotToKinesisFirehoseToS3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotToKinesisFirehoseToS3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotToKinesisFirehoseToS3Props struct {
	// User provided CfnTopicRuleProps to override the defaults.
	IotTopicRuleProps *awsiot.CfnTopicRuleProps `json:"iotTopicRuleProps"`
	// User provided props to override the default props for the S3 Bucket.
	BucketProps *awss3.BucketProps `json:"bucketProps"`
	// Existing instance of S3 Bucket object, providing both this and `bucketProps` will cause an error.
	ExistingBucketObj awss3.IBucket `json:"existingBucketObj"`
	// Optional user provided props to override the default props.
	KinesisFirehoseProps interface{} `json:"kinesisFirehoseProps"`
	// Optional user provided props to override the default props for the S3 Logging Bucket.
	LoggingBucketProps *awss3.BucketProps `json:"loggingBucketProps"`
	// User provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
	// Whether to turn on Access Logs for the S3 bucket with the associated storage costs.
	//
	// Enabling Access Logging is a best practice.
	LogS3AccessLogs *bool `json:"logS3AccessLogs"`
}

