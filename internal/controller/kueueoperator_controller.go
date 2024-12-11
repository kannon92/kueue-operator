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

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logr "sigs.k8s.io/controller-runtime/pkg/log"

	cachev1 "github.com/kannon92/kueue-operator/api/v1"
	"github.com/kannon92/kueue-operator/internal/configmap"
)

// KueueOperatorReconciler reconciles a KueueOperator object
type KueueOperatorReconciler struct {
	client.Client
	Scheme         *runtime.Scheme
	kueueNamespace string
}

//+kubebuilder:rbac:groups=cache.kannon92,resources=kueueoperators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cache.kannon92,resources=kueueoperators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cache.kannon92,resources=kueueoperators/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;create;update;patch;delete

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

	if kueueOperator.Spec.Kueue == nil {
		return ctrl.Result{}, nil
	}

	// Reconcile differences of config maps
	configMap, cfgMapErr := r.fetchConfigMap(ctx)
	if cfgMapErr != nil {
		return ctrl.Result{}, cfgMapErr
	}
	if !configmap.IsConfigMapDifferentFromConfiguration(configMap, kueueOperator.Spec.Kueue) {
		log.V(4).Info("Kueue config maps are the same. skipping")
		return ctrl.Result{}, nil
	}
	// scale down kueue deployment
	if downScaleErr := r.scaleDeployment(ctx, 0); downScaleErr != nil {
		return ctrl.Result{}, downScaleErr
	}

	// set new config map
	newCfgMap := configmap.BuildConfigMap(&kueueOperator.Spec.Kueue.Config)
	if cfgMapErr := r.updateConfigMap(ctx, newCfgMap); cfgMapErr != nil {
		return ctrl.Result{}, cfgMapErr
	}
	// scale up kueue deployment
	if upScaleErr := r.scaleDeployment(ctx, 1); upScaleErr != nil {
		return ctrl.Result{}, upScaleErr
	}

	return ctrl.Result{}, nil
}

func (r *KueueOperatorReconciler) fetchConfigMap(ctx context.Context) (*corev1.ConfigMap, error) {
	log := logr.FromContext(ctx)
	configMap := &corev1.ConfigMap{}
	configMapErr := r.Get(ctx, types.NamespacedName{Name: "kueue-manager-config", Namespace: r.kueueNamespace}, configMap)
	if configMapErr != nil {
		if errors.IsNotFound(configMapErr) {
			log.Info("Kueue Config map not found. Ignoring")
			return nil, nil
		}
		log.Error(configMapErr, "Failed to get config map from kueue")
		return nil, configMapErr
	}
	return configMap, nil
}

func (r *KueueOperatorReconciler) updateConfigMap(ctx context.Context, newCfgMap *corev1.ConfigMap) error {
	log := logr.FromContext(ctx)
	configMap := &corev1.ConfigMap{}
	configMapErr := r.Get(ctx, types.NamespacedName{Name: "kueue-manager-config", Namespace: r.kueueNamespace}, configMap)
	if configMapErr != nil {
		if errors.IsNotFound(configMapErr) {
			log.Info("Kueue Config map not found. Ignoring")
			return nil
		}
		log.Error(configMapErr, "Failed to get config map from kueue")
		return configMapErr
	}
	configMap = newCfgMap
	return r.Update(ctx, configMap, &client.UpdateOptions{})
}

func (r *KueueOperatorReconciler) scaleDeployment(ctx context.Context, replicas int32) error {
	log := logr.FromContext(ctx)
	deployment := &appsv1.Deployment{}
	deploymentErr := r.Get(ctx, types.NamespacedName{Name: "kueue-controller-manager", Namespace: r.kueueNamespace}, deployment)
	if deploymentErr != nil {
		if errors.IsNotFound(deploymentErr) {
			log.Info("Kueue Deployment not found. Ignoring")
			return nil
		}
		log.Error(deploymentErr, "Failed to get deployment from kueue namespace")
		return deploymentErr
	}
	deployment.Spec.Replicas = ptr.To(replicas)
	if uErr := r.Update(ctx, deployment, &client.UpdateOptions{}); uErr != nil {
		log.Error(uErr, "Unable to update deployment")
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KueueOperatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1.KueueOperator{}).
		Complete(r)
}
