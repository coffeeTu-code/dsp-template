package abtesting

type AbInfos []*AbInfo

type AbInfo struct {
	FlowControl    FlowControl    `json:"flow_control"`
	ExperimentType string         `json:"experiment_type"`
	HashType       string         `json:"hash_type"`
	Layer          string         `json:"layer"`
	StartTime      string         `json:"start_time"`
	EndTime        string         `json:"end_time"`
	Experiment     ExperimentInfo `json:"experiment"`
}

type FlowControl struct {
	Region      []string `json:"region"`
	AdType      []string `json:"ad_type"`
	Adx         []string `json:"adx"`
	Country     []string `json:"country"`
	Device      []string `json:"device"`
	PackageName []string `json:"package_name"`
}

type ExperimentInfo struct {
	Name    string         `json:"name"`
	Content map[string]int `json:"content"`
}

type ContentInfo struct {
	Name string `json:"name"`
	Num  int    `json:"num"`
}
