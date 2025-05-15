package cache

import (
	"gotask/sqlc/db_generated"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type BannerCacheKey struct {
	Geo     string
	Feature int32
}

type cacheItem struct {
	banner    *db_generated.Banner
	expiresAt time.Time
}

type bannerMemoryCache struct {
	cache      map[BannerCacheKey]cacheItem
	bannerKeys map[pgtype.UUID][]BannerCacheKey
	mutex      sync.RWMutex
	defaultTTL time.Duration
}

func NewBannerMemoryCache(defaultTTL time.Duration) BannerCache {
	cache := &bannerMemoryCache{
		cache:      make(map[BannerCacheKey]cacheItem),
		bannerKeys: make(map[pgtype.UUID][]BannerCacheKey),
		defaultTTL: defaultTTL,
	}

	go cache.cleanupLoop()

	return cache
}

func (c *bannerMemoryCache) Get(geo string, feature int32) (*db_generated.Banner, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	key := BannerCacheKey{Geo: geo, Feature: feature}
	item, exists := c.cache[key]

	if !exists || time.Now().After(item.expiresAt) {
		if exists {
			delete(c.cache, key)
		}
		return nil, false
	}

	return item.banner, true
}

func (c *bannerMemoryCache) Set(geo string, feature int32, banner *db_generated.Banner, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if ttl == 0 {
		ttl = c.defaultTTL
	}

	key := BannerCacheKey{Geo: geo, Feature: feature}
	c.cache[key] = cacheItem{
		banner:    banner,
		expiresAt: time.Now().Add(ttl),
	}

	c.bannerKeys[banner.ID] = append(c.bannerKeys[banner.ID], key)
}

func (c *bannerMemoryCache) Invalidate(id pgtype.UUID) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if keys, exists := c.bannerKeys[id]; exists {
		for _, key := range keys {
			delete(c.cache, key)
		}
		delete(c.bannerKeys, id)
	}
}

func (c *bannerMemoryCache) InvalidateAll() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache = make(map[BannerCacheKey]cacheItem)
	c.bannerKeys = make(map[pgtype.UUID][]BannerCacheKey)
}

func (c *bannerMemoryCache) cleanupLoop() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.cleanup()
	}
}

func (c *bannerMemoryCache) cleanup() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()

	removedBannerKeys := make(map[pgtype.UUID]bool)

	for key, item := range c.cache {
		if now.After(item.expiresAt) {
			delete(c.cache, key)
			removedBannerKeys[item.banner.ID] = true
		}
	}

	for bannerID := range removedBannerKeys {
		keys := c.bannerKeys[bannerID]
		updatedKeys := make([]BannerCacheKey, 0)

		for _, key := range keys {
			if _, exists := c.cache[key]; exists {
				updatedKeys = append(updatedKeys, key)
			}
		}

		if len(updatedKeys) == 0 {
			delete(c.bannerKeys, bannerID)
		} else {
			c.bannerKeys[bannerID] = updatedKeys
		}
	}
}
