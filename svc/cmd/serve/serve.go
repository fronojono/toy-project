package serve

import (
	"fmt"
	"log"
	"os"
	"toy-project/svc/configs"
	"toy-project/svc/server"

	"github.com/spf13/cobra"
)

var (
	Cmd *cobra.Command

	argAddress   string
	argGRPCAddr  string
	argDsn       string
	argCORSHosts string
)

func init() {
	Cmd = &cobra.Command{
		Use:   "serve",
		Short: "Connect to the storage and begin serving requests.",
		Long:  ``,
		Run: func(Cmd *cobra.Command, args []string) {
			if err := serve(Cmd, args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}

	Cmd.Flags().StringVarP(&argAddress, "address", "a", ":8080", "address to listen on")
	Cmd.Flags().StringVar(&argGRPCAddr, "grpc-address", ":9301", "GRPC address of configurator")
	Cmd.Flags().StringVar(&argDsn, "dsn", "postgres://postgres@localhost:5432/toy_project_db?sslmode=disable", "db url")
	Cmd.Flags().StringVar(&argCORSHosts, "cors-hosts", "*", "cors hosts, separated by comma")
}

func serve(cmd *cobra.Command, args []string) error {
	svr, err := server.NewServer(&configs.Config{
		HostPort:    argAddress,
		GRPCAddress: argGRPCAddr,
		Dsn:         argDsn,
		CORSHosts:   argCORSHosts,
	})
	if err != nil {
		return err
	}

	log.Fatalln(svr.Run())

	return nil
}
