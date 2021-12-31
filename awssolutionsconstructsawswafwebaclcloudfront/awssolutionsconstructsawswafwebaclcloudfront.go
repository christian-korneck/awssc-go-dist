// CDK constructs for defining an AWS web WAF connected to Amazon CloudFront.
package awssolutionsconstructsawswafwebaclcloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawswafwebaclcloudfront/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawswafwebaclcloudfront/v2/internal"
)

type WafwebaclToCloudFront interface {
	constructs.Construct
	CloudFrontWebDistribution() awscloudfront.Distribution
	Node() constructs.Node
	Webacl() awswafv2.CfnWebACL
	ToString() *string
}

// The jsii proxy struct for WafwebaclToCloudFront
type jsiiProxy_WafwebaclToCloudFront struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_WafwebaclToCloudFront) CloudFrontWebDistribution() awscloudfront.Distribution {
	var returns awscloudfront.Distribution
	_jsii_.Get(
		j,
		"cloudFrontWebDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafwebaclToCloudFront) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafwebaclToCloudFront) Webacl() awswafv2.CfnWebACL {
	var returns awswafv2.CfnWebACL
	_jsii_.Get(
		j,
		"webacl",
		&returns,
	)
	return returns
}


func NewWafwebaclToCloudFront(scope constructs.Construct, id *string, props *WafwebaclToCloudFrontProps) WafwebaclToCloudFront {
	_init_.Initialize()

	j := jsiiProxy_WafwebaclToCloudFront{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-wafwebacl-cloudfront.WafwebaclToCloudFront",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewWafwebaclToCloudFront_Override(w WafwebaclToCloudFront, scope constructs.Construct, id *string, props *WafwebaclToCloudFrontProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-wafwebacl-cloudfront.WafwebaclToCloudFront",
		[]interface{}{scope, id, props},
		w,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func WafwebaclToCloudFront_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-wafwebacl-cloudfront.WafwebaclToCloudFront",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_WafwebaclToCloudFront) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WafwebaclToCloudFrontProps struct {
	// The existing CloudFront instance that will be protected with the WAF web ACL.
	//
	// This construct changes the CloudFront distribution by directly manipulating
	// the CloudFormation output, so this must be the Construct and cannot be
	// changed to the Interface (IDistribution)
	ExistingCloudFrontWebDistribution awscloudfront.Distribution `json:"existingCloudFrontWebDistribution"`
	// Existing instance of a WAF web ACL, an error will occur if this and props is set.
	ExistingWebaclObj awswafv2.CfnWebACL `json:"existingWebaclObj"`
	// Optional user-provided props to override the default props for the AWS WAF web ACL.
	WebaclProps *awswafv2.CfnWebACLProps `json:"webaclProps"`
}

