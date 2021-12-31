package awssolutionsconstructsawseventbridgekinesisstreams

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-eventbridge-kinesisstreams.EventbridgeToKinesisStreams",
		reflect.TypeOf((*EventbridgeToKinesisStreams)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchAlarms", GoGetter: "CloudwatchAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "eventBus", GoGetter: "EventBus"},
			_jsii_.MemberProperty{JsiiProperty: "eventsRole", GoGetter: "EventsRole"},
			_jsii_.MemberProperty{JsiiProperty: "eventsRule", GoGetter: "EventsRule"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisStream", GoGetter: "KinesisStream"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EventbridgeToKinesisStreams{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-eventbridge-kinesisstreams.EventbridgeToKinesisStreamsProps",
		reflect.TypeOf((*EventbridgeToKinesisStreamsProps)(nil)).Elem(),
	)
}
