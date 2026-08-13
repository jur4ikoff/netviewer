package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	host string
	from int
	to   int
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan TCP ports",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"Scanning %s ports %d-%d\n",
			host,
			from,
			to,
		)
	},
}

func init() {
	scanCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host to scan")
	scanCmd.Flags().IntVar(&from, "from", 1, "First port")
	scanCmd.Flags().IntVar(&to, "to", 1024, "Last port")

	rootCmd.AddCommand(scanCmd)
}
