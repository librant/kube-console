package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog/v2"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kube-console",
	Short: "kube-console for multi-cluster container login",
	Run: func(cmd *cobra.Command, args []string) {
		klog.Infof("kube-console run begin args: %v", args)
	},
}

// Execute cmd execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		runtime.HandleError(err)
		os.Exit(1)
	}
}
