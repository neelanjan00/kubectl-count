package cmd

import (
	"context"
	"fmt"

	"github.com/neelanjan00/kubectl-count/client"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// nodesCmd represents the nodes command
var nodesCmd = &cobra.Command{
	Use:     "nodes",
	Aliases: []string{"no", "node"},
	Short:   "Count nodes in a cluster.",
	Long:    `Count nodes in a cluster, optionally filtered by a label.`,
	Run: func(cmd *cobra.Command, args []string) {

		listOptions := metav1.ListOptions{}
		if selector != "" {
			listOptions.LabelSelector = selector
		}

		list, err := client.GetClient().CoreV1().Nodes().List(context.Background(), listOptions)
		if err != nil {
			panic(err)
		}

		fmt.Println(len(list.Items))
	},
}

func init() {
	countCmd.AddCommand(nodesCmd)

	nodesCmd.PersistentFlags().StringVarP(&selector, "selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
}
