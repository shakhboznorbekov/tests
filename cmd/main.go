package main

import (
	"fmt"
	"log"
	auth2 "github.com/tests/internal/auth"
	"github.com/tests/internal/controller/http/v1/auth"

	item_controller "github.com/tests/internal/controller/http/v1/item"
	customer_controller "github.com/tests/internal/controller/http/v1/customer"
	transaction_controller "github.com/tests/internal/controller/http/v1/transaction"
	transactionview_controller "github.com/tests/internal/controller/http/v1/transactionview"

	user_controller "github.com/tests/internal/controller/http/v1/user"
	"github.com/tests/internal/pkg/config"
	"github.com/tests/internal/pkg/repository/postgres"
	"github.com/tests/internal/pkg/script"
	item_repo "github.com/tests/internal/repository/postgres/item"
	customer_repo "github.com/tests/internal/repository/postgres/customer"
	transaction_repo "github.com/tests/internal/repository/postgres/transaction"
	transactionview_repo "github.com/tests/internal/repository/postgres/transactionview"

	user_repo "github.com/tests/internal/repository/postgres/user"

	"github.com/tests/internal/router"
)

func main() {
	// config
	cfg := config.GetConf()

	// databases
	postgresDB := postgres.New(cfg.DBUsername, cfg.DBPassword, cfg.DBPort, cfg.DBName, config.GetConf().DefaultLang, config.GetConf().BaseUrl)

	//migration
	script.MigrateUP(postgresDB)

	// authenticator
	authenticator := auth2.New(postgresDB)

	//repository
	userRepo := user_repo.NewRepository(postgresDB)
	itemRepo := item_repo.NewRepository(postgresDB)
	customerRepo := customer_repo.NewRepository(postgresDB)
	transactionRepo := transaction_repo.NewRepository(postgresDB)
	transactionViewRepo := transactionview_repo.NewRepository(postgresDB)


	//controller
	userController := user_controller.NewController(userRepo, authenticator)
	itemController := item_controller.NewController(itemRepo)
	authController := auth.NewController(userRepo, authenticator)
	customerController := customer_controller.NewController(customerRepo)
	transactionController := transaction_controller.NewController(transactionRepo)
	transactionViewController := transactionview_controller.NewController(transactionViewRepo)

	// router
	r := router.New(authenticator, userController, authController, itemController, customerController, transactionController, transactionViewController)
	log.Fatalln(r.Init(fmt.Sprintf(":%s", cfg.Port)))

}
