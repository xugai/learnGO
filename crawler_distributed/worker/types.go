package worker

import (
	"fmt"
	"learnGO/crawler/engine"
	parser2 "learnGO/crawler/lianjia/parser"
	"learnGO/crawler_distributed/config"
)

type SerializedParser struct {
	FuncName string
	Args [] interface{}
}

type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct {
	Items [] engine.Item
	Requests [] Request
}

func SerializeRequest(request engine.Request) Request {
	name, args := request.Parser.Serialize()
	return Request{
		Url: request.Url,
		Parser: SerializedParser{
			FuncName: name,
			Args: args,
		},
	}
}

func SerializeParseResult(r engine.ParseResult) ParseResult {
	parseResult := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		parseResult.Requests = append(parseResult.Requests, SerializeRequest(req))
	}
	return parseResult
}

func DeserializeRequest(request Request) (engine.Request, error) {
	parser, err := DeserializeParser(request.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url: request.Url,
		Parser: parser,
	}, nil
}

func DeserializeParseResult(parseResult ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: parseResult.Items,
	}

	for _, req := range parseResult.Requests {
		engineRequest, err := DeserializeRequest(req)
		if err != nil {
			panic(err)
		}
		result.Requests = append(result.Requests, engineRequest)
	}
	return result
}

func DeserializeParser(parser SerializedParser) (engine.Parser, error) {
	switch parser.FuncName {
	case config.PARSECITYLIST:
		return engine.NewFuncParser(parser2.ParseCityList, config.PARSECITYLIST), nil
	case config.PARSECITY:
		return engine.NewFuncParser(parser2.ParseCity, config.PARSECITY), nil
	case config.PARSEHOUSE:
		return parser2.NewHouseParser(parser.Args[0].(string),
										parser.Args[1].(string),
										parser.Args[2].(string),
										parser.Args[3].(string)), nil
	case config.NILPARSE:
		return engine.NilParser{}, nil
	default:
		return nil, fmt.Errorf("unknow parser given: %v\n", parser)
	}
}


