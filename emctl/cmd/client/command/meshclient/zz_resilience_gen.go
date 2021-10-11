/*
Copyright (c) 2017, MegaEase
All rights reserved.

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
// code generated by github.com/megaease/easemeshctl/cmd/generator, DO NOT EDIT.
package meshclient

import (
	"context"
	"encoding/json"
	"fmt"
	v1alpha1 "github.com/megaease/easemesh-api/v1alpha1"
	resource "github.com/megaease/easemeshctl/cmd/client/resource"
	client "github.com/megaease/easemeshctl/cmd/common/client"
	errors "github.com/pkg/errors"
	"net/http"
)

type resilienceGetter struct {
	client *meshClient
}
type resilienceInterface struct {
	client *meshClient
}

func (r *resilienceGetter) Resilience() ResilienceInterface {
	return &resilienceInterface{client: r.client}
}
func (r *resilienceInterface) Get(args0 context.Context, args1 string) (*resource.Resilience, error) {
	url := fmt.Sprintf("http://"+r.client.server+apiURL+"/mesh/"+"services/%s/resilience", args1)
	r0, err := client.NewHTTPJSON().GetByContext(args0, url, nil, nil).HandleResponse(func(buff []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "get Resilience %s", args1)
		}
		if statusCode >= 300 {
			return nil, errors.Errorf("call %s failed, return status code %d text %+v", url, statusCode, buff)
		}
		Resilience := &v1alpha1.Resilience{}
		err := json.Unmarshal(buff, Resilience)
		if err != nil {
			return nil, errors.Wrapf(err, "unmarshal data to v1alpha1.Resilience")
		}
		return resource.ToResilience(args1, Resilience), nil
	})
	if err != nil {
		return nil, err
	}
	return r0.(*resource.Resilience), nil
}
func (r *resilienceInterface) Patch(args0 context.Context, args1 *resource.Resilience) error {
	url := fmt.Sprintf("http://"+r.client.server+apiURL+"/mesh/"+"services/%s/resilience", args1)
	object := args1.ToV1Alpha1()
	_, err := client.NewHTTPJSON().PutByContext(args0, url, object, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "patch Resilience %s", args1.Name())
		}
		if statusCode < 300 && statusCode >= 200 {
			return nil, nil
		}
		return nil, errors.Errorf("call PUT %s failed, return statuscode %d text %+v", url, statusCode, b)
	})
	return err
}
func (r *resilienceInterface) Create(args0 context.Context, args1 *resource.Resilience) error {
	url := fmt.Sprintf("http://"+r.client.server+apiURL+"/mesh/"+"services/%s/resilience", args1)
	_, err := client.NewHTTPJSON().PostByContext(args0, url, nil, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusConflict {
			return nil, errors.Wrapf(ConflictError, "create Resilience %s", args1.Name())
		}
		if statusCode < 300 && statusCode >= 200 {
			return nil, nil
		}
		return nil, errors.Errorf("call Post %s failed, return statuscode %d text %+v", url, statusCode, b)
	})
	return err
}
func (r *resilienceInterface) Delete(args0 context.Context, args1 string) error {
	url := fmt.Sprintf("http://"+r.client.server+apiURL+"/mesh/"+"services/%s/resilience", args1)
	_, err := client.NewHTTPJSON().DeleteByContext(args0, url, nil, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "Delete Resilience %s", args1)
		}
		if statusCode < 300 && statusCode >= 200 {
			return nil, nil
		}
		return nil, errors.Errorf("call Delete %s failed, return statuscode %d text %+v", url, statusCode, b)
	})
	return err
}
func (r *resilienceInterface) List(args0 context.Context) ([]*resource.Resilience, error) {
	url := "http://" + r.client.server + apiURL + "/mesh/"
	result, err := client.NewHTTPJSON().GetByContext(args0, url, nil, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "list service")
		}
		if statusCode >= 300 && statusCode < 200 {
			return nil, errors.Errorf("call GET %s failed, return statuscode %d text %+v", url, statusCode, b)
		}
		services := []v1alpha1.Service{}
		err := json.Unmarshal(b, &services)
		if err != nil {
			return nil, errors.Wrapf(err, "unmarshal data to v1alpha1.")
		}
		results := []*resource.Resilience{}
		for _, service := range services {
			if service.Resilience != nil {
				results = append(results, resource.ToResilience(service.Name, service.Resilience))
			}
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]*resource.Resilience), nil
}
