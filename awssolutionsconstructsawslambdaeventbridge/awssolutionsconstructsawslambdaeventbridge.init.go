package awssolutionsconstructsawslambdaeventbridge

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-lambda-eventbridge.LambdaToEventbridge",
		reflect.TypeOf((*LambdaToEventbridge)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "eventBus", GoGetter: "EventBus"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaToEventbridge{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-lambda-eventbridge.LambdaToEventbridgeProps",
		reflect.TypeOf((*LambdaToEventbridgeProps)(nil)).Elem(),
	)
}
