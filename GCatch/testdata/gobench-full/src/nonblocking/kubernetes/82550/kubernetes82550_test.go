package kubernetes82550

import (
	"testing"
)

type DockerConfig map[string]DockerConfigEntry

type DockerConfigEntry struct{}

type CachingDockerConfigProvider struct {
	cacheDockerConfig DockerConfig
}

func (d *CachingDockerConfigProvider) Provide() DockerConfig {
	return DockerConfig{}
}

type lazyEcrProvider struct {
	actualProvider *CachingDockerConfigProvider
}

func (p *lazyEcrProvider) LazyProvide() *DockerConfigEntry {
	if p.actualProvider == nil {
		p.actualProvider = &CachingDockerConfigProvider{}
	}
	entry := p.actualProvider.Provide()["0"]
	return &entry
}

func TestKubernetes82550(t *testing.T) {
	provider := &lazyEcrProvider{}
	for i := 0; i < 10; i++ {
		go provider.LazyProvide()
	}
}
