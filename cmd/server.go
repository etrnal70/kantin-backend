package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	router "github.com/etrnal70/kantin-backend/pkg/routers"
	"github.com/etrnal70/kantin-backend/pkg/storage"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	// gin.SetMode(gin.ReleaseMode)

	db, err := storage.InitDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot connect to database: %s", err)
		// TODO: Should try to reconnect
		os.Exit(1)
	}
	defer db.Close()

	userServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("USER_PORT")),
		Handler: router.UserRouter(db),
	}

	sellerServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SELLER_PORT")),
		Handler: router.SellerRouter(db),
	}

	g.Go(func() error {
		return userServer.ListenAndServe()
	})

	g.Go(func() error {
		return sellerServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
