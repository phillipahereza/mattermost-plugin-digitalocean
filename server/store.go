package main

// Store handles any data we might need to persist
type Store interface {
	StoreUserDOToken(token string, key string) error
	LoadUserDOToken(key string) (string, error)
	DeleteUserDOToken(key string) error

	// StoreSubscription()
	// LoadSubscription()
	// DeleteSubscription()
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
		s.plugin.API.LogError("Failed to set token", "key", key, "Err", apiErr.Error())
		return apiErr
	}

	return nil
}

// LoadUserDOToken is
func (s *PluginStore) LoadUserDOToken(key string) (string, error) {
	data, apiErr := s.plugin.API.KVGet(key)
	if apiErr != nil {
		s.plugin.API.LogError("Failed to get token", "key", key, "Err", apiErr.Error())
		return "", apiErr
	}

	return string(data), nil
}

// DeleteUserDOToken is
func (s *PluginStore) DeleteUserDOToken(key string) error {
	apiErr := s.plugin.API.KVDelete(key)
	if apiErr != nil {
		s.plugin.API.LogError("Failed to delete token", "key", key, "Err", apiErr.Error())
		return apiErr
	}
	return nil
}
