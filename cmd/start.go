package cmd

import (
	"fmt"
	"sync"

	"github.com/ambiyansyah/go-boilerplate/pkg/server"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  "Start server",
}

var startHTTPCmd = &cobra.Command{
	Use:   "http",
	Short: "Start HTTP server",
	Long:  "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		wg := sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer wg.Done()
			server.NewHTTPServer().Start()
		}()

		// Wait All services to end
		wg.Wait()

	},
}

var startGRPCCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start GRPC server",
	Long:  "Start GRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start-grpc called")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.AddCommand(startHTTPCmd)
	startCmd.AddCommand(startGRPCCmd)
}
