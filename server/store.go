package main

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Store handles any data we might need to persist
type Store interface {
	StoreUserDOToken(token string, key string) error
	LoadUserDOToken(key string) (string, error)
	DeleteUserDOToken(key string) error

	StoreSubscription(sub Subscription) error
	LoadSubscription() (Subscription, error)
	AddChannelToSubcription(id string) error
	RemoveChannelFromSubscription(id string) error
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

// StoreSubscription adds an empty subcription on plugin activation
func (s *PluginStore) StoreSubscription(sub Subscription) error {
	data, err := json.Marshal(sub)
	if err != nil {
		return err
	}
	apiErr := s.plugin.API.KVSet(subscriptionKVKey, data)
	if apiErr != nil {
		return apiErr
	}

	return nil
}

// LoadSubscription adds an empty subcription on plugin activation
func (s *PluginStore) LoadSubscription() (Subscription, error) {
	data, apiErr := s.plugin.API.KVGet(subscriptionKVKey)
	if apiErr != nil {
		s.plugin.API.LogError("Failed to get subscription", "key", subscriptionKVKey, "Err", apiErr.Error())
		return Subscription{}, apiErr
	}

	subscription := &Subscription{}
	err := json.Unmarshal(data, subscription)
	if err != nil {
		return Subscription{}, err
	}

	return *subscription, nil
}

// AddChannelToSubcription adds a channel
func (s *PluginStore) AddChannelToSubcription(id string) error {
	sub, err := s.LoadSubscription()
	if err != nil {
		return err
	}
	channels := sub.Channels
	i := indexOf(id, channels)

	// Channel is already subscribed
	if i != -1 {
		return errors.New("already_subcribed")
	}
	channels = append(channels, id)
	sub.Channels = channels

	e := s.StoreSubscription(sub)
	if e != nil {
		return e
	}
	return nil
}

// RemoveChannelFromSubscription removes a channel from subcriptions
func (s *PluginStore) RemoveChannelFromSubscription(id string) error {
	sub, err := s.LoadSubscription()
	if err != nil {
		return err
	}

	channels := sub.Channels
	i := indexOf(id, channels)
	channels = remove(channels, i)
	sub.Channels = channels

	e := s.StoreSubscription(sub)
	if e != nil {
		return e
	}
	return nil

}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
