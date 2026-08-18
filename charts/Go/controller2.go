package main

import (
	"flag"
	"fmt"
	"path/filepath"

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

func (c *Controller) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		fmt.Printf("Failed to get key: %v\n", err)
		return
	}

	fmt.Printf("Key: %s\n", key)
	c.queue.Add(key)
}

func main() {
	// Lấy đường dẫn tới file kubeconfig (thường ở ~/.kube/config)
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Tạo cấu hình kết nối tới Kubernetes cluster
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(fmt.Sprintf("Error building kubeconfig: %v", err))
	}

	// Tạo Kubernetes clientset
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("Error building kubernetes clientset: %v", err))
	}

	// Tạo Dynamic client
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("Error building dynamic client: %v", err))
	}

	// Khởi tạo Controller của bạn
	controller := NewController(kubeClient, dynamicClient)

	// Ví dụ: Giả lập gọi thử nghiệm hàm enqueue để đẩy một đối tượng vào hàng đợi queue
	// (Trong thực tế, bạn sẽ dùng hàm này trong Informer/EventHandler)
	fmt.Println("Đang khởi động Controller và test queue...")

	// Bạn có thể giả lập một đối tượng bứt phá hoặc kiểm tra hàng đợi:
	// controller.queue.Add("default/my-postgres-cluster")

	// Vòng đời cơ bản của Worker lấy dữ liệu từ Queue
	stopCh := make(chan struct{})
	defer close(stopCh)

	go func() {
		for {
			// Lấy key ra khỏi hàng đợi
			item, shutdown := controller.queue.Get()
			if shutdown {
				return
			}

			// Xử lý key (ở đây chỉ in ra để kiểm tra)
			fmt.Printf("Đang xử lý item từ queue: %v\n", item)

			// Đánh dấu đã xử lý xong item này
			controller.queue.Done(item)
		}
	}()

	// Giữ chương trình chạy (nhấn Ctrl+C để thoát)
	select {}
}
