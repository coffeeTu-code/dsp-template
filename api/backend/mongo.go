package backend

type MongoConfig struct {
	TryTimes       int
	URI            string
	DB             string
	Collection     string
	ConnectTimeout int
	ReadTimeout    int
}
