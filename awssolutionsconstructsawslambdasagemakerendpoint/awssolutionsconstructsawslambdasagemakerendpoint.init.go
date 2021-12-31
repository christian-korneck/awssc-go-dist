package awssolutionsconstructsawslambdasagemakerendpoint

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-lambda-sagemakerendpoint.LambdaToSagemakerEndpoint",
		reflect.TypeOf((*LambdaToSagemakerEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerEndpoint", GoGetter: "SagemakerEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerEndpointConfig", GoGetter: "SagemakerEndpointConfig"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerModel", GoGetter: "SagemakerModel"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaToSagemakerEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-lambda-sagemakerendpoint.LambdaToSagemakerEndpointProps",
		reflect.TypeOf((*LambdaToSagemakerEndpointProps)(nil)).Elem(),
	)
}
