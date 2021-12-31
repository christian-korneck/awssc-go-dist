// CDK constructs for defining an AWS web WAF connected to an Application Load Balancer.
package awssolutionsconstructsawswafwebaclalb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawswafwebaclalb/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawswafwebaclalb/v2/internal"
)

type WafwebaclToAlb interface {
	constructs.Construct
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	Node() constructs.Node
	Webacl() awswafv2.CfnWebACL
	ToString() *string
}

// The jsii proxy struct for WafwebaclToAlb
type jsiiProxy_WafwebaclToAlb struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_WafwebaclToAlb) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafwebaclToAlb) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafwebaclToAlb) Webacl() awswafv2.CfnWebACL {
	var returns awswafv2.CfnWebACL
	_jsii_.Get(
		j,
		"webacl",
		&returns,
	)
	return returns
}


func NewWafwebaclToAlb(scope constructs.Construct, id *string, props *WafwebaclToAlbProps) WafwebaclToAlb {
	_init_.Initialize()

	j := jsiiProxy_WafwebaclToAlb{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-wafwebacl-alb.WafwebaclToAlb",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewWafwebaclToAlb_Override(w WafwebaclToAlb, scope constructs.Construct, id *string, props *WafwebaclToAlbProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-wafwebacl-alb.WafwebaclToAlb",
		[]interface{}{scope, id, props},
		w,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func WafwebaclToAlb_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-wafwebacl-alb.WafwebaclToAlb",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_WafwebaclToAlb) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WafwebaclToAlbProps struct {
	// The existing Application Load Balancer instance that will be protected with the WAF web ACL.
	ExistingLoadBalancerObj awselasticloadbalancingv2.ApplicationLoadBalancer `json:"existingLoadBalancerObj"`
	// Existing instance of a WAF web ACL, an error will occur if this and props is set.
	ExistingWebaclObj awswafv2.CfnWebACL `json:"existingWebaclObj"`
	// Optional user-provided props to override the default props for the AWS WAF web ACL.
	WebaclProps *awswafv2.CfnWebACLProps `json:"webaclProps"`
}

