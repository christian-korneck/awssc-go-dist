package awssolutionsconstructsawscloudfrontapigateway

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-cloudfront-apigateway.CloudFrontToApiGateway",
		reflect.TypeOf((*CloudFrontToApiGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGateway", GoGetter: "ApiGateway"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontFunction", GoGetter: "CloudFrontFunction"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontLoggingBucket", GoGetter: "CloudFrontLoggingBucket"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontWebDistribution", GoGetter: "CloudFrontWebDistribution"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudFrontToApiGateway{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-cloudfront-apigateway.CloudFrontToApiGatewayProps",
		reflect.TypeOf((*CloudFrontToApiGatewayProps)(nil)).Elem(),
	)
}
