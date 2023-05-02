package tsn

import (
	"context"
	"github.com/anaregdesign/papaya/cache/graph"
	"github.com/anaregdesign/papaya/nlp"
	"time"
)

type TokenCache struct {
	window int
	dic    *nlp.Dictionary
	c      *graph.GraphCache[int, string]
}

func NewTokenCache(ctx context.Context, window int) *TokenCache {
	return &TokenCache{
		window: window,
		dic:    nlp.NewDictionary(),
		c:      graph.NewGraphCache[int, string](ctx, 30*time.Minute),
	}
}

func (c *TokenCache) Add(words []string) {
	cbow := c.dic.Words2CBOW(words, c.window)
	for _, w := range cbow {
		c.c.AddVertex(w.Source, c.dic.ID2Word[w.Source])
		for k, v := range w.Bow {
			c.c.AddVertex(k, c.dic.ID2Word[k])
			c.c.AddEdge(w.Source, k, float64(v))
		}
	}
}

func (c *TokenCache) Neighbor(word string, step int, k int) *graph.Graph[int, string] {
	return c.c.Neighbor(c.dic.Word2ID[word], step, k, true)
}
