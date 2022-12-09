package cmd

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog/v2"

	"github.com/librant/kube-console/handler"
)

// kubeConfig kube-config path
var kubeConfig string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kube-console",
	Short: "kube-console for multi-cluster container login",
	Run: func(cmd *cobra.Command, args []string) {
		klog.Infof("kube-console run begin args: %v", args)

		router := gin.Default()
		// 静态资源加载，本例为 css,js 以及资源图片
		router.StaticFS("/static/", http.Dir("."))
		router.StaticFile("/favicon.ico", "/static/favicon.ico")

		router.GET("/login", handler.IndexHandler)
		klog.Fatal(router.Run(":8080"))
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&kubeConfig, "config", "",
		"config file (default is ./.kube/kubeconfig)")

	rootCmd.AddCommand(versionCmd)
}

// Execute cmd execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		runtime.HandleError(err)
		os.Exit(1)
	}
}
