/*
Copyright The KCP Authors.

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


// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"github.com/kcp-dev/kcp/sdk/client/informers/externalversions/internalinterfaces"
)

type ClusterInterface interface {
// Workspaces returns a WorkspaceClusterInformer
	Workspaces() WorkspaceClusterInformer
// WorkspaceTypes returns a WorkspaceTypeClusterInformer
	WorkspaceTypes() WorkspaceTypeClusterInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new ClusterInterface.
func New(f internalinterfaces.SharedInformerFactory, tweakListOptions internalinterfaces.TweakListOptionsFunc) ClusterInterface {
	return &version{factory: f, tweakListOptions: tweakListOptions}
}

// Workspaces returns a WorkspaceClusterInformer
func (v *version) Workspaces() WorkspaceClusterInformer {
	return &workspaceClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
// WorkspaceTypes returns a WorkspaceTypeClusterInformer
func (v *version) WorkspaceTypes() WorkspaceTypeClusterInformer {
	return &workspaceTypeClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
type Interface interface {
// Workspaces returns a WorkspaceInformer
	Workspaces() WorkspaceInformer
// WorkspaceTypes returns a WorkspaceTypeInformer
	WorkspaceTypes() WorkspaceTypeInformer
}

type scopedVersion struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace string
}

// New returns a new ClusterInterface.
func NewScoped(f internalinterfaces.SharedScopedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &scopedVersion{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Workspaces returns a WorkspaceInformer
func (v *scopedVersion) Workspaces() WorkspaceInformer {
	return &workspaceScopedInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
// WorkspaceTypes returns a WorkspaceTypeInformer
func (v *scopedVersion) WorkspaceTypes() WorkspaceTypeInformer {
	return &workspaceTypeScopedInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

