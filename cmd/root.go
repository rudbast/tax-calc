package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/rudbast/tax-calc/database"
	"github.com/rudbast/tax-calc/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootRouter *mux.Router

	rootCmd = &cobra.Command{
		Use:   "taxcalc",
		Short: "Tax-Calc is a tool to calculate your tax objects.",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Initiate config.
			viper.SetConfigType("toml")

			// Search config in home directory with name "config" (without extension).
			viper.AddConfigPath("./files/data/tax-calc")
			viper.SetConfigName("config")

			// Read in environment variables that match.
			viper.AutomaticEnv()

			// If a config file is found, read it in.
			err := viper.ReadInConfig()
			if err != nil {
				log.Fatalln("Read config file error:", err)
			}

			// Initiate database.
			db, err := database.Connect(database.Option{
				Username: viper.GetString("database.username"),
				Password: viper.GetString("database.password"),
				Host:     viper.GetString("database.host"),
				Port:     viper.GetInt("database.port"),
				Name:     viper.GetString("database.name"),
			})
			if err != nil {
				log.Fatalln(err)
			}

			// Initiate routes.
			rootRouter = mux.NewRouter()

			billHandler := handler.NewBillModule(db)
			rootRouter.Handle("/bills", handler.HandlerFunc(billHandler.InsertBill)).Methods(http.MethodPost)
			rootRouter.Handle("/bills", handler.HandlerFunc(billHandler.GetBillsSummary)).Methods(http.MethodGet)
		},
		Run: func(cmd *cobra.Command, args []string) {
			port := viper.GetInt("app.port")

			srv := &http.Server{
				Addr:    fmt.Sprintf(":%d", port),
				Handler: rootRouter,
			}

			idleConnsClosed := make(chan struct{})
			go func() {
				sigint := make(chan os.Signal, 1)
				signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
				<-sigint

				// We received an interrupt signal, shut down.
				if err := srv.Shutdown(context.Background()); err != nil {
					// Error from closing listeners, or context timeout.
					log.Println("Server shutdown error.")
				}

				log.Println("Server shutdown.")
				close(idleConnsClosed)
			}()

			log.Printf("Server listening on :%d.\n", port)
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalln("Listen and serve error.")
			}

			<-idleConnsClosed
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
