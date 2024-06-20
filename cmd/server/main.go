package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	goapp "goapp/internal/app/server"
	"goapp/pkg/util"

	"github.com/gorilla/csrf"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmsgprefix | log.Lshortfile)
}

func main() {
	secureKey := util.GenerateSecureKey()

	CSRF := csrf.Protect(
		secureKey,
		csrf.Secure(false), // Set to true in production
	)

	// Debug.
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	// Register signal handlers for exiting
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGTERM)

	// Start.
	if err := goapp.Start(exitChannel, CSRF); err != nil {
		log.Fatalf("fatal: %+v\n", err)
	}
}
