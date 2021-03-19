package enum

import "strconv"

const (
	ResourceTypeBasicCreative     uint8 = iota + 1 //1
	ResourceTypeImage                              //2
	ResourceTypeVideo                              //3
	ResourceTypeEndcard                            //4
	ResourceTypeRichMediaPlayable                  //5
	ResourceTypeRichMediaShow                      //6
	ResourceTypeCapture                            //7 新增截屏图
)

type SubResourceType uint8

const (
	SubResourceTypeImage                SubResourceType = iota + 1 //1	static_image		|image
	SubResourceTypeVideo                                           //2	video				|video
	DeprecatedMraidCreative                                        //3	rich_media_show		|mraid
	DeprecatedJsImageCreative                                      //4	static_image		|js
	SubResourceTypeStaticEndCard                                   //5	static_endcard		|image
	SubResourceTypeDynamicImage                                    //6	dynamic_image		|dynamic_image
	SubResourceTypeAppName                                         //7	app_name			|common
	SubResourceTypeAppDesc                                         //8	app_desc			|common
	SubResourceTypeAppRate                                         //9	app_rate			|common
	SubResourceTypeCtaButton                                       //10	cta_button			|common
	SubResourceTypeIcon                                            //11	icon				|common
	SubResourceTypeNormalPlayable                                  //12	normal_playable		|mraid
	SubResourceTypeDynamicEndCard                                  //13	dynamic_endcard		|js
	SubResourceTypeRichMediaShow                                   //14	rich_media_show		|js,mraid
	SubResourceTypeTplPhotoGallery                                 //15	tpl_photogallery	|js
	SubResourceTypeTplAnimation                                    //16	tpl_animation		|js
	SubResourceTypeTplVideo                                        //17	tpl_video			|js
	SubResourceTypeTplFallingObjectGame                            //18	tpl_falling_object_game|mraid
	SubResourceTypeAR                                              //19	AR					|mraid
	SubResourceTypeFullViewGame                                    //20	full_view_game		|mraid
	SubResourceTypeShortTitle                                      //21	short title			|common
	SubResourceTypeNumberRate                                      //22	number rate			|common
	SubResourceTypeCapturePortrait                                 //23	竖屏截屏图			|image
	SubResourceTypeCaptureLandscape                                //24	横屏截屏图			|image
	SubResourceTypeAdWords                                         //25   AdWords           |common
)

var SubResourceType_name = map[SubResourceType]string{
	SubResourceTypeImage:                "StaticImage",
	SubResourceTypeVideo:                "Video",
	DeprecatedMraidCreative:             "MraidImage",
	DeprecatedJsImageCreative:           "JsImage",
	SubResourceTypeStaticEndCard:        "StaticEndCard",
	SubResourceTypeDynamicImage:         "DynamicImage",
	SubResourceTypeAppName:              "AppName",
	SubResourceTypeAppDesc:              "Desc",
	SubResourceTypeAppRate:              "Rate",
	SubResourceTypeCtaButton:            "CTA",
	SubResourceTypeIcon:                 "Icon",
	SubResourceTypeNormalPlayable:       "NormalPlayable",
	SubResourceTypeDynamicEndCard:       "DynamicEndCard",
	SubResourceTypeRichMediaShow:        "RichMediaShow",
	SubResourceTypeTplPhotoGallery:      "TplPhotoCallery",
	SubResourceTypeTplAnimation:         "TplAnimation",
	SubResourceTypeTplVideo:             "TplVideo",
	SubResourceTypeTplFallingObjectGame: "TplFallingObjectGame",
	SubResourceTypeAR:                   "AR",
	SubResourceTypeFullViewGame:         "FullViewGame",
	SubResourceTypeShortTitle:           "ShortTitle",
	SubResourceTypeNumberRate:           "NumberRate",
	SubResourceTypeCapturePortrait:      "ProtraitCapture",
	SubResourceTypeCaptureLandscape:     "LandscapeCapture",
	SubResourceTypeAdWords:              "AdWords",
}

var SubResourceType_value = map[string]SubResourceType{
	"StaticImage":          SubResourceTypeImage,
	"Video":                SubResourceTypeVideo,
	"MraidImage":           DeprecatedMraidCreative,
	"JsImage":              DeprecatedJsImageCreative,
	"StaticEndCard":        SubResourceTypeStaticEndCard,
	"DynamicImage":         SubResourceTypeDynamicImage,
	"AppName":              SubResourceTypeAppName,
	"Desc":                 SubResourceTypeAppDesc,
	"Rate":                 SubResourceTypeAppRate,
	"CTA":                  SubResourceTypeCtaButton,
	"Icon":                 SubResourceTypeIcon,
	"NormalPlayable":       SubResourceTypeNormalPlayable,
	"DynamicEndCard":       SubResourceTypeDynamicEndCard,
	"RichMediaShow":        SubResourceTypeRichMediaShow,
	"TplPhotoCallery":      SubResourceTypeTplPhotoGallery,
	"TplAnimation":         SubResourceTypeTplAnimation,
	"TplVideo":             SubResourceTypeTplVideo,
	"TplFallingObjectGame": SubResourceTypeTplFallingObjectGame,
	"AR":                   SubResourceTypeAR,
	"FullViewGame":         SubResourceTypeFullViewGame,
	"ShortTitle":           SubResourceTypeShortTitle,
	"NumberRate":           SubResourceTypeNumberRate,
	"ProtraitCapture":      SubResourceTypeCapturePortrait,
	"LandscapeCapture":     SubResourceTypeCaptureLandscape,
	"AdWords":              SubResourceTypeAdWords,
}

func (x SubResourceType) String() string {
	return SubResourceType_name[x]
}

func (x SubResourceType) Enum() uint8 {
	return uint8(x)
}

func SubResourceTypeEnum(value string) SubResourceType {
	return SubResourceType_value[value]
}

type FormatType uint8

const (
	FormatTypeImage        FormatType = iota + 1 // 1
	FormatTypeDynamicImage                       // 2
	FormatTypeVideo                              // 3
	FormatTypeJsTag                              // 4
	FormatTypeMraid                              // 5 crType = mraid playable
	FormatTypeCommon                             // 6
	TencentGroupCreative                         // 7 分组素材。
)

var FormatType_name = map[FormatType]string{
	FormatTypeImage:        "Image",
	FormatTypeDynamicImage: "DyImage",
	FormatTypeVideo:        "Video",
	FormatTypeJsTag:        "JsImage",
	FormatTypeMraid:        "MraidImage",
	FormatTypeCommon:       "Common",
	TencentGroupCreative:   "GroupCreative",
}

var FormatType_value = map[string]FormatType{
	"Image":         FormatTypeImage,
	"DyImage":       FormatTypeDynamicImage,
	"Video":         FormatTypeVideo,
	"JsImage":       FormatTypeJsTag,
	"MraidImage":    FormatTypeMraid,
	"Common":        FormatTypeCommon,
	"GroupCreative": TencentGroupCreative,
}

func (x FormatType) String() string {
	return FormatType_name[x]
}

func (x FormatType) Enum() uint8 {
	return uint8(x)
}

func FormatTypeEnum(value string) FormatType {
	return FormatType_value[value]
}

//********************* CreativeType *********************
// 下表罗列了 基础素材的 定义
type CreativeType int

const (
	CreativeType_APP_NAME   CreativeType = 401
	CreativeType_APP_DESC   CreativeType = 402
	CreativeType_APP_RATE   CreativeType = 403
	CreativeType_CTA_BUTTON CreativeType = 404
	CreativeType_ICON       CreativeType = 405
	CreativeType_NumberRate CreativeType = 406
	CreativeType_ShortTitle CreativeType = 407
	CreativeType_AD_WORDS   CreativeType = 409
)

func (x CreativeType) String() string {
	return strconv.Itoa(int(x))
}

//********************* Template Click Area *********************
// 下表罗列了Polaris召回模版元素的前缀。
// Polaris、Pioneer共同约定。Polaris按照约定为每个元素添加固定前缀，Pioneer按照约定解析前缀后使用（关联CDN）
const (
	ElementPrefix_Image                = "Image_"
	ElementPrefix_DyImage              = "Gif_"
	ElementPrefix_Video                = "Video_"
	ElementPrefix_NonPlayableRichMedia = "NonPlayable_"
	ElementPrefix_PlayableRichMedia    = "Playable_"
	ElementPrefix_APPName              = "Name_"
	ElementPrefix_APPDesc              = "Desc_"
	ElementPrefix_APPRate              = "Rate_"
	ElementPrefix_CtaButton            = "Cta_"
	ElementPrefix_Icon                 = "Icon_"
	ElementPrefix_NumberRating         = "Comment_"
	ElementPrefix_ShortTitle           = "Title_"
	ElementPrefix_QrCode               = "QrCode_"
)

func CreativeTypeMapping(crtype string) *CreativeTypeInfo {
	return nil
}

type CreativeTypeInfo struct {
	CreativeType string

	ResourceType    int
	SubResourceType int
	FormatType      int

	Resolution string
}
