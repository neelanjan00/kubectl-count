package cmd

import (
	"context"
	"fmt"

	"github.com/neelanjan00/kubectl-count/client"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// persistentvolumesCmd represents the persistentvolumes command
var persistentvolumesCmd = &cobra.Command{
	Use:     "persistentvolumes",
	Aliases: []string{"pv", "persistentvolume"},
	Short:   "Count persistentvolumes in a cluster.",
	Long:    `Count persistentvolumes in a cluster, optionally filtered by a label.`,
	Run: func(cmd *cobra.Command, args []string) {

		listOptions := metav1.ListOptions{}
		if selector != "" {
			listOptions.LabelSelector = selector
		}

		list, err := client.GetClient().CoreV1().PersistentVolumes().List(context.Background(), listOptions)
		if err != nil {
			panic(err)
		}

		fmt.Println(len(list.Items))
	},
}

func init() {
	countCmd.AddCommand(persistentvolumesCmd)

	persistentvolumesCmd.PersistentFlags().StringVarP(&selector, "selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
}
