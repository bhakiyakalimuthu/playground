package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"playground/k8s/worker"
	"syscall"
	"time"
)

func main() {
	c := worker.NewKubeClient()
	ctx, cancel := context.WithCancel(context.Background())
	//namespace := "goerli"
	//resource := "bundle-simulation-backend"
	//namespace := "default"
	//resource := "proxyd"
	//c.GetDeployments(context.Background(),namespace)
	//c.GetDeployment(context.Background(), namespace, resource)
	//c.GetPods(context.Background(), namespace, resource)
	done := make(chan os.Signal, 1)
	go func() {
		<-time.After(time.Second * 60)
		done <- syscall.SIGINT
	}()

	go func() {
		signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
		<-done
		fmt.Println("shutting down")
		cancel()
	}()
	//c.ListPods(ctx, namespace, nil)
	c.Start(ctx)
	//c.GetDeployment(ctx, namespace, resource)
}
