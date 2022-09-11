package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// stores value of the all-namespaces flag
var allNamespaces bool

// stores value of the namespace flag
var namespace string

// stores value of the selector flag
var selector string

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count Kubernetes resource instances.",
	Long:  `Count Kubernetes resource instances.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := countCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
