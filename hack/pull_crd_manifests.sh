#!/usr/bin/env bash

set -eou pipefail

KUEUE_MANIFEST_LIST=${KUEUE_MANIFEST_LIST-"https://github.com/kubernetes-sigs/kueue/releases/download/v0.9.1/manifests.yaml"}
JOBSET_MANIFEST_LIST=${JOBSET_MANIFEST_LIST-"https://github.com/kubernetes-sigs/jobset/releases/download/v0.7.0/manifests.yaml"}

wget $KUEUE_MANIFEST_LIST -O kueue_manifest.yaml
wget $JOBSET_MANIFEST_LIST -O jobset_manifest.yaml

podman run --rm -v "${PWD}":/workdir:z mikefarah/yq 'select(.kind == "CustomResourceDefinition")' /workdir/jobset_manifest.yaml > config/crd/bases/jobset.yaml
podman run --rm -v "${PWD}":/workdir:z mikefarah/yq 'select(.kind == "CustomResourceDefinition")' /workdir/kueue_manifest.yaml > config/crd/bases/kueue.yaml

rm kueue_manifest.yaml
rm jobset_manifest.yaml
