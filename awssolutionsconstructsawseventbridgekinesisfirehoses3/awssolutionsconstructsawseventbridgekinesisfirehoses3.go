// CDK Constructs for Amazon CloudWatch Events Rule to Amazon Kinesis Firehose to Amazon S3 integration.
package awssolutionsconstructsawseventbridgekinesisfirehoses3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgekinesisfirehoses3/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgekinesisfirehoses3/v2/internal"
)

type EventbridgeToKinesisFirehoseToS3 interface {
	constructs.Construct
	EventBus() awsevents.IEventBus
	EventsRole() awsiam.Role
	EventsRule() awsevents.Rule
	KinesisFirehose() awskinesisfirehose.CfnDeliveryStream
	KinesisFirehoseLogGroup() awslogs.LogGroup
	KinesisFirehoseRole() awsiam.Role
	Node() constructs.Node
	S3Bucket() awss3.Bucket
	S3BucketInterface() awss3.IBucket
	S3LoggingBucket() awss3.Bucket
	ToString() *string
}

// The jsii proxy struct for EventbridgeToKinesisFirehoseToS3
type jsiiProxy_EventbridgeToKinesisFirehoseToS3 struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EventbridgeToKinesisFirehoseToS3) EventBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"eventBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisFirehoseToS3) EventsRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"eventsRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisFirehoseToS3) EventsRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisFirehoseToS3) KinesisFirehose() awskinesisfirehose.CfnDeliveryStream {
	var returns awskinesisfirehose.CfnDeliveryStream
	_jsii_.Get(
		j,
		"kinesisFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisFirehoseToS3) KinesisFirehoseLogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"kinesisFirehoseLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisFirehoseToS3) KinesisFirehoseRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"kinesisFirehoseRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisFirehoseToS3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisFirehoseToS3) S3Bucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisFirehoseToS3) S3BucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3BucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisFirehoseToS3) S3LoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3LoggingBucket",
		&returns,
	)
	return returns
}


func NewEventbridgeToKinesisFirehoseToS3(scope constructs.Construct, id *string, props *EventbridgeToKinesisFirehoseToS3Props) EventbridgeToKinesisFirehoseToS3 {
	_init_.Initialize()

	j := jsiiProxy_EventbridgeToKinesisFirehoseToS3{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-kinesisfirehose-s3.EventbridgeToKinesisFirehoseToS3",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEventbridgeToKinesisFirehoseToS3_Override(e EventbridgeToKinesisFirehoseToS3, scope constructs.Construct, id *string, props *EventbridgeToKinesisFirehoseToS3Props) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-kinesisfirehose-s3.EventbridgeToKinesisFirehoseToS3",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EventbridgeToKinesisFirehoseToS3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-eventbridge-kinesisfirehose-s3.EventbridgeToKinesisFirehoseToS3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EventbridgeToKinesisFirehoseToS3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventbridgeToKinesisFirehoseToS3Props struct {
	// User provided eventRuleProps to override the defaults.
	EventRuleProps *awsevents.RuleProps `json:"eventRuleProps"`
	// User provided props to override the default props for the S3 Bucket.
	BucketProps *awss3.BucketProps `json:"bucketProps"`
	// A new custom EventBus is created with provided props.
	EventBusProps *awsevents.EventBusProps `json:"eventBusProps"`
	// Existing instance of S3 Bucket object, providing both this and `bucketProps` will cause an error.
	ExistingBucketObj awss3.IBucket `json:"existingBucketObj"`
	// Existing instance of a custom EventBus.
	ExistingEventBusInterface awsevents.IEventBus `json:"existingEventBusInterface"`
	// User provided props to override the default props for the Kinesis Firehose.
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

