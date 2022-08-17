package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

type KubeClient struct {
	client *kubernetes.Clientset
	mu     sync.Mutex
}

func NewKubeClient() *KubeClient {

	config := ctrl.GetConfigOrDie()
	clientset := kubernetes.NewForConfigOrDie(config)

	return &KubeClient{client: clientset}
}

func (k *KubeClient) GetDeployments(ctx context.Context,
	namespace string) ([]v1.Deployment, error) {

	list, err := k.client.AppsV1().Deployments(namespace).
		List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range list.Items {
		b, _ := json.MarshalIndent(item, "", " ")

		fmt.Printf("%+v\n", string(b))
	}
	return list.Items, nil
}

func (k *KubeClient) GetDeployment(ctx context.Context,
	namespace, resource string) (*v1.Deployment, error) {

	list, err := k.client.AppsV1().Deployments(namespace).Get(ctx, resource, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	//b, _ := json.MarshalIndent(list, "", " ")
	fmt.Printf("%+v\n", *list.Spec.Replicas)
	return list, nil
}

func (k *KubeClient) GetPods(ctx context.Context, namespace, resource string) (*v12.Pod, error) {

	pods, err := k.client.CoreV1().Pods(namespace).Get(ctx, resource, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	b, _ := json.MarshalIndent(pods, "", " ")
	fmt.Printf("%+v\n", string(b))
	//podName := strings.Split(pods.Items[0].Name, "-")
	//fmt.Printf("pod name %s\n", podName[len(podName)-1])
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	return pods, nil
}

func (k *KubeClient) ListPods(ctx context.Context, namespace string, quit chan uuid.UUID) (pods *v12.PodList) {
	select {
	case <-time.After(time.Second * 2):
		var err error
		pods, err = k.client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		//podName := strings.Split(pods.Items[0].Name, "-")
		//fmt.Printf("pod name %s\n", podName[len(podName)-1])
		//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		//b, _ := json.MarshalIndent(pods, "", " ")
		//fmt.Printf("%+v\n", string(b))
	case val := <-quit:
		fmt.Printf("received quit workerId %v", val)
		if shouldWorkerQuit(ctx, val) {
			fmt.Printf("quitting workerId %v", val)
			return nil
		}

	case <-ctx.Done():
		return nil
	}
	return
}

//func (k *KubeClient) ListPods(ctx context.Context, namespace string) (pods *v12.PodList) {
//	fmt.Printf("received namespace %s\n", namespace)
//	if namespace != "default" {
//		return nil
//	}
//	var err error
//	pods, err = k.client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
//	if err != nil {
//		panic(err.Error())
//	}
//	podName := strings.Split(pods.Items[0].Name, "-")
//	fmt.Printf("pod name %s\n", podName[len(podName)-1])
//	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
//	//b, _ := json.MarshalIndent(pods, "", " ")
//	//fmt.Printf("%+v\n", string(b))
//	return pods
//}

func (k *KubeClient) Start(ctx context.Context) {
	//wg := new(sync.WaitGroup)
	//desiredWorkers := 4
	nWorkers := 3
	dWorkers := 5
	//go func() {
	//	time.After(time.Second * 5)
	//	nWorkers++
	//}()
	workers := make(map[uuid.UUID]struct{})
	quit := make(chan uuid.UUID)
	go k.startWorker(ctx, workers, nWorkers, quit)
	go func() {
		for {
			<-time.After(time.Second)
			fmt.Printf("workers length %d \n", len(workers))
		}

	}()

	go func() {
		for {
			<-time.After(time.Second * 10)
			fmt.Printf("descrease dWorker length to 2 %d \n", len(workers))
			dWorkers = dWorkers - 2
			return
			//<-time.After(time.Second * 5)
			//fmt.Printf("descrease dWorker length to 2 %d \n", len(workers))
			//dWorkers = dWorkers + 5
		}

	}()
	go func() {

		for {
			<-time.After(time.Second * 5)
			fmt.Printf("starting background check \n")
			count := len(workers) < dWorkers
			if count {
				scalingCount := dWorkers - len(workers)
				fmt.Printf("scaling worker up, count  %d\n", scalingCount)
				k.startWorker(ctx, workers, scalingCount, quit)
				//nWorkers += scalingCount
			} else if len(workers) > dWorkers {
				diff := len(workers) - dWorkers
				k.mu.Lock()
				for id, _ := range workers {
					if diff == 0 {
						fmt.Printf("breaking inner loop, workers length %d \n", len(workers))
						break
					}

					fmt.Printf("removing worker, id %v workers %d \n", id, len(workers))
					delete(workers, id)
					diff--
					quit <- id
				}
				k.mu.Unlock()
			}
		}

	}()
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (k *KubeClient) startWorker(ctx context.Context, workers map[uuid.UUID]struct{}, nWorkers int, quit chan uuid.UUID) {
	for i := 0; i < nWorkers; i++ {
		go func() {
			workerId := uuid.New()
			k.mu.Lock()
			workers[workerId] = struct{}{}
			k.mu.Unlock()
			ctxWithVal := context.WithValue(ctx, workerId, struct{}{})
			for pods := k.ListPods(ctxWithVal, "default", quit); pods != nil; pods = k.ListPods(ctxWithVal, "default", quit) {
				//for nWorkers != desiredWorkers {
				//	fmt.Printf("continue")
				//	 continue
				//}
				fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
			}
		}()
	}
}
func shouldWorkerQuit(ctx context.Context, workerId uuid.UUID) bool {
	_, ok := ctx.Value(workerId).(uuid.UUID)
	return ok
}
