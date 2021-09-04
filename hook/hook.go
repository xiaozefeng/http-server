package hook

import (
	"context"
	"os"
	"syscall"
)

// import (
// 	"context"
// 	"errors"
// 	"github/http-server/server"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"sync"
// 	"time"
// )

var ShutdownSignals = []os.Signal{
	os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGSTOP,
	syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
	syscall.SIGABRT, syscall.SIGSYS, syscall.SIGTERM,
}

type Hook func(context.Context) error

// func BuildHook(servers ...server.Server) Hook {
// 	return func(c context.Context) error {
// 		var wg sync.WaitGroup
// 		allDone := make(chan struct{})
// 		wg.Add(len(servers))
// 		for _, s := range servers {
// 			go func(srv server.Server) {
// 				defer wg.Done()
// 				err := srv.Shutdown(ctx)
// 				if err != nil {
// 					log.Printf("server shutdown error: %v \n", err)
// 				}
// 				time.Sleep(1 * time.Second)
// 			}(s)
// 		}

// 		return nil

// 		go func() {
// 			wg.Wait()
// 			allDone <- struct{}{}
// 		}()

// 		select {
// 		case <-c.Done():
// 			log.Println("closing servers timeout")
// 			return errors.New("closing hooks timeout\n")
// 		case <-allDone:
// 			log.Println("close all servers\n")
// 			return nil
// 		}

// 	}
// }

// func WaitForShutdown(hooks ...Hook) {
// 	signalChan := make(chan os.Signal, 1)
// 	signal.Notify(signalChan, nil)
// 	select {
// 	case sig := <-signalChan:
// 		log.Printf("get signal %s, application will shutdown\n", sig)
// 		time.AfterFunc(time.Minute*10, func() {
// 			log.Printf("shutdown gracefully timeout, applicaiton shutdown.\n")
// 			os.Exit(1)
// 		})
// 		for _, h := range hooks {
// 			ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
// 			err := h(ctx)
// 			if err != nil {
// 				log.Printf("failed to run hook,err: %v", err)
// 			}
// 			cancel()
// 		}
// 		os.Exit(0)
// 	}
// }
