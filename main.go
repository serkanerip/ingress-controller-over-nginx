/*
Copyright 2016 The Kubernetes Authors.

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

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"fmt"

	client "github.com/serkanerip/hello-k8s-client/internal"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {

	k8sClient := client.NewK8SClient()
	clientset := k8sClient.GetClientSet()

	w, _ := clientset.CoreV1().Services("default").Watch(context.TODO(), metav1.ListOptions{})
	go func() {
		for event := range w.ResultChan() {
			svc := event.Object.(*v1.Service)
			switch event.Type {
			case watch.Added:
				fmt.Printf("Service %s/%s added\n", svc.ObjectMeta.Namespace, svc.ObjectMeta.Name)
			case watch.Modified:
				fmt.Printf("Service %s/%s modified\n", svc.ObjectMeta.Namespace, svc.ObjectMeta.Name)
			case watch.Deleted:
				fmt.Printf("Service %s/%s deleted\n", svc.ObjectMeta.Namespace, svc.ObjectMeta.Name)
			}
		}
	}()

	go func() {
		for event := range w.ResultChan() {
			svc := event.Object.(*v1.Service)
			switch event.Type {
			case watch.Added:
				fmt.Printf("Service %s/%s added\n", svc.ObjectMeta.Namespace, svc.ObjectMeta.Name)
			case watch.Modified:
				fmt.Printf("Service %s/%s modified\n", svc.ObjectMeta.Namespace, svc.ObjectMeta.Name)
			case watch.Deleted:
				fmt.Printf("Service %s/%s deleted\n", svc.ObjectMeta.Namespace, svc.ObjectMeta.Name)
			}
		}
	}()

	forever := make(chan (int))

	<-forever
	// list, _ := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	// for _, srv := range list.Items {
	// 	fmt.Printf("Selectors of service: %s\n", srv.GetName())
	// 	for k, v := range srv.Spec.Selector {
	// 		fmt.Printf("\t%s:%s, ", k, v)
	// 	}
	// 	fmt.Println("\n-----------")
	// }
	// os.Exit(1)

	// list, err = clientset.ExtensionsV1beta1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
	// for _, ingress := range list.Items {
	// 	fmt.Println(ingress.Labels)
	// }
}
