// +build !ignore_autogenerated

/*
Copyright 2020 Rancher Labs.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	genericcondition "github.com/rancher/wrangler/pkg/genericcondition"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscaleConfig) DeepCopyInto(out *AutoscaleConfig) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscaleConfig.
func (in *AutoscaleConfig) DeepCopy() *AutoscaleConfig {
	if in == nil {
		return nil
	}
	out := new(AutoscaleConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildRevision) DeepCopyInto(out *BuildRevision) {
	*out = *in
	if in.Commits != nil {
		in, out := &in.Commits, &out.Commits
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildRevision.
func (in *BuildRevision) DeepCopy() *BuildRevision {
	if in == nil {
		return nil
	}
	out := new(BuildRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerPort) DeepCopyInto(out *ContainerPort) {
	*out = *in
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerPort.
func (in *ContainerPort) DeepCopy() *ContainerPort {
	if in == nil {
		return nil
	}
	out := new(ContainerPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentWrangler) DeepCopyInto(out *DeploymentWrangler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentWrangler.
func (in *DeploymentWrangler) DeepCopy() *DeploymentWrangler {
	if in == nil {
		return nil
	}
	out := new(DeploymentWrangler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentWrangler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentWranglerList) DeepCopyInto(out *DeploymentWranglerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeploymentWrangler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentWranglerList.
func (in *DeploymentWranglerList) DeepCopy() *DeploymentWranglerList {
	if in == nil {
		return nil
	}
	out := new(DeploymentWranglerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentWranglerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentWranglerSpec) DeepCopyInto(out *DeploymentWranglerSpec) {
	*out = *in
	in.WranglerSpec.DeepCopyInto(&out.WranglerSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentWranglerSpec.
func (in *DeploymentWranglerSpec) DeepCopy() *DeploymentWranglerSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentWranglerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentWranglerStatus) DeepCopyInto(out *DeploymentWranglerStatus) {
	*out = *in
	in.WranglerStatus.DeepCopyInto(&out.WranglerStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentWranglerStatus.
func (in *DeploymentWranglerStatus) DeepCopy() *DeploymentWranglerStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentWranglerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalService) DeepCopyInto(out *ExternalService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalService.
func (in *ExternalService) DeepCopy() *ExternalService {
	if in == nil {
		return nil
	}
	out := new(ExternalService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalServiceList) DeepCopyInto(out *ExternalServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalServiceList.
func (in *ExternalServiceList) DeepCopy() *ExternalServiceList {
	if in == nil {
		return nil
	}
	out := new(ExternalServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalServiceSpec) DeepCopyInto(out *ExternalServiceSpec) {
	*out = *in
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalServiceSpec.
func (in *ExternalServiceSpec) DeepCopy() *ExternalServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalServiceStatus) DeepCopyInto(out *ExternalServiceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]genericcondition.GenericCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalServiceStatus.
func (in *ExternalServiceStatus) DeepCopy() *ExternalServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fault) DeepCopyInto(out *Fault) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fault.
func (in *Fault) DeepCopy() *Fault {
	if in == nil {
		return nil
	}
	out := new(Fault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderMatch) DeepCopyInto(out *HeaderMatch) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(StringMatch)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderMatch.
func (in *HeaderMatch) DeepCopy() *HeaderMatch {
	if in == nil {
		return nil
	}
	out := new(HeaderMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderOperations) DeepCopyInto(out *HeaderOperations) {
	*out = *in
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = make([]NameValue, len(*in))
		copy(*out, *in)
	}
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = make([]NameValue, len(*in))
		copy(*out, *in)
	}
	if in.Remove != nil {
		in, out := &in.Remove, &out.Remove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderOperations.
func (in *HeaderOperations) DeepCopy() *HeaderOperations {
	if in == nil {
		return nil
	}
	out := new(HeaderOperations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuildSpec) DeepCopyInto(out *ImageBuildSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuildSpec.
func (in *ImageBuildSpec) DeepCopy() *ImageBuildSpec {
	if in == nil {
		return nil
	}
	out := new(ImageBuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Match) DeepCopyInto(out *Match) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(StringMatch)
		**out = **in
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Match.
func (in *Match) DeepCopy() *Match {
	if in == nil {
		return nil
	}
	out := new(Match)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameValue) DeepCopyInto(out *NameValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameValue.
func (in *NameValue) DeepCopy() *NameValue {
	if in == nil {
		return nil
	}
	out := new(NameValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Permission) DeepCopyInto(out *Permission) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Permission.
func (in *Permission) DeepCopy() *Permission {
	if in == nil {
		return nil
	}
	out := new(Permission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Question) DeepCopyInto(out *Question) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subquestions != nil {
		in, out := &in.Subquestions, &out.Subquestions
		*out = make([]SubQuestion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Question.
func (in *Question) DeepCopy() *Question {
	if in == nil {
		return nil
	}
	out := new(Question)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redirect) DeepCopyInto(out *Redirect) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redirect.
func (in *Redirect) DeepCopy() *Redirect {
	if in == nil {
		return nil
	}
	out := new(Redirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Retry) DeepCopyInto(out *Retry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Retry.
func (in *Retry) DeepCopy() *Retry {
	if in == nil {
		return nil
	}
	out := new(Retry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rewrite) DeepCopyInto(out *Rewrite) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rewrite.
func (in *Rewrite) DeepCopy() *Rewrite {
	if in == nil {
		return nil
	}
	out := new(Rewrite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutConfig) DeepCopyInto(out *RolloutConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutConfig.
func (in *RolloutConfig) DeepCopy() *RolloutConfig {
	if in == nil {
		return nil
	}
	out := new(RolloutConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
	in.Match.DeepCopyInto(&out.Match)
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]WeightedDestination, len(*in))
		copy(*out, *in)
	}
	if in.Redirect != nil {
		in, out := &in.Redirect, &out.Redirect
		*out = new(Redirect)
		**out = **in
	}
	if in.Rewrite != nil {
		in, out := &in.Rewrite, &out.Rewrite
		*out = new(Rewrite)
		**out = **in
	}
	if in.Retry != nil {
		in, out := &in.Retry, &out.Retry
		*out = new(Retry)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = new(HeaderOperations)
		(*in).DeepCopyInto(*out)
	}
	if in.Fault != nil {
		in, out := &in.Fault, &out.Fault
		*out = new(Fault)
		**out = **in
	}
	if in.Mirror != nil {
		in, out := &in.Mirror, &out.Mirror
		*out = new(Destination)
		**out = **in
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Router) DeepCopyInto(out *Router) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Router.
func (in *Router) DeepCopy() *Router {
	if in == nil {
		return nil
	}
	out := new(Router)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Router) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouterList) DeepCopyInto(out *RouterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Router, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouterList.
func (in *RouterList) DeepCopy() *RouterList {
	if in == nil {
		return nil
	}
	out := new(RouterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouterSpec) DeepCopyInto(out *RouterSpec) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RouteSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouterSpec.
func (in *RouterSpec) DeepCopy() *RouterSpec {
	if in == nil {
		return nil
	}
	out := new(RouterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouterStatus) DeepCopyInto(out *RouterStatus) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]genericcondition.GenericCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouterStatus.
func (in *RouterStatus) DeepCopy() *RouterStatus {
	if in == nil {
		return nil
	}
	out := new(RouterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleStatus) DeepCopyInto(out *ScaleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleStatus.
func (in *ScaleStatus) DeepCopy() *ScaleStatus {
	if in == nil {
		return nil
	}
	out := new(ScaleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stack) DeepCopyInto(out *Stack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stack.
func (in *Stack) DeepCopy() *Stack {
	if in == nil {
		return nil
	}
	out := new(Stack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackBuild) DeepCopyInto(out *StackBuild) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackBuild.
func (in *StackBuild) DeepCopy() *StackBuild {
	if in == nil {
		return nil
	}
	out := new(StackBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackList) DeepCopyInto(out *StackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackList.
func (in *StackList) DeepCopy() *StackList {
	if in == nil {
		return nil
	}
	out := new(StackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSpec) DeepCopyInto(out *StackSpec) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(StackBuild)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]Permission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalGroupVersionKinds != nil {
		in, out := &in.AdditionalGroupVersionKinds, &out.AdditionalGroupVersionKinds
		*out = make([]schema.GroupVersionKind, len(*in))
		copy(*out, *in)
	}
	if in.Answers != nil {
		in, out := &in.Answers, &out.Answers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSpec.
func (in *StackSpec) DeepCopy() *StackSpec {
	if in == nil {
		return nil
	}
	out := new(StackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackStatus) DeepCopyInto(out *StackStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]genericcondition.GenericCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackStatus.
func (in *StackStatus) DeepCopy() *StackStatus {
	if in == nil {
		return nil
	}
	out := new(StackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetWrangler) DeepCopyInto(out *StatefulSetWrangler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetWrangler.
func (in *StatefulSetWrangler) DeepCopy() *StatefulSetWrangler {
	if in == nil {
		return nil
	}
	out := new(StatefulSetWrangler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatefulSetWrangler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetWranglerList) DeepCopyInto(out *StatefulSetWranglerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StatefulSetWrangler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetWranglerList.
func (in *StatefulSetWranglerList) DeepCopy() *StatefulSetWranglerList {
	if in == nil {
		return nil
	}
	out := new(StatefulSetWranglerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatefulSetWranglerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetWranglerSpec) DeepCopyInto(out *StatefulSetWranglerSpec) {
	*out = *in
	in.WranglerSpec.DeepCopyInto(&out.WranglerSpec)
	if in.VolumeTemplates != nil {
		in, out := &in.VolumeTemplates, &out.VolumeTemplates
		*out = make([]VolumeTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetWranglerSpec.
func (in *StatefulSetWranglerSpec) DeepCopy() *StatefulSetWranglerSpec {
	if in == nil {
		return nil
	}
	out := new(StatefulSetWranglerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetWranglerStatus) DeepCopyInto(out *StatefulSetWranglerStatus) {
	*out = *in
	in.WranglerStatus.DeepCopyInto(&out.WranglerStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetWranglerStatus.
func (in *StatefulSetWranglerStatus) DeepCopy() *StatefulSetWranglerStatus {
	if in == nil {
		return nil
	}
	out := new(StatefulSetWranglerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringMatch) DeepCopyInto(out *StringMatch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch.
func (in *StringMatch) DeepCopy() *StringMatch {
	if in == nil {
		return nil
	}
	out := new(StringMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubQuestion) DeepCopyInto(out *SubQuestion) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubQuestion.
func (in *SubQuestion) DeepCopy() *SubQuestion {
	if in == nil {
		return nil
	}
	out := new(SubQuestion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateMeta) DeepCopyInto(out *TemplateMeta) {
	*out = *in
	if in.Questions != nil {
		in, out := &in.Questions, &out.Questions
		*out = make([]Question, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateMeta.
func (in *TemplateMeta) DeepCopy() *TemplateMeta {
	if in == nil {
		return nil
	}
	out := new(TemplateMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeTemplate) DeepCopyInto(out *VolumeTemplate) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]corev1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	if in.VolumeMode != nil {
		in, out := &in.VolumeMode, &out.VolumeMode
		*out = new(corev1.PersistentVolumeMode)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeTemplate.
func (in *VolumeTemplate) DeepCopy() *VolumeTemplate {
	if in == nil {
		return nil
	}
	out := new(VolumeTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeightedDestination) DeepCopyInto(out *WeightedDestination) {
	*out = *in
	out.Destination = in.Destination
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeightedDestination.
func (in *WeightedDestination) DeepCopy() *WeightedDestination {
	if in == nil {
		return nil
	}
	out := new(WeightedDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WranglerSpec) DeepCopyInto(out *WranglerSpec) {
	*out = *in
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]ContainerPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageBuild != nil {
		in, out := &in.ImageBuild, &out.ImageBuild
		*out = new(ImageBuildSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Autoscale != nil {
		in, out := &in.Autoscale, &out.Autoscale
		*out = new(AutoscaleConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RolloutDuration != nil {
		in, out := &in.RolloutDuration, &out.RolloutDuration
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.RolloutConfig != nil {
		in, out := &in.RolloutConfig, &out.RolloutConfig
		*out = new(RolloutConfig)
		**out = **in
	}
	if in.ServiceMesh != nil {
		in, out := &in.ServiceMesh, &out.ServiceMesh
		*out = new(bool)
		**out = **in
	}
	if in.RequestTimeoutSeconds != nil {
		in, out := &in.RequestTimeoutSeconds, &out.RequestTimeoutSeconds
		*out = new(int)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]Permission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GlobalPermissions != nil {
		in, out := &in.GlobalPermissions, &out.GlobalPermissions
		*out = make([]Permission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WranglerSpec.
func (in *WranglerSpec) DeepCopy() *WranglerSpec {
	if in == nil {
		return nil
	}
	out := new(WranglerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WranglerStatus) DeepCopyInto(out *WranglerStatus) {
	*out = *in
	if in.ScaleStatus != nil {
		in, out := &in.ScaleStatus, &out.ScaleStatus
		*out = new(ScaleStatus)
		**out = **in
	}
	if in.ComputedReplicas != nil {
		in, out := &in.ComputedReplicas, &out.ComputedReplicas
		*out = new(int)
		**out = **in
	}
	if in.ComputedWeight != nil {
		in, out := &in.ComputedWeight, &out.ComputedWeight
		*out = new(int)
		**out = **in
	}
	if in.ContainerRevision != nil {
		in, out := &in.ContainerRevision, &out.ContainerRevision
		*out = make(map[string]BuildRevision, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.GeneratedServices != nil {
		in, out := &in.GeneratedServices, &out.GeneratedServices
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.GitCommits != nil {
		in, out := &in.GitCommits, &out.GitCommits
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ShouldClean != nil {
		in, out := &in.ShouldClean, &out.ShouldClean
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]genericcondition.GenericCondition, len(*in))
		copy(*out, *in)
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AppEndpoints != nil {
		in, out := &in.AppEndpoints, &out.AppEndpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WranglerStatus.
func (in *WranglerStatus) DeepCopy() *WranglerStatus {
	if in == nil {
		return nil
	}
	out := new(WranglerStatus)
	in.DeepCopyInto(out)
	return out
}
