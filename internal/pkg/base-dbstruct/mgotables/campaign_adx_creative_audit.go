package mgotables

import (
	"strings"

	"dsp-template/api/enum"
)

type CampaignAdxCreativeAudit struct {
	CampaignId  int64  `bson:"campaignId,omitempty" json:"campaignId"`
	CountryCode string `bson:"countryCode,omitempty" json:"countryCode"`
	PackageName string `bson:"packageName,omitempty" json:"packageName"`
	CompanyId   int    `bson:"companyId,omitempty" json:"companyId"` //审核资质时返回的广告主id
	Status      int8   `bson:"status,omitempty" json:"status"`
	Updated     int64  `bson:"updated,omitempty" json:"updated"`

	AdxCreativeAuditSlice            []AdxCreativeAudit                        `bson:"adx,omitempty" json:"adx"`
	AdxCreativeSpecGroupCreativesIdx map[string]*CreativeSpecGroupCreativesIdx // key=adx

	Ext interface{} // 提供给外部包的扩展字段，非通用的。如果是多个模块都需要的扩展字段，建议还是新增字段好一些
}

type (
	AdxCreativeAudit struct {
		Name              string         `bson:"name,omitempty" json:"name"` //adx 名字
		Industry          string         `bson:"industry,omitempty" json:"industry"`
		CreativeSpecSlice []CreativeSpec `bson:"adCreative,omitempty" json:"adCreative"` //每个adx对应多个AdCreative,即多个广告位unit
	}
	CreativeSpecGroupCreativesIdx struct {
		Industry        string
		CreativeSpecMap map[string]*CreativeSpec
	}

	CreativeSpec struct {
		Id            string              `bson:"creativeSpec,omitempty" json:"creativeSpec"`
		CreativeGroup []SGroupCreativeIdx `bson:"creativeGroup,omitempty" json:"creativeGroup"`
		AuditInfo     AuditInfo           // 自定义结构体，审核内容，快速获取 blockGroupId，加速实时检索

		NoActiveGroup   []SGroupCreativeIdx // 审核未通过的素材组
		NoCompleteGroup []SGroupCreativeIdx // 缺失素材id的素材组
	}
	SGroupCreativeIdx struct {
		Id        string         `bson:"id,omitempty" json:"id"`
		Creatives []SCreativeAdt `bson:"creatives,omitempty" json:"creatives"`
		Status    uint8          `bson:"status,omitempty" json:"status"`

		Attribute string `bson:"attribute,omitempty" json:"attribute"`
		Industry  string `bson:"industry,omitempty" json:"industry"`
	}
	SCreativeAdt struct {
		Id            int64 `bson:"id,omitempty" json:"id"`
		VideoLength   int
		CreativeDocId string
	}
	AuditInfo struct {
		Values    map[string][]string // key=creative values, value=groupId
		Attribute map[string][]string // key=group Attribute, value=groupId
		Industry  map[string][]string // key=group Industry, value=groupId
	}
)

// ----------------------------------------------- Get Attr ------------------------------------------------------------

//@GetGroupCreativeIndex
// adx ：流量方
// creativeSpec ：流量方广告位
func (campAdxAudit *CampaignAdxCreativeAudit) GetGroupCreativeIndex(adx, creativeSpec string) *CreativeSpec {
	if specAudit, exist := campAdxAudit.AdxCreativeSpecGroupCreativesIdx[adx]; exist {
		if groupCreative, exist := specAudit.CreativeSpecMap[creativeSpec]; exist {
			return groupCreative
		}
	}
	return nil
}

//@IsActive
func (group *SGroupCreativeIdx) IsActive() bool {
	//status 枚举值： 0准备中，1待审核，2通过，3拒绝，4信息变更，5送审失败，6预审通过，11重新送审，44删除。 其中先审后投2是通过，先投后审2和6通过
	return group.Status == 2 || group.Status == 6
}

// ----------------------------------------------- Parser --------------------------------------------------------------

func (campAdxAudit *CampaignAdxCreativeAudit) Parser(CreativeDocIdCC, CreativeDocIdALL map[int64]*SCreativeAttr) {
	campAdxAudit.AdxCreativeSpecGroupCreativesIdx = map[string]*CreativeSpecGroupCreativesIdx{}

	for _, adx := range campAdxAudit.AdxCreativeAuditSlice {
		if campAdxAudit.AdxCreativeSpecGroupCreativesIdx[adx.Name] == nil {
			campAdxAudit.AdxCreativeSpecGroupCreativesIdx[adx.Name] = &CreativeSpecGroupCreativesIdx{
				Industry:        adx.Industry,
				CreativeSpecMap: map[string]*CreativeSpec{},
			}
		}

		for i := range adx.CreativeSpecSlice {
			var creativeSpec = adx.CreativeSpecSlice[i]
			creativeSpec.AuditInfo.Values = map[string][]string{}
			creativeSpec.AuditInfo.Attribute = map[string][]string{}
			creativeSpec.AuditInfo.Industry = map[string][]string{}

			groups := make([]SGroupCreativeIdx, 0, len(creativeSpec.CreativeGroup))
			for _, group := range creativeSpec.CreativeGroup {
				if !group.IsActive() {
					creativeSpec.NoActiveGroup = append(creativeSpec.NoActiveGroup, group)
					continue // 审核中，不投放
				}
				value, complete := group.checkCreative(CreativeDocIdCC, CreativeDocIdALL)
				if !complete {
					creativeSpec.NoCompleteGroup = append(creativeSpec.NoCompleteGroup, group)
					continue // 素材组不完整，不投放
				}
				groups = append(groups, group)
				for _, v := range value {
					creativeSpec.AuditInfo.Values[v] = append(creativeSpec.AuditInfo.Values[v], group.Id)
				}
				if len(group.Attribute) > 0 {
					for _, attr := range strings.Split(group.Attribute, ",") {
						creativeSpec.AuditInfo.Attribute[attr] = append(creativeSpec.AuditInfo.Attribute[attr], group.Id)
					}
				}
				if len(group.Industry) > 0 {
					for _, industry := range strings.Split(group.Industry, ",") {
						creativeSpec.AuditInfo.Industry[industry] = append(creativeSpec.AuditInfo.Industry[industry], group.Id)
					}
				}
			}
			if len(groups) == 0 {
				continue
			}
			creativeSpec.CreativeGroup = groups
			campAdxAudit.AdxCreativeSpecGroupCreativesIdx[adx.Name].CreativeSpecMap[creativeSpec.Id] = &creativeSpec
		}
	}
}

func (group *SGroupCreativeIdx) checkCreative(CreativeDocIdCC, CreativeDocIdALL map[int64]*SCreativeAttr) (value []string, complete bool) {
	complete = true
	for i := range group.Creatives {
		crId := group.Creatives[i].Id
		crAttr, exist := CreativeDocIdCC[crId]
		if !exist {
			crAttr, exist = CreativeDocIdALL[crId]
			if !exist {
				complete = false
				return // 素材组投放素材的ID未在内存中找到，素材组无效
			}
		}
		group.Creatives[i].CreativeDocId = crAttr.DocId
		switch enum.SubResourceType(crAttr.SubResourceType) {
		case enum.SubResourceTypeVideo:
			if crAttr.VideoLength == 0 {
				complete = false
				return // 视频素材咋能没有length呢
			}
			group.Creatives[i].VideoLength = crAttr.VideoLength
		case enum.SubResourceTypeAppName, enum.SubResourceTypeAppDesc, enum.SubResourceTypeAdWords:
			if len(crAttr.Value) > 0 {
				value = append(value, crAttr.Value)
			}
		}
	}
	return
}
