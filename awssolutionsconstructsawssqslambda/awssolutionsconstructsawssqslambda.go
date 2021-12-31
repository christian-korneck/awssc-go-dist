// CDK constructs for defining an interaction between an Amazon SQS queue and an AWS Lambda function.
package awssolutionsconstructsawssqslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawssqslambda/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawssqslambda/v2/internal"
)

type SqsToLambda interface {
	constructs.Construct
	DeadLetterQueue() *awssqs.DeadLetterQueue
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	SqsQueue() awssqs.Queue
	ToString() *string
}

// The jsii proxy struct for SqsToLambda
type jsiiProxy_SqsToLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SqsToLambda) DeadLetterQueue() *awssqs.DeadLetterQueue {
	var returns *awssqs.DeadLetterQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsToLambda) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsToLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsToLambda) SqsQueue() awssqs.Queue {
	var returns awssqs.Queue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}


func NewSqsToLambda(scope constructs.Construct, id *string, props *SqsToLambdaProps) SqsToLambda {
	_init_.Initialize()

	j := jsiiProxy_SqsToLambda{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-sqs-lambda.SqsToLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSqsToLambda_Override(s SqsToLambda, scope constructs.Construct, id *string, props *SqsToLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-sqs-lambda.SqsToLambda",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SqsToLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-sqs-lambda.SqsToLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SqsToLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SqsToLambdaProps struct {
	// Optional user provided properties for the dead letter queue.
	DeadLetterQueueProps *awssqs.QueueProps `json:"deadLetterQueueProps"`
	// Whether to deploy a secondary queue to be used as a dead letter queue.
	DeployDeadLetterQueue *bool `json:"deployDeadLetterQueue"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing instance of SQS queue object, Providing both this and queueProps will cause an error.
	ExistingQueueObj awssqs.Queue `json:"existingQueueObj"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// The number of times a message can be unsuccessfully dequeued before being moved to the dead-letter queue.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
	// Optional user provided properties.
	QueueProps *awssqs.QueueProps `json:"queueProps"`
	// Optional user provided properties for the queue event source.
	SqsEventSourceProps *awslambdaeventsources.SqsEventSourceProps `json:"sqsEventSourceProps"`
}

