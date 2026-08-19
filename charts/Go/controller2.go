package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
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
		Namespace(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("Cluster: %v\n", cluster)
	if cluster != nil {
		fmt.Printf("Cluster exists\n", cluster, namespace, name)
	}
	return nil
}

func main() {

}
