/*
Copyright 2016 The Rook Authors. All rights reserved.

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

// Package main for a sample operator
package main

import (
        "flag"
        "fmt"
        "os"
        "os/signal"
        "syscall"
        "time"
        opkit "github.com/rook/operator-kit"
        sample "cosmosdb/pkg/apis/dbprovision/v1"
        sampleclient "cosmosdb/pkg/client/clientset/versioned/typed/dbprovision/v1"
        "k8s.io/api/core/v1"
        apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
        "k8s.io/client-go/kubernetes"
        "k8s.io/client-go/rest"
        "k8s.io/client-go/tools/clientcmd"
)

var (
        kubeconfig string
)

func init() {
        flag.StringVar(&kubeconfig, "kubeconfig", "/opt/ees/kubespray/inventory/ees/artifacts/admin.conf", "Path to KUBECONFIG for running out of cluster. (Default: null)")
}

func main() {
        flag.Parse()
        fmt.Println("Getting kubernetes context")
        context, sampleClientset, clientset, err := createContext(kubeconfig)
        if err != nil {
                fmt.Println("failed to create context")
                fmt.Printf("failed to create context. %+v\n", err)
                os.Exit(1)
        }

        // Create and wait for CRD resources
        fmt.Println("Registering the sample resource")
        resources := []opkit.CustomResource{sample.CosmosResource}
        err = opkit.CreateCustomResources(*context, resources)
        if err != nil {
                fmt.Println("Failed to create custom resource")
                fmt.Printf("failed to create custom resource. %+v\n", err)
                os.Exit(1)
        }

        // create signals to stop watching the resources
        signalChan := make(chan os.Signal, 1)
        stopChan := make(chan struct{})
        signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

        // start watching the sample resource
        fmt.Println("Watching the sample resource")
        controller := newSampleController(context, sampleClientset, clientset)
        controller.StartWatch(v1.NamespaceAll, stopChan)

        for {
                select {
                case <-signalChan:
                        fmt.Println("shutdown signal received, exiting...")
                        close(stopChan)
                        return
                }
        }
}

func getClientConfig(kubeconfig string) (*rest.Config, error) {
        if kubeconfig != "" {
                fmt.Println("Getting kubeconfig from local")
                return clientcmd.BuildConfigFromFlags("", kubeconfig)
        }
        fmt.Println("Getting Incluster Kubeconfig")
        return rest.InClusterConfig()
}

func createContext(kubeconfig string) (*opkit.Context, sampleclient.DbprovisionV1Interface, *kubernetes.Clientset,   error) {
        config, err := getClientConfig(kubeconfig)
        if err != nil {
                fmt.Println("failed to get k8s config")
                return nil, nil, nil,fmt.Errorf("failed to get k8s config. %+v", err)
        }

        clientset, err := kubernetes.NewForConfig(config)
        if err != nil {
                fmt.Println("failed to get k8s client")
                return nil, nil,nil, fmt.Errorf("failed to get k8s client. %+v", err)
        }

        apiExtClientset, err := apiextensionsclient.NewForConfig(config)
        if err != nil {
                fmt.Println("failed to create k8s API extension clientset")
                return nil, nil,nil, fmt.Errorf("failed to create k8s API extension clientset. %+v", err)
        }

        sampleClientset, err := sampleclient.NewForConfig(config)
        if err != nil {
                fmt.Println("failed to create sample clientset")
                return nil, nil,nil, fmt.Errorf("failed to create sample clientset. %+v", err)
        }

        context := &opkit.Context{
                Clientset:             clientset,
                APIExtensionClientset: apiExtClientset,
                Interval:              500 * time.Millisecond,
                Timeout:               60 * time.Second,
        }
        return context, sampleClientset, clientset, nil

}


