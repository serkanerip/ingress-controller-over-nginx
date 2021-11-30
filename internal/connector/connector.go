package connector

import (
	"context"
	"fmt"
	"time"

	"k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8SClient) GetIngresses() ([]v1beta1.Ingress, error) {
	list, err := c.clientset.ExtensionsV1beta1().Ingresses("").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return list.Items, nil
}

func (c *K8SClient) WatchIngressForChanges() (chan bool, error) {
	changeChan := make(chan bool)

	go func() {
		for {
			err := c.watcher(changeChan, c.clientset.ExtensionsV1beta1().Ingresses(""))
			fmt.Println(err)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	return changeChan, nil
}
