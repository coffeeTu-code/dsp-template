package base_bidcache

func Init(cfg *Config) error {
	return defaultBidCache.Init(cfg)
}

func Open() bool {
	return defaultBidCache.cfg.Open
}

func App() *BidCache {
	return defaultBidCache
}

func ReadExchangeConfig(exchange string) *ExchangeConfig {
	return defaultBidCache.cfg.AdxConfig[exchange]
}
