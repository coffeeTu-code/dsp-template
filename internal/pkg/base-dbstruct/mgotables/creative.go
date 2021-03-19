package mgotables

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"dsp-template/api/enum"
)

type Creative struct {
	DocId        string `bson:"_id,omitempty" json:"_id"`
	CampaignId   int64  `bson:"campaignId,omitempty" json:"campaignId"`
	PackageName  string `bson:"packageName,omitempty" json:"packageName"`
	CountryCode  string `bson:"countryCode,omitempty" json:"countryCode"`
	Status       int64  `bson:"status,omitempty" json:"status"`
	Updated      int64  `bson:"updated,omitempty" json:"updated"`
	AdvertiserId int64  `bson:"advertiserId,omitempty" json:"advertiserId"`

	//creative_type、creative_id、字段名、字段值
	Content  map[string]map[string]map[string]interface{} `bson:"content,omitempty" json:"content"`
	Creative SCreative

	Ext interface{} // 提供给外部包的扩展字段，非通用的。如果是多个模块都需要的扩展字段，建议还是新增字段好一些
}

type (
	SCreative struct {
		Default    map[string][]SCreativeAttr //key=CreativeType
		Image      map[string][]SCreativeAttr //key=Resolution
		DyImage    map[string][]SCreativeAttr //key=Resolution
		JsImage    map[string][]SCreativeAttr //key=Resolution
		MraidImage map[string][]SCreativeAttr //key=Resolution
		Video      map[string][]SCreativeAttr //key=VideoResolution
		Common     map[string][]SCreativeAttr //key=CreativeType
	}
	SCreativeAttr struct {
		DocId           string
		CountryCode     string
		PackageName     string
		CreativeType    int
		ResourceType    int //系统自定义类型
		SubResourceType int //系统自定义类型
		FormatType      int //系统自定义类型
		CreativeKey     int64
		CreativeId      int64
		AdvCreativeId   string // TsCreativeId
		Mime            string
		Attribute       string
		FMd5            string
		Source          int
		Ctime           int64
		Utime           int64
		UniqCid         int64
		VideoResolution string
		VideoSize       string
		VideoLength     int
		Bitrate         int
		Clarity         int
		Width           int
		Height          int
		Orientation     int
		MinOs           int64
		Platform        int32
		Resolution      string
		Url             string
		TagCode         string
		Value           string
		AbtestCpId      []int64
		Feature         []string
		WebGl           int
		ScreenShot      string

		ShowType     uint8
		TemplateType int64
		Score        int64
		Cname        string
		//add for as
		CpId          int64
		DimensionType int
		Induced       int //1: horizontal, 2:vertical, 3:both

		Ext interface{} // 提供给外部包的扩展字段，非通用的。如果是多个模块都需要的扩展字段，建议还是新增字段好一些
	}
)

// ----------------------------------------------- Get Attr ------------------------------------------------------------

//@GetCreative
func (scr *SCreative) GetCreative(formatType int) map[string][]SCreativeAttr {
	switch enum.FormatType(formatType) {
	case enum.FormatTypeImage:
		return scr.Image
	case enum.FormatTypeDynamicImage:
		return scr.DyImage
	case enum.FormatTypeJsTag:
		return scr.JsImage
	case enum.FormatTypeMraid:
		return scr.MraidImage
	case enum.FormatTypeVideo:
		return scr.Video
	case enum.FormatTypeCommon:
		return scr.Common
	default:
		return nil
	}
}

func (scr *SCreative) StoreCreative(cr *SCreativeAttr) {
	switch enum.FormatType(cr.FormatType) {
	case enum.FormatTypeImage:
		scr.Image[cr.StoreKey()] = append(scr.Image[cr.StoreKey()], *cr)
	case enum.FormatTypeDynamicImage:
		scr.DyImage[cr.StoreKey()] = append(scr.DyImage[cr.StoreKey()], *cr)
	case enum.FormatTypeJsTag:
		scr.JsImage[cr.StoreKey()] = append(scr.JsImage[cr.StoreKey()], *cr)
	case enum.FormatTypeMraid:
		scr.MraidImage[cr.StoreKey()] = append(scr.MraidImage[cr.StoreKey()], *cr)
	case enum.FormatTypeVideo:
		scr.Video[cr.StoreKey()] = append(scr.Video[cr.StoreKey()], *cr)
	case enum.FormatTypeCommon:
		scr.Common[cr.StoreKey()] = append(scr.Common[cr.StoreKey()], *cr)
	default:
		scr.Default[cr.StoreKey()] = append(scr.Default[cr.StoreKey()], *cr)
	}
}

func (cr *SCreativeAttr) StoreKey() string {
	switch enum.FormatType(cr.FormatType) {
	case enum.FormatTypeImage, enum.FormatTypeDynamicImage, enum.FormatTypeJsTag, enum.FormatTypeMraid:
		return cr.Resolution
	case enum.FormatTypeVideo:
		return cr.VideoResolution
	case enum.FormatTypeCommon:
		return strconv.Itoa(cr.CreativeType)
	default:
		return strconv.Itoa(cr.CreativeType)
	}
}

// ----------------------------------------------- Parser --------------------------------------------------------------

func (creativeDoc *Creative) Parser() {
	docId := bson.ObjectId(creativeDoc.DocId).Hex()
	creativeDoc.DocId = docId

	creativeDoc.Creative.Default = make(map[string][]SCreativeAttr)
	creativeDoc.Creative.Common = make(map[string][]SCreativeAttr)
	creativeDoc.Creative.Image = make(map[string][]SCreativeAttr)
	creativeDoc.Creative.DyImage = make(map[string][]SCreativeAttr)
	creativeDoc.Creative.JsImage = make(map[string][]SCreativeAttr)
	creativeDoc.Creative.Video = make(map[string][]SCreativeAttr)
	creativeDoc.Creative.MraidImage = make(map[string][]SCreativeAttr)

	for creativeType, sub := range creativeDoc.Content {
		creativeTypeInfo := enum.CreativeTypeMapping(creativeType)
		creativeDoc.parseSCreative(sub, creativeTypeInfo)
	}

	creativeDoc.Content = nil
}

func (creativeDoc *Creative) parseSCreative(mapCreative map[string]map[string]interface{}, cTypeInfo *enum.CreativeTypeInfo) {
	creativeTypeId, _ := strconv.Atoi(cTypeInfo.CreativeType)

	for id, creativeKV := range mapCreative {
		var creative = &SCreativeAttr{
			DocId:           creativeDoc.DocId,
			CountryCode:     creativeDoc.CountryCode,
			PackageName:     creativeDoc.PackageName,
			CreativeType:    creativeTypeId,
			ResourceType:    cTypeInfo.ResourceType,
			SubResourceType: cTypeInfo.SubResourceType,
			FormatType:      cTypeInfo.FormatType,
			CreativeKey:     parseIndexId(id),
			CreativeId:      parseInt64(creativeKV, "creativeId"),
			AdvCreativeId:   parseString(creativeKV, "advCreativeId"),
			Mime:            parseString(creativeKV, "mime"),
			Attribute:       parseString(creativeKV, "attribute"),
			FMd5:            parseString(creativeKV, "fMd5"),
			Source:          int(parseInt64(creativeKV, "source")),
			Ctime:           parseInt64(creativeKV, "ctime"),
			Utime:           parseInt64(creativeKV, "utime"),
			UniqCid:         parseInt64(creativeKV, "uniqCid"),
			VideoResolution: parseVideoResolution(creativeKV),
			VideoSize:       strconv.Itoa(int(parseInt64(creativeKV, "videoSize"))),
			VideoLength:     int(parseInt64(creativeKV, "videoLength")),
			Bitrate:         int(parseInt64(creativeKV, "bitrate")),
			Clarity:         parseClarity(creativeKV, creativeTypeId),
			Width:           int(parseInt64(creativeKV, "width")),
			Height:          int(parseInt64(creativeKV, "height")),
			Orientation:     int(parseInt64(creativeKV, "orientation")),
			MinOs:           parseInt64(creativeKV, "minOs"),
			Platform:        int32(parseInt64(creativeKV, "platform")),
			Resolution:      parseString(creativeKV, "resolution"),
			Url:             parseUrl(creativeKV),
			TagCode:         replaceCreativeDomain(parseString(creativeKV, "tagCode")),
			Value:           parseValue(creativeKV, creativeTypeId),
			AbtestCpId:      parseInt64Slice(creativeKV, "abtestCpId"),
			Feature:         parseStringSlice(creativeKV, "feature"),
			WebGl:           int(parseInt64(creativeKV, "webgl")),
			ScreenShot:      replaceCreativeDomain(parseString(creativeKV, "screenShot")),
			ShowType:        uint8(parseInt64(creativeKV, "showType")),
			TemplateType:    parseInt64(creativeKV, "templateType"),
			Score:           parseInt64(creativeKV, "score"),
			Cname:           parseString(creativeKV, "cname"),
			CpId:            parseInt64(creativeKV, "cpId"),
			DimensionType:   int(parseInt64(creativeKV, "webgl")),
		}

		if creative.CreativeType == 12000 {
			if creative.Width > creative.Height {
				creative.Induced = 1
			} else if creative.Width < creative.Height {
				creative.Induced = 2
			} else {
				creative.Induced = 3
			}
		}

		// default resolution
		if cTypeInfo.Resolution != "" {
			creative.Resolution = cTypeInfo.Resolution
		}

		// default width | height
		index := strings.IndexByte(creative.Resolution, 'x')
		if index != -1 {
			creative.Width, _ = strconv.Atoi(creative.Resolution[0:index])
			creative.Height, _ = strconv.Atoi(creative.Resolution[index+1:])
		}

		// add one
		creativeDoc.Creative.StoreCreative(creative)
	}
}

func parseIndexId(id string) (creativeId int64) {
	indexId, _ := strconv.Atoi(id)
	return int64(indexId)
}

//视频清晰度： 1=>高清 0=>标清 2=>低清。缺省1=>高清
func parseClarity(iface map[string]interface{}, cTypeId int) (clarity int) {
	if value, ok := iface["clarity"]; ok {
		v, _ := assertInter2Int64(value)
		return int(v)
	}
	if cTypeId == 201 || cTypeId == 8201 {
		return 1 // default HD. 视频清晰度： 1=>高清 0=>标清 2=>低清。缺省1=>高清，
	}
	return 0
}

func parseVideoResolution(iface map[string]interface{}) (videoResolution string) {
	if value, ok := iface["videoResolution"]; ok {
		switch vt := value.(type) {
		case string:
			return vt
		default:
		}
	}

	width, widthOk := iface["width"]

	height, heightOk := iface["height"]
	if !widthOk || !heightOk {
		return ""
	} else {
		widthV, _ := assertInter2Int64(width)
		heightV, _ := assertInter2Int64(height)
		return fmt.Sprintf("%dx%d", widthV, heightV)
	}

}

func parseInt64Slice(iface map[string]interface{}, key string) (abtestCpId []int64) {
	if value, ok := iface[key]; ok {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice, reflect.Array:
			s := reflect.ValueOf(value)
			for i := 0; i < s.Len(); i++ {
				v, _ := assertInter2Int64(s.Index(i).Interface())
				abtestCpId = append(abtestCpId, v)
			}
		default:
		}
	}
	return abtestCpId
}

func parseStringSlice(iface map[string]interface{}, key string) (feature []string) {
	if value, ok := iface[key]; ok {
		switch vt := value.(type) {
		case []interface{}:
			for _, str := range vt {
				if cFeature, ok := str.(string); ok {
					feature = append(feature, cFeature)
				}
			}
			return feature
		default:
		}
	}
	return nil
}

func parseUrl(iface map[string]interface{}) (url string) {
	if value, ok := iface["url"]; ok {
		switch vt := value.(type) {
		case string:
			return replaceCreativeDomain(vt)
		default:
		}
	}

	if value, ok := iface["list"]; ok {
		switch vt := value.(type) {
		case []interface{}:
			us := []string{}
			for _, u := range vt {
				us = append(us, u.(string))
			}
			return replaceCreativeDomain(strings.Join(us, "<#>"))
		default:
		}
	}

	return ""
}

func parseValue(iface map[string]interface{}, cTypeId int) (value string) {
	if sub2, ok := iface["value"]; ok {
		switch enum.CreativeType(cTypeId) {
		case enum.CreativeType_APP_RATE:
			val, _ := sub2.(float64)
			return strconv.FormatFloat(val, 'f', -1, 64)
		case enum.CreativeType_NumberRate:
			val, _ := assertInter2Int64(sub2) // 兼容 int32、 int64
			return strconv.FormatInt(val, 10)
		case enum.CreativeType_ICON:
			val, _ := sub2.(string)
			return replaceCreativeDomain(val)
		case enum.CreativeType_CTA_BUTTON: //cta
			val, _ := sub2.(string)
			return strings.TrimSpace(val)
		default:
			value, _ = sub2.(string)
			return value
		}
	}
	return ""
}

func parseString(iface map[string]interface{}, key string) (s string) {
	if value, ok := iface[key]; ok {
		switch vt := value.(type) {
		case string:
			return vt
		default:
		}
	}
	return ""
}

func parseInt64(iface map[string]interface{}, key string) (i int64) {
	if value, ok := iface[key]; ok {
		v, _ := assertInter2Int64(value)
		return v
	}
	return 0
}

// ----------------------------------------------- Base ----------------------------------------------------------------

func replaceCreativeDomain(srcDomain string) (retDomain string) {
	return strings.Replace(srcDomain, "cdn-adn.rayjump.com", "assets.dspunion.com", -1)
}

func assertInter2Int64(inter interface{}) (int64, error) {
	if inter == nil {
		return 0, nil
	}

	switch num := inter.(type) {
	case int64:
		return num, nil
	case int32:
		return int64(num), nil
	case int:
		return int64(num), nil
	case float64:
		return int64(num), nil
	case float32:
		return int64(num), nil
	case string:
		return strconv.ParseInt(num, 10, 64)
	default:
		return 0, errors.New("Invalid Number" + reflect.TypeOf(inter).String())
	}
}
