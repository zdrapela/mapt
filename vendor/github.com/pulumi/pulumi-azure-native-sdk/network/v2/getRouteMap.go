// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the details of a RouteMap.
//
// Uses Azure REST API version 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
func LookupRouteMap(ctx *pulumi.Context, args *LookupRouteMapArgs, opts ...pulumi.InvokeOption) (*LookupRouteMapResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteMapResult
	err := ctx.Invoke("azure-native:network:getRouteMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteMapArgs struct {
	// The resource group name of the RouteMap's resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the RouteMap.
	RouteMapName string `pulumi:"routeMapName"`
	// The name of the VirtualHub containing the RouteMap.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// The RouteMap child resource of a Virtual hub.
type LookupRouteMapResult struct {
	// List of connections which have this RoutMap associated for inbound traffic.
	AssociatedInboundConnections []string `pulumi:"associatedInboundConnections"`
	// List of connections which have this RoutMap associated for outbound traffic.
	AssociatedOutboundConnections []string `pulumi:"associatedOutboundConnections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `pulumi:"name"`
	// The provisioning state of the RouteMap resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// List of RouteMap rules to be applied.
	Rules []RouteMapRuleResponse `pulumi:"rules"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupRouteMapOutput(ctx *pulumi.Context, args LookupRouteMapOutputArgs, opts ...pulumi.InvokeOption) LookupRouteMapResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRouteMapResultOutput, error) {
			args := v.(LookupRouteMapArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getRouteMap", args, LookupRouteMapResultOutput{}, options).(LookupRouteMapResultOutput), nil
		}).(LookupRouteMapResultOutput)
}

type LookupRouteMapOutputArgs struct {
	// The resource group name of the RouteMap's resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the RouteMap.
	RouteMapName pulumi.StringInput `pulumi:"routeMapName"`
	// The name of the VirtualHub containing the RouteMap.
	VirtualHubName pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupRouteMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteMapArgs)(nil)).Elem()
}

// The RouteMap child resource of a Virtual hub.
type LookupRouteMapResultOutput struct{ *pulumi.OutputState }

func (LookupRouteMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteMapResult)(nil)).Elem()
}

func (o LookupRouteMapResultOutput) ToLookupRouteMapResultOutput() LookupRouteMapResultOutput {
	return o
}

func (o LookupRouteMapResultOutput) ToLookupRouteMapResultOutputWithContext(ctx context.Context) LookupRouteMapResultOutput {
	return o
}

// List of connections which have this RoutMap associated for inbound traffic.
func (o LookupRouteMapResultOutput) AssociatedInboundConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteMapResult) []string { return v.AssociatedInboundConnections }).(pulumi.StringArrayOutput)
}

// List of connections which have this RoutMap associated for outbound traffic.
func (o LookupRouteMapResultOutput) AssociatedOutboundConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteMapResult) []string { return v.AssociatedOutboundConnections }).(pulumi.StringArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupRouteMapResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupRouteMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupRouteMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the RouteMap resource.
func (o LookupRouteMapResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of RouteMap rules to be applied.
func (o LookupRouteMapResultOutput) Rules() RouteMapRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteMapResult) []RouteMapRuleResponse { return v.Rules }).(RouteMapRuleResponseArrayOutput)
}

// Resource type.
func (o LookupRouteMapResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteMapResultOutput{})
}
