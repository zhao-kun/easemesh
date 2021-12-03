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

//go:generate go run github.com/megaease/easemeshctl/cmd/transformer Global ServiceCanary=servicecanaries/%s

package meshclient

import (
	"context"

	"github.com/megaease/easemeshctl/cmd/client/resource"
)

// ServiceCanaryGetter represents a ServiceCanary resource accessor.
type ServiceCanaryGetter interface {
	ServiceCanary() ServiceCanaryInterface
}

// ServiceCanaryInterface captures the set of operations for interacting
// with the EaseMesh REST apis of the service canaray resource.
type ServiceCanaryInterface interface {
	Get(context.Context, string) (*resource.ServiceCanary, error)
	Patch(context.Context, *resource.ServiceCanary) error
	Create(context.Context, *resource.ServiceCanary) error
	Delete(context.Context, string) error
	List(context.Context) ([]*resource.ServiceCanary, error)
}
