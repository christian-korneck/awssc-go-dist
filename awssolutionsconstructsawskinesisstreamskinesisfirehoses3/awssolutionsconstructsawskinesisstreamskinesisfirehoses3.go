// CDK constructs for defining an interaction between an Amazon Kinesis Data Stream (KDS), Amazon Kinesis Data Firehose (KDF) delivery stream and an Amazon S3 bucket.
package awssolutionsconstructsawskinesisstreamskinesisfirehoses3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawskinesisstreamskinesisfirehoses3/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawskinesisstreamskinesisfirehoses3/v2/internal"
)

type KinesisStreamsToKinesisFirehoseToS3 interface {
	constructs.Construct
	CloudwatchAlarms() *[]awscloudwatch.Alarm
	KinesisFirehose() awskinesisfirehose.CfnDeliveryStream
	KinesisFirehoseLogGroup() awslogs.LogGroup
	KinesisFirehoseRole() awsiam.Role
	KinesisStream() awskinesis.Stream
	KinesisStreamRole() awsiam.Role
	Node() constructs.Node
	S3Bucket() awss3.Bucket
	S3BucketInterface() awss3.IBucket
	S3LoggingBucket() awss3.Bucket
	ToString() *string
}

// The jsii proxy struct for KinesisStreamsToKinesisFirehoseToS3
type jsiiProxy_KinesisStreamsToKinesisFirehoseToS3 struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_KinesisStreamsToKinesisFirehoseToS3) CloudwatchAlarms() *[]awscloudwatch.Alarm {
	var returns *[]awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"cloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToKinesisFirehoseToS3) KinesisFirehose() awskinesisfirehose.CfnDeliveryStream {
	var returns awskinesisfirehose.CfnDeliveryStream
	_jsii_.Get(
		j,
		"kinesisFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToKinesisFirehoseToS3) KinesisFirehoseLogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"kinesisFirehoseLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToKinesisFirehoseToS3) KinesisFirehoseRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"kinesisFirehoseRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToKinesisFirehoseToS3) KinesisStream() awskinesis.Stream {
	var returns awskinesis.Stream
	_jsii_.Get(
		j,
		"kinesisStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToKinesisFirehoseToS3) KinesisStreamRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"kinesisStreamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToKinesisFirehoseToS3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToKinesisFirehoseToS3) S3Bucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToKinesisFirehoseToS3) S3BucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3BucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToKinesisFirehoseToS3) S3LoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3LoggingBucket",
		&returns,
	)
	return returns
}


func NewKinesisStreamsToKinesisFirehoseToS3(scope constructs.Construct, id *string, props *KinesisStreamsToKinesisFirehoseToS3Props) KinesisStreamsToKinesisFirehoseToS3 {
	_init_.Initialize()

	j := jsiiProxy_KinesisStreamsToKinesisFirehoseToS3{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-kinesisstreams-kinesisfirehose-s3.KinesisStreamsToKinesisFirehoseToS3",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewKinesisStreamsToKinesisFirehoseToS3_Override(k KinesisStreamsToKinesisFirehoseToS3, scope constructs.Construct, id *string, props *KinesisStreamsToKinesisFirehoseToS3Props) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-kinesisstreams-kinesisfirehose-s3.KinesisStreamsToKinesisFirehoseToS3",
		[]interface{}{scope, id, props},
		k,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func KinesisStreamsToKinesisFirehoseToS3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-kinesisstreams-kinesisfirehose-s3.KinesisStreamsToKinesisFirehoseToS3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (k *jsiiProxy_KinesisStreamsToKinesisFirehoseToS3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the KinesisStreamsToKinesisFirehoseToS3 class.
type KinesisStreamsToKinesisFirehoseToS3Props struct {
	// Optional user provided props to override the default props for the S3 Bucket.
	BucketProps *awss3.BucketProps `json:"bucketProps"`
	// Optional whether to create recommended CloudWatch alarms.
	CreateCloudWatchAlarms *bool `json:"createCloudWatchAlarms"`
	// Optional existing instance of S3 Bucket, providing both this and bucketProps will cause an error.
	//
	// Providing both this and bucketProps will cause an error.
	ExistingBucketObj awss3.IBucket `json:"existingBucketObj"`
	// Optional existing instance of logging S3 Bucket for the S3 Bucket created by the pattern.
	ExistingLoggingBucketObj awss3.IBucket `json:"existingLoggingBucketObj"`
	// Optional existing instance of Kinesis Stream, providing both this and `kinesisStreamProps` will cause an error.
	ExistingStreamObj awskinesis.Stream `json:"existingStreamObj"`
	// Optional user provided props to override the default props.
	KinesisFirehoseProps interface{} `json:"kinesisFirehoseProps"`
	// Optional user-provided props to override the default props for the Kinesis stream.
	KinesisStreamProps *awskinesis.StreamProps `json:"kinesisStreamProps"`
	// Optional user provided props to override the default props for the S3 Logging Bucket.
	LoggingBucketProps *awss3.BucketProps `json:"loggingBucketProps"`
	// Optional user provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
	// Whether to turn on Access Logs for the S3 bucket with the associated storage costs.
	//
	// Enabling Access Logging is a best practice.
	LogS3AccessLogs *bool `json:"logS3AccessLogs"`
}

