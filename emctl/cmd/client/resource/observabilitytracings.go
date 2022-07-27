/*
 * Copyright (c) 2021, MegaEase
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package resource

import (
	"github.com/megaease/easemesh-api/v2alpha1"
	"github.com/megaease/easemeshctl/cmd/client/resource/meta"
)

type (
	// ObservabilityTracings describes observability tracings resource of the EaseMesh
	ObservabilityTracings struct {
		meta.MeshResource `yaml:",inline"`
		Spec              *v2alpha1.ObservabilityTracings `yaml:"spec" jsonschema:"required"`
	}
)

// ToV2Alpha1 converts a ObservabilityTracings resource to v2alpha1.ObservabilityTracings
func (r *ObservabilityTracings) ToV2Alpha1() (result *v2alpha1.ObservabilityTracings) {
	return r.Spec
}

// ToObservabilityTracings converts a v2alpha1.ObservabilityTracings resource to a ObservabilityTracings resource
func ToObservabilityTracings(serviceID string, tracing *v2alpha1.ObservabilityTracings) *ObservabilityTracings {
	result := &ObservabilityTracings{
		Spec: &v2alpha1.ObservabilityTracings{},
	}
	result.MeshResource = NewObservabilityTracingsResource(DefaultAPIVersion, serviceID)
	result.Spec = tracing
	return result
}
