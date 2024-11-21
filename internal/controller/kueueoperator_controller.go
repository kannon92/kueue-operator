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

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"
	cachev1 "github.com/kannon92/kueue-operator/api/v1"
	"github.com/kannon92/kueue-operator/internal/install"
)

// KueueOperatorReconciler reconciles a KueueOperator object
type KueueOperatorReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cache.kannon92,resources=kueueoperators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cache.kannon92,resources=kueueoperators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cache.kannon92,resources=kueueoperators/finalizers,verbs=update
//+kubebuilder:rbac:groups=cache.kannon92,resources=deployments,verbs=get;list;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=secrets,verbs=create
//+kubebuilder:rbac:groups=core,resources=deployments,verbs=get;list;create;update;patch;delete

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
	_ = log.FromContext(ctx)

	log := r.Log.WithValues("KueueOperator", req.NamespacedName)
	// Fetch the Memcached instance
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

	// Check for JobSet
	if kueueOperator.Spec.JobSet != nil {
		deployment, err := install.BuildJobSetDeployment(kueueOperator.Spec.JobSet)
		if err != nil {
			log.Error(err, "Failed to build jobset deployment")
			return ctrl.Result{}, err
		}
		log.Info("Creating a jobset deployment")
		if err := r.Create(ctx, deployment, &client.CreateOptions{}); err != nil {
			log.Error(err, "Failed to create deployment")
		}

	}

	// Check for Kueue
	if kueueOperator.Spec.Kueue != nil {

	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KueueOperatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1.KueueOperator{}).
		Complete(r)
}
