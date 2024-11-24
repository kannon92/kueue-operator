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
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logr "sigs.k8s.io/controller-runtime/pkg/log"

	cachev1 "github.com/kannon92/kueue-operator/api/v1"
	"github.com/kannon92/kueue-operator/internal/install"
)

// KueueOperatorReconciler reconciles a KueueOperator object
type KueueOperatorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	hack   bool
}

//+kubebuilder:rbac:groups=cache.kannon92,resources=kueueoperators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cache.kannon92,resources=kueueoperators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cache.kannon92,resources=kueueoperators/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=secrets,verbs=create;delete
//+kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=create;delete
//+kubebuilder:rbac:groups="",resources=events,verbs=create;watch;update;patch
//+kubebuilder:rbac:groups=jobset.x-k8s.io,resources=jobsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=jobset.x-k8s.io,resources=jobsets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=jobset.x-k8s.io,resources=jobsets/finalizers,verbs=update
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch,resources=jobs/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=events,verbs=create;watch;update;patch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KueueOperator object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *KueueOperatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logr.FromContext(ctx)
	// Fetch the KueueOperator instance
	kueueOperator := &cachev1.KueueOperator{}
	err := r.Get(ctx, req.NamespacedName, kueueOperator)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("Kueue Operator resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get Kueue Operator")
		return ctrl.Result{}, err
	}

	if r.hack {
		log.Info("TODO: no difference detected so skipping installation")
		return ctrl.Result{}, nil
	}
	if kueueOperator.Status.JobSetVersion == "" {
		r.hack = true
		// Check for JobSet
		if kueueOperator.Spec.JobSet != nil {
			log.Info("Installing JobSetOperator")
			err := r.installJobSet(ctx, kueueOperator.Spec.JobSet, req.Namespace)
			if err != nil {
				log.Error(err, "error detected when installing jobset")
				return ctrl.Result{}, err
			} else {
				kueueOperator.Status.JobSetVersion = kueueOperator.Spec.JobSet.JobSetImage
				if err := r.Status().Update(ctx, kueueOperator, &client.SubResourceUpdateOptions{}); err != nil {
					return ctrl.Result{}, err
				}
			}
		}
	}

	// Check for Kueue
	if kueueOperator.Spec.Kueue != nil {

	}
	return ctrl.Result{}, nil
}

func (r *KueueOperatorReconciler) installJobSet(ctx context.Context, jobSetSpec *cachev1.JobSetSpec, namespace string) error {
	log := logr.FromContext(ctx)

	// read secret and create on cluster
	secrets, err := install.ReadSecret("assets/jobset/secret.yaml", namespace)
	if err != nil {
		return fmt.Errorf("failed to read secrets: %v", err)
	}
	if err := r.Create(ctx, secrets, &client.CreateOptions{}); err != nil {
		return fmt.Errorf("failed to create secrets %w", err)
	}
	log.Info("created jobset secret")

	// read service metrics and create on cluster
	serviceMetrics, err := install.ReadService("assets/jobset/service_metrics.yaml", namespace)
	if err != nil {
		return fmt.Errorf("failed to read service: %v", err)
	}
	if err := r.Create(ctx, serviceMetrics, &client.CreateOptions{}); err != nil {
		return fmt.Errorf("failed to create service %w", err)
	}
	log.Info("created jobset service metrics")

	// read service metrics and create on cluster
	serviceWebhook, err := install.ReadService("assets/jobset/service_webhook.yaml", namespace)
	if err != nil {
		return fmt.Errorf("failed to read service webhook: %v", err)
	}
	if err := r.Create(ctx, serviceWebhook, &client.CreateOptions{}); err != nil {
		return fmt.Errorf("failed to create service webhook %w", err)
	}
	log.Info("created jobset service webhook")

	// read deploy and create on cluster
	deployment, err := install.ReadJobSetDeployment(jobSetSpec, "assets/jobset/deployment.yaml", namespace)
	if err != nil {
		return fmt.Errorf("failed to read jobset deployment: %v", err)
	}
	if err := r.Create(ctx, deployment, &client.CreateOptions{}); err != nil {
		return fmt.Errorf("failed to create deployment: %w", err)
	}
	log.Info("all jobset components are installed")

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KueueOperatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1.KueueOperator{}).
		Complete(r)
}
