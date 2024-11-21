//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSetSpec) DeepCopyInto(out *JobSetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSetSpec.
func (in *JobSetSpec) DeepCopy() *JobSetSpec {
	if in == nil {
		return nil
	}
	out := new(JobSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kueue) DeepCopyInto(out *Kueue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kueue.
func (in *Kueue) DeepCopy() *Kueue {
	if in == nil {
		return nil
	}
	out := new(Kueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KueueOperator) DeepCopyInto(out *KueueOperator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KueueOperator.
func (in *KueueOperator) DeepCopy() *KueueOperator {
	if in == nil {
		return nil
	}
	out := new(KueueOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KueueOperator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KueueOperatorList) DeepCopyInto(out *KueueOperatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KueueOperator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KueueOperatorList.
func (in *KueueOperatorList) DeepCopy() *KueueOperatorList {
	if in == nil {
		return nil
	}
	out := new(KueueOperatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KueueOperatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KueueOperatorSpec) DeepCopyInto(out *KueueOperatorSpec) {
	*out = *in
	if in.JobSet != nil {
		in, out := &in.JobSet, &out.JobSet
		*out = new(JobSetSpec)
		**out = **in
	}
	if in.LeaderWorkerSet != nil {
		in, out := &in.LeaderWorkerSet, &out.LeaderWorkerSet
		*out = new(LeaderWorkerSet)
		**out = **in
	}
	if in.Kueue != nil {
		in, out := &in.Kueue, &out.Kueue
		*out = new(Kueue)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KueueOperatorSpec.
func (in *KueueOperatorSpec) DeepCopy() *KueueOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(KueueOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KueueOperatorStatus) DeepCopyInto(out *KueueOperatorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KueueOperatorStatus.
func (in *KueueOperatorStatus) DeepCopy() *KueueOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(KueueOperatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderWorkerSet) DeepCopyInto(out *LeaderWorkerSet) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderWorkerSet.
func (in *LeaderWorkerSet) DeepCopy() *LeaderWorkerSet {
	if in == nil {
		return nil
	}
	out := new(LeaderWorkerSet)
	in.DeepCopyInto(out)
	return out
}