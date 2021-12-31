// CDK constructs for defining an interaction between an Amazon Route53 domain and an Application Load Balancer.
package awssolutionsconstructsawsroute53alb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsroute53alb/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsroute53alb/v2/internal"
)

type Route53ToAlb interface {
	constructs.Construct
	HostedZone() awsroute53.IHostedZone
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	Node() constructs.Node
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for Route53ToAlb
type jsiiProxy_Route53ToAlb struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Route53ToAlb) HostedZone() awsroute53.IHostedZone {
	var returns awsroute53.IHostedZone
	_jsii_.Get(
		j,
		"hostedZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ToAlb) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ToAlb) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ToAlb) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewRoute53ToAlb(scope constructs.Construct, id *string, props *Route53ToAlbProps) Route53ToAlb {
	_init_.Initialize()

	j := jsiiProxy_Route53ToAlb{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-route53-alb.Route53ToAlb",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRoute53ToAlb_Override(r Route53ToAlb, scope constructs.Construct, id *string, props *Route53ToAlbProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-route53-alb.Route53ToAlb",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Route53ToAlb_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-route53-alb.Route53ToAlb",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_Route53ToAlb) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type Route53ToAlbProps struct {
	// Whether to create a public or private API.
	//
	// This value has implications
	// for the VPC, the type of Hosted Zone and the Application Load Balancer
	PublicApi *bool `json:"publicApi"`
	// Optional properties to customize the bucket used to store the ALB Access Logs.
	//
	// Supplying this and setting logAccessLogs to false is an error.
	AlbLoggingBucketProps *awss3.BucketProps `json:"albLoggingBucketProps"`
	// Existing Public or Private Hosted Zone.
	//
	// If a Private Hosted Zone, must
	// exist in the same VPC specified in existingVpc
	ExistingHostedZoneInterface awsroute53.IHostedZone `json:"existingHostedZoneInterface"`
	// An existing Application Load Balancer.
	//
	// Providing both this and loadBalancerProps
	// is an error. This ALB must exist in the same VPC specified in existingVPC
	ExistingLoadBalancerObj awselasticloadbalancingv2.ApplicationLoadBalancer `json:"existingLoadBalancerObj"`
	// An existing VPC.
	//
	// Providing both this and vpcProps is an error. If an existingAlb or existing
	// Private Hosted Zone is provided, this value must be the VPC associated with those resources.
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// Custom properties for a new ALB.
	//
	// Providing both this and existingLoadBalancerObj
	// is an error. These properties cannot include a VPC.
	LoadBalancerProps interface{} `json:"loadBalancerProps"`
	// Whether to turn on Access Logs for the Application Load Balancer.
	//
	// Uses an S3 bucket
	// with associated storage costs. Enabling Access Logging is a best practice.
	LogAlbAccessLogs *bool `json:"logAlbAccessLogs"`
	// Custom properties for a new Private Hosted Zone.
	//
	// Cannot be specified for a
	// public API. Cannot specify a VPC
	PrivateHostedZoneProps interface{} `json:"privateHostedZoneProps"`
	// Custom properties for a new VPC.
	//
	// Providing both this and existingVpc is
	// an error. If an existingAlb or existing Private Hosted Zone is provided, those
	// already exist in a VPC so this value cannot be provided.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

