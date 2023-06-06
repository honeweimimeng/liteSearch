package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"liteSearch/collector"
	"liteSearch/query"
	"liteSearch/reader"
	"liteSearch/schema"
	"net/http"
	"runtime"
)

type QueryResult struct {
	Hits  []*DocResult `json:"hits"`
	Count int          `json:"total"`
}

type DocResult struct {
	TopScore *collector.TopScoreResult `json:"topScore"`
	Doc      *schema.Document          `json:"doc"`
}

func ToQueryResult(indexReader *reader.IndexReader,
	tupleResult *collector.TupleResult[[]*collector.TopScoreResult, int]) *QueryResult {
	var result QueryResult
	hits := tupleResult.Left
	result.Count = tupleResult.Right
	for _, hit := range hits {
		result.Hits = append(result.Hits, &DocResult{Doc: indexReader.GetDoc(hit.DocAddress), TopScore: hit})
		fmt.Printf("docAddress: %+v, score: %v\n", hit.DocAddress.DocID, hit.Score)
	}
	return &result
}

type QueryWord struct {
	Field string `json:"field"`
	Query string `json:"query"`
}

func Query(c *gin.Context) {
	appCtx := GetAppCtx(c)
	data := GetData(c)
	var qWord QueryWord
	err := json.Unmarshal(data, &qWord)
	if err != nil {
		panic(err)
	}
	indexReader, err := reader.NewIndexReader(appCtx.idx)
	if err != nil {
		panic(err)
	}
	defer indexReader.Close()
	queryParser := query.NewParser(appCtx.schema, appCtx.schema.FieldIDs())
	word := qWord.Field + ":" + qWord.Query
	if qWord.Field == "" {
		word = ""
		for i, item := range appCtx.schema.Fields {
			if i != 0 && i != len(appCtx.schema.Fields) {
				word = word + " OR "
			}
			word = word + item.Name + ":" + qWord.Query
		}
	}
	q, err := queryParser.Parse(word)
	if err != nil {
		panic(err)
	}
	tupleCollector := collector.NewTupleCollector(
		collector.NewTopScoreCollector(10, 0),
		collector.NewCountCollector(),
	)

	searcher := indexReader.Searcher()
	tupleResult, err := reader.Search(
		searcher,
		q,
		tupleCollector,
		reader.SearchOptionConcurrent(runtime.GOMAXPROCS(0)),
	)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": ToQueryResult(indexReader, tupleResult)})
}

func QueryCount(c *gin.Context) {
	appCtx := GetAppCtx(c)
	indexReader, err := reader.NewIndexReader(appCtx.idx)
	if err != nil {
		panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"result": indexReader.DocTotal()})
}
