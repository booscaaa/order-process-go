/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/booscaa/order-process-go/microservices/order-service/internal/adapter/database/postgres"
	"github.com/booscaa/order-process-go/microservices/order-service/internal/adapter/kafka"
	"github.com/booscaa/order-process-go/microservices/order-service/internal/di"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		kafkaWriter := kafka.Initialize("kafka:9092")
		database := postgres.Initialize()

		orderController := di.ConfigOrderDIController(database, kafkaWriter)

		mux := http.NewServeMux()
		mux.HandleFunc("POST /order", orderController.Create)

		http.ListenAndServe(":3001", mux)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

}
