// CDK Constructs for AWS SNS to AWS Lambda integration
package awssolutionsconstructsawssnslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawssnslambda/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawssnslambda/v2/internal"
)

type SnsToLambda interface {
	constructs.Construct
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	SnsTopic() awssns.Topic
	ToString() *string
}

// The jsii proxy struct for SnsToLambda
type jsiiProxy_SnsToLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SnsToLambda) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsToLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsToLambda) SnsTopic() awssns.Topic {
	var returns awssns.Topic
	_jsii_.Get(
		j,
		"snsTopic",
		&returns,
	)
	return returns
}


func NewSnsToLambda(scope constructs.Construct, id *string, props *SnsToLambdaProps) SnsToLambda {
	_init_.Initialize()

	j := jsiiProxy_SnsToLambda{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-sns-lambda.SnsToLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSnsToLambda_Override(s SnsToLambda, scope constructs.Construct, id *string, props *SnsToLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-sns-lambda.SnsToLambda",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SnsToLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-sns-lambda.SnsToLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SnsToLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SnsToLambdaProps struct {
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing instance of SNS Topic object, providing both this and topicProps will cause an error..
	ExistingTopicObj awssns.Topic `json:"existingTopicObj"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Optional user provided properties to override the default properties for the SNS topic.
	TopicProps *awssns.TopicProps `json:"topicProps"`
}

