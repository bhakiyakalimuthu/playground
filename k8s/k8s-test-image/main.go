package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	c := NewKubeClient()
	ctx, cancel := context.WithCancel(context.Background())
	namespace := "default"
	//resource := "k8s-test"
	resource := "proxyd"
	done := make(chan os.Signal, 1)
	exit := make(chan struct{})
	go func() {
		signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
		<-done
		fmt.Println("shutting down")
		cancel()
		close(exit)
	}()
	go c.ListPods(ctx, namespace, nil)
	go c.GetDeployment(ctx, namespace, resource)
	<-exit
}

type KubeClient struct {
	client *kubernetes.Clientset
	mu     sync.Mutex
}

func NewKubeClient() *KubeClient {

	config := ctrl.GetConfigOrDie()
	clientset := kubernetes.NewForConfigOrDie(config)

	return &KubeClient{client: clientset}
}

func (k *KubeClient) GetDeployment(ctx context.Context,
	namespace, resource string) {

	for {

		select {
		case <-time.After(time.Second * 2):
			list, err := k.client.AppsV1().Deployments(namespace).Get(ctx, resource, metav1.GetOptions{})
			if err != nil {
				fmt.Printf("failed to get replicas %+v\n", err)
				continue
			}
			fmt.Printf("number of replicas %+v\n", *list.Spec.Replicas)

		case <-ctx.Done():
			return
		}
	}
}

func (k *KubeClient) ListPods(ctx context.Context, namespace string, quit chan uuid.UUID) {
	for {

		select {
		case <-time.After(time.Second * 2):
			pods, err := k.client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				fmt.Printf("failed to get pod list %+v\n", err)
				continue
			}

			fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		case <-ctx.Done():
			return
		}
	}
}
