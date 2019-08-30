package kabaneroplatform

import (
	"context"
	mf "github.com/jcrossley3/manifestival"
	kabanerov1alpha1 "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"k8s.io/apimachinery/pkg/api/errors"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func reconcile_appsody(ctx context.Context, k *kabanerov1alpha1.Kabanero, c client.Client) error {
	filename := "config/reconciler/appsody-operator/appsody-0.1.0.yaml"
	m, err := mf.NewManifest(filename, true, c)
	if err != nil {
		return err
	}

	transforms := []mf.Transformer{
		mf.InjectOwner(k),
		mf.InjectNamespace(k.GetNamespace()),
	}

	err = m.Transform(transforms...)
	if err != nil {
		return err
	}

	// Enables Appsody
	err = m.ApplyAll()

	return err
}

// Retrieves the Appsody deployment status.
func getAppsodyStatus(k *kabanerov1alpha1.Kabanero, c client.Client, reqLogger logr.Logger) (bool, error) {

	ready := false
	message := "The Appsody application deployment does not have condition indicating it is available"

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "Deployment",
		Group:   "apps",
		Version: "v1",
	})
	err := c.Get(context.Background(), client.ObjectKey{
		Namespace: k.ObjectMeta.Namespace,
		Name:      "appsody-operator",
	}, u) 
	if err == nil {
		conditionsMap, ok, err := unstructured.NestedFieldCopy(u.Object, "status", "conditions")
		if err == nil && ok {
			conditions, ok := conditionsMap.([]interface{})
			if ok {
				for _, conditionObject := range conditions {
					aCondition, ok := conditionObject.(map[string]interface{})
					if ok {
						typeValue , ok, err := unstructured.NestedString(aCondition, "type")
						if err == nil && ok {
							statusValue , ok, err := unstructured.NestedString(aCondition, "status")
							if err == nil && ok {
								if typeValue == "Available" && statusValue == "True" {
									ready = true
								}
							}
						}
					} else {
						message = "An error occurred using an Apposdy deployment status condition"
					}
				}
			} else {
				message = "An error occurred using the Apposdy deployment status conditions"
			}
		} else {
			message = "An error occurred getting the Apposdy deployment status conditions"
                        if err != nil {
				reqLogger.Error(err, message)
				message = message + ": " + err.Error()
			}
		}
	} else {
		if errors.IsNotFound(err) {
			message = "The Appsody deployment was not found"
		} else {
			message = "An error occurred retrieving the Apposdy deployment"
		}
		reqLogger.Error(err, message)
		message = message + ": " + err.Error()
		ready = false
	}
	if ready == true {
		k.Status.Appsody.Ready = "True"
		k.Status.Appsody.ErrorMessage = ""
		err = nil				
	} else {
		k.Status.Appsody.Ready = "False"
		k.Status.Appsody.ErrorMessage = message
	}

	return ready, err
}