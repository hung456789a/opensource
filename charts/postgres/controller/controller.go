package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
		queue: workqueue.NewRateLimitingQueue(
			workqueue.DefaultControllerRateLimiter(),
		),
	}
}

func (c *Controller) enqueue(obj interface{}) {

	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		fmt.Printf("failed to get key: %v\n", err)
		return
	}

	fmt.Printf("Event received: %s\n", key)

	c.queue.Add(key)
}

func (c *Controller) run(ctx context.Context) {

	for {
		item, shutdown := c.queue.Get()

		if shutdown {
			return
		}

		func() {
			defer c.queue.Done(item)

			key := item.(string)

			if err := c.reconcile(ctx, key); err != nil {
				fmt.Printf("reconcile failed: %v\n", err)

				c.queue.AddRateLimited(key)
				return
			}

			c.queue.Forget(item)
		}()
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

	fmt.Printf(
		"Reconciling PostgresCluster %s/%s\n",
		namespace,
		name,
	)

	cluster, err := c.dynamicClient.
		Resource(postgresClusterGVR).
		Namespace(namespace).
		Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		return err
	}

	fmt.Printf(
		"PostgresCluster found: %s/%s\n",
		namespace,
		name,
	)

	instances, found, err := unstructured.NestedInt64(
		cluster.Object,
		"spec",
		"instances",
	)

	if err != nil {
		return err
	}

	if !found {
		instances = 1
	}

	fmt.Printf(
		"Desired instances: %d\n",
		instances,
	)

	return c.ensurePod(ctx, cluster)
}

func (c *Controller) ensurePod(
	ctx context.Context,
	cluster *unstructured.Unstructured,
) error {

	namespace := cluster.GetNamespace()
	name := cluster.GetName()

	podName := name + "-0"

	_, err := c.kubeClient.
		CoreV1().
		Pods(namespace).
		Get(ctx, podName, metav1.GetOptions{})

	if err == nil {
		fmt.Printf(
			"Pod %s already exists\n",
			podName,
		)

		return nil
	}

	pod := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Pod",

			"metadata": map[string]interface{}{
				"name":      podName,
				"namespace": namespace,

				"labels": map[string]interface{}{
					"app": name,
				},
			},

			"spec": map[string]interface{}{
				"containers": []interface{}{
					map[string]interface{}{
						"name":  "postgres",
						"image": "nginx:latest",
					},
				},
			},
		},
	}

	fmt.Printf(
		"Creating Pod %s/%s\n",
		namespace,
		podName,
	)

	_, err = c.dynamicClient.
		Resource(schema.GroupVersionResource{
			Version:  "v1",
			Resource: "pods",
		}).
		Namespace(namespace).
		Create(
			ctx,
			pod,
			metav1.CreateOptions{},
		)

	return err
}
