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

package configmap

import (
	corev1 "k8s.io/api/core/v1"

	configapi "sigs.k8s.io/kueue/apis/config/v1beta1"

	cachev1 "github.com/kannon92/kueue-operator/api/v1"
)

func BuildConfigMap(cfgMap *configapi.Configuration) *corev1.ConfigMap {
	if cfgMap == nil {
		return nil
	}
	return nil
}

func IsConfigMapDifferentFromConfiguration(kueueCfgMap *corev1.ConfigMap, kueueCfg *cachev1.Kueue) bool {
	return true
}
