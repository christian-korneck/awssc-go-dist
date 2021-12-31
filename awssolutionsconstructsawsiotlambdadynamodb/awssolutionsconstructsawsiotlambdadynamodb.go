// CDK Constructs for AWS IoT to AWS Lambda to AWS DyanmoDB integration.
package awssolutionsconstructsawsiotlambdadynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsiotlambdadynamodb/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsiotlambdadynamodb/v2/internal"
)

type IotToLambdaToDynamoDB interface {
	constructs.Construct
	DynamoTable() awsdynamodb.Table
	IotTopicRule() awsiot.CfnTopicRule
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for IotToLambdaToDynamoDB
type jsiiProxy_IotToLambdaToDynamoDB struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IotToLambdaToDynamoDB) DynamoTable() awsdynamodb.Table {
	var returns awsdynamodb.Table
	_jsii_.Get(
		j,
		"dynamoTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToLambdaToDynamoDB) IotTopicRule() awsiot.CfnTopicRule {
	var returns awsiot.CfnTopicRule
	_jsii_.Get(
		j,
		"iotTopicRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToLambdaToDynamoDB) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToLambdaToDynamoDB) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewIotToLambdaToDynamoDB(scope constructs.Construct, id *string, props *IotToLambdaToDynamoDBProps) IotToLambdaToDynamoDB {
	_init_.Initialize()

	j := jsiiProxy_IotToLambdaToDynamoDB{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-iot-lambda-dynamodb.IotToLambdaToDynamoDB",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIotToLambdaToDynamoDB_Override(i IotToLambdaToDynamoDB, scope constructs.Construct, id *string, props *IotToLambdaToDynamoDBProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-iot-lambda-dynamodb.IotToLambdaToDynamoDB",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func IotToLambdaToDynamoDB_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-iot-lambda-dynamodb.IotToLambdaToDynamoDB",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotToLambdaToDynamoDB) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotToLambdaToDynamoDBProps struct {
	// User provided props to override the default props.
	IotTopicRuleProps *awsiot.CfnTopicRuleProps `json:"iotTopicRuleProps"`
	// Optional user provided props to override the default props.
	DynamoTableProps *awsdynamodb.TableProps `json:"dynamoTableProps"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Optional table permissions to grant to the Lambda function.
	//
	// One of the following may be specified: "All", "Read", "ReadWrite", "Write".
	TablePermissions *string `json:"tablePermissions"`
}

