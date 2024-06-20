package goapp

import (
	"fmt"
	"goapp/internal/pkg/httpsrv"
	"goapp/internal/pkg/strgen"
	"log"
	"net/http"
	"os"
)

func Start(exitChannel chan os.Signal, csrfMiddleware func(http.Handler) http.Handler) error {
	var (
		strChan = make(chan string, 100) // String channel with max parallel counter processes.
		strCli  = strgen.New(strChan)    // String generator.
		httpSrv = httpsrv.New(strChan)   // HTTP server.
	)

	// Start String Generator.
	if err := strCli.Start(); err != nil {
		return fmt.Errorf("failed to start string generator: %w", err)
	}
	defer strCli.Stop()

	// Start HTTP server.
	if err := httpSrv.Start(csrfMiddleware); err != nil {
		return fmt.Errorf("failed to start HTTP server: %w", err)
	}
	defer httpSrv.Stop()

	log.Println("GoApp Started")
	defer log.Println("GoApp Stopped")

	<-exitChannel

	return nil
}
