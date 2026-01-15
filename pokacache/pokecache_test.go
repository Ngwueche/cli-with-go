package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Errorf("Expected cache to be created, got nil")
	}

}

func TestAddGetCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	key := "key1"
	data := []byte("val1")

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: key,
			inputVal: data,
		},
	}
	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("Expected to retrieve data, but got not found")
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("Expected %s, got %s", cas.inputVal, actual)
			continue
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	key := "key1"
	data := []byte("val1")
	cache.Add(key, data)

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("Expected cache entry to be reaped, but it still exists")
	}
}
func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	key := "key1"
	data := []byte("val1")
	cache.Add(key, data)

	time.Sleep(interval /2)

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("Cache does not still exist")
	}
}
