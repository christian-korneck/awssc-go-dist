// CDK Constructs for Application Load Balancer to AWS Lambda integration
package awssolutionsconstructsawsalblambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsalblambda/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsalblambda/v2/internal"
)

type AlbToLambda interface {
	constructs.Construct
	LambdaFunction() awslambda.Function
	Listener() awselasticloadbalancingv2.ApplicationListener
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	Node() constructs.Node
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for AlbToLambda
type jsiiProxy_AlbToLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AlbToLambda) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbToLambda) Listener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbToLambda) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbToLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbToLambda) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewAlbToLambda(scope constructs.Construct, id *string, props *AlbToLambdaProps) AlbToLambda {
	_init_.Initialize()

	j := jsiiProxy_AlbToLambda{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-alb-lambda.AlbToLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAlbToLambda_Override(a AlbToLambda, scope constructs.Construct, id *string, props *AlbToLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-alb-lambda.AlbToLambda",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AlbToLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-alb-lambda.AlbToLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AlbToLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AlbToLambdaProps struct {
	// Whether the construct is deploying a private or public API.
	//
	// This has implications for the VPC and ALB.
	PublicApi *bool `json:"publicApi"`
	// Optional properties to customize the bucket used to store the ALB Access Logs.
	//
	// Supplying this and setting logAccessLogs to false is an error.
	AlbLoggingBucketProps *awss3.BucketProps `json:"albLoggingBucketProps"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing Application Load Balancer to incorporate into the construct architecture.
	//
	// Providing both this and loadBalancerProps is an
	// error. The VPC containing this loadBalancer must match the VPC provided in existingVpc.
	ExistingLoadBalancerObj awselasticloadbalancingv2.ApplicationLoadBalancer `json:"existingLoadBalancerObj"`
	// An existing VPC in which to deploy the construct.
	//
	// Providing both this and
	// vpcProps is an error. If the client provides an existing load balancer and/or
	// existing Private Hosted Zone, those constructs must exist in this VPC.
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Props to define the listener.
	//
	// Must be provided when adding the listener
	// to an ALB (eg - when creating the alb), may not be provided when adding
	// a second target to an already established listener. When provided, must include
	// either a certificate or protocol: HTTP
	ListenerProps interface{} `json:"listenerProps"`
	// Optional custom properties for a new loadBalancer.
	//
	// Providing both this and
	// existingLoadBalancer is an error. This cannot specify a VPC, it will use the VPC
	// in existingVpc or the VPC created by the construct.
	LoadBalancerProps interface{} `json:"loadBalancerProps"`
	// Whether to turn on Access Logs for the Application Load Balancer.
	//
	// Uses an S3 bucket
	// with associated storage costs. Enabling Access Logging is a best practice.
	LogAlbAccessLogs *bool `json:"logAlbAccessLogs"`
	// Rules for directing traffic to the target being created.
	//
	// May not be specified
	// for the first listener added to an ALB, and must be specified for the second
	// target added to a listener. Add a second target by instantiating this construct a
	// second time and providing the existingAlb from the first instantiation.
	RuleProps *awselasticloadbalancingv2.AddRuleProps `json:"ruleProps"`
	// Optional custom properties for a new target group.
	//
	// While this is a standard
	// attribute of props for ALB constructs, there are few pertinent properties for a Lambda target.
	TargetProps *awselasticloadbalancingv2.ApplicationTargetGroupProps `json:"targetProps"`
	// Optional custom properties for a VPC the construct will create.
	//
	// This VPC will
	// be used by the new ALB and any Private Hosted Zone the construct creates (that's
	// why loadBalancerProps and privateHostedZoneProps can't include a VPC). Providing
	// both this and existingVpc is an error.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

