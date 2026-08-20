package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
)

var postgresClusterGVR = schema.GroupVersionResource{
	Group:    "database.example.com",
	Version:  "v1",
	Resource: "postgresclusters",
}

type Controller struct {
	kubeClient    kubernetes.Interface
	dynamicClient dynamic.Interface
	queue         workqueue.RateLimitingInterface
}

func NewController(
	kubeClient kubernetes.Interface,
	dynamicClient dynamic.Interface,
) *Controller {
	return &Controller{
		kubeClient:    kubeClient,
		dynamicClient: dynamicClient,
		queue:         workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
	}
}

func (c *Controller) reconcile(
	ctx context.Context,
	key string,
) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}

	fmt.Printf("Reconciling %s/%s\n", namespace, name)

	cluster, err := c.dynamicClient.
		Resource(postgresClusterGVR).
		Namespace(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	data, _ := json.MarshalIndent(cluster, "", "   ")
	fmt.Println(string(data))

	return nil
}

func main() {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	controller := NewController(kubeClient, client)
	if err := controller.reconcile(context.Background(), "default"); err != nil {
		panic(err)
	}
}
