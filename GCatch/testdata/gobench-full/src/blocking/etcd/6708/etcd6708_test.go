package etcd6708

import (
	"context"
	"sync"
	"testing"
)

type EndpointSelectionMode int

const (
	EndpointSelectionRandom EndpointSelectionMode = iota
	EndpointSelectionPrioritizeLeader
)

type MembersAPI interface {
	Leader(ctx context.Context)
}

type Client interface {
	Sync(ctx context.Context)
	SetEndpoints()
	httpClient
}

type httpClient interface {
	Do(context.Context)
}

type httpClusterClient struct {
	sync.RWMutex
	selectionMode EndpointSelectionMode
}

func (c *httpClusterClient) getLeaderEndpoint() {
	mAPI := NewMembersAPI(c)
	mAPI.Leader(context.Background())
}

func (c *httpClusterClient) SetEndpoints() {
	switch c.selectionMode {
	case EndpointSelectionRandom:
	case EndpointSelectionPrioritizeLeader:
		c.getLeaderEndpoint()
	}
}

func (c *httpClusterClient) Do(ctx context.Context) {
	c.RLock() // block here
	c.RUnlock()
}

func (c *httpClusterClient) Sync(ctx context.Context) {
	c.Lock()
	defer c.Unlock()

	c.SetEndpoints()
}

type httpMembersAPI struct {
	client httpClient
}

func (m *httpMembersAPI) Leader(ctx context.Context) {
	m.client.Do(ctx)
}

func NewMembersAPI(c Client) MembersAPI {
	return &httpMembersAPI{
		client: c,
	}
}
func TestEtcd6708(t *testing.T) {
	hc := &httpClusterClient{
		selectionMode: EndpointSelectionPrioritizeLeader,
	}
	hc.Sync(context.Background())
}
