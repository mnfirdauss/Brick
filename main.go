package main

import (
	"context"
	"flag"
	"fmt"

	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mnfirdauss/Brick/config"
	"github.com/mnfirdauss/Brick/db"
	bankHandler "github.com/mnfirdauss/Brick/internal/handler/bank"
	bankRepo "github.com/mnfirdauss/Brick/internal/repository/bank"
	bankUsecase "github.com/mnfirdauss/Brick/internal/usecase/bank"

	transactionHandler "github.com/mnfirdauss/Brick/internal/handler/transaction"
	transactionRepo "github.com/mnfirdauss/Brick/internal/repository/transaction"
	transactionUsecase "github.com/mnfirdauss/Brick/internal/usecase/transaction"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	flags := flag.NewFlagSet("db:migrate", flag.ExitOnError)
	flags.Parse(os.Args)

	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err)
	}
	var configFileName string
	flag.StringVar(&configFileName, "c", "config.yml", "Config file name")

	flag.Parse()

	cfg := config.DefaultConfig()
	cfg.LoadFromEnv()

	if len(configFileName) > 0 {
		err := config.LoadConfigFromFile(configFileName, &cfg)
		if err != nil {
			log.Warn().Str("file", configFileName).Err(err).Msg("cannot load config file, use defaults")
		}
	}

	log.Debug().Any("config", cfg).Msg("config loaded")
	args := flags.Args()

	// Migrate
	if len(args) > 1 {
		if args[1] == "up" {
			db.Migrate(cfg.DBConfig.ConnStr())
		}
	}

	// Create database connection
	connPool, err := pgxpool.NewWithConfig(context.Background(), db.Config(cfg.DBConfig.ConnStr()))
	if err != nil {
		log.Fatal().Err(err)
	}

	fmt.Println("Connected to the database!!")

	bankRepo := bankRepo.NewIAccountRepository(cfg.BaseURL.BankURL)
	bankUsecase := bankUsecase.NewBankUseCase(bankRepo)

	transactionRepo := transactionRepo.NewITransactionRepository(connPool)
	transactionUsecase := transactionUsecase.NewBanktUseCase(transactionRepo, bankRepo)

	// Initialize the router and account handler
	r := mux.NewRouter()
	bankHandler.NewAccountHandler(r, bankUsecase)
	transactionHandler.NewAccountHandler(r, transactionUsecase)

	// Start the server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		log.Info().Msg("Starting the server on :8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal().Err(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server shutdown error")
	}
	log.Info().Msg("Server gracefully stopped")
}
