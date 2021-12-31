// CDK Constructs for AWS S3 to AWS Step Functions integration
package awssolutionsconstructsawss3stepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawss3stepfunctions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawss3stepfunctions/v2/internal"
)

type S3ToStepfunctions interface {
	constructs.Construct
	CloudwatchAlarms() *[]awscloudwatch.Alarm
	Node() constructs.Node
	S3Bucket() awss3.Bucket
	S3BucketInterface() awss3.IBucket
	S3LoggingBucket() awss3.Bucket
	StateMachine() awsstepfunctions.StateMachine
	StateMachineLogGroup() awslogs.ILogGroup
	ToString() *string
}

// The jsii proxy struct for S3ToStepfunctions
type jsiiProxy_S3ToStepfunctions struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_S3ToStepfunctions) CloudwatchAlarms() *[]awscloudwatch.Alarm {
	var returns *[]awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"cloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToStepfunctions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToStepfunctions) S3Bucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToStepfunctions) S3BucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3BucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToStepfunctions) S3LoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3LoggingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToStepfunctions) StateMachine() awsstepfunctions.StateMachine {
	var returns awsstepfunctions.StateMachine
	_jsii_.Get(
		j,
		"stateMachine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToStepfunctions) StateMachineLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"stateMachineLogGroup",
		&returns,
	)
	return returns
}


func NewS3ToStepfunctions(scope constructs.Construct, id *string, props *S3ToStepfunctionsProps) S3ToStepfunctions {
	_init_.Initialize()

	j := jsiiProxy_S3ToStepfunctions{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-s3-stepfunctions.S3ToStepfunctions",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewS3ToStepfunctions_Override(s S3ToStepfunctions, scope constructs.Construct, id *string, props *S3ToStepfunctionsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-s3-stepfunctions.S3ToStepfunctions",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func S3ToStepfunctions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-s3-stepfunctions.S3ToStepfunctions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_S3ToStepfunctions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type S3ToStepfunctionsProps struct {
	// User provided StateMachineProps to override the defaults.
	StateMachineProps *awsstepfunctions.StateMachineProps `json:"stateMachineProps"`
	// Optional user provided props to override the default props for the S3 Bucket.
	BucketProps *awss3.BucketProps `json:"bucketProps"`
	// Whether to create recommended CloudWatch alarms.
	CreateCloudWatchAlarms *bool `json:"createCloudWatchAlarms"`
	// Whether to deploy a Trail in AWS CloudTrail to log API events in Amazon S3.
	DeployCloudTrail *bool `json:"deployCloudTrail"`
	// Optional user provided eventRuleProps to override the defaults.
	EventRuleProps *awsevents.RuleProps `json:"eventRuleProps"`
	// Existing instance of S3 Bucket object, providing both this and `bucketProps` will cause an error.
	ExistingBucketObj awss3.IBucket `json:"existingBucketObj"`
	// Optional user provided props to override the default props for the S3 Logging Bucket.
	LoggingBucketProps *awss3.BucketProps `json:"loggingBucketProps"`
	// Optional user provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
	// Whether to turn on Access Logs for the S3 bucket with the associated storage costs.
	//
	// Enabling Access Logging is a best practice.
	LogS3AccessLogs *bool `json:"logS3AccessLogs"`
}

