package awssolutionsconstructsawseventbridgestepfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-eventbridge-stepfunctions.EventbridgeToStepfunctions",
		reflect.TypeOf((*EventbridgeToStepfunctions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchAlarms", GoGetter: "CloudwatchAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "eventBus", GoGetter: "EventBus"},
			_jsii_.MemberProperty{JsiiProperty: "eventsRule", GoGetter: "EventsRule"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachine", GoGetter: "StateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachineLogGroup", GoGetter: "StateMachineLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EventbridgeToStepfunctions{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-eventbridge-stepfunctions.EventbridgeToStepfunctionsProps",
		reflect.TypeOf((*EventbridgeToStepfunctionsProps)(nil)).Elem(),
	)
}
