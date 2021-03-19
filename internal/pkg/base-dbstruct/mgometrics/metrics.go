package mgometrics

import (
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

//*************** Metrics const ******************
type Metrics int

const (
	DataUpdated Metrics = iota //mongo数据表最后更新的时间
	DataNumber                 //mongo数据表最后更新的数据条数
	CreativeNum
	GroupCreativeNum
)

//************** SetMetrics *******************
//监控指标维度
type Labels struct {
	TableName    string
	Adx          string
	CreativeSpec string
	CreativeType string
	FormatType   string
	Status       string
}

var metricsAddFunction = map[Metrics]func(labels Labels, value float64){}

func SetMetrics(item Metrics, labels Labels, value float64) {
	addMetrics, ok := metricsAddFunction[item]
	if !ok {
		return
	}
	addMetrics(labels, value)
}

var tableUpdatedTimeM sync.Map

func DataUpdatedTime(TableName string) {
	tableUpdatedTimeM.Store(TableName, time.Now().Unix())
}

func init() {
	go func() {
		ticker := time.NewTicker(time.Minute * 1)
		for _ = range ticker.C {
			now := time.Now().Unix()
			tableUpdatedTimeM.Range(func(key, value interface{}) bool {
				tablename, ok := key.(string)
				if !ok {
					return true
				}
				updatedTime, ok := value.(int64)
				if !ok {
					return true
				}
				if tablename != "" && updatedTime > 0 && now-updatedTime > 0 {
					SetMetrics(DataUpdated, Labels{TableName: tablename}, float64(now-updatedTime))
				}
				return true
			})
		}
	}()
}

//************** InitMetrics *******************

var (
	NameSpace = "Voyager"
	SubSystem = "MgoFetcher"
)

var (
	Region = ""
	IP     = ""
)

func InitMgoExtractorMetrics(items ...string) {
	if len(items) > 0 {
		NameSpace = items[0]
	}
	if len(items) > 1 {
		SubSystem = items[1]
	}
	if len(items) > 2 {
		Region = items[2]
	}
	if len(items) > 3 {
		IP = items[3]
	}

	DataUpdatedGauge := newGaugeMetrics("data_updated", []string{"region", "ip", "table"})
	metricsAddFunction[DataUpdated] = func(labels Labels, value float64) {
		DataUpdatedGauge.WithLabelValues(Region, IP, labels.TableName).Set(value)
	}

	DataNumberGauge := newGaugeMetrics("data_number", []string{"region", "ip", "table"})
	metricsAddFunction[DataNumber] = func(labels Labels, value float64) {
		DataNumberGauge.WithLabelValues(Region, IP, labels.TableName).Set(value)
	}

	CreativeNumGauge := newGaugeMetrics("creative_num", []string{"creativetype", "formattype"})
	metricsAddFunction[CreativeNum] = func(labels Labels, value float64) {
		CreativeNumGauge.WithLabelValues(labels.CreativeType, labels.FormatType).Set(value)
	}

	GroupCreativeNumGauge := newGaugeMetrics("group_creative_num", []string{"adx", "spec", "status"})
	metricsAddFunction[GroupCreativeNum] = func(labels Labels, value float64) {
		GroupCreativeNumGauge.WithLabelValues(labels.Adx, labels.CreativeSpec, labels.Status).Set(value)
	}
}

//************* base function ******************

func newCounterMetrics(name string, labels []string) *prometheus.CounterVec {
	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: NameSpace,
		Subsystem: SubSystem,
		Name:      name,
		Help:      "Count < labels=" + strings.Join(labels, ",") + " >",
	}, labels)
	prometheus.MustRegister(counter)
	return counter
}

func newHistogramMetrics(name string, labels []string, buckets []float64) *prometheus.HistogramVec {
	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: NameSpace,
		Subsystem: SubSystem,
		Name:      name,
		Help:      "Histogram < labels=" + strings.Join(labels, ",") + " >",
		Buckets:   buckets,
	}, labels)
	prometheus.MustRegister(histogram)
	return histogram
}

func newGaugeMetrics(name string, labels []string) *prometheus.GaugeVec {
	gauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: NameSpace,
		Subsystem: SubSystem,
		Name:      name,
		Help:      "Gauge < labels=" + strings.Join(labels, ",") + " >",
	}, labels)
	prometheus.MustRegister(gauge)
	return gauge
}
