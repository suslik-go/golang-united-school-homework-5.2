package cache

import "time"

type Data struct {
	key      string
	value    string
	deadline time.Time
}

type Cache struct {
	CacheData []Data
}

func NewCache() Cache {
	return Cache{}
}

func (c *Cache) Get(key string) (string, bool) {
	for _, data := range c.CacheData {
		if data.key == key && data.deadline.Before(time.Now()) {
			return data.value, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	newData := Data{key, value, time.Now()}
	c.CacheData = append(c.CacheData, newData)
}

func (c *Cache) Keys() []string {
	var notExpiredKeys []string
	for _, data := range c.CacheData {
		if data.deadline.Before(time.Now()) {
			notExpiredKeys = append(notExpiredKeys, data.key)
		}
	}
	return notExpiredKeys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	newData := Data{key, value, deadline}
	c.CacheData = append(c.CacheData, newData)
}
