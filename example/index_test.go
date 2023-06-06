package main

import (
	"fmt"
	"liteSearch/analyzer"
	"liteSearch/collector"
	"liteSearch/index"
	"liteSearch/indexer"
	"liteSearch/query"
	"liteSearch/reader"
	"liteSearch/schema"
	"runtime"
	"testing"
)

func TestStore(t *testing.T) {
	indexSchema := schema.NewSchema()
	analyzer.Register("en_stem", analyzer.NewEnglishAnalyzer())
	poet := indexSchema.AddTextField("poet", "en_stem")
	content := indexSchema.AddTextField("content", "en_stem")

	idx, err := index.NewBuilder(indexSchema).OpenOrCreate("C:\\Users\\Administrator\\Desktop\\liteSearch\\temp")
	if err != nil {
		panic(err)
	}

	indexWriter, err := indexer.NewIndexWriter(idx, 1000000000)
	if err != nil {
		panic(err)
	}
	defer indexWriter.Close()

	docs := []*schema.Document{
		{FieldValues: []*schema.FieldValue{
			{
				FieldID: poet,
				Value:   "张三",
			},
			{
				FieldID: content,
				Value:   "人们 常说 这是撒啊啊啊",
			},
		}},
		{FieldValues: []*schema.FieldValue{
			{
				FieldID: poet,
				Value:   "李四",
			},
			{
				FieldID: content,
				Value:   "不要担心，不要 得瑟得瑟 常说",
			},
		}},
	}

	for _, doc := range docs {
		indexWriter.AddDocument(doc)
	}
	if _, err := indexWriter.Commit(); err != nil {
		panic(err)
	}
}

func TestSearch(t *testing.T) {
	indexSchema := schema.NewSchema()
	analyzer.Register("en_stem", analyzer.NewEnglishAnalyzer())
	idx, err := index.NewBuilder(indexSchema).OpenOrCreate("C:\\Users\\Administrator\\Desktop\\liteSearch\\temp")
	if err != nil {
		panic(err)
	}
	indexReader, err := reader.NewIndexReader(idx)
	if err != nil {
		panic(err)
	}
	defer indexReader.Close()

	queryParser := query.NewParser(idx.Schema(), idx.Schema().FieldIDs())
	q, err := queryParser.Parse("content:常说")
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
	hits := tupleResult.Left
	count := tupleResult.Right
	fmt.Println("total hit:", count)
	for _, hit := range hits {
		doc := indexReader.GetDoc(hit.DocAddress)
		println(doc.FieldValues[1].Value.(string))
		fmt.Printf("docAddress: %+v, score: %v\n", hit.DocAddress.DocID, hit.Score)
	}
}

func TestOffset(t *testing.T) {
}
