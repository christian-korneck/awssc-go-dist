// CDK constructs for defining an interaction between an AWS Lambda function and an Amazon SQS queue.
package awssolutionsconstructsawslambdasqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdasqs/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdasqs/v2/internal"
)

type LambdaToSqs interface {
	constructs.Construct
	DeadLetterQueue() *awssqs.DeadLetterQueue
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	SqsQueue() awssqs.Queue
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for LambdaToSqs
type jsiiProxy_LambdaToSqs struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaToSqs) DeadLetterQueue() *awssqs.DeadLetterQueue {
	var returns *awssqs.DeadLetterQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSqs) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSqs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSqs) SqsQueue() awssqs.Queue {
	var returns awssqs.Queue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSqs) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewLambdaToSqs(scope constructs.Construct, id *string, props *LambdaToSqsProps) LambdaToSqs {
	_init_.Initialize()

	j := jsiiProxy_LambdaToSqs{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-sqs.LambdaToSqs",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaToSqs_Override(l LambdaToSqs, scope constructs.Construct, id *string, props *LambdaToSqsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-sqs.LambdaToSqs",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaToSqs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-lambda-sqs.LambdaToSqs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaToSqs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaToSqsProps struct {
	// Optional user provided properties for the dead letter queue.
	DeadLetterQueueProps *awssqs.QueueProps `json:"deadLetterQueueProps"`
	// Whether to deploy a secondary queue to be used as a dead letter queue.
	DeployDeadLetterQueue *bool `json:"deployDeadLetterQueue"`
	// Whether to deploy a new VPC.
	DeployVpc *bool `json:"deployVpc"`
	// Whether to grant additional permissions to the Lambda function enabling it to purge the SQS queue.
	EnableQueuePurging *bool `json:"enableQueuePurging"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing instance of SQS queue object, Providing both this and queueProps will cause an error.
	ExistingQueueObj awssqs.Queue `json:"existingQueueObj"`
	// An existing VPC for the construct to use (construct will NOT create a new VPC in this case).
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// The number of times a message can be unsuccessfully dequeued before being moved to the dead-letter queue.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
	// Optional Name for the SQS queue URL environment variable set for the Lambda function.
	QueueEnvironmentVariableName *string `json:"queueEnvironmentVariableName"`
	// Optional user-provided props to override the default props for the SQS queue.
	QueueProps *awssqs.QueueProps `json:"queueProps"`
	// Properties to override default properties if deployVpc is true.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

