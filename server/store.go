package main

import "encoding/json"

const currentInstanceKVKey string = "current_do_instance"

// Store handles any data we might need to persist
type Store interface {
	StoreInstance(doi DOInstance) error
	LoadInstance() (DOInstance, error)
}

type store struct {
	plugin *Plugin
}

func CreateStore(p *Plugin) *store {
	return &store{
		plugin: p,
	}
}

func (s *store) StoreInstance(doi DOInstance) error {
	data, err := json.Marshal(doi)
	if err != nil {
		return err
	}

	apiErr := s.plugin.API.KVSet(currentInstanceKVKey, data)
	if apiErr != nil {
		return apiErr
	}

	// Set the current instance into configuration
	cloneConfig := s.plugin.getConfiguration().Clone()
	cloneConfig.currentDOInstance = doi
	s.plugin.setConfiguration(cloneConfig)

	return nil
}

func (s *store) LoadInstance() (DOInstance, error) {
	data, err := s.plugin.API.KVGet(currentInstanceKVKey)
	if err != nil {
		return nil, err
	}
	instance := &DOAccountInstance{}
	e := json.Unmarshal(data, instance)
	if e != nil {
		return nil, e
	}
	return instance, nil
}
