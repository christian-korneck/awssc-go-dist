package awssolutionsconstructsawswafwebaclapigateway

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-wafwebacl-apigateway.WafwebaclToApiGateway",
		reflect.TypeOf((*WafwebaclToApiGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGateway", GoGetter: "ApiGateway"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webacl", GoGetter: "Webacl"},
		},
		func() interface{} {
			j := jsiiProxy_WafwebaclToApiGateway{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-wafwebacl-apigateway.WafwebaclToApiGatewayProps",
		reflect.TypeOf((*WafwebaclToApiGatewayProps)(nil)).Elem(),
	)
}
