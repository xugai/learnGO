package engine

type Request struct {
	Url string
	Parser Parser
}

type Parser interface {
	Parse([] byte) ParseResult
	Serialize() (name string, args [] interface{})
}

type ParserFunc func ([] byte) ParseResult

type ParseResult struct {
	Requests [] Request
	Items [] Item
}

type Item struct {
	Url string
	Id string
	Payload interface{}
}

type NilParser struct {}

// {"ParseCityList", nil}, {"ParseCity", ["houseName", "rent", "url", "id"]}
type FuncParser struct {
	Func ParserFunc
	Name string
}

func (f *FuncParser) Parse(bytes []byte) ParseResult {
	return f.Func(bytes)
}

func (f *FuncParser) Serialize() (name string, args []interface{}) {
	return f.Name, nil
}

func (n NilParser) Parse(_ []byte) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args [] interface{}) {
	return "NilParser", nil
}

func NewFuncParser(parseFunc ParserFunc, name string) *FuncParser{
	return &FuncParser{
		Func: parseFunc,
		Name: name,
	}
}






