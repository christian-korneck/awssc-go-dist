// CDK constructs for defining an interaction between an Amazon Kinesis Data Stream and an AWS Lambda function.
package awssolutionsconstructsawskinesisstreamslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawskinesisstreamslambda/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawskinesisstreamslambda/v2/internal"
)

type KinesisStreamsToLambda interface {
	constructs.Construct
	CloudwatchAlarms() *[]awscloudwatch.Alarm
	KinesisStream() awskinesis.Stream
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for KinesisStreamsToLambda
type jsiiProxy_KinesisStreamsToLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_KinesisStreamsToLambda) CloudwatchAlarms() *[]awscloudwatch.Alarm {
	var returns *[]awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"cloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToLambda) KinesisStream() awskinesis.Stream {
	var returns awskinesis.Stream
	_jsii_.Get(
		j,
		"kinesisStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToLambda) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisStreamsToLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewKinesisStreamsToLambda(scope constructs.Construct, id *string, props *KinesisStreamsToLambdaProps) KinesisStreamsToLambda {
	_init_.Initialize()

	j := jsiiProxy_KinesisStreamsToLambda{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-kinesisstreams-lambda.KinesisStreamsToLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewKinesisStreamsToLambda_Override(k KinesisStreamsToLambda, scope constructs.Construct, id *string, props *KinesisStreamsToLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-kinesisstreams-lambda.KinesisStreamsToLambda",
		[]interface{}{scope, id, props},
		k,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func KinesisStreamsToLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-kinesisstreams-lambda.KinesisStreamsToLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (k *jsiiProxy_KinesisStreamsToLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the KinesisStreamsToLambda class.
type KinesisStreamsToLambdaProps struct {
	// Whether to create recommended CloudWatch alarms.
	CreateCloudWatchAlarms *bool `json:"createCloudWatchAlarms"`
	// Whether to deploy a SQS dead letter queue when a data record reaches the Maximum Retry Attempts or Maximum Record Age, its metadata like shard ID and stream ARN will be sent to an SQS queue.
	DeploySqsDlqQueue *bool `json:"deploySqsDlqQueue"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing instance of Kinesis Stream, providing both this and `kinesisStreamProps` will cause an error.
	ExistingStreamObj awskinesis.Stream `json:"existingStreamObj"`
	// Optional user-provided props to override the default props for the Lambda event source mapping.
	KinesisEventSourceProps interface{} `json:"kinesisEventSourceProps"`
	// Optional user-provided props to override the default props for the Kinesis stream.
	KinesisStreamProps *awskinesis.StreamProps `json:"kinesisStreamProps"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Optional user provided properties for the SQS dead letter queue.
	SqsDlqQueueProps *awssqs.QueueProps `json:"sqsDlqQueueProps"`
}

