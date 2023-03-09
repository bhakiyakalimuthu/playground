package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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

func (k *KubeClient) ListPodsSingle(ctx context.Context, namespace string) (pods *v12.PodList) {
	fmt.Printf("received namespace %s\n", namespace)
	if namespace != "default" {
		return nil
	}
	var err error
	pods, err = k.client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for item, _ := range pods {

	}
	podName := strings.Split(pods.Items[0].Name, "-")
	fmt.Printf("pod name %s\n", podName[len(podName)-1])
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	//b, _ := json.MarshalIndent(pods, "", " ")
	//fmt.Printf("%+v\n", string(b))
	return pods
}

func (k *KubeClient) ListPods(ctx context.Context, namespace string, quit chan struct{}) (pods *v12.PodList) {
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
	case <-quit:
		fmt.Printf("received quit signal, quitting")
		return nil
	case <-ctx.Done():
		return nil
	}
	return
}

func (k *KubeClient) Start(ctx context.Context) {
	//wg := new(sync.WaitGroup)
	//desiredWorkers := 4
	runningWorkers := 3
	desiredWorkers := 5
	go func() {
		for {
			<-time.After(time.Second * 2)
			k.mu.Lock()
			fmt.Printf("running workers %d desired workers %d \n", runningWorkers, desiredWorkers)
			k.mu.Unlock()
		}

	}()
	quit := make(chan struct{})
	go k.startWorker(ctx, runningWorkers, quit)

	go func() {
		for {
			<-time.After(time.Second * 5)
			fmt.Printf("descrease dWorker length to 2 %d \n", runningWorkers)
			k.mu.Lock()
			desiredWorkers = desiredWorkers - 2
			k.mu.Unlock()
			return
			//<-time.After(time.Second * 5)
			//fmt.Printf("descrease dWorker length to 2 %d \n", len(workers))
			//desiredWorkers = desiredWorkers + 5
		}

	}()
	go func() {

		for {
			<-time.After(time.Second * 2)
			fmt.Printf("starting background check \n")
			k.mu.Lock()
			if runningWorkers < desiredWorkers {
				scalingCount := desiredWorkers - runningWorkers
				fmt.Printf("scaling worker up, scaling count  %d\n", scalingCount)
				k.startWorker(ctx, scalingCount, quit)
				runningWorkers += scalingCount
			} else if runningWorkers > desiredWorkers {
				diff := runningWorkers - desiredWorkers
				for i := 0; i < diff; i++ {
					fmt.Printf("removing worker id : %d", i)
					runningWorkers--
					quit <- struct{}{}
				}
			}
			k.mu.Unlock()
		}

	}()
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (k *KubeClient) startWorker(ctx context.Context, nWorkers int, quit chan struct{}) {
	for i := 0; i < nWorkers; i++ {
		go func() {
			for pods := k.ListPods(ctx, "default", quit); pods != nil; pods = k.ListPods(ctx, "default", quit) {
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
