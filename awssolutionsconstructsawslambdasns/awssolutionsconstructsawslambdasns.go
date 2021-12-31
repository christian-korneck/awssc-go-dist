// CDK constructs for defining an interaction between an AWS Lambda function and an Amazon SNS topic.
package awssolutionsconstructsawslambdasns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdasns/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdasns/v2/internal"
)

type LambdaToSns interface {
	constructs.Construct
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	SnsTopic() awssns.Topic
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for LambdaToSns
type jsiiProxy_LambdaToSns struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaToSns) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSns) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSns) SnsTopic() awssns.Topic {
	var returns awssns.Topic
	_jsii_.Get(
		j,
		"snsTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSns) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewLambdaToSns(scope constructs.Construct, id *string, props *LambdaToSnsProps) LambdaToSns {
	_init_.Initialize()

	j := jsiiProxy_LambdaToSns{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-sns.LambdaToSns",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaToSns_Override(l LambdaToSns, scope constructs.Construct, id *string, props *LambdaToSnsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-sns.LambdaToSns",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaToSns_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-lambda-sns.LambdaToSns",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaToSns) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaToSnsProps struct {
	// Whether to deploy a new VPC.
	DeployVpc *bool `json:"deployVpc"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing instance of SNS Topic object, providing both this and topicProps will cause an error..
	ExistingTopicObj awssns.Topic `json:"existingTopicObj"`
	// An existing VPC for the construct to use (construct will NOT create a new VPC in this case).
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Optional Name for the SNS topic arn environment variable set for the Lambda function.
	TopicArnEnvironmentVariableName *string `json:"topicArnEnvironmentVariableName"`
	// Optional Name for the SNS topic name environment variable set for the Lambda function.
	TopicNameEnvironmentVariableName *string `json:"topicNameEnvironmentVariableName"`
	// Optional user provided properties to override the default properties for the SNS topic.
	TopicProps *awssns.TopicProps `json:"topicProps"`
	// Properties to override default properties if deployVpc is true.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

