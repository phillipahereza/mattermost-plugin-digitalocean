package main

// Store handles any data we might need to persist
type Store interface {
	StoreUserDOToken(token string, key string) error
	LoadUserDOToken(key string) (string, error)
}

// PluginStore is
type PluginStore struct {
	plugin *Plugin
}

// CreateStore is
func CreateStore(p *Plugin) *PluginStore {
	return &PluginStore{
		plugin: p,
	}
}

// StoreUserDOToken is
func (s *PluginStore) StoreUserDOToken(token, key string) error {
	apiErr := s.plugin.API.KVSet(key, []byte(token))
	if apiErr != nil {
		return apiErr
	}

	return nil
}

// LoadUserDOToken is
func (s *PluginStore) LoadUserDOToken(key string) (string, error) {
	data, apiErr := s.plugin.API.KVGet(key)
	if apiErr != nil {
		return "", apiErr
	}

	return string(data), nil
}
