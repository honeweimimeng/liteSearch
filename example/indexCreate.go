package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"liteSearch/indexer"
	"liteSearch/schema"
	"net/http"
)

func Create(c *gin.Context) {
	appCtx := GetAppCtx(c)
	data := GetData(c)
	var doc []*schema.Document
	err := json.Unmarshal(data, &doc)
	if err != nil {
		panic(err)
	}
	err = doIndex(appCtx, doc)
	if err != nil {
		panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"message": "create success"})
}

func doIndex(appCtx *AppServer, docs []*schema.Document) error {
	indexWriter, err := indexer.NewIndexWriter(appCtx.idx, 100_000_000)
	for _, item := range docs {
		for _, field := range item.FieldValues {
			field.FieldID = appCtx.schema.FieldByName(field.FieldName).ID
		}
	}
	for _, doc := range docs {
		indexWriter.AddDocument(doc)
	}
	if _, err := indexWriter.Commit(); err != nil {
		panic(err)
	}
	return err
}

func GetAppCtx(c *gin.Context) *AppServer {
	v, _ := c.Get(AppCtx)
	return v.(*AppServer)
}
