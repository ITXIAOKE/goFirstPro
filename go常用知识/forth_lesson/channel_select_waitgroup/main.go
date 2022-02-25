package channel_select_waitgroup

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Server struct {
}

func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

type Hook func(ctx context.Context) error

func BuildCloseServerHook(servers ...Server) Hook {
	return func(ctx context.Context) error {
		wg := sync.WaitGroup{}
		doneCh := make(chan struct{})
		wg.Add(len(servers))
		for _, s := range servers {
			go func(srv Server) {
				err := srv.Shutdown(ctx)
				if err != nil {
					fmt.Printf("server shutdown error: %v \n", err)
				}
				time.Sleep(time.Second)
				wg.Done()
			}(s)
		}
		go func() {
			wg.Wait()
			doneCh <- struct{}{}
		}()
		select {
		case <-ctx.Done():
			fmt.Printf("closing servers timeout")
			return errors.New("timeout")
		case <-doneCh:
			fmt.Printf("close all servers")
			return nil
		}
	}
}
