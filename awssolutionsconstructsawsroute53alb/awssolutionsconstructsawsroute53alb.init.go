package awssolutionsconstructsawsroute53alb

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-route53-alb.Route53ToAlb",
		reflect.TypeOf((*Route53ToAlb)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "hostedZone", GoGetter: "HostedZone"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancer", GoGetter: "LoadBalancer"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_Route53ToAlb{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-route53-alb.Route53ToAlbProps",
		reflect.TypeOf((*Route53ToAlbProps)(nil)).Elem(),
	)
}
