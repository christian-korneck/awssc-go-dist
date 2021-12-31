// CDK Constructs for streaming data from AWS Kinesis Data Stream for Glue ETL custom Job processing
package awssolutionsconstructsawskinesisstreamsgluejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawskinesisstreamsgluejob/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawskinesisstreamsgluejob/v2/internal"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructscore/v2"
)

type KinesisstreamsToGluejob interface {
	constructs.Construct
	CloudwatchAlarms() *[]awscloudwatch.Alarm
	Database() awsglue.CfnDatabase
	GlueJob() awsglue.CfnJob
	GlueJobRole() awsiam.IRole
	KinesisStream() awskinesis.Stream
	Node() constructs.Node
	OutputBucket() *map[string]interface{}
	Table() awsglue.CfnTable
	ToString() *string
}

// The jsii proxy struct for KinesisstreamsToGluejob
type jsiiProxy_KinesisstreamsToGluejob struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_KinesisstreamsToGluejob) CloudwatchAlarms() *[]awscloudwatch.Alarm {
	var returns *[]awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"cloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisstreamsToGluejob) Database() awsglue.CfnDatabase {
	var returns awsglue.CfnDatabase
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisstreamsToGluejob) GlueJob() awsglue.CfnJob {
	var returns awsglue.CfnJob
	_jsii_.Get(
		j,
		"glueJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisstreamsToGluejob) GlueJobRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"glueJobRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisstreamsToGluejob) KinesisStream() awskinesis.Stream {
	var returns awskinesis.Stream
	_jsii_.Get(
		j,
		"kinesisStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisstreamsToGluejob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisstreamsToGluejob) OutputBucket() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"outputBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisstreamsToGluejob) Table() awsglue.CfnTable {
	var returns awsglue.CfnTable
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}


// Constructs a new instance of KinesisstreamsToGluejob.Based on the values set in the @props.
func NewKinesisstreamsToGluejob(scope constructs.Construct, id *string, props *KinesisstreamsToGluejobProps) KinesisstreamsToGluejob {
	_init_.Initialize()

	j := jsiiProxy_KinesisstreamsToGluejob{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-kinesisstreams-gluejob.KinesisstreamsToGluejob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of KinesisstreamsToGluejob.Based on the values set in the @props.
func NewKinesisstreamsToGluejob_Override(k KinesisstreamsToGluejob, scope constructs.Construct, id *string, props *KinesisstreamsToGluejobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-kinesisstreams-gluejob.KinesisstreamsToGluejob",
		[]interface{}{scope, id, props},
		k,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func KinesisstreamsToGluejob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-kinesisstreams-gluejob.KinesisstreamsToGluejob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (k *jsiiProxy_KinesisstreamsToGluejob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KinesisstreamsToGluejobProps struct {
	// Whether to create recommended CloudWatch alarms.
	CreateCloudWatchAlarms *bool `json:"createCloudWatchAlarms"`
	// The props for the Glue database that the construct should use to create.
	//
	// If @database is set
	// then this property is ignored. If none of @database and @databaseprops is provided, the
	// construct will define a GlueDatabase resoruce.
	DatabaseProps *awsglue.CfnDatabaseProps `json:"databaseProps"`
	// Glue Database for this construct.
	//
	// If not provided the construct will create a new Glue Database.
	// The database is where the schema for the data in Kinesis Data Streams is stored
	ExistingDatabase awsglue.CfnDatabase `json:"existingDatabase"`
	// Existing GlueJob configuration.
	//
	// If this property is provided, any properties provided through @glueJobProps is ignored
	ExistingGlueJob awsglue.CfnJob `json:"existingGlueJob"`
	// Existing instance of Kineses Data Stream.
	//
	// If not set, it will create an instance
	ExistingStreamObj awskinesis.Stream `json:"existingStreamObj"`
	// Glue Table for this construct, If not provided the construct will create a new Table in the database.
	//
	// This table should define the schema for the records in the Kinesis Data Streams.
	// One of @tableprops or @table or @fieldSchema is mandatory. If @tableprops is provided then
	ExistingTable awsglue.CfnTable `json:"existingTable"`
	// Structure of the records in the Amazon Kinesis Data Streams.
	//
	// An example of such a  definition is as below.
	// Either @table or @fieldSchema is mandatory. If @table is provided then @fieldSchema is ignored
	//  	"FieldSchema": [{
	//   	"name": "id",
	//   	"type": "int",
	//     "comment": "Identifier for the record"
	//   }, {
	//     "name": "name",
	//     "type": "string",
	//     "comment": "The name of the record"
	//   }, {
	//     "name": "type",
	//     "type": "string",
	//     "comment": "The type of the record"
	//   }, {
	//     "name": "numericvalue",
	//     "type": "int",
	//     "comment": "Some value associated with the record"
	//   },
	FieldSchema *[]*awsglue.CfnTable_ColumnProperty `json:"fieldSchema"`
	// User provides props to override the default props for Glue ETL Jobs.
	//
	// Providing both this and
	// existingGlueJob will cause an error.
	//
	// This parameter is defined as `any` to not enforce passing the Glue Job role which is a mandatory parameter
	// for CfnJobProps. If a role is not passed, the construct creates one for you and attaches the appropriate
	// role policies
	//
	// The default props will set the Glue Version 2.0, with 2 Workers and WorkerType as G1.X. For details on
	// defining a Glue Job, please refer the following link for documentation - https://docs.aws.amazon.com/glue/latest/webapi/API_Job.html
	GlueJobProps interface{} `json:"glueJobProps"`
	// User provided props to override the default props for the Kinesis Stream.
	KinesisStreamProps interface{} `json:"kinesisStreamProps"`
	// The output data stores where the transformed data should be written.
	//
	// Current supported data stores
	// include only S3, other potential stores may be added in the future.
	OutputDataStore *awssolutionsconstructscore.SinkDataStoreProps `json:"outputDataStore"`
	// The table properties for the construct to create the table.
	//
	// One of @tableprops or @table
	// or @fieldSchema is mandatory. If @tableprops is provided then @table and @fieldSchema
	// are ignored. If @table is provided, @fieldSchema is ignored
	TableProps *awsglue.CfnTableProps `json:"tableProps"`
}

