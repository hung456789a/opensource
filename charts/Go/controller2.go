package main

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

var postgresClusterGVR = schema.GroupVersionResource{
	Group:    "database.example.com",
	Version:  "v1",
	Resource: "postgresclusters",
}

type Controller struct {
	kubeClient    kubernetes.Interface
	dynamicClient dynamic.Interface
}
