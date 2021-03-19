package mgotable_load

import (
	"sync"

	"dsp-template/internal/pkg/base-dbstruct/mgotables"
)

var (
	CampaignInvertedIndex = NewCampInvertedIndex()
)

func NewCampInvertedIndex() *CampInvertedIndex {
	return &CampInvertedIndex{
		Mu:                  &sync.RWMutex{},
		InvertedIndexMap:    make(map[string]map[int64]struct{}),
		CampIdRecallKeysMap: make(map[int64]map[string]struct{}),
	}
}

func UpdateCampInvertedIndex(camp *mgotables.Campaign) {
	recallKeys := GetCampaignRecallKeys(camp)
	deleteKeys := make([]string, 0)
	campId := camp.CampaignId
	if camp.Status != 1 {
		// 单子已停, 删除正排中的key
		deleteKeys = CampaignInvertedIndex.DeleteCampIdRecallKeysMap(campId)
	} else {
		// 更新正排表, 获取因更新而无用的key
		deleteKeys = CampaignInvertedIndex.UpdateCampIdRecallKeysMap(recallKeys, campId)
		// 更新倒排表, 插入因更新产生新的key
		CampaignInvertedIndex.UpdateInvertedIndexMap(recallKeys, campId)
	}
	// 删除倒排表中无用的key
	CampaignInvertedIndex.DeleteInvertedIndexMap(deleteKeys, campId)
}

type CampInvertedIndex struct {
	Mu *sync.RWMutex
	// 倒排表: map[recallKey]map[campId]struct{}
	InvertedIndexMap map[string]map[int64]struct{}
	// 正排表: map[campId]map[recallKey]struct{}
	CampIdRecallKeysMap map[int64]map[string]struct{}
}

func (index *CampInvertedIndex) UpdateInvertedIndexMap(recallKeys map[string]struct{}, campId int64) {
	index.Mu.Lock()
	defer index.Mu.Unlock()

	for recallKey, _ := range recallKeys {
		if _, ok := index.InvertedIndexMap[recallKey]; !ok {
			campIdMap := make(map[int64]struct{})
			campIdMap[campId] = struct{}{}
			index.InvertedIndexMap[recallKey] = campIdMap
		} else {
			index.InvertedIndexMap[recallKey][campId] = struct{}{}
		}
	}
}

func (index *CampInvertedIndex) DeleteInvertedIndexMap(deleteKeys []string, campId int64) {
	index.Mu.Lock()
	defer index.Mu.Unlock()

	for _, deleteKey := range deleteKeys {
		_, ok := index.InvertedIndexMap[deleteKey]
		if !ok {
			continue
		}
		delete(index.InvertedIndexMap[deleteKey], campId)
		if len(index.InvertedIndexMap[deleteKey]) == 0 {
			delete(index.InvertedIndexMap, deleteKey)
		}
	}
}

func (index *CampInvertedIndex) DeleteCampIdRecallKeysMap(campId int64) []string {
	index.Mu.Lock()
	defer index.Mu.Unlock()

	deleteKeys := make([]string, 0)
	if recallKeys, ok := index.CampIdRecallKeysMap[campId]; ok {
		for val := range recallKeys {
			deleteKeys = append(deleteKeys, val)
		}
		delete(index.CampIdRecallKeysMap, campId)
	}
	return deleteKeys
}

func (index *CampInvertedIndex) UpdateCampIdRecallKeysMap(recallKeys map[string]struct{}, campId int64) []string {
	deleteKeys := make([]string, 0)
	if _, ok := index.CampIdRecallKeysMap[campId]; !ok {
		index.CampIdRecallKeysMap[campId] = make(map[string]struct{})
		for k := range recallKeys {
			index.CampIdRecallKeysMap[campId][k] = struct{}{}
		}
	} else {
		//字段更新后清掉无用的key
		for k, _ := range index.CampIdRecallKeysMap[campId] {
			if _, ok := recallKeys[k]; !ok {
				delete(index.CampIdRecallKeysMap[campId], k)
				deleteKeys = append(deleteKeys, k)
			}
		}
		//字段更新后插入新的key
		for k, _ := range recallKeys {
			if _, ok := index.CampIdRecallKeysMap[campId][k]; !ok {
				index.CampIdRecallKeysMap[campId][k] = struct{}{}
			}
		}
	}
	return deleteKeys
}

func (index *CampInvertedIndex) GetCampIds(recallKey string) []int64 {
	index.Mu.RLock()
	defer index.Mu.RUnlock()

	campIds := make([]int64, 0)
	campIdMap, ok := index.InvertedIndexMap[recallKey]
	if !ok {
		return campIds
	}
	for campId, _ := range campIdMap {
		campIds = append(campIds, campId)
	}
	return campIds
}

func (index *CampInvertedIndex) GetCampIdList() []int64 {
	index.Mu.RLock()
	defer index.Mu.RUnlock()

	campIds := make([]int64, 0)
	for campId, _ := range index.CampIdRecallKeysMap {
		campIds = append(campIds, campId)
	}
	return campIds
}

func (index *CampInvertedIndex) GetAllRecallKey() []string {
	index.Mu.RLock()
	defer index.Mu.RUnlock()

	res := make([]string, 0, len(index.InvertedIndexMap))
	for key, _ := range index.InvertedIndexMap {
		res = append(res, key)
	}
	return res
}

func (index *CampInvertedIndex) Length() int {
	return len(index.InvertedIndexMap)
}
