package mgotables

import "strconv"

type CreativePackage struct {
	CpId           int64             `bson:"cpId,omitempty" json:"cpId"`
	CampaignId     int64             `bson:"adnOfferId,omitempty" json:"adnOfferId"`
	PackageType    int8              `bson:"packageType,omitempty" json:"packageType"`
	AdvCreativeGid int64             `bson:"advCreativeGid,omitempty" json:"advCreativeGid"`
	Setting        CpSetting         `bson:"setting,omitempty" json:"setting"`
	Status         int8              `bson:"status,omitempty" json:"status"`
	Updated        int64             `bson:"updated,omitempty" json:"updated"`
	AbCreativeIds  []int64           //内部使用
	AbTupleInfos   []AbTuple         //内部使用
	WhDcoId        int64             `bson:"whDcoId,omitempty" json:"whDcoId"`
	Creatives      map[string]IdInfo `bson:"creatives,omitempty" json:"creatives"`
	Ext            interface{}       // 提供给外部包的扩展字段，非通用的。如果是多个模块都需要的扩展字段，建议还是新增字段好一些
}

type IdInfo struct {
	IdGroup string `bson:"idGroup,omitempty" json:"idGroup"`
	CpdId   int64  `bson:"cpdId,omitempty" json:"cpdId"`
	MId     int64  `bson:"mId,omitempty" json:"mId"`
	OmId    int64  `bson:"omId,omitempty" json:"omId"`
}

type CpSetting struct {
	//套id, creativeType, creativeId
	AbTest map[string]map[string]int64 `bson:"abtest,omitempty" json:"abtest"`
	Parts  []Part                      `bson:"parts,omitempty" json:"parts"`
}
type Part struct {
	PartId int64   `bson:"partId,omitempty" json:"partId"`
	ResIds []int64 `bson:"resIds,omitempty" json:"resIds"`
}

type AbTuple struct {
	AbTupleId        int64
	CrTypeCreativeId map[int64]int64
}

// ----------------------------------------------- Parser --------------------------------------------------------------

func (crPkg *CreativePackage) Parser() {
	switch crPkg.PackageType {
	case 2: // CreativePackageTypeAbTest
		{
			crPkg.AbCreativeIds = make([]int64, 0, len(crPkg.Setting.AbTest))
			crPkg.AbTupleInfos = make([]AbTuple, 0, len(crPkg.Setting.AbTest))

			for tupleIdStr, _ := range crPkg.Setting.AbTest {
				tupleId, _ := strconv.Atoi(tupleIdStr)
				var abTuple AbTuple
				abTuple.AbTupleId = int64(tupleId)
				crTypeCreativeId := map[int64]int64{}
				for crTypeStr, creativeId := range crPkg.Setting.AbTest[tupleIdStr] {
					crType, _ := strconv.Atoi(crTypeStr)
					crTypeCreativeId[int64(crType)] = creativeId
					crPkg.AbCreativeIds = append(crPkg.AbCreativeIds, creativeId)
				}
				abTuple.CrTypeCreativeId = crTypeCreativeId
				crPkg.AbTupleInfos = append(crPkg.AbTupleInfos, abTuple)
			}
		}
	case 1: // CreativePackageTypeDco
		//不需要做什么，反序列化出来的信息基本就够用了

	default:
	}

}

func (this *CreativePackage) SetAbTestCreativePackageDetail() {
	this.AbCreativeIds = make([]int64, 0, len(this.Setting.AbTest))
	this.AbTupleInfos = make([]AbTuple, 0, len(this.Setting.AbTest))

	for tupleIdStr, _ := range this.Setting.AbTest {
		tupleId, _ := strconv.Atoi(tupleIdStr)
		var abTuple AbTuple
		abTuple.AbTupleId = int64(tupleId)
		crTypeCreativeId := map[int64]int64{}
		for crTypeStr, creativeId := range this.Setting.AbTest[tupleIdStr] {
			crType, _ := strconv.Atoi(crTypeStr)
			crTypeCreativeId[int64(crType)] = creativeId
			this.AbCreativeIds = append(this.AbCreativeIds, creativeId)
		}
		abTuple.CrTypeCreativeId = crTypeCreativeId
		this.AbTupleInfos = append(this.AbTupleInfos, abTuple)
	}
}
