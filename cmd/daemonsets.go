package cmd

import (
	"context"
	"fmt"

	"github.com/neelanjan00/kubectl-count/client"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// daemonsetsCmd represents the daemonsets command
var daemonsetsCmd = &cobra.Command{
	Use:     "daemonsets",
	Aliases: []string{"ds", "daemonset"},
	Short:   "Count daemonsets in a namespace.",
	Long:    `Count daemonsets in a namespace, optionally filtered by a label.`,
	Run: func(cmd *cobra.Command, args []string) {

		if allNamespaces {
			namespace = ""
		}

		listOptions := metav1.ListOptions{}
		if selector != "" {
			listOptions.LabelSelector = selector
		}

		list, err := client.GetClient().AppsV1().DaemonSets(namespace).List(context.Background(), listOptions)
		if err != nil {
			panic(err)
		}

		fmt.Println(len(list.Items))
	},
}

func init() {
	countCmd.AddCommand(daemonsetsCmd)

	daemonsetsCmd.PersistentFlags().BoolVarP(&allNamespaces, "all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.")
	daemonsetsCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "resource namespace")
	daemonsetsCmd.PersistentFlags().StringVarP(&selector, "selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
}
