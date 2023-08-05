package main

import (
	"log"

	"gorm.io/gorm"
)

type TLDRDBCached struct {
	Provider TLDRProvider
	DB       *gorm.DB
}

func (t *TLDRDBCached) Retrieve(key string) string {
	// Attempt to read from cache
	item := TLDREntity{}
	err := t.DB.Where("Key = ?", key).First(&item).Error
	if err != nil {
		log.Println(err)
		// read from provider
		res := t.Provider.Retrieve(key)
		// insert to cache
		item.Key = key
		item.Val = res
		t.DB.Create(&item)
		return res
	}
	return item.Val
}

func (t *TLDRDBCached) List() []string {
	items := make([]TLDREntity, 0)
	t.DB.Find(&items)

	res := make([]string, 0)
	mp := make(map[string]struct{}, 0)
	for _, v := range items {
		mp[v.Key] = struct{}{}
		res = append(res, v.Key)
	}

	nonCachedItems := t.Provider.List()
	for _, v := range nonCachedItems {
		if _, ok := mp[v]; !ok {
			res = append(res, v)
		}
	}

	return res
}

func NewTLDRDBCached(nonCachedProvider TLDRProvider) TLDRProvider {
	return &TLDRDBCached{
		Provider: nonCachedProvider,
		DB:       GetConnection(),
	}
}
