package main

import (
    "context"
    "os"
    "os/signal"
    "time"
    "fmt"
    "net/http"
    "github.com/pkg/errors"
    "github.com/pkg/sync/errgroup"
    "syscall"
)


func main() {
    g, ctx := errgroup.WithContext(context.Background())
    
    mux := http.NewServeMux()
    mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("pong"))
    })

    serverOut := make(chan struct{})
    mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
        serverOut <- struct{}{}
    })

    server := http.Server{
        Handler: mux,
        Addr:    ":9003",
    }

    g.Go(func() error {
        return server.ListenAndServe()
    })
    
    g.Go(func() error {
        select {
        case <-ctx.Done():
            fmt.Println("errgroup exit...")
        case <-serverOut:
            fmt.Println("server will out...")
        }

        timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
        defer cancel()

        fmt.Println("shutting down server...")
        return server.Shutdown(timeoutCtx)
    })
    
    g.Go(func() error {
        quit := make(chan os.Signal, 0)
        signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

        select {
        case <-ctx.Done():
            return ctx.Err()
        case sig := <-quit:
            return errors.Errorf("get os signal: %v", sig)
        }
    })

    fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}


