package abtesting

import (
	"context"
	"errors"
	"github.com/hashicorp/consul/api"
	jsoniter "github.com/json-iterator/go"
	"hash/crc32"
	"log"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
	"unsafe"
)

type AbTesting struct {
	path   string
	client *api.Client
}

type FlowInfo struct {
	Device      string
	Region      string
	Adx         string
	AdType      string
	Country     string
	PackageName string
	AppId       int64
	UnitId      int64
	HashKey     string
}

type Info struct {
	Region      map[string]struct{} `json:"region"`
	AdType      map[string]struct{} `json:"ad_type"`
	Adx         map[string]struct{} `json:"adx"`
	Country     map[string]struct{} `json:"country"`
	Device      map[string]struct{} `json:"device"`
	PackageName map[string]struct{} `json:"package_name"`
	AppId       map[int64]struct{}  `json:"app_id"`
	UnitId      map[int64]struct{}  `json:"unit_id"`
	Name        string              `json:"name"`
	Content     []ContentInfo       `json:"content"`
	IsHashKey   bool                `json:"is_hash_key"`
}

type Experiments struct {
	Exclusive  []Info
	MultiLayer map[string][]Info
}

const HashNum = 10000
const ExclusiveFlow = "exclusive_flow"

var ExperimentsInfo *Experiments

func NewAbTesting(addr, path string) (*AbTesting, error) {
	config := api.DefaultConfig()
	config.Address = addr
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &AbTesting{
		path:   path,
		client: client,
	}, nil
}

func (at *AbTesting) Init(ctx context.Context) error {
	ExperimentsInfo = &Experiments{}
	err := at.updateAb()
	if err != nil {
		log.Printf("get data failed. err: %s", err.Error())
		return err
	}
	go func() {
		update := time.After(time.Minute)
		for {
			select {
			case <-ctx.Done():
				return
			case <-update:
				err := at.updateAb()
				update = time.After(time.Minute)
				if err != nil {
					log.Printf("update failed. err: %s", err.Error())
				}
			}
		}
	}()
	return nil
}

func (at *AbTesting) GetAbFlow(flow *FlowInfo) *AbFlow {
	allAbflags := at.GetAbTesting(flow)
	if allAbflags == nil {
		allAbflags = map[string]string{}
	}
	return &AbFlow{
		AllAbflags: allAbflags,
		UsedAbflag: map[string]string{},
	}
}

func (at *AbTesting) GetAbTesting(flow *FlowInfo) map[string]string {
	if ExperimentsInfo == nil {
		return nil
	}

	tmpExperiments := *(*Experiments)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&ExperimentsInfo))))
	sum := 0
	randomSalt := ExclusiveFlow + "-" + strconv.Itoa(int(time.Now().UnixNano()))
	for _, info := range tmpExperiments.Exclusive {
		if !checkFlow(flow, info) {
			continue
		}
		for _, content := range info.Content {
			sum += content.Num
		}
	}

	if key, value, ok := abTestHash(randomSalt, sum, tmpExperiments.Exclusive); ok {
		return map[string]string{key: value}
	}

	value := make(map[string]string)
	for layer, info := range tmpExperiments.MultiLayer {
		if len(info) == 0 {
			continue
		}
		sum := 0
		for _, info := range info {
			if !checkFlow(flow, info) {
				continue
			}
			for _, content := range info.Content {
				sum += content.Num
			}
		}
		var hashSalt string
		if info[0].IsHashKey && flow.HashKey != "" {
			hashSalt = flow.HashKey + layer
		} else {
			hashSalt = randomSalt + layer
		}
		if k, v, ok := abTestHash(hashSalt, sum, info); ok {
			value[k] = v
		}
	}
	return value
}

func ToString(m map[string]string) string {
	if m == nil || len(m) == 0 {
		return ""
	}
	sb := strings.Builder{}
	for k, v := range m {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(v)
		sb.WriteString(";")
	}
	return sb.String()[0 : sb.Len()-1]
}

func (at *AbTesting) updateAb() error {
	kv := at.client.KV()
	if kv == nil {
		return errors.New("consul kv info get failed")
	}

	pair, _, err := kv.Get(at.path, &api.QueryOptions{})
	if err != nil {
		return err
	}
	var abInfos []*AbInfo

	err = jsoniter.Unmarshal(pair.Value, &abInfos)
	if err != nil {
		return errors.New("json unmarshal failed")
	}
	if len(abInfos) == 0 {
		return errors.New("ab info is nil")
	}

	experiments := &Experiments{
		Exclusive:  make([]Info, 0),
		MultiLayer: make(map[string][]Info, 0),
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	for _, abInfo := range abInfos {

		end, _ := time.ParseInLocation("2006-01-02 15:04:05", abInfo.EndTime, loc)
		now, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().String(), loc)
		if abInfo.EndTime != "" && now.Unix() > end.Unix() {
			continue
		}

		experiment := &Info{
			Region:      map[string]struct{}{},
			AdType:      map[string]struct{}{},
			Adx:         map[string]struct{}{},
			Country:     map[string]struct{}{},
			Device:      map[string]struct{}{},
			PackageName: map[string]struct{}{},
			AppId:       map[int64]struct{}{},
			UnitId:      map[int64]struct{}{},
			Name:        "",
			Content:     make([]ContentInfo, 0),
		}
		flowControl(experiment, *abInfo)

		if abInfo.ExperimentType == "exclusive" {
			experiments.Exclusive = append(experiments.Exclusive, *experiment)
		} else {
			experiments.MultiLayer[abInfo.Layer] = append(experiments.MultiLayer[abInfo.Layer], *experiment)
		}
	}
	ExperimentsInfo = (*Experiments)(atomic.SwapPointer((*unsafe.Pointer)((unsafe.Pointer)(&experiments)), unsafe.Pointer(&ExperimentsInfo)))
	return nil
}

func abTestHash(hashSalt string, sum int, infos []Info) (string, string, bool) {
	if sum == 0 {
		return "", "", false
	}

	v, num := 0, HashNum
	if sum > HashNum {
		num = sum
	}
	hashNum := int(crc32.ChecksumIEEE(*(*[]byte)(unsafe.Pointer(&hashSalt))) % uint32(num))

	for _, info := range infos {
		for _, value := range info.Content {
			v += value.Num
			if hashNum <= v {
				return info.Name, value.Name, true
			}
		}
	}
	return "", "", false
}

func flowControl(exclusive *Info, abInfo AbInfo) {

	for _, v := range abInfo.FlowControl.Device {
		exclusive.Device[v] = struct{}{}
	}
	for _, v := range abInfo.FlowControl.AdType {
		exclusive.AdType[v] = struct{}{}
	}
	for _, v := range abInfo.FlowControl.Adx {
		exclusive.Adx[v] = struct{}{}
	}
	for _, v := range abInfo.FlowControl.Country {
		exclusive.Country[v] = struct{}{}
	}
	for _, v := range abInfo.FlowControl.Region {
		exclusive.Region[v] = struct{}{}
	}
	for _, v := range abInfo.FlowControl.PackageName {
		exclusive.PackageName[v] = struct{}{}
	}
	for _, v := range abInfo.FlowControl.AppId {
		exclusive.AppId[v] = struct{}{}
	}
	for _, v := range abInfo.FlowControl.UnitId {
		exclusive.UnitId[v] = struct{}{}
	}

	exclusive.Name = abInfo.Experiment.Name
	exclusive.IsHashKey = abInfo.HashType == "hash"

	keys := make([]string, 0, len(abInfo.Experiment.Content))
	for k := range abInfo.Experiment.Content {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		exclusive.Content = append(exclusive.Content, ContentInfo{
			Name: key,
			Num:  abInfo.Experiment.Content[key],
		})
	}
}

func checkFlow(flow *FlowInfo, info Info) bool {
	if flow == nil {
		return false
	}

	if len(info.AdType) > 0 {
		if _, ok := info.AdType[flow.AdType]; !ok {
			return false
		}
	}

	if len(info.Adx) > 0 {
		if _, ok := info.Adx[flow.Adx]; !ok {
			return false
		}
	}

	if len(info.Country) > 0 {
		if _, ok := info.Country[flow.Country]; !ok {
			return false
		}
	}

	if len(info.Device) > 0 {
		if _, ok := info.Device[flow.Device]; !ok {
			return false
		}
	}

	if len(info.Region) > 0 {
		if _, ok := info.Region[flow.Region]; !ok {
			return false
		}
	}

	if len(info.PackageName) > 0 {
		if _, ok := info.PackageName[flow.PackageName]; !ok {
			return false
		}
	}

	if len(info.AppId) > 0 {
		if _, ok := info.AppId[flow.AppId]; !ok {
			return false
		}
	}

	if len(info.UnitId) > 0 {
		if _, ok := info.UnitId[flow.UnitId]; !ok {
			return false
		}
	}

	return true
}
