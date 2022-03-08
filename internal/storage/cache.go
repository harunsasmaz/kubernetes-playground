package storage

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

func (s *Storage) CacheGet(ctx context.Context, key string) (interface{}, error) {
	resp, err := s.cache.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var raw interface{}
	if err := json.Unmarshal([]byte(resp), &raw); err != nil {
		log.Printf("got error while unmarshalling: %v", err)
	}

	return raw, nil
}

func (s *Storage) CacheSet(ctx context.Context, key string, todo []byte) error {
	_, err := s.cache.Set(ctx, key, todo, 5*time.Minute).Result()
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CacheDelete(ctx context.Context, key string) {
	s.cache.Del(ctx, key)
}
