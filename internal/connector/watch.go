package connector

import (
	"errors"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type watchable interface {
	Watch(v1.ListOptions) (watch.Interface, error)
}

func (k *K8SClient) watcher(changeChan chan bool, object watchable) error {
	w, err := object.Watch(v1.ListOptions{})
	if err != nil {
		return err
	}

	for {
		event := <-w.ResultChan()
		if event.Type == watch.Error || event.Object == nil {
			err = errors.New(fmt.Sprintln(event.Object))
			break
		}
		changeChan <- true
	}

	w.Stop()

	return err
}
