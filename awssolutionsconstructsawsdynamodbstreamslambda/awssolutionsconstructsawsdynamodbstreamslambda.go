// CDK Constructs for AWS DynamoDB Streams to AWS Lambda integration.
package awssolutionsconstructsawsdynamodbstreamslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsdynamodbstreamslambda/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsdynamodbstreamslambda/v2/internal"
)

type DynamoDBStreamsToLambda interface {
	constructs.Construct
	DynamoTable() awsdynamodb.Table
	DynamoTableInterface() awsdynamodb.ITable
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for DynamoDBStreamsToLambda
type jsiiProxy_DynamoDBStreamsToLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DynamoDBStreamsToLambda) DynamoTable() awsdynamodb.Table {
	var returns awsdynamodb.Table
	_jsii_.Get(
		j,
		"dynamoTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambda) DynamoTableInterface() awsdynamodb.ITable {
	var returns awsdynamodb.ITable
	_jsii_.Get(
		j,
		"dynamoTableInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambda) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewDynamoDBStreamsToLambda(scope constructs.Construct, id *string, props *DynamoDBStreamsToLambdaProps) DynamoDBStreamsToLambda {
	_init_.Initialize()

	j := jsiiProxy_DynamoDBStreamsToLambda{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-dynamodbstreams-lambda.DynamoDBStreamsToLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDynamoDBStreamsToLambda_Override(d DynamoDBStreamsToLambda, scope constructs.Construct, id *string, props *DynamoDBStreamsToLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-dynamodbstreams-lambda.DynamoDBStreamsToLambda",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DynamoDBStreamsToLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-dynamodbstreams-lambda.DynamoDBStreamsToLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DynamoDBStreamsToLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DynamoDBStreamsToLambdaProps struct {
	// Whether to deploy a SQS dead letter queue when a data record reaches the Maximum Retry Attempts or Maximum Record Age, its metadata like shard ID and stream ARN will be sent to an SQS queue.
	DeploySqsDlqQueue *bool `json:"deploySqsDlqQueue"`
	// Optional user provided props to override the default props.
	DynamoEventSourceProps interface{} `json:"dynamoEventSourceProps"`
	// Optional user provided props to override the default props.
	DynamoTableProps *awsdynamodb.TableProps `json:"dynamoTableProps"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing instance of DynamoDB table object, providing both this and `dynamoTableProps` will cause an error.
	ExistingTableInterface awsdynamodb.ITable `json:"existingTableInterface"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Optional user provided properties for the SQS dead letter queue.
	SqsDlqQueueProps *awssqs.QueueProps `json:"sqsDlqQueueProps"`
}

