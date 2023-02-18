package serving3148

import (
	"sync"
	"testing"
)

type PodAutoscalerInterface interface {
	Create()
}

type PodAutoscalersGetter interface {
	PodAutoscalers() PodAutoscalerInterface
}

type AutoscalingV1alpha1Interface interface {
	PodAutoscalersGetter
}

type clientset_Interface interface {
	AutoscalingV1alpha1() AutoscalingV1alpha1Interface
}

type FakeAutoscalingV1alpha1 struct {
	*Fake
}

func (c *FakeAutoscalingV1alpha1) PodAutoscalers() PodAutoscalerInterface {
	return &FakePodAutoscalers{c}
}

type Clientset struct {
	Fake
}

func (c *Clientset) AutoscalingV1alpha1() AutoscalingV1alpha1Interface {
	return &FakeAutoscalingV1alpha1{Fake: &c.Fake}
}

type FakePodAutoscalers struct {
	Fake *FakeAutoscalingV1alpha1
}

func (c *FakePodAutoscalers) Create() {
	c.Fake.Invokes()
}

type Reconciler struct {
	ServingClientSet clientset_Interface
}

func (c *Reconciler) Reconcile() {
	c.reconcile()
}

func (c *Reconciler) reconcile() {
	phases := []struct {
		name string
		f    func()
	}{{
		name: "KPA",
		f:    c.reconcileKPA,
	}}
	for _, phase := range phases {
		phase.f()
	}
}

func (c *Reconciler) reconcileKPA() {
	c.createKPA()
}

func (c *Reconciler) createKPA() {
	c.ServingClientSet.AutoscalingV1alpha1().PodAutoscalers().Create()
}

type controller_Reconciler interface {
	Reconcile()
}

type Impl struct {
	controller_Reconciler controller_Reconciler
}

func (c *Impl) Run(threadiness int) {
	sg := sync.WaitGroup{}
	defer sg.Wait()

	for i := 0; i < threadiness; i++ {
		sg.Add(1)
		go func() {
			defer sg.Done()
			c.processNextWorkItem()
		}()
	}
}

func (c *Impl) processNextWorkItem() {
	c.controller_Reconciler.Reconcile()
}

func NewImpl(r controller_Reconciler) *Impl {
	return &Impl{
		controller_Reconciler: r,
	}
}

func NewController() *Impl {
	c := &Reconciler{}
	return NewImpl(c)
}

type Group struct {
	wg      sync.WaitGroup
	errOnce sync.Once
}

func (g *Group) Wait() {
	g.wg.Wait()
}

func (g *Group) Go(f func()) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		f()
	}()
}

type Hooks struct{}

func NewHooks() *Hooks {
	return &Hooks{}
}
func (h *Hooks) OnUpdate(fake *Fake) {
	fake.PrependReactor()
}

type Reactor interface{}

type SimpleReactor struct{}

type Fake struct {
	ReactionChain []Reactor
}

func (c *Fake) Invokes() {
	for _ = range c.ReactionChain {
	}
}

func (c *Fake) PrependReactor() {
	c.ReactionChain = append([]Reactor{&SimpleReactor{}})
}

func TestServing3148(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cs := &Clientset{}
		controller := NewController()
		controller.controller_Reconciler.(*Reconciler).ServingClientSet = cs
		eg := &Group{}
		defer func() {
			eg.Wait()
		}()
		eg.Go(func() { controller.Run(1) })
		h := NewHooks()
		h.OnUpdate(&cs.Fake)
	}()
	wg.Wait()
}
