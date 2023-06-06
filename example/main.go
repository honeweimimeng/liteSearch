package main

import (
	"liteSearch/index"
	"liteSearch/schema"
)

type Demo struct {
	indexSchema *schema.Schema
	idx         *index.Index
}

func main() {
	server := NewAppServer()
	server.IndexSearch().Start()
}
