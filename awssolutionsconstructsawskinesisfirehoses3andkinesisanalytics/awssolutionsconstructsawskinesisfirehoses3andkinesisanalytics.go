// CDK constructs for defining an interaction between an Amazon Kinesis Data Firehose delivery stream and (1) an Amazon S3 bucket, and (2) an Amazon Kinesis Data Analytics application.
package awssolutionsconstructsawskinesisfirehoses3andkinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawskinesisfirehoses3andkinesisanalytics/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawskinesisfirehoses3andkinesisanalytics/v2/internal"
)

type KinesisFirehoseToAnalyticsAndS3 interface {
	constructs.Construct
	KinesisAnalytics() awskinesisanalytics.CfnApplication
	KinesisFirehose() awskinesisfirehose.CfnDeliveryStream
	KinesisFirehoseLogGroup() awslogs.LogGroup
	KinesisFirehoseRole() awsiam.Role
	Node() constructs.Node
	S3Bucket() awss3.Bucket
	S3BucketInterface() awss3.IBucket
	S3LoggingBucket() awss3.Bucket
	ToString() *string
}

// The jsii proxy struct for KinesisFirehoseToAnalyticsAndS3
type jsiiProxy_KinesisFirehoseToAnalyticsAndS3 struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_KinesisFirehoseToAnalyticsAndS3) KinesisAnalytics() awskinesisanalytics.CfnApplication {
	var returns awskinesisanalytics.CfnApplication
	_jsii_.Get(
		j,
		"kinesisAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseToAnalyticsAndS3) KinesisFirehose() awskinesisfirehose.CfnDeliveryStream {
	var returns awskinesisfirehose.CfnDeliveryStream
	_jsii_.Get(
		j,
		"kinesisFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseToAnalyticsAndS3) KinesisFirehoseLogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"kinesisFirehoseLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseToAnalyticsAndS3) KinesisFirehoseRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"kinesisFirehoseRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseToAnalyticsAndS3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseToAnalyticsAndS3) S3Bucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseToAnalyticsAndS3) S3BucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3BucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseToAnalyticsAndS3) S3LoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3LoggingBucket",
		&returns,
	)
	return returns
}


func NewKinesisFirehoseToAnalyticsAndS3(scope constructs.Construct, id *string, props *KinesisFirehoseToAnalyticsAndS3Props) KinesisFirehoseToAnalyticsAndS3 {
	_init_.Initialize()

	j := jsiiProxy_KinesisFirehoseToAnalyticsAndS3{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-kinesisfirehose-s3-and-kinesisanalytics.KinesisFirehoseToAnalyticsAndS3",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewKinesisFirehoseToAnalyticsAndS3_Override(k KinesisFirehoseToAnalyticsAndS3, scope constructs.Construct, id *string, props *KinesisFirehoseToAnalyticsAndS3Props) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-kinesisfirehose-s3-and-kinesisanalytics.KinesisFirehoseToAnalyticsAndS3",
		[]interface{}{scope, id, props},
		k,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func KinesisFirehoseToAnalyticsAndS3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-kinesisfirehose-s3-and-kinesisanalytics.KinesisFirehoseToAnalyticsAndS3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (k *jsiiProxy_KinesisFirehoseToAnalyticsAndS3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the KinesisFirehoseToAnalyticsAndS3 class.
type KinesisFirehoseToAnalyticsAndS3Props struct {
	// User provided props to override the default props for the S3 Bucket.
	BucketProps *awss3.BucketProps `json:"bucketProps"`
	// Existing instance of S3 Bucket object, providing both this and `bucketProps` will cause an error.
	ExistingBucketObj awss3.IBucket `json:"existingBucketObj"`
	// Optional user-provided props to override the default props for the Kinesis Analytics application.
	KinesisAnalyticsProps *awskinesisanalytics.CfnApplicationProps `json:"kinesisAnalyticsProps"`
	// Optional user-provided props to override the default props for the Kinesis Firehose delivery stream.
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

