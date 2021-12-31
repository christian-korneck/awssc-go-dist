package awssolutionsconstructsawseventbridgesns

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-eventbridge-sns.EventbridgeToSns",
		reflect.TypeOf((*EventbridgeToSns)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "eventBus", GoGetter: "EventBus"},
			_jsii_.MemberProperty{JsiiProperty: "eventsRule", GoGetter: "EventsRule"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "snsTopic", GoGetter: "SnsTopic"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EventbridgeToSns{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-eventbridge-sns.EventbridgeToSnsProps",
		reflect.TypeOf((*EventbridgeToSnsProps)(nil)).Elem(),
	)
}
