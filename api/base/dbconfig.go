package base

type RedisConfig struct {
	Host               string
	Password           string
	DialTimeout        int
	ReadTimeout        int
	MaxRetries         int
	PoolSize           int
	PoolTimeout        int
	MinIdleConns       int
	MaxConnAge         int
	Expiration         int
	ReadOnly           bool
	RouteRandomly      bool
	MaxRedirects       int
	IdleTimeout        int
	IdleCheckFrequency int
	RouteByLatency     bool
	BeginExpiration    int
}

type RedisHystrixConfig struct {
	Use                   bool
	Timeout               int
	MaxConcurrentRequests int
	ErrorPercentThreshold int
	SleepWindow           int
}

type MongoConfig struct {
	Url          string
	DB           string
	Table        string
	BaseInterval int //全量更新数据周期，单位 s
	IncInterval  int //增量更新数据周期，单位 s
}
