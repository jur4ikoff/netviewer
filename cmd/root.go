package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "port-scanner",
	Short: "A concurrent TCP port scanner",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}