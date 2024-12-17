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

package v1alpha1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	k8sservicev1alpha1 "example.cn/api/v1alpha1"
)

// nolint:unused
// log is for logging in this package.
var getservicelog = logf.Log.WithName("getservice-resource")

// SetupGetServiceWebhookWithManager registers the webhook for GetService in the manager.
func SetupGetServiceWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&k8sservicev1alpha1.GetService{}).
		WithValidator(&GetServiceCustomValidator{}).
		WithDefaulter(&GetServiceCustomDefaulter{
			Namespace: "default",
		}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-k8sservice-example-cn-v1alpha1-getservice,mutating=true,failurePolicy=fail,sideEffects=None,groups=k8sservice.example.cn,resources=getservices,verbs=create;update,versions=v1alpha1,name=mgetservice-v1alpha1.kb.io,admissionReviewVersions=v1

// GetServiceCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind GetService when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type GetServiceCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
	Namespace string
}

var _ webhook.CustomDefaulter = &GetServiceCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind GetService.
func (d *GetServiceCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	getservice, ok := obj.(*k8sservicev1alpha1.GetService)

	if !ok {
		return fmt.Errorf("expected an GetService object but got %T", obj)
	}
	getservicelog.Info("Defaulting for GetService", "name", getservice.GetName())

	if getservice.Spec.Namespace == "" {
		getservice.Spec.Namespace = d.Namespace
	}
	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-k8sservice-example-cn-v1alpha1-getservice,mutating=false,failurePolicy=fail,sideEffects=None,groups=k8sservice.example.cn,resources=getservices,verbs=create;update,versions=v1alpha1,name=vgetservice-v1alpha1.kb.io,admissionReviewVersions=v1

// GetServiceCustomValidator struct is responsible for validating the GetService resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type GetServiceCustomValidator struct {
	//TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &GetServiceCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type GetService.
func (v *GetServiceCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	getservice, ok := obj.(*k8sservicev1alpha1.GetService)
	if !ok {
		return nil, fmt.Errorf("expected a GetService object but got %T", obj)
	}
	getservicelog.Info("Validation for GetService upon creation", "name", getservice.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type GetService.
func (v *GetServiceCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	getservice, ok := newObj.(*k8sservicev1alpha1.GetService)
	if !ok {
		return nil, fmt.Errorf("expected a GetService object for the newObj but got %T", newObj)
	}
	getservicelog.Info("Validation for GetService upon update", "name", getservice.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type GetService.
func (v *GetServiceCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	getservice, ok := obj.(*k8sservicev1alpha1.GetService)
	if !ok {
		return nil, fmt.Errorf("expected a GetService object but got %T", obj)
	}
	getservicelog.Info("Validation for GetService upon deletion", "name", getservice.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
