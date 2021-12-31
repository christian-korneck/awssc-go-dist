// CDK Constructs for AWS IoT to AWS Lambda integration
package awssolutionsconstructsawsiotlambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsiotlambda/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsiotlambda/v2/internal"
)

type IotToLambda interface {
	constructs.Construct
	IotTopicRule() awsiot.CfnTopicRule
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for IotToLambda
type jsiiProxy_IotToLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IotToLambda) IotTopicRule() awsiot.CfnTopicRule {
	var returns awsiot.CfnTopicRule
	_jsii_.Get(
		j,
		"iotTopicRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToLambda) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewIotToLambda(scope constructs.Construct, id *string, props *IotToLambdaProps) IotToLambda {
	_init_.Initialize()

	j := jsiiProxy_IotToLambda{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-iot-lambda.IotToLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIotToLambda_Override(i IotToLambda, scope constructs.Construct, id *string, props *IotToLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-iot-lambda.IotToLambda",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func IotToLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-iot-lambda.IotToLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotToLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotToLambdaProps struct {
	// User provided CfnTopicRuleProps to override the defaults.
	IotTopicRuleProps *awsiot.CfnTopicRuleProps `json:"iotTopicRuleProps"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
}

