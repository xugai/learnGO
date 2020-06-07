package config

const (
	// Service config
	SERVICE_PROTOCOL = "tcp"
	//SERVICE_ENDPOINT_PORT = ":1234"
	//CRAWLER_SERVICE_ENDPOINT_PORT = ":9000"
	ITEMSAVER_SERVICE = "ItemSaverService.Save"
	CRAWLER_SERVICE = "CrawlService.Process"
	// Concurrent worker count
	WORKER_COUNT = 5

	// Parser
	PARSECITYLIST = "ParseCityList"
	PARSECITY = "ParseCity"
	PARSEHOUSE = "ParseHouse"
	NILPARSE = "NilParser"

	// Rate limit
	QPS = 20
)
