// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/authentication/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Action) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionIORef) DeepCopyInto(out *ActionIORef) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionIORef.
func (in *ActionIORef) DeepCopy() *ActionIORef {
	if in == nil {
		return nil
	}
	out := new(ActionIORef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionList) DeepCopyInto(out *ActionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Action, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionList.
func (in *ActionList) DeepCopy() *ActionList {
	if in == nil {
		return nil
	}
	out := new(ActionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionSpec) DeepCopyInto(out *ActionSpec) {
	*out = *in
	if in.InputRef != nil {
		in, out := &in.InputRef, &out.InputRef
		*out = new(ActionIORef)
		**out = **in
	}
	if in.AdvancedRendering != nil {
		in, out := &in.AdvancedRendering, &out.AdvancedRendering
		*out = new(AdvancedRendering)
		(*in).DeepCopyInto(*out)
	}
	if in.RenderedActionOverride != nil {
		in, out := &in.RenderedActionOverride, &out.RenderedActionOverride
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Run != nil {
		in, out := &in.Run, &out.Run
		*out = new(bool)
		**out = **in
	}
	if in.Cancel != nil {
		in, out := &in.Cancel, &out.Cancel
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionSpec.
func (in *ActionSpec) DeepCopy() *ActionSpec {
	if in == nil {
		return nil
	}
	out := new(ActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionStatus) DeepCopyInto(out *ActionStatus) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Runner != nil {
		in, out := &in.Runner, &out.Runner
		*out = new(RunnerStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.OutputRef != nil {
		in, out := &in.OutputRef, &out.OutputRef
		*out = new(ActionIORef)
		**out = **in
	}
	if in.RenderedAction != nil {
		in, out := &in.RenderedAction, &out.RenderedAction
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.AdvancedRendering != nil {
		in, out := &in.AdvancedRendering, &out.AdvancedRendering
		*out = new(AdvancedRenderingStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(v1beta1.UserInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.RunBy != nil {
		in, out := &in.RunBy, &out.RunBy
		*out = new(v1beta1.UserInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.CancelledBy != nil {
		in, out := &in.CancelledBy, &out.CancelledBy
		*out = new(v1beta1.UserInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionStatus.
func (in *ActionStatus) DeepCopy() *ActionStatus {
	if in == nil {
		return nil
	}
	out := new(ActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedRendering) DeepCopyInto(out *AdvancedRendering) {
	*out = *in
	if in.RenderingIteration != nil {
		in, out := &in.RenderingIteration, &out.RenderingIteration
		*out = new(RenderingIteration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedRendering.
func (in *AdvancedRendering) DeepCopy() *AdvancedRendering {
	if in == nil {
		return nil
	}
	out := new(AdvancedRendering)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedRenderingStatus) DeepCopyInto(out *AdvancedRenderingStatus) {
	*out = *in
	if in.RenderingIteration != nil {
		in, out := &in.RenderingIteration, &out.RenderingIteration
		*out = new(RenderingIterationStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedRenderingStatus.
func (in *AdvancedRenderingStatus) DeepCopy() *AdvancedRenderingStatus {
	if in == nil {
		return nil
	}
	out := new(AdvancedRenderingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputArtifact) DeepCopyInto(out *InputArtifact) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputArtifact.
func (in *InputArtifact) DeepCopy() *InputArtifact {
	if in == nil {
		return nil
	}
	out := new(InputArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputArtifactToProvide) DeepCopyInto(out *InputArtifactToProvide) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputArtifactToProvide.
func (in *InputArtifactToProvide) DeepCopy() *InputArtifactToProvide {
	if in == nil {
		return nil
	}
	out := new(InputArtifactToProvide)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenderingIteration) DeepCopyInto(out *RenderingIteration) {
	*out = *in
	if in.InputArtifacts != nil {
		in, out := &in.InputArtifacts, &out.InputArtifacts
		*out = new([]InputArtifact)
		if **in != nil {
			in, out := *in, *out
			*out = make([]InputArtifact, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenderingIteration.
func (in *RenderingIteration) DeepCopy() *RenderingIteration {
	if in == nil {
		return nil
	}
	out := new(RenderingIteration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenderingIterationStatus) DeepCopyInto(out *RenderingIterationStatus) {
	*out = *in
	if in.InputArtifactsToProvide != nil {
		in, out := &in.InputArtifactsToProvide, &out.InputArtifactsToProvide
		*out = new([]InputArtifact)
		if **in != nil {
			in, out := *in, *out
			*out = make([]InputArtifact, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenderingIterationStatus.
func (in *RenderingIterationStatus) DeepCopy() *RenderingIterationStatus {
	if in == nil {
		return nil
	}
	out := new(RenderingIterationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerStatus) DeepCopyInto(out *RunnerStatus) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerStatus.
func (in *RunnerStatus) DeepCopy() *RunnerStatus {
	if in == nil {
		return nil
	}
	out := new(RunnerStatus)
	in.DeepCopyInto(out)
	return out
}