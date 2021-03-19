package backend

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
