// Code generated by github.com/megaease/easemeshctl/cmd/generator, DO NOT EDIT.
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

type tenantGetter struct {
	client *meshClient
}
type tenantInterface struct {
	client *meshClient
}

func (t *tenantGetter) Tenant() TenantInterface {
	return &tenantInterface{client: t.client}
}
func (t *tenantInterface) Get(args_0 context.Context, args_1 string) (*resource.Tenant, error) {
	url := fmt.Sprintf("http://"+t.client.server+apiURL+"/mesh/services/%s/"+"Tenant", args_1)
	r, err := client.NewHTTPJSON().GetByContext(args_0, url, nil, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "get Tenant %s", args_1)
		}
		if statusCode >= 300 {
			return nil, errors.Errorf("call %s failed, return status code %d text %+v", url, statusCode, b)
		}
		Tenant := &v1alpha1.Tenant{}
		err := json.Unmarshal(b, Tenant)
		if err != nil {
			return nil, errors.Wrapf(err, "unmarshal data to v1alpha1.Tenant")
		}
		return resource.ToTenant(Tenant), nil
	})
	if err != nil {
		return nil, err
	}
	return r.(*resource.Tenant), nil
}
func (t *tenantInterface) Patch(args_0 context.Context, args_1 *resource.Tenant) error {
	url := fmt.Sprintf("http://"+t.client.server+apiURL+"/mesh/services/%s/"+"Tenant", args_1)
	object := args_1.ToV1Alpha1()
	_, err := client.NewHTTPJSON().PutByContext(args_0, url, object, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "patch Tenant %s", args_1.Name())
		}
		if statusCode < 300 && statusCode >= 200 {
			return nil, nil
		}
		return nil, errors.Errorf("call PUT %s failed, return statuscode %d text %+v", url, statusCode, b)
	})
	return err
}
func (t *tenantInterface) Create(args_0 context.Context, args_1 *resource.Tenant) error {
	url := fmt.Sprintf("http://"+t.client.server+apiURL+"/mesh/services/%s/"+"Tenant", args_1)
	_, err := client.NewHTTPJSON().PostByContext(args_0, url, nil, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusConflict {
			return nil, errors.Wrapf(ConflictError, "create Tenant %s", args_1.Name())
		}
		if statusCode < 300 && statusCode >= 200 {
			return nil, nil
		}
		return nil, errors.Errorf("call Post %s failed, return statuscode %d text %+v", url, statusCode, b)
	})
	return err
}
func (t *tenantInterface) Delete(args_0 context.Context, args_1 string) error {
	url := fmt.Sprintf("http://"+t.client.server+apiURL+"/mesh/services/%s/"+"Tenant", args_1)
	_, err := client.NewHTTPJSON().DeleteByContext(args_0, url, nil, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "Delete Tenant %s", args_1)
		}
		if statusCode < 300 && statusCode >= 200 {
			return nil, nil
		}
		return nil, errors.Errorf("call Delete %s failed, return statuscode %d text %+v", url, statusCode, b)
	})
	return err
}
func (t *tenantInterface) List(args_0 context.Context) ([]*resource.Tenant, error) {
	url := "http://" + t.client.server + apiURL + "/mesh/services"
	result, err := client.NewHTTPJSON().GetByContext(args_0, url, nil, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "list service")
		}
		if statusCode >= 300 && statusCode < 200 {
			return nil, errors.Errorf("call GET %s failed, return statuscode %d text %+v", url, statusCode, b)
		}
		tenant := []v1alpha1.Tenant{}
		err := json.Unmarshal(b, &tenant)
		if err != nil {
			return nil, errors.Wrapf(err, "unmarshal data to v1alpha1.")
		}
		results := []*resource.Tenant{}
		for _, item := range tenant {
			copy := item
			results = append(results, resource.ToTenant(&copy))
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]*resource.Tenant), nil
}
