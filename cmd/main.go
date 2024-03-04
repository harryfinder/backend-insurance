package main

import (
	"context"
	"fmt"
	"github.com/harryfinder/backend-Insurance/cmd/app/controller/http"
	"github.com/harryfinder/backend-Insurance/internal/config"
	"github.com/harryfinder/backend-Insurance/internal/database/pgx"
	"github.com/harryfinder/backend-Insurance/internal/entity"
	"github.com/harryfinder/backend-Insurance/internal/usecase"
	pkghttp "github.com/harryfinder/backend-Insurance/pkg/controller/http"
	gomssqldb "github.com/harryfinder/backend-Insurance/pkg/storage/mssql/go-mssqldb"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	ctx, cancelFunc := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	cfg := config.InitConfig()
	mssql, err := gomssqldb.NewMSSQLClient(ctx, cfg)
	if err != nil {
		panic(err)
	}
	database := pgx.New(mssql)
	entities := entity.New(database)
	useCase := usecase.New(entities)
	srv := pkghttp.NewServer()
	controller := http.NewController(useCase, srv)

	wg.Add(1)
	go func() {
		defer wg.Done()

		quitCh := make(chan os.Signal, 1)
		signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		sig := <-quitCh
		log.Println(sig)
		cancelFunc()

		ctx, cancelFunc = context.WithTimeout(ctx, time.Second*10)
		defer cancelFunc()

		if err := controller.Shutdown(ctx); err != nil {
			log.Println()
		}
		log.Println("Server - < finished goroutines")
	}()
	fmt.Println("Serve Started on port:", cfg.Port)
	if err = controller.Serve(ctx, &cfg); err != nil {
		panic(err)
	}
	wg.Wait()

	log.Println("Server - finished main goroutines")
}
