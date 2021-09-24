package cmd

import (
	"fmt"
	"time"

	"github.com/Insperias/lvl2-golang/develop/dev11/go-calendar/api"
	"github.com/Insperias/lvl2-golang/develop/dev11/go-calendar/internal/domain"
	"github.com/spf13/cobra"
)

var (
	paramHost           string
	paramPort           string
	paramReadTimeout    uint8
	paramWriteTimeout   uint8
	paramMaxHeaderBytes int

	apiServerCmd = &cobra.Command{
		Use:   "api-server",
		Short: "Run REST API server",
		Run: func(cmd *cobra.Command, args []string) {
			config := &api.ServerConfig{
				Host:           paramHost,
				Port:           paramPort,
				ReadTimeout:    time.Duration(paramReadTimeout) * time.Second,
				WriteTimeout:   time.Duration(paramWriteTimeout) * time.Second,
				MaxHeaderBytes: paramMaxHeaderBytes,
			}

			storage := domain.NewStorage()

			fmt.Println("Starting API sever...")
			fmt.Printf("  - Options: %+v\n", config)
			api.StartServer(storage, config)
		},
	}
)

func init() {
	apiServerCmd.PersistentFlags().StringVar(&paramHost, "host", "", "Listening host")
	apiServerCmd.PersistentFlags().StringVar(&paramPort, "port", "8080", "Listening port")
	apiServerCmd.PersistentFlags().Uint8Var(&paramReadTimeout, "read-timeout", 10, "Read timeout in sec")
	apiServerCmd.PersistentFlags().Uint8Var(&paramWriteTimeout, "write-timeout", 10, "Write timeout in sec")
	apiServerCmd.PersistentFlags().IntVar(&paramMaxHeaderBytes, "max-header-size", 1<<20, "Max header size in bytes")
}
