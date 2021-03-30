package main

import (
	"log"
	"os"

	"github.com/nuthankumar/az_cosmos_cassandra/operations"
	"github.com/nuthankumar/az_cosmos_cassandra/utils"

	sample "github.com/nuthankumar/cosmosdb/pkg/apis/dbprovision/v1"
	sampleclient "github.com/nuthankumar/cosmosdb/pkg/client/clientset/versioned/typed/dbprovision/v1"
	opkit "github.com/rook/operator-kit"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

var (
	cosmosCassandraContactPoint string
	cosmosCassandraPort         string
	cosmosCassandraUser         string
	cosmosCassandraPassword     string
)

type SampleController struct {
	context         *opkit.Context
	sampleClientset sampleclient.DbprovisionV1Interface
	clientset       *kubernetes.Clientset
}

func init() {
	cosmosCassandraContactPoint = os.Getenv("COSMOSDB_CASSANDRA_CONTACT_POINT")
	cosmosCassandraPort = os.Getenv("COSMOSDB_CASSANDRA_PORT")
	cosmosCassandraUser = os.Getenv("COSMOSDB_CASSANDRA_USER")
	cosmosCassandraPassword = os.Getenv("COSMOSDB_CASSANDRA_PASSWORD")

	if cosmosCassandraContactPoint == "" || cosmosCassandraUser == "" || cosmosCassandraPassword == "" || cosmosCassandraPort == "" {
		log.Fatal("missing mandatory environment variables")
	}
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
	data := obj.(*sample.DBProvisioning).DeepCopy()
	valid := utils.ValidateClientSession(data.Spec.CosmosCassandraContactPoint, cosmosCassandraContactPoint, data.Spec.CosmosCassandraPassword, cosmosCassandraPassword)
	crd_valid := utils.ValidateCrd(data)
	if valid && crd_valid {
		session := utils.GetSession(cosmosCassandraContactPoint, cosmosCassandraPort, cosmosCassandraUser, cosmosCassandraPassword)
		defer session.Close()
		log.Println("Connected to Azure Cosmos DB")
		operations.CreateCRD(data, session, c.clientset)
	}
}

func (c *SampleController) onUpdate(oldObj, newObj interface{}) {
	data := oldObj.(*sample.DBProvisioning).DeepCopy()
	newdata := newObj.(*sample.DBProvisioning).DeepCopy()
	valid := utils.ValidateClientSession(data.Spec.CosmosCassandraContactPoint, newdata.Spec.CosmosCassandraContactPoint, data.Spec.CosmosCassandraPassword, newdata.Spec.CosmosCassandraPassword)
	if valid {
		session := utils.GetSession(cosmosCassandraContactPoint, cosmosCassandraPort, cosmosCassandraUser, cosmosCassandraPassword)
		defer session.Close()
		log.Println("Nothing to do on update")
	}
}

func (c *SampleController) onDelete(obj interface{}) {
	data := obj.(*sample.DBProvisioning).DeepCopy()
	valid := utils.ValidateClientSession(data.Spec.CosmosCassandraContactPoint, cosmosCassandraContactPoint, data.Spec.CosmosCassandraPassword, cosmosCassandraPassword)
	crd_valid := utils.ValidateCrd(data)
	if valid && crd_valid {
		session := utils.GetSession(cosmosCassandraContactPoint, cosmosCassandraPort, cosmosCassandraUser, cosmosCassandraPassword)
		defer session.Close()
		log.Println("Connected to Azure Cosmos DB")
		operations.DeleteCRD(data, session, data.Spec.Keyspace, data.Spec.ClientID, c.clientset)
	}
}
