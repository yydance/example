/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	k8sservicev1alpha1 "example.cn/api/v1alpha1"
	"example.cn/internal/pkg"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// GetServiceReconciler reconciles a GetService object
type GetServiceReconciler struct {
	client.Client
	Scheme  *runtime.Scheme
	Recoder record.EventRecorder
}

// 注意：rbac定义只能加在这里生效
// +kubebuilder:rbac:groups=k8sservice.example.cn,resources=getservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=k8sservice.example.cn,resources=getservices/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=k8sservice.example.cn,resources=getservices/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the GetService object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.1/pkg/reconcile

var ctrlLog = log.Log.WithName("getservice controller")

func (r *GetServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	// TODO(user): your logic here
	var getSvc k8sservicev1alpha1.GetService
	err := r.Get(ctx, req.NamespacedName, &getSvc)

	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			ctrlLog.Info("GetService resource not found. Ignoring since object must be deleted", "name", req.NamespacedName)
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}
	outputSvc, err := r.outputServiceList(ctx, getSvc.Spec.Namespace)
	if err != nil {
		return ctrl.Result{}, err
	}
	ctrlLog.Info("output service list", "outputSvc", outputSvc)
	/*
		if err := r.watchService(ctx, getSvc.Namespace); err != nil {
			return ctrl.Result{}, err
		}
	*/
	// TODO: 创建资源后，需要启动一个server deployment，该server服务对外提供service查询服务

	if getSvc.Status.Status == "" {
		getSvc.Status.Status = metav1.StatusSuccess
		getSvc.Status.Complated = true
		if err := r.Status().Update(ctx, &getSvc); err != nil {
			return ctrl.Result{}, err
		}
	}

	if !getSvc.DeletionTimestamp.IsZero() {
		r.Recoder.Eventf(&getSvc, corev1.EventTypeNormal, "Deleting", "Custom Resource %s is being deleted from the namespace %s", req.Name, req.NamespacedName)
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GetServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {

	return ctrl.NewControllerManagedBy(mgr).
		For(&k8sservicev1alpha1.GetService{}).
		Named("getservice").
		Complete(r)
}

// outputServiceList 获取所有service列表，service格式: "namespace/serviceName:portName"
func (r *GetServiceReconciler) outputServiceList(ctx context.Context, namespace string) ([]string, error) {
	var svcList corev1.ServiceList
	if err := r.List(ctx, &svcList, &client.ListOptions{Namespace: namespace}); err != nil {
		ctrlLog.Error(err, "unable to list services")
		return nil, err
	}
	if len(svcList.Items) == 0 {
		ctrlLog.Info("no native services found")
		return nil, nil
	}
	ctrlLog.Info("found native services", "services", len(svcList.Items))
	var outputSvcList []string
	for i := range svcList.Items {
		var outputSvc k8sservicev1alpha1.OutputServiceSpec
		outputSvc.Name = svcList.Items[i].Name
		outputSvc.Namespace = svcList.Items[i].Namespace
		for j := range svcList.Items[i].Spec.Ports {
			if svcList.Items[i].Spec.Ports[j].Name != "" {
				outputSvc.PortName = svcList.Items[i].Spec.Ports[j].Name
				outputSvcList = append(outputSvcList, pkg.ApisixK8sSvcName(
					outputSvc.Namespace,
					outputSvc.Name,
					outputSvc.PortName,
				))
				break
			}
			if svcList.Items[i].Spec.Ports[j].TargetPort.String() != "" {
				outputSvc.TargetPort = svcList.Items[i].Spec.Ports[j].TargetPort.IntVal
				outputSvcList = append(outputSvcList, pkg.ApisixK8sSvcName(
					outputSvc.Namespace,
					outputSvc.Name,
					outputSvc.TargetPort,
				))
				break
			}
			outputSvc.Port = svcList.Items[i].Spec.Ports[j].Port
			outputSvcList = append(outputSvcList, pkg.ApisixK8sSvcName(
				outputSvc.Namespace,
				outputSvc.Name,
				outputSvc.Port,
			))
		}
	}

	return outputSvcList, nil
}
