package cmd

import (
	"context"
	"fmt"

	"github.com/neelanjan00/kubectl-count/client"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// servicesCmd represents the services command
var servicesCmd = &cobra.Command{
	Use:     "services",
	Aliases: []string{"svc", "service"},
	Short:   "Count services in a namespace.",
	Long:    `Count services in a namespace, optionally filtered by a label.`,
	Run: func(cmd *cobra.Command, args []string) {

		if allNamespaces {
			namespace = ""
		}

		listOptions := metav1.ListOptions{}
		if selector != "" {
			listOptions.LabelSelector = selector
		}

		list, err := client.GetClient().CoreV1().Services(namespace).List(context.Background(), listOptions)
		if err != nil {
			panic(err)
		}

		fmt.Println(len(list.Items))
	},
}

func init() {
	countCmd.AddCommand(servicesCmd)

	servicesCmd.PersistentFlags().BoolVarP(&allNamespaces, "all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.")
	servicesCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "resource namespace")
	servicesCmd.PersistentFlags().StringVarP(&selector, "selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
}
