// CDK constructs for defining an AWS web WAF connected to Amazon API Gateway REST API.
package awssolutionsconstructsawswafwebaclapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawswafwebaclapigateway/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawswafwebaclapigateway/v2/internal"
)

type WafwebaclToApiGateway interface {
	constructs.Construct
	ApiGateway() awsapigateway.IRestApi
	Node() constructs.Node
	Webacl() awswafv2.CfnWebACL
	ToString() *string
}

// The jsii proxy struct for WafwebaclToApiGateway
type jsiiProxy_WafwebaclToApiGateway struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_WafwebaclToApiGateway) ApiGateway() awsapigateway.IRestApi {
	var returns awsapigateway.IRestApi
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafwebaclToApiGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafwebaclToApiGateway) Webacl() awswafv2.CfnWebACL {
	var returns awswafv2.CfnWebACL
	_jsii_.Get(
		j,
		"webacl",
		&returns,
	)
	return returns
}


func NewWafwebaclToApiGateway(scope constructs.Construct, id *string, props *WafwebaclToApiGatewayProps) WafwebaclToApiGateway {
	_init_.Initialize()

	j := jsiiProxy_WafwebaclToApiGateway{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-wafwebacl-apigateway.WafwebaclToApiGateway",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewWafwebaclToApiGateway_Override(w WafwebaclToApiGateway, scope constructs.Construct, id *string, props *WafwebaclToApiGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-wafwebacl-apigateway.WafwebaclToApiGateway",
		[]interface{}{scope, id, props},
		w,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func WafwebaclToApiGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-wafwebacl-apigateway.WafwebaclToApiGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_WafwebaclToApiGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WafwebaclToApiGatewayProps struct {
	// The existing API Gateway instance that will be protected with the WAF web ACL.
	ExistingApiGatewayInterface awsapigateway.IRestApi `json:"existingApiGatewayInterface"`
	// Existing instance of a WAF web ACL, an error will occur if this and props is set.
	ExistingWebaclObj awswafv2.CfnWebACL `json:"existingWebaclObj"`
	// Optional user-provided props to override the default props for the AWS WAF web ACL.
	WebaclProps *awswafv2.CfnWebACLProps `json:"webaclProps"`
}

