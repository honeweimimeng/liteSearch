package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"liteSearch/analyzer"
	"liteSearch/index"
	"liteSearch/schema"
)

const AppCtx = "appCtx"

type AppServer struct {
	idx    *index.Index
	schema *schema.Schema
	engine *gin.Engine
}

func NewAppServer() *AppServer {
	app := &AppServer{
		engine: gin.Default(),
	}
	app.engine.Use(Cors())
	return app.Registry()
}

func (s *AppServer) IndexSearch() *AppServer {
	indexSchema := schema.NewSchema()
	_ = indexSchema.AddTextField("title", "en_stem")
	_ = indexSchema.AddTextField("content", "en_stem")
	analyzer.Register("en_stem", analyzer.NewEnglishAnalyzer())
	idx, err := index.NewBuilder(indexSchema).OpenOrCreate("C:\\Users\\Administrator\\Desktop\\liteSearch\\temp")
	if err != nil {
		panic(err)
	}
	s.idx = idx
	s.schema = indexSchema
	return s
}

func (s *AppServer) AppCtx(c *gin.Context) {
	c.Set(AppCtx, s)
}

func (s *AppServer) Registry() *AppServer {
	s.engine.POST("/query", s.AppCtx, Query)
	s.engine.POST("/queryCount", Cors(), s.AppCtx, QueryCount)
	s.engine.POST("/create", Cors(), s.AppCtx, Create)
	return s
}

func (s *AppServer) Start() {
	err := s.engine.Run()
	if err != nil {
		panic(err.Error())
	}
}

func GetData(c *gin.Context) []byte {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err.Error())
	}
	return data
}
