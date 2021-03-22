package main

import (
        opkit "github.com/rook/operator-kit"
        sample "github.com/nuthankumar/cosmosdb/pkg/apis/dbprovision/v1"
        sampleclient "github.com/nuthankumar/cosmosdb/pkg/client/clientset/versioned/typed/dbprovision/v1"
        "k8s.io/client-go/kubernetes"
        "k8s.io/client-go/tools/cache"
)

type SampleController struct {
        context         *opkit.Context
        sampleClientset sampleclient.DbprovisionV1Interface
        clientset       *kubernetes.Clientset
}

func newSampleController(context *opkit.Context, sampleClientset sampleclient.DbprovisionV1Interface, clientset *kubernetes.Clientset) *SampleController {
        return &SampleController{
                context:         context,
                sampleClientset: sampleClientset,
                clientset:       clientset,
        }
}

func (c *SampleController) StartWatch(namespace string, stopCh chan struct{}) error {

        resourceHandlers := cache.ResourceEventHandlerFuncs{
                AddFunc:    c.onAdd,
                UpdateFunc: c.onUpdate,
                DeleteFunc: c.onDelete,
        }
        restClient := c.sampleClientset.RESTClient()
        watcher := opkit.NewWatcher(sample.CosmosResource, namespace, resourceHandlers, restClient)
        go watcher.Watch(&sample.DBProvisioning{}, stopCh)
        return nil
}

func (c *SampleController) onAdd(obj interface{}) {
}

func (c *SampleController) onUpdate(oldObj, newObj interface{}) {
}

func (c *SampleController) onDelete(obj interface{}) {

}

