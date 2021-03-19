package mgotable_load

import (
	"sync"

	"dsp-template/internal/pkg/base-dbstruct/mgotables"
)

var (
	CreativeIndex = NewCrIndex()
)

func NewCrIndex() *CrIndex {
	return &CrIndex{
		Mu:                  &sync.RWMutex{},
		CreativeIdIndexMap:  map[string]map[int64]*mgotables.SCreativeAttr{},
		CreativeKeyIndexMap: map[string]map[int64]*mgotables.SCreativeAttr{},
	}
}

func UpdateCrIndex(cr *mgotables.Creative) {

	if cr.Status != 1 {
		CreativeIndex.DeleteIndexMap(cr.DocId)
	} else {
		CreativeIndex.UpdateIndexMap(cr.DocId, &cr.Creative)
	}
}

type CrIndex struct {
	Mu *sync.RWMutex
	// 正排表: map[docId]map[creativeId]*creativeAttr
	CreativeIdIndexMap map[string]map[int64]*mgotables.SCreativeAttr
	// 正排表: map[docId]map[creativeKey]*creativeAttr
	CreativeKeyIndexMap map[string]map[int64]*mgotables.SCreativeAttr
}

func (index *CrIndex) UpdateIndexMap(docId string, creative *mgotables.SCreative) {
	crIdMap := map[int64]*mgotables.SCreativeAttr{}
	crKeyMap := map[int64]*mgotables.SCreativeAttr{}

	for key, _ := range creative.Image {
		for i, _ := range creative.Image[key] {
			crKeyMap[creative.Image[key][i].CreativeKey] = &creative.Image[key][i]
			crIdMap[creative.Image[key][i].CreativeId] = &creative.Image[key][i]
		}
	}
	for key, _ := range creative.DyImage {
		for i, _ := range creative.DyImage[key] {
			crKeyMap[creative.DyImage[key][i].CreativeKey] = &creative.DyImage[key][i]
			crIdMap[creative.DyImage[key][i].CreativeId] = &creative.DyImage[key][i]
		}
	}
	for key, _ := range creative.JsImage {
		for i, _ := range creative.JsImage[key] {
			crKeyMap[creative.JsImage[key][i].CreativeKey] = &creative.JsImage[key][i]
			crIdMap[creative.JsImage[key][i].CreativeId] = &creative.JsImage[key][i]
		}
	}
	for key, _ := range creative.MraidImage {
		for i, _ := range creative.MraidImage[key] {
			crKeyMap[creative.MraidImage[key][i].CreativeKey] = &creative.MraidImage[key][i]
			crIdMap[creative.MraidImage[key][i].CreativeId] = &creative.MraidImage[key][i]
		}
	}
	for key, _ := range creative.Video {
		for i, _ := range creative.Video[key] {
			crKeyMap[creative.Video[key][i].CreativeKey] = &creative.Video[key][i]
			crIdMap[creative.Video[key][i].CreativeId] = &creative.Video[key][i]
		}
	}
	for key, _ := range creative.Common {
		for i, _ := range creative.Common[key] {
			crKeyMap[creative.Common[key][i].CreativeKey] = &creative.Common[key][i]
			crIdMap[creative.Common[key][i].CreativeId] = &creative.Common[key][i]
		}
	}
	for key, _ := range creative.Default {
		for i, _ := range creative.Default[key] {
			crKeyMap[creative.Default[key][i].CreativeKey] = &creative.Default[key][i]
			crIdMap[creative.Default[key][i].CreativeId] = &creative.Default[key][i]
		}
	}

	index.Mu.Lock()
	defer index.Mu.Unlock()

	index.CreativeIdIndexMap[docId] = crIdMap
	index.CreativeKeyIndexMap[docId] = crKeyMap
}

func (index *CrIndex) DeleteIndexMap(docId string) {
	index.Mu.Lock()
	defer index.Mu.Unlock()

	delete(index.CreativeIdIndexMap, docId)
	delete(index.CreativeKeyIndexMap, docId)
}

func (index *CrIndex) GetCreativeAttrWithCrKey(docId string, creativeKey int64) *mgotables.SCreativeAttr {
	index.Mu.Lock()
	defer index.Mu.Unlock()

	if _, exist := index.CreativeKeyIndexMap[docId]; exist {
		return index.CreativeKeyIndexMap[docId][creativeKey]
	}
	return nil
}

func (index *CrIndex) GetCreativeAttrWithCrId(docId string, creativeId int64) *mgotables.SCreativeAttr {
	index.Mu.Lock()
	defer index.Mu.Unlock()

	if _, exist := index.CreativeIdIndexMap[docId]; exist {
		return index.CreativeIdIndexMap[docId][creativeId]
	}
	return nil
}
