package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/koifish082/golang-api-layered/src/app/config"
	"github.com/koifish082/golang-api-layered/src/app/interfaces/api"
	"github.com/sirupsen/logrus"

	"github.com/koifish082/golang-api-layered/src/app/library/log"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts golang-api server",
	Long: `serve' Starts golang-api server, serving 
	 my sample API`,
	Run: runServer,
}

const (
	serverTimeout = 10 * time.Second
)

var requiredEnvironments = []string{
	config.Env,
	config.EnvPort,
	config.EnvLogLevel,
}

var optionalEnvironments = []string{}

func init() {
	rootCmd.AddCommand(serverCmd)
	for _, e := range requiredEnvironments {
		err := viper.BindEnv(e)
		if err != nil {
			log.Fatalf("Failed to bind required environment variable: %s, %v", e, err)
		}
		if v := viper.GetString(e); v == "" {
			log.Fatalln(errors.Errorf("Required environment value is empty : " + e))
		}
	}
	for _, e := range optionalEnvironments {
		err := viper.BindEnv(e)
		if err != nil {
			log.Fatalf("Failed to bind optional environment variable: %s, %v", e, err)
		}
	}
}

func runServer(cmd *cobra.Command, args []string) {
	fmt.Print("Hello world")
	initLogger()
	startAPIServer()
}

func initLogger() {
	level, err := logrus.ParseLevel(viper.GetString(config.EnvLogLevel))
	if err != nil {
		log.Fatalf("invalid log level. level:%s", level)
	}
	details := map[string]string{
		"env": viper.GetString(config.Env),
	}
	cfg := &log.Config{
		ServiceName:    "golang-api-layered",
		LogLevel:       level,
		ServiceDetails: details,
	}
	log.InitLogger(cfg)
}

func startAPIServer() {
	server := http.Server{
		Addr:         viper.GetString(config.EnvPort),
		Handler:      api.NewApp().Router(),
		ReadTimeout:  serverTimeout,
		WriteTimeout: serverTimeout,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start running server")
	}
}
