package services

import (
	awsparams "github.com/redhat-developer/mapt/cmd/mapt/cmd/aws/constants"
	params "github.com/redhat-developer/mapt/cmd/mapt/cmd/constants"
	maptContext "github.com/redhat-developer/mapt/pkg/manager/context"
	awsEKS "github.com/redhat-developer/mapt/pkg/provider/aws/action/eks"
	"github.com/redhat-developer/mapt/pkg/util/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	cmdEKS     = "eks"
	cmdEKSDesc = "eks operations"

	paramVersion              = "version"
	paramVersionDesc          = "EKS K8s cluster version"
	paramVMSizeDesc           = "VMSize to be used on the user pool. Typically this is used to provision spot node pools" // Is this actually used in K8S cluster creation?
	defaultVersion            = "1.31"
	paramOnlySystemPool       = "only-system-pool"
	paramOnlySystemPoolDesc   = "if we do not need bunch of resources we can run only the systempool. More info https://learn.microsoft.com/es-es/azure/aks/use-system-pools?tabs=azure-cli#system-and-user-node-pools"
	// paramEnableAppRouting     = "enable-app-routing"
	// paramEnableAppRoutingDesc = "enable application routing add-on with NGINX"
)

func GetEKSCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   cmdEKS,
		Short: cmdEKSDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
	}
	c.AddCommand(getCreateEKS(), getDestroyEKS())
	return c
}

func getCreateEKS() *cobra.Command {
	c := &cobra.Command{
		Use:   params.CreateCmdName,
		Short: params.CreateCmdName,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			// // ParseEvictionRate
			// var spotToleranceValue = spotAzure.DefaultEvictionRate
			// if viper.IsSet(azparams.ParamSpotTolerance) {
			// 	var ok bool
			// 	spotToleranceValue, ok = spotAzure.ParseEvictionRate(
			// 		viper.GetString(azparams.ParamSpotTolerance))
			// 	if !ok {
			// 		return fmt.Errorf("%s is not a valid spot tolerance value", viper.GetString(azparams.ParamSpotTolerance))
			// 	}
			// }

			if err := awsEKS.Create(
				&maptContext.ContextArgs{
					ProjectName:   viper.GetString(params.ProjectName),
					BackedURL:     viper.GetString(params.BackedURL),
					ResultsOutput: viper.GetString(params.ConnectionDetailsOutput),
					Debug:         viper.IsSet(params.Debug),
					DebugLevel:    viper.GetUint(params.DebugLevel),
					Tags:          viper.GetStringMapString(params.Tags),
				},
				&awsEKS.EKSRequest{
					Prefix:              viper.GetString(params.ProjectName),
					Location:            viper.GetString(awsparams.ParamLocation),
					// Spot:                viper.IsSet(azparams.ParamSpot),
					// SpotTolerance:       spotToleranceValue,
					// SpotExcludedRegions: viper.GetStringSlice(azparams.ParamSpotExcludedRegions),
					VMSize:              viper.GetString(awsparams.ParamVMSize),
					KubernetesVersion:   viper.GetString(paramVersion)}); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
	flagSet := pflag.NewFlagSet(params.CreateCmdName, pflag.ExitOnError)
	params.AddCommonFlags(flagSet)
	flagSet.StringP(awsparams.ParamLocation, "", awsparams.DefaultLocation, awsparams.ParamLocationDesc)
	flagSet.StringP(awsparams.ParamVMSize, "", awsparams.DefaultVMSize, paramVMSizeDesc)
	flagSet.StringP(paramVersion, "", defaultVersion, paramVersionDesc)
	flagSet.Bool(paramOnlySystemPool, false, paramOnlySystemPoolDesc)
	// flagSet.Bool(paramEnableAppRouting, false, paramEnableAppRoutingDesc)
	// flagSet.StringP(awsparams.ParamSpotTolerance, "", awsparams.DefaultSpotTolerance, awsparams.ParamSpotToleranceDesc)
	// flagSet.StringSliceP(awsparams.ParamSpotExcludedRegions, "", []string{}, azparams.ParamSpotExcludedRegionsDesc)
	c.PersistentFlags().AddFlagSet(flagSet)
	return c
}

func getDestroyEKS() *cobra.Command {
	return &cobra.Command{
		Use:   params.DestroyCmdName,
		Short: params.DestroyCmdName,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			if err := awsEKS.Destroy(
				&maptContext.ContextArgs{
					ProjectName: viper.GetString(params.ProjectName),
					BackedURL:   viper.GetString(params.BackedURL),
					Debug:       viper.IsSet(params.Debug),
					DebugLevel:  viper.GetUint(params.DebugLevel),
				}); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
}
