// CDK constructs for defining an interaction between an AWS Lambda function and an Amazon S3 bucket.
package awssolutionsconstructsawslambdas3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdas3/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdas3/v2/internal"
)

type LambdaToS3 interface {
	constructs.Construct
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	S3Bucket() awss3.Bucket
	S3BucketInterface() awss3.IBucket
	S3LoggingBucket() awss3.Bucket
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for LambdaToS3
type jsiiProxy_LambdaToS3 struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaToS3) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToS3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToS3) S3Bucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToS3) S3BucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3BucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToS3) S3LoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3LoggingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToS3) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewLambdaToS3(scope constructs.Construct, id *string, props *LambdaToS3Props) LambdaToS3 {
	_init_.Initialize()

	j := jsiiProxy_LambdaToS3{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-s3.LambdaToS3",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaToS3_Override(l LambdaToS3, scope constructs.Construct, id *string, props *LambdaToS3Props) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-s3.LambdaToS3",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaToS3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-lambda-s3.LambdaToS3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaToS3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaToS3Props struct {
	// Optional name for the S3 bucket environment variable set for the Lambda function.
	BucketEnvironmentVariableName *string `json:"bucketEnvironmentVariableName"`
	// Optional bucket permissions to grant to the Lambda function.
	//
	// One or more of the following may be specified: "Delete", "Put", "Read", "ReadWrite", "Write".
	BucketPermissions *[]*string `json:"bucketPermissions"`
	// Optional user provided props to override the default props for the S3 Bucket.
	BucketProps *awss3.BucketProps `json:"bucketProps"`
	// Whether to deploy a new VPC.
	DeployVpc *bool `json:"deployVpc"`
	// Existing instance of S3 Bucket object, providing both this and `bucketProps` will cause an error.
	ExistingBucketObj awss3.IBucket `json:"existingBucketObj"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// An existing VPC for the construct to use (construct will NOT create a new VPC in this case).
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// Optional user provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Optional user provided props to override the default props for the S3 Logging Bucket.
	LoggingBucketProps *awss3.BucketProps `json:"loggingBucketProps"`
	// Whether to turn on Access Logs for the S3 bucket with the associated storage costs.
	//
	// Enabling Access Logging is a best practice.
	LogS3AccessLogs *bool `json:"logS3AccessLogs"`
	// Properties to override default properties if deployVpc is true.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

