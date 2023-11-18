package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	awscdkapigw "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/jsii-runtime-go"

	awsapigwintegrations "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2"
	awscdklambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"

	"github.com/aws/constructs-go/constructs/v10"
)

type AwsShopBackendGoStackProps struct {
	awscdk.StackProps
}

func NewAwsShopBackendGoStack(scope constructs.Construct, id string, props *AwsShopBackendGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	listProducts := awscdklambdago.NewGoFunction(stack,
		jsii.String("GetProductsListGoFunc"),
		&awscdklambdago.GoFunctionProps{
			FunctionName: jsii.String("GetProductsListGo"),
			Description:  jsii.String("get products list"),
			Entry:        jsii.String("../productslist"),
		})

	productById := awscdklambdago.NewGoFunction(stack,
		jsii.String("GetProductByIdGoFunc"),
		&awscdklambdago.GoFunctionProps{
			FunctionName: jsii.String("GetProductByIdGo"),
			Description:  jsii.String("get product by id"),
			Entry:        jsii.String("../productbyid"),
		})

	productsApi := awscdkapigw.NewHttpApi(stack, jsii.String("ProductsApiGo"), nil)

	productsApi.AddRoutes(&awscdkapigw.AddRoutesOptions{
		Path:        jsii.String("/products"),
		Methods:     &[]awscdkapigw.HttpMethod{awscdkapigw.HttpMethod_GET},
		Integration: awsapigwintegrations.NewHttpLambdaIntegration(jsii.String("ProductsListIntegration"), listProducts, nil),
	})
	productsApi.AddRoutes(&awscdkapigw.AddRoutesOptions{
		Path:        jsii.String("/products/{productId}"),
		Methods:     &[]awscdkapigw.HttpMethod{awscdkapigw.HttpMethod_GET},
		Integration: awsapigwintegrations.NewHttpLambdaIntegration(jsii.String("ProductByIdIntegration"), productById, nil),
	})

	awscdk.NewCfnOutput(stack, jsii.String("ProductsApi"), &awscdk.CfnOutputProps{
		Value:       productsApi.ApiEndpoint(),
		Description: jsii.String("the URL to the Products API"),
		ExportName:  jsii.String("ProductsAPIURL"),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewAwsShopBackendGoStack(app, "AwsShopBackendGoStack", &AwsShopBackendGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
