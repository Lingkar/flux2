package build

import (
	"bytes"
	"fmt"
	// "io"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"log"
	"os"
	"strings"

	k8syaml "k8s.io/apimachinery/pkg/util/yaml"

	helmv2 "github.com/fluxcd/helm-controller/api/v2"
	// "helm.sh/helm/v3/pkg/action"
	// "helm.sh/helm/v3/pkg/cli"
	// "helm.sh/helm/v3/pkg/kube"
)

func CheckIfFileIsHR(path string) error {
	hr, err := unMarshallHelmRelease(path)
	if err != nil {
		return err
	}

	log.Printf("helmrelease found: %v", hr.Name)
	return nil
}

func unMarshallHelmRelease(path string) (*helmv2.HelmRelease, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read helmrelease file %s: %w", path, err)
	}
	hr := &helmv2.HelmRelease{}
	decoder := k8syaml.NewYAMLOrJSONDecoder(bytes.NewBuffer(data), len(data))
	// check for helmrelease in yaml with the same name and namespace
	for {
		err = decoder.Decode(hr)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshall helmrelease file %s: %w", path, err)
		}

		if strings.HasPrefix(hr.APIVersion, helmv2.GroupVersion.Group+"/") &&
			hr.Kind == helmv2.HelmReleaseKind {
			break
		}
	}
	return hr, nil
}

func isHelmRelease(object *unstructured.Unstructured) bool {
	return strings.HasPrefix(object.GetAPIVersion(), helmv2.GroupVersion.Group+"/") &&
		object.GetKind() == helmv2.HelmReleaseKind
}

func toHelmRelease(object *unstructured.Unstructured) (*helmv2.HelmRelease, error) {
	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(object)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to unstructured: %w", err)
	}
	k := &helmv2.HelmRelease{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj, k)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to HelmRelease: %w", err)
	}
	return k, nil
}
