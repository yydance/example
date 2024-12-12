package controller

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/cache"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func WatchService(ctx context.Context, obj client.ObjectList, opts ...client.ListOption) error {
	config := ctrl.GetConfigOrDie()
	withWatch, err := client.NewWithWatch(config, client.Options{})
	if err != nil {
		return err
	}
	watcher, err := withWatch.Watch(ctx, obj, opts...)
	if err != nil {
		ctrlLog.Error(err, "unable to watch services")
	}
	for ev := range watcher.ResultChan() {
		switch v := ev.Object.(type) {
		case *corev1.Service:
			ctrlLog.Info("Watch service event", "type", ev.Type, "object", v.GetName())
		case *appsv1.Deployment:
			ctrlLog.Info("Watch service event", "type", ev.Type, "object", v.GetName())
		case *metav1.Status:
			ctrlLog.Info("Watch service event", "type", ev.Type, "object", v.Reason)
		}
	}

	return nil
}

func InformerGetService(ctx context.Context) {

	config := ctrl.GetConfigOrDie()
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	gvr := schema.GroupVersionResource{
		Group:    "k8sservice.example.cn",
		Version:  "v1alpha1",
		Resource: "getservices",
	}
	restClient := dynamicClient.Resource(gvr)
	lw := &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return restClient.List(ctx, options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return restClient.Watch(ctx, options)
		},
	}
	// cache.NewListWatchFromClient()只能监听结构化type
	//lw := cache.NewListWatchFromClient(restClient.Namespace(""), "GetServices", metav1.NamespaceAll, fields.Everything())
	//informer := cache.NewSharedIndexInformer(lw, &k8sservicev1alpha1.GetService{}, 0, cache.Indexers{})
	informer := cache.NewSharedInformer(lw, &unstructured.Unstructured{}, 0)

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ctrlLog.Info("add GetService", "object", obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			ctrlLog.Info("update GetService", "oldObject", oldObj, "newObject", newObj)
		},
		DeleteFunc: func(obj interface{}) {
			ctrlLog.Info("delete GetService", "object", obj)
		},
	})
	stopCh := make(chan struct{})
	go informer.Run(stopCh)
	go func() {
		<-ctx.Done()
		close(stopCh)
		ctrlLog.Info("Stopping watcher due to context cancellation")
	}()
	<-ctx.Done()
}
