package engine

type Request struct {
	Url string
	ParserFunc func([] byte) ParseResult
}

type ParseResult struct {
	Requests [] Request
	Items [] Item
}

type Item struct {
	Url string
	Id string
	Payload interface{}
}

func NilParserFunc([] byte) ParseResult {
	return ParseResult{}
}