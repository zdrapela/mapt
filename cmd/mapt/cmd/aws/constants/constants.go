package constants

import (
	"github.com/redhat-developer/mapt/pkg/integrations/cirrus"
	"github.com/redhat-developer/mapt/pkg/provider/aws/action/mac"
)

const (
	MACRequestCmd      = "request"
	MACRRequestCmdDesc = "request mac machine"
	MACReleaseCmd      = "release"
	MACReleaseCmdDesc  = "release mac machine"

	MACArch              string = "arch"
	MACArchDesc          string = "MAC architecture allowed values x86, m1, m2"
	MACArchDefault       string = mac.DefaultArch
	MACOSVersion         string = "version"
	MACOSVersionDesc     string = "MACos operating system version 11, 12 on x86 and m1/m2; 13, 14, 15 on all archs"
	MACOSVersionDefault  string = mac.DefaultOSVersion
	MACFixedLocation     string = "fixed-location"
	MACFixedLocationDesc string = "if this flag is set the host will be created only on the region set by the AWS Env (AWS_DEFAULT_REGION)"
	MACDHID              string = "dedicated-host-id"
	MACDHIDDesc          string = "id for the dedicated host"

	ParamLocation                = "location"
	ParamLocationDesc            = "location for created resources in case spot flag (if available) is not passed"
	DefaultLocation              = "West US"
	ParamVMSize                  = "vmsize"
	ParamVMSizeDesc              = "size for the VM"
	DefaultVMSize                = "Standard_D8as_v5"
	// ParamSpot                    = "spot"
	// ParamSpotDesc                = "if spot is set the spot prices across all regions will be cheked and machine will be started on best spot option (price / eviction)"
	// ParamSpotTolerance           = "spot-eviction-tolerance"
	// ParamSpotToleranceDesc       = "if spot is enable we can define the minimum tolerance level of eviction. Allowed value are: lowest, low, medium, high or highest"
	// DefaultSpotTolerance         = "lowest"
	// ParamSpotExcludedRegions     = "spot-excluded-regions"
	// ParamSpotExcludedRegionsDesc = "this params allows to pass a comma separated list of regions to avoid when searching for best spot option"
)

func MACArchAsCirrusArch(arch string) *cirrus.Arch {
	switch arch {
	case "x86":
		return &cirrus.Amd64
	}
	return &cirrus.Arm64
}
