package tool

import (
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	"yinwhm.com/yin/catw/models"
)

var (
	searcher = engine.Engine{}
	docId    uint64
)

func init()  {
	searcher.Init(types.EngineInitOptions{
		IndexerInitOptions: &types.IndexerInitOptions{
			IndexType: types.DocIdsIndex,
		},
		SegmenterDictionaries: "../../../github.com/huichen/wukong/data/dictionary.txt",
		StopTokenFile:         "../../../github.com/huichen/wukong/data/stop_tokens.txt",
	})
	//defer searcher.Close()
}

func Search(articles []models.Article,keyString string)(ids []interface{})  {
	//searcher.Init(types.EngineInitOptions{
	//	IndexerInitOptions: &types.IndexerInitOptions{
	//		IndexType: types.DocIdsIndex,
	//	},
	//	SegmenterDictionaries: "../../../github.com/huichen/wukong/data/dictionary.txt",
	//	StopTokenFile:         "../../../github.com/huichen/wukong/data/stop_tokens.txt",
	//})
	defer searcher.Close()
	//length := len(articles)
	for _, s := range articles{
		docId = uint64(s.Tid)
		searcher.IndexDocument(docId, types.DocumentIndexData{Content: s.TextContent}, false)
	}

	searcher.FlushIndex()

	resp := searcher.Search(types.SearchRequest{Text:keyString})

	for _, s := range resp.Docs{
		ids =append(ids,s.DocId)
	}
	return
}
