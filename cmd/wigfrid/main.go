package main

import (
	"context"
	"fmt"
	internal "github.com/dstgo/wigfrid/server"
	"github.com/dstgo/wigfrid/server/conf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

var (
	Version    string
	BuildTime  string
	ConfigFile string
)

var rootCmd = &cobra.Command{
	Use:          "wigfrid [command] [-flags]",
	Short:        "wigfrid is the gRPC daemon of wendy panel, responsible for managing local docker containers.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		// read in config file
		appConf := conf.App{
			Version:   Version,
			BuildTime: BuildTime,
		}

		v := viper.New()
		v.SetConfigFile(ConfigFile)
		if err := v.ReadInConfig(); err != nil {
			return err
		}
		// map to appConf
		if err := v.Unmarshal(&appConf); err != nil {
			return err
		}

		// initialize logger
		appConf.Log.Prompt = "[wigfrid]"
		logger, err := internal.NewLogger(appConf.Log)
		if err != nil {
			return err
		}
		defer logger.Close()
		slog.SetDefault(logger.Slog())

		// print banner
		if err := internal.PrintBanner(os.Stdout); err != nil {
			return err
		}

		// initialize app server
		app, err := internal.NewApp(ctx, &appConf)
		if err != nil {
			return err
		}
		defer app.Stop()

		return app.Run()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&ConfigFile, "config", "f", "/etc/wigfrid/conf.yaml", "server config file")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
