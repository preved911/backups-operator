//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DgraphBackup) DeepCopyInto(out *DgraphBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DgraphBackup.
func (in *DgraphBackup) DeepCopy() *DgraphBackup {
	if in == nil {
		return nil
	}
	out := new(DgraphBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DgraphBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DgraphBackupList) DeepCopyInto(out *DgraphBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DgraphBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DgraphBackupList.
func (in *DgraphBackupList) DeepCopy() *DgraphBackupList {
	if in == nil {
		return nil
	}
	out := new(DgraphBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DgraphBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DgraphBackupSchedule) DeepCopyInto(out *DgraphBackupSchedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DgraphBackupSchedule.
func (in *DgraphBackupSchedule) DeepCopy() *DgraphBackupSchedule {
	if in == nil {
		return nil
	}
	out := new(DgraphBackupSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DgraphBackupSchedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DgraphBackupScheduleList) DeepCopyInto(out *DgraphBackupScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DgraphBackupSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DgraphBackupScheduleList.
func (in *DgraphBackupScheduleList) DeepCopy() *DgraphBackupScheduleList {
	if in == nil {
		return nil
	}
	out := new(DgraphBackupScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DgraphBackupScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DgraphBackupScheduleSpec) DeepCopyInto(out *DgraphBackupScheduleSpec) {
	*out = *in
	in.Backup.DeepCopyInto(&out.Backup)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DgraphBackupScheduleSpec.
func (in *DgraphBackupScheduleSpec) DeepCopy() *DgraphBackupScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(DgraphBackupScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DgraphBackupScheduleStatus) DeepCopyInto(out *DgraphBackupScheduleStatus) {
	*out = *in
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DgraphBackupScheduleStatus.
func (in *DgraphBackupScheduleStatus) DeepCopy() *DgraphBackupScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(DgraphBackupScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DgraphBackupSpec) DeepCopyInto(out *DgraphBackupSpec) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DgraphBackupSpec.
func (in *DgraphBackupSpec) DeepCopy() *DgraphBackupSpec {
	if in == nil {
		return nil
	}
	out := new(DgraphBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DgraphBackupStatus) DeepCopyInto(out *DgraphBackupStatus) {
	*out = *in
	in.ExportResponse.DeepCopyInto(&out.ExportResponse)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DgraphBackupStatus.
func (in *DgraphBackupStatus) DeepCopy() *DgraphBackupStatus {
	if in == nil {
		return nil
	}
	out := new(DgraphBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DgraphBackupStatusExportResonse) DeepCopyInto(out *DgraphBackupStatusExportResonse) {
	*out = *in
	if in.ExportedFiles != nil {
		in, out := &in.ExportedFiles, &out.ExportedFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DgraphBackupStatusExportResonse.
func (in *DgraphBackupStatusExportResonse) DeepCopy() *DgraphBackupStatusExportResonse {
	if in == nil {
		return nil
	}
	out := new(DgraphBackupStatusExportResonse)
	in.DeepCopyInto(out)
	return out
}
