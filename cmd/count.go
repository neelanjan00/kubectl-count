package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count instances of a Kubernetes resource.",
	Long:  `Count instances of a Kubernetes resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("Invalid Command: ", cmd.CommandPath()+" "+strings.Join(args, " "))
		}
		cmd.Usage()
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}
