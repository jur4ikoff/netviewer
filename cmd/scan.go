package cmd

import (
	scanner "github.com/jur4ikoff/netviewer/internal/commands"
	"github.com/rs/zerolog/log"
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
		ctx := cmd.Context()
		portScanner := scanner.NewScanner()

		log.Ctx(ctx).Info().Msgf("Scanning %s ports %d-%d", host, from, to)
		err := portScanner.Scan(ctx, &scanner.ScanRequest{Host: host, From: from, To: to})
		if err != nil {
			log.Ctx(ctx).Fatal().Err(err).Msgf("failed to scan host %s", host)
		}
	},
}

func init() {
	scanCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host to scan")
	scanCmd.Flags().IntVar(&from, "from", 1, "First port")
	scanCmd.Flags().IntVar(&to, "to", 1024, "Last port")

	rootCmd.AddCommand(scanCmd)
}
