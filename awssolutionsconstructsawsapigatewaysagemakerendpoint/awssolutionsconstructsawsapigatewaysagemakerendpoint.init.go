package awssolutionsconstructsawsapigatewaysagemakerendpoint

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-apigateway-sagemakerendpoint.ApiGatewayToSageMakerEndpoint",
		reflect.TypeOf((*ApiGatewayToSageMakerEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGateway", GoGetter: "ApiGateway"},
			_jsii_.MemberProperty{JsiiProperty: "apiGatewayCloudWatchRole", GoGetter: "ApiGatewayCloudWatchRole"},
			_jsii_.MemberProperty{JsiiProperty: "apiGatewayLogGroup", GoGetter: "ApiGatewayLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiGatewayRole", GoGetter: "ApiGatewayRole"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayToSageMakerEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-apigateway-sagemakerendpoint.ApiGatewayToSageMakerEndpointProps",
		reflect.TypeOf((*ApiGatewayToSageMakerEndpointProps)(nil)).Elem(),
	)
}
