package model

import "learnGO/crawler/engine"

type SearchResult struct {
	Query string
	Start int
	Hits int64
	PrevFrom int
	NextFrom int
	Items [] engine.Item
}
