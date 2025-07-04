package data

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"fmt"
	"sync"
)

type KeyStore struct {
	mu sync.RWMutex
	keys map[int64]any
}

func NewKeyStore() *KeyStore{
	return &KeyStore{keys: make(map[int64]any)}
}

func (s *KeyStore) Set(keyID int64, key any) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.keys[keyID] = key
}

func (s *KeyStore) Get(keyID int64) (any, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	key, ok := s.keys[keyID]
	return key, ok
}

func (s *KeyStore) Delete(keyID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.keys, keyID)
}

func AssertPrivateKey(key any) (any, error) {
	switch k := key.(type) {
	case *ecdsa.PrivateKey:
		return k, nil
	case *rsa.PrivateKey:
		return k, nil
	case ed25519.PrivateKey:
		return k, nil
	default:
		return nil, fmt.Errorf("unsupported private key type")
	}
}
