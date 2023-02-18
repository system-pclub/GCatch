package istio16742

import (
	"sync"
	"testing"
)

var (
	adsClients      = map[string]*XdsConnection{}
	adsClientsMutex sync.RWMutex
)

type Collection []struct{}

func BuildSidecarVirtualHostsFromConfigAndRegistry(proxyLabels Collection) {}

type ConfigGenerator interface {
	BuildHTTPRoutes(node *Proxy)
}

type ConfigGeneratorImpl struct{}

func (configgen *ConfigGeneratorImpl) BuildHTTPRoutes(node *Proxy) {
	configgen.buildSidecarOutboundHTTPRouteConfig(node)
}

func (configgen *ConfigGeneratorImpl) buildSidecarOutboundHTTPRouteConfig(node *Proxy) {
	BuildSidecarVirtualHostsFromConfigAndRegistry(node.WorkloadLabels)
}

type Proxy struct {
	WorkloadLabels Collection
}

type XdsConnection struct {
	modelNode *Proxy
}

func newXdsConnection() *XdsConnection {
	return &XdsConnection{
		modelNode: &Proxy{},
	}
}

type DiscoveryServer struct {
	ConfigGenerator ConfigGenerator
}

func (s *DiscoveryServer) addCon(con *XdsConnection) {
	adsClientsMutex.Lock()
	defer adsClientsMutex.Unlock()
	adsClients["1"] = con
}

func (s *DiscoveryServer) StreamAggregatedResources() {
	con := newXdsConnection()
	s.addCon(con)
	s.pushRoute(con)
}

func (s *DiscoveryServer) generateRawRoutes(con *XdsConnection) {
	s.ConfigGenerator.BuildHTTPRoutes(con.modelNode)
}

func (s *DiscoveryServer) pushRoute(con *XdsConnection) {
	s.generateRawRoutes(con)
}

func (s *DiscoveryServer) WorkloadUpdate() {
	adsClientsMutex.RLock()
	for _, connection := range adsClients {
		connection.modelNode.WorkloadLabels = nil
	}
	adsClientsMutex.RUnlock()
}

type XDSUpdater interface {
	WorkloadUpdate()
}

type MemServiceDiscovery struct {
	EDSUpdater XDSUpdater
}

func (sd *MemServiceDiscovery) AddWorkload() {
	sd.EDSUpdater.WorkloadUpdate()
}

func TestIstio16742(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		registry := &MemServiceDiscovery{
			EDSUpdater: &DiscoveryServer{
				ConfigGenerator: &ConfigGeneratorImpl{},
			},
		}
		go func() {
			defer wg.Done()
			registry.EDSUpdater.(*DiscoveryServer).StreamAggregatedResources()
		}()
		go func() {
			defer wg.Done()
			registry.AddWorkload()
		}()
	}()
	wg.Wait()
}
