package kubernetes30872

import (
	"sync"
	"testing"
)

type PopProcessFunc func()

type ProcessFunc func()

func Util(f func(), stopCh <-chan struct{}) {
	JitterUntil(f, stopCh)
}

func JitterUntil(f func(), stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		default:
		}
		func() {
			f()
		}()
	}
}

type Queue interface {
	HasSynced()
	Pop(PopProcessFunc)
}

type Config struct {
	Queue
	Process ProcessFunc
}

type Controller struct {
	config Config
}

func (c *Controller) Run(stopCh <-chan struct{}) {
	Util(c.processLoop, stopCh)
}

func (c *Controller) HasSynced() {
	c.config.Queue.HasSynced()
}

func (c *Controller) processLoop() {
	c.config.Queue.Pop(PopProcessFunc(c.config.Process))
}

type ControllerInterface interface {
	Run(<-chan struct{})
	HasSynced()
}

type ResourceEventHandler interface {
	OnAdd()
}

type ResourceEventHandlerFuncs struct {
	AddFunc func()
}

func (r ResourceEventHandlerFuncs) OnAdd() {
	if r.AddFunc != nil {
		r.AddFunc()
	}
}

type informer struct {
	controller ControllerInterface

	stopChan chan struct{}
}

type federatedInformerImpl struct {
	sync.Mutex
	clusterInformer informer
}

func (f *federatedInformerImpl) ClustersSynced() {
	f.Lock()
	defer f.Unlock()
	f.clusterInformer.controller.HasSynced()
}

func (f *federatedInformerImpl) addCluster() {
	f.Lock()
	defer f.Unlock()
}

func (f *federatedInformerImpl) Start() {
	f.Lock()
	defer f.Unlock()

	f.clusterInformer.stopChan = make(chan struct{})
	go f.clusterInformer.controller.Run(f.clusterInformer.stopChan)
}

func (f *federatedInformerImpl) Stop() {
	f.Lock()
	defer f.Unlock()
	close(f.clusterInformer.stopChan)
}

type DelayingDeliverer struct{}

func (d *DelayingDeliverer) StartWithHandler(handler func()) {
	go func() {
		handler()
	}()
}

type FederationView interface {
	ClustersSynced()
}

type FederatedInformer interface {
	FederationView
	Start()
	Stop()
}

type NamespaceController struct {
	namespaceDeliverer         *DelayingDeliverer
	namespaceFederatedInformer FederatedInformer
}

func (nc *NamespaceController) isSynced() {
	nc.namespaceFederatedInformer.ClustersSynced()
}

func (nc *NamespaceController) reconcileNamespace() {
	nc.isSynced()
}

func (nc *NamespaceController) Run(stopChan <-chan struct{}) {
	nc.namespaceFederatedInformer.Start()
	go func() {
		<-stopChan
		nc.namespaceFederatedInformer.Stop()
	}()
	nc.namespaceDeliverer.StartWithHandler(func() {
		nc.reconcileNamespace()
	})
}

type DeltaFIFO struct {
	lock sync.RWMutex
}

func (f *DeltaFIFO) HasSynced() {
	f.lock.Lock()
	defer f.lock.Unlock()
}

func (f *DeltaFIFO) Pop(process PopProcessFunc) {
	f.lock.Lock()
	defer f.lock.Unlock()
	process()
}

func NewFederatedInformer() FederatedInformer {
	federatedInformer := &federatedInformerImpl{}
	federatedInformer.clusterInformer.controller = NewInformer(
		ResourceEventHandlerFuncs{
			AddFunc: func() {
				federatedInformer.addCluster()
			},
		})
	return federatedInformer
}

func NewInformer(h ResourceEventHandler) *Controller {
	fifo := &DeltaFIFO{}
	cfg := &Config{
		Queue: fifo,
		Process: func() {
			h.OnAdd()
		},
	}
	return &Controller{config: *cfg}
}

func NewNamespaceController() *NamespaceController {
	nc := &NamespaceController{}
	nc.namespaceDeliverer = &DelayingDeliverer{}
	nc.namespaceFederatedInformer = NewFederatedInformer()
	return nc
}

func TestKubernetes30872_bad_test(t *testing.T) {
	namespaceController := NewNamespaceController()
	stop := make(chan struct{})
	namespaceController.Run(stop)
	close(stop)
}
