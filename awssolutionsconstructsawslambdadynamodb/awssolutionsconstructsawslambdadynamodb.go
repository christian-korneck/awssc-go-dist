// CDK Constructs for AWS Lambda to AWS DynamoDB integration.
package awssolutionsconstructsawslambdadynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdadynamodb/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdadynamodb/v2/internal"
)

type LambdaToDynamoDB interface {
	constructs.Construct
	DynamoTable() awsdynamodb.Table
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for LambdaToDynamoDB
type jsiiProxy_LambdaToDynamoDB struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaToDynamoDB) DynamoTable() awsdynamodb.Table {
	var returns awsdynamodb.Table
	_jsii_.Get(
		j,
		"dynamoTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToDynamoDB) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToDynamoDB) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToDynamoDB) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewLambdaToDynamoDB(scope constructs.Construct, id *string, props *LambdaToDynamoDBProps) LambdaToDynamoDB {
	_init_.Initialize()

	j := jsiiProxy_LambdaToDynamoDB{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-dynamodb.LambdaToDynamoDB",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaToDynamoDB_Override(l LambdaToDynamoDB, scope constructs.Construct, id *string, props *LambdaToDynamoDBProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-dynamodb.LambdaToDynamoDB",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaToDynamoDB_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-lambda-dynamodb.LambdaToDynamoDB",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaToDynamoDB) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaToDynamoDBProps struct {
	// Whether to deploy a new VPC.
	DeployVpc *bool `json:"deployVpc"`
	// Optional user provided props to override the default props.
	DynamoTableProps *awsdynamodb.TableProps `json:"dynamoTableProps"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing instance of DynamoDB table object, providing both this and `dynamoTableProps` will cause an error.
	ExistingTableObj awsdynamodb.Table `json:"existingTableObj"`
	// An existing VPC for the construct to use (construct will NOT create a new VPC in this case).
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Optional Name for the DynamoDB table environment variable set for the Lambda function.
	TableEnvironmentVariableName *string `json:"tableEnvironmentVariableName"`
	// Optional table permissions to grant to the Lambda function.
	//
	// One of the following may be specified: "All", "Read", "ReadWrite", "Write".
	TablePermissions *string `json:"tablePermissions"`
	// Properties to override default properties if deployVpc is true.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

