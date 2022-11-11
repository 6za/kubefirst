/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	//"k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/kubefirst/kubefirst/pkg/k1/informer"
)

//https://firehydrant.com/blog/stay-informed-with-kubernetes-informers/
// podinformerCmd represents the podinformer command
// this doesn't use logs, as the goal is to send to kubectl logs..(stdout)
var k1InformerCmd = &cobra.Command{
	Use:   "k1-informer",
	Short: "To be used in-cluster",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("podinformer called")
		informer.StartWatcher()
	},
}

func init() {
	rootCmd.AddCommand(k1InformerCmd)

}
