package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(ctx context.Context, n int) error {
	defer log.Println("return")

	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("server shutdown:", n)
			return nil
		default:
			log.Println("default", n)
		}
	}

}
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return worker(gCtx, 1)
	})

	time.Sleep(2 * time.Second)

	g.Go(func() error {
		return worker(gCtx, 2)
	})
	g.Go(func() error {
		time.Sleep(25 * time.Second)
		log.Println("20 sec off")
		return nil
	})
	// g.Go(func() error {
	// 	time.Sleep(20 * time.Second)
	// 	return errors.New("Error Big")
	// })

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}

	fmt.Println("Grace")
}
