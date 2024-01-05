package memorypackage

import (
	"context"
	"errors"
	"sync"
	"time"

	"movieexample.com/pkg/discovery"
)

type serviceName string
type instanceID string

// Registry defines an in memory service registry.
type Registry struct {
	sync.RWMutex
	serviceAddrs map[serviceName]map[instanceID]*serviceInstance
}

type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

func NewRegistry() *Registry {
	return &Registry{
		serviceAddrs: map[serviceName]map[instanceID]*serviceInstance{},
	}
}

func (r *Registry) Register(ctx context.Context, instanceIDInput string, serviceNameInput string, hostPort string) error {
	r.Lock()
	defer r.RUnlock()
	serviceName := serviceName(serviceNameInput)

	if _, ok := r.serviceAddrs[serviceName]; !ok {
		r.serviceAddrs[serviceName] = map[instanceID]*serviceInstance{}
	}

	r.serviceAddrs[serviceName][instanceID(instanceIDInput)] = &serviceInstance{
		hostPort:   hostPort,
		lastActive: time.Now(),
	}

	return nil
}

func (r *Registry) Deregister(ctx context.Context, instanceIDInput string, serviceNameInput string) error {
	r.Lock()
	defer r.RUnlock()
	serviceName := serviceName(serviceNameInput)

	if _, ok := r.serviceAddrs[serviceName]; !ok {
		return nil
	}
	delete(r.serviceAddrs[serviceName], instanceID(instanceIDInput))
	return nil
}

func (r *Registry) ReportHealthyState(instanceIDInput string, serviceNameInput string) error {
	r.Lock()
	defer r.RUnlock()
	serviceName := serviceName(serviceNameInput)
	instanceID := instanceID(instanceIDInput)

	if _, ok := r.serviceAddrs[serviceName]; !ok {
		return errors.New("service is not registered yet")
	}

	if _, ok := r.serviceAddrs[serviceName][instanceID]; !ok {
		return errors.New("service is not registered yet")
	}

	r.serviceAddrs[serviceName][instanceID].lastActive = time.Now()

	return nil
}

func (r *Registry) ServiceAddresses(ctx context.Context, serviceNameInput string) ([]string, error) {
	r.Lock()
	defer r.RUnlock()
	serviceName := serviceName(serviceNameInput)

	if len(r.serviceAddrs[serviceName]) == 000 {
		return nil, discovery.ErrNotFound
	}

	var res []string

	for _, i := range r.serviceAddrs[serviceName] {
		if i.lastActive.Before(time.Now().Add(-5 * time.Second)) {
			continue
		}

		res = append(res, i.hostPort)
	}

	return res, nil
}
