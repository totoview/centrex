package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/go-kit/kit/log/level"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/totoview/centrex/dgw"
	"github.com/totoview/centrex/log"
)

var configFile string

func init() {
	centrexCmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "centrex configuration")
	viper.SetDefault("centrex.addr", "127.0.0.1:3333")
	viper.SetDefault("centrex.secure", true)
	viper.SetDefault("centrex.keypair.key", "centrex.key")
	viper.SetDefault("centrex.keypair.cert", "centrex.pem")
}

var centrexCmd = &cobra.Command{
	Use:   "centrex",
	Short: "Run centrex data exchange",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		filePath, _ := filepath.Abs(configFile)
		viper.SetConfigFile(filePath)

		logger := log.Logger()
		if err = viper.ReadInConfig(); err != nil {
			level.Error(logger).Log("error", err)
			return
		}

		var config map[string]interface{}
		if err := viper.Unmarshal(&config); err != nil {
			level.Error(logger).Log("error", err)
			return
		}

		if err := VerifyConfig(config); err != nil {
			level.Error(logger).Log("error", err)
			return
		}

		addr := viper.GetString("centrex.addr")

		var dgwTCPServer *dgw.TCPServer

		if viper.GetBool("centrex.secure") {
			if dgwTCPServer, err = dgw.NewSecureTCPServer(addr, viper.GetString("centrex.keypair.cert"), viper.GetString("centrex.keypair.key")); err != nil {
				level.Error(logger).Log("error", err)
				return
			}
		} else {
			if dgwTCPServer, err = dgw.NewTCPServer(addr); err != nil {
				level.Error(logger).Log("error", err)
				return
			}
		}

		if err = dgwTCPServer.Start(); err != nil {
			level.Error(logger).Log("error", err)
			return
		}

		errc := make(chan error)

		go func() {
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			errc <- fmt.Errorf("%s", <-c)
		}()

		err = <-errc
		level.Info(logger).Log("exit", err)
		dgwTCPServer.Stop()
	},
}
