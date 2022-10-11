package main

import (
	"encoding/json"
	"sync"
	"time"

	cache "github.com/patrickmn/go-cache"
)

const key = "1"

var userCache = cache.New(time.Second*5, time.Second*10)

var m PgAttrs

func init() {
	a := `{
		"category": "reference",
		"author": "Nigel Rees",
		"title": "Sayings of the Century"
	}`
	err := json.Unmarshal([]byte(a), &m)
	if err != nil {
		panic(err.Error())
	}
	userCache.Set(key, m, 0)
}

type PgAttrs map[string]string

func getAttrs(uid string) PgAttrs {
	if v, ok := userCache.Get(uid); ok {
		attrs := v.(PgAttrs)
		return clonePgAttrs(attrs)
	} else {
		time.Sleep(time.Second * 2)
		setAttrs(uid, m)
		return m
	}
}

func setAttrs(uid string, attrs PgAttrs) {
	userCache.Set(uid, attrs, 0)
}

var mux sync.RWMutex

func clonePgAttrs(attrs PgAttrs) PgAttrs {
	cloneAttrs := make(PgAttrs)
	mux.RLock()
	for k, v := range attrs {
		cloneAttrs[k] = v
	}
	mux.RUnlock()
	return cloneAttrs
}
