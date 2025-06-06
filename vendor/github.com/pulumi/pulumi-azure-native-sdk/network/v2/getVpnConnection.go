// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the details of a vpn connection.
//
// Uses Azure REST API version 2023-02-01.
//
// Other available API versions: 2018-07-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
func LookupVpnConnection(ctx *pulumi.Context, args *LookupVpnConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVpnConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVpnConnectionResult
	err := ctx.Invoke("azure-native:network:getVpnConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnConnectionArgs struct {
	// The name of the vpn connection.
	ConnectionName string `pulumi:"connectionName"`
	// The name of the gateway.
	GatewayName string `pulumi:"gatewayName"`
	// The resource group name of the VpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// VpnConnection Resource.
type LookupVpnConnectionResult struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidth *int `pulumi:"connectionBandwidth"`
	// The connection status.
	ConnectionStatus string `pulumi:"connectionStatus"`
	// DPD timeout in seconds for vpn connection.
	DpdTimeoutSeconds *int `pulumi:"dpdTimeoutSeconds"`
	// Egress bytes transferred.
	EgressBytesTransferred float64 `pulumi:"egressBytesTransferred"`
	// EnableBgp flag.
	EnableBgp *bool `pulumi:"enableBgp"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// EnableBgp flag.
	EnableRateLimiting *bool `pulumi:"enableRateLimiting"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Ingress bytes transferred.
	IngressBytesTransferred float64 `pulumi:"ingressBytesTransferred"`
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies []IpsecPolicyResponse `pulumi:"ipsecPolicies"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the VPN connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Id of the connected vpn site.
	RemoteVpnSite *SubResourceResponse `pulumi:"remoteVpnSite"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration *RoutingConfigurationResponse `pulumi:"routingConfiguration"`
	// Routing weight for vpn connection.
	RoutingWeight *int `pulumi:"routingWeight"`
	// SharedKey for the vpn connection.
	SharedKey *string `pulumi:"sharedKey"`
	// The Traffic Selector Policies to be considered by this connection.
	TrafficSelectorPolicies []TrafficSelectorPolicyResponse `pulumi:"trafficSelectorPolicies"`
	// Use local azure ip to initiate connection.
	UseLocalAzureIpAddress *bool `pulumi:"useLocalAzureIpAddress"`
	// Enable policy-based traffic selectors.
	UsePolicyBasedTrafficSelectors *bool `pulumi:"usePolicyBasedTrafficSelectors"`
	// Connection protocol used for this connection.
	VpnConnectionProtocolType *string `pulumi:"vpnConnectionProtocolType"`
	// List of all vpn site link connections to the gateway.
	VpnLinkConnections []VpnSiteLinkConnectionResponse `pulumi:"vpnLinkConnections"`
}

func LookupVpnConnectionOutput(ctx *pulumi.Context, args LookupVpnConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVpnConnectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpnConnectionResultOutput, error) {
			args := v.(LookupVpnConnectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getVpnConnection", args, LookupVpnConnectionResultOutput{}, options).(LookupVpnConnectionResultOutput), nil
		}).(LookupVpnConnectionResultOutput)
}

type LookupVpnConnectionOutputArgs struct {
	// The name of the vpn connection.
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The name of the gateway.
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	// The resource group name of the VpnGateway.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVpnConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnConnectionArgs)(nil)).Elem()
}

// VpnConnection Resource.
type LookupVpnConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVpnConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnConnectionResult)(nil)).Elem()
}

func (o LookupVpnConnectionResultOutput) ToLookupVpnConnectionResultOutput() LookupVpnConnectionResultOutput {
	return o
}

func (o LookupVpnConnectionResultOutput) ToLookupVpnConnectionResultOutputWithContext(ctx context.Context) LookupVpnConnectionResultOutput {
	return o
}

// Expected bandwidth in MBPS.
func (o LookupVpnConnectionResultOutput) ConnectionBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *int { return v.ConnectionBandwidth }).(pulumi.IntPtrOutput)
}

// The connection status.
func (o LookupVpnConnectionResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// DPD timeout in seconds for vpn connection.
func (o LookupVpnConnectionResultOutput) DpdTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *int { return v.DpdTimeoutSeconds }).(pulumi.IntPtrOutput)
}

// Egress bytes transferred.
func (o LookupVpnConnectionResultOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVpnConnectionResult) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

// EnableBgp flag.
func (o LookupVpnConnectionResultOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

// Enable internet security.
func (o LookupVpnConnectionResultOutput) EnableInternetSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *bool { return v.EnableInternetSecurity }).(pulumi.BoolPtrOutput)
}

// EnableBgp flag.
func (o LookupVpnConnectionResultOutput) EnableRateLimiting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *bool { return v.EnableRateLimiting }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVpnConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupVpnConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Ingress bytes transferred.
func (o LookupVpnConnectionResultOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVpnConnectionResult) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

// The IPSec Policies to be considered by this connection.
func (o LookupVpnConnectionResultOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) []IpsecPolicyResponse { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupVpnConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the VPN connection resource.
func (o LookupVpnConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Id of the connected vpn site.
func (o LookupVpnConnectionResultOutput) RemoteVpnSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *SubResourceResponse { return v.RemoteVpnSite }).(SubResourceResponsePtrOutput)
}

// The Routing Configuration indicating the associated and propagated route tables on this connection.
func (o LookupVpnConnectionResultOutput) RoutingConfiguration() RoutingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *RoutingConfigurationResponse { return v.RoutingConfiguration }).(RoutingConfigurationResponsePtrOutput)
}

// Routing weight for vpn connection.
func (o LookupVpnConnectionResultOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

// SharedKey for the vpn connection.
func (o LookupVpnConnectionResultOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

// The Traffic Selector Policies to be considered by this connection.
func (o LookupVpnConnectionResultOutput) TrafficSelectorPolicies() TrafficSelectorPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) []TrafficSelectorPolicyResponse { return v.TrafficSelectorPolicies }).(TrafficSelectorPolicyResponseArrayOutput)
}

// Use local azure ip to initiate connection.
func (o LookupVpnConnectionResultOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

// Enable policy-based traffic selectors.
func (o LookupVpnConnectionResultOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

// Connection protocol used for this connection.
func (o LookupVpnConnectionResultOutput) VpnConnectionProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *string { return v.VpnConnectionProtocolType }).(pulumi.StringPtrOutput)
}

// List of all vpn site link connections to the gateway.
func (o LookupVpnConnectionResultOutput) VpnLinkConnections() VpnSiteLinkConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) []VpnSiteLinkConnectionResponse { return v.VpnLinkConnections }).(VpnSiteLinkConnectionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnConnectionResultOutput{})
}
