// CDK construct that provisions (1) an AWS Lambda function that is configured to send messages to a queue; (2) an Amazon SQS queue; and (3) an AWS Lambda function configured to consume messages from the queue.
package awssolutionsconstructsawslambdasqslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdasqslambda/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdasqslambda/v2/internal"
)

type LambdaToSqsToLambda interface {
	constructs.Construct
	ConsumerLambdaFunction() awslambda.Function
	DeadLetterQueue() *awssqs.DeadLetterQueue
	Node() constructs.Node
	ProducerLambdaFunction() awslambda.Function
	SqsQueue() awssqs.Queue
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for LambdaToSqsToLambda
type jsiiProxy_LambdaToSqsToLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaToSqsToLambda) ConsumerLambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"consumerLambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSqsToLambda) DeadLetterQueue() *awssqs.DeadLetterQueue {
	var returns *awssqs.DeadLetterQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSqsToLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSqsToLambda) ProducerLambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"producerLambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSqsToLambda) SqsQueue() awssqs.Queue {
	var returns awssqs.Queue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSqsToLambda) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewLambdaToSqsToLambda(scope constructs.Construct, id *string, props *LambdaToSqsToLambdaProps) LambdaToSqsToLambda {
	_init_.Initialize()

	j := jsiiProxy_LambdaToSqsToLambda{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-sqs-lambda.LambdaToSqsToLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaToSqsToLambda_Override(l LambdaToSqsToLambda, scope constructs.Construct, id *string, props *LambdaToSqsToLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-sqs-lambda.LambdaToSqsToLambda",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaToSqsToLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-lambda-sqs-lambda.LambdaToSqsToLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaToSqsToLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaToSqsToLambdaProps struct {
	// Optional user-provided properties to override the default properties for the consumer Lambda function.
	ConsumerLambdaFunctionProps *awslambda.FunctionProps `json:"consumerLambdaFunctionProps"`
	// Optional user-provided props to override the default props for the dead letter queue.
	//
	// Only used if the
	// `deployDeadLetterQueue` property is set to true.
	DeadLetterQueueProps *awssqs.QueueProps `json:"deadLetterQueueProps"`
	// Whether to create a secondary queue to be used as a dead letter queue.
	//
	// Defaults to `true`.
	DeployDeadLetterQueue *bool `json:"deployDeadLetterQueue"`
	// Whether to deploy a new VPC.
	DeployVpc *bool `json:"deployVpc"`
	// An optional, existing Lambda function to be used instead of the default function for receiving/consuming messages from the queue.
	//
	// Providing both this and `consumerLambdaFunctionProps` will cause an error.
	ExistingConsumerLambdaObj awslambda.Function `json:"existingConsumerLambdaObj"`
	// An optional, existing Lambda function to be used instead of the default function for sending messages to the queue.
	//
	// Providing both this and `producerLambdaFunctionProps` property will cause an error.
	ExistingProducerLambdaObj awslambda.Function `json:"existingProducerLambdaObj"`
	// An optional, existing SQS queue to be used instead of the default queue.
	//
	// Providing both this and `queueProps`
	// will cause an error.
	ExistingQueueObj awssqs.Queue `json:"existingQueueObj"`
	// An existing VPC for the construct to use (construct will NOT create a new VPC in this case).
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// The number of times a message can be unsuccessfully dequeued before being moved to the dead letter queue.
	//
	// Defaults to `15`.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
	// Optional user-provided properties to override the default properties for the producer Lambda function.
	ProducerLambdaFunctionProps *awslambda.FunctionProps `json:"producerLambdaFunctionProps"`
	// Optional Name for the SQS queue URL environment variable set for the producer Lambda function.
	QueueEnvironmentVariableName *string `json:"queueEnvironmentVariableName"`
	// Optional user-provided properties to override the default properties for the SQS queue.
	QueueProps *awssqs.QueueProps `json:"queueProps"`
	// Optional user provided properties for the queue event source.
	SqsEventSourceProps *awslambdaeventsources.SqsEventSourceProps `json:"sqsEventSourceProps"`
	// Properties to override default properties if deployVpc is true.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

