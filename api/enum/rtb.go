package enum

//****************** Table 5.2 Banner Ad Type *********************
// Banner Ad TypesThe following table indicates the types of ads that can be accepted by the exchange unless restricted by
// publisher site settings.
// 下表展示了除被展示者站点设置限制的交易平台可以接受的广告类型。
const (
	BannerAdType_XHTML_Text_Ad   = 1 //XHTML 文本广告（通常用于移动手机）
	BannerAdType_XHTML_Banner_Ad = 2 //XHTML Banner 广告 （通常用于移动手机）
	BannerAdType_JavaScript_Ad   = 3 //javascript 广告， 必须是合法的XHTML(即用script标签包裹)
	BannerAdType_Iframe          = 4 //iframe广告
)

//****************** Table 5.3 Creative Attribute *********************
// The following table specifies a standard list of creative attributes that can describe an ad being served or
// serve as restrictions of thereof.
// 下表展示了广告物料属性的标准列表， 用于描述一个被提供的广告或者作为限制提供。
const (
	Attribute_UNKNOWN                        = 0
	Attribute_AUDIO_AUTO_PLAY                = 1  //音频广告（自动播放）
	Attribute_AUDIO_USER_INITIATED           = 2  //音频广告 （用户触发）
	Attribute_EXPANDABLE_AUTOMATIC           = 3  //可伸展 （自动）
	Attribute_EXPANDABLE_CLICK_INITIATED     = 4  //可伸展 （用户点击触发）
	Attribute_EXPANDABLE_ROLLOVER_INITIATED  = 5  //可伸展 （用户翻转触发）
	Attribute_VIDEO_IN_BANNER_AUTO_PLAY      = 6  //Banner内视频广告 （自动播放）
	Attribute_VIDEO_IN_BANNER_USER_INITIATED = 7  //Banner内视频广告 （用户触发）
	Attribute_POP                            = 8  //弹出广告 （例如， 在退出之前，之后或者当退出的时候）
	Attribute_PROVOCATIVE_OR_SUGGESTIVE      = 9  //Provocative or Suggestive Imagery
	Attribute_ANNOYING                       = 10 // Defined as "Shaky, Flashing, Flickering, Extreme Animation, Smileys".
	Attribute_SURVEYS                        = 11 //调查问卷
	Attribute_TEXT_ONLY                      = 12 //文本广告
	Attribute_USER_INTERACTIVE               = 13 //用户互动 （比如，嵌入式游戏）
	Attribute_WINDOWS_DIALOG_OR_ALERT_STYLE  = 14 //窗口对话框或弹出提示框
	Attribute_HAS_AUDIO_ON_OFF_BUTTON        = 15 //具有音频开关按钮
	Attribute_AD_CAN_BE_SKIPPED              = 16 //广告可以被跳过（例如，片头广告视频的跳过按钮）
	Attribute_FLASH                          = 17
)

//****************** Table 5.6 API Framework *********************
// The following table is a list of API frameworks supported by the publisher.
// 下表罗列了展示者支持的API框架。
// 注意MRAID-1是MRAID-2的子集。在OpenRTB2.1及之前版本，“3”对应的是"MRAID", 然而并不是所有MARID兼容的API可以理解MRAID-2的特性，
// 所以最安全的方式是将3理解为MRAID-1.在OpenRTB2.2中， 该值已经被显式标明了， MRAID-2也使用值5被加入。
const (
	Api_UNKNOWN = 0
	Api_VPAID_1 = 1 // [AdX: attribute 30/VPAID]
	Api_VPAID_2 = 2 // [AdX: attribute 30/VPAID]
	Api_MRAID_1 = 3 // [AdX: attribute 32/MRAID]
	Api_ORMMA   = 4
	Api_MRAID_2 = 5 // [AdX: attribute 32/MRAID (OpenRTB 2.3+)]
	Api_MRAID_3 = 6
)

//****************** Table 5.14 Companion Types *********************
// The following table lists the options to indicate markup types allowed for companionadsthat apply to
// video and audio ads.  This table is derived from VAST 2.0+and DAAST 1.0 specifications.  Refer to
// www.iab.com/guidelines/digital-video-suitefor more information.
// 下表罗列了允许的视频伴随广告的类型。表格继承自IAB VAST2.0+。
const (
	CompanionTypes_Static = 1
	CompanionTypes_HTML   = 2
	CompanionTypes_Iframe = 3
)

//****************** Table 5.21 Device Type *********************
// The following table lists the type of device from which the impression originated.
// OpenRTB version 2.2 of the specificationadded distinct values for Mobile and Tablet.  It is
// recommended that any bidder adding support for 2.2 treat a value of 1 as an acceptable alias of 4 & 5.
// This OpenRTB table has values derived from the InventoryQuality Guidelines (IQG). Practitioners should
// keep in syncwith updates to the IQG values.
// 下表罗列了广告展示所在的设备类型。
// OpenRTB规范版本2.2区分了手机和平板， 建议任何竞价者添加对2.2的支持， 可以将1等价4&5解析。 该表继承自QAG,实现者应当与QAG的相关描述保持一致。
const ()

//****************** Table 5.22 Connection Type *********************
// The following table lists the various options for the type of device connectivity.
// 下表罗列了设备网络连接方式的各种选项。
const (
	ConnectionType_Unknown                 = 0
	ConnectionType_Ethernet                = 1 //以太网
	ConnectionType_WIFI                    = 2 //WIFI
	ConnectionType_CellularNetwork_Unknown = 3 //移动网络，未知类型
	ConnectionType_CellularNetwork_2G      = 4 //移动网络-2G
	ConnectionType_CellularNetwork_3G      = 5 //移动网络-3G
	ConnectionType_CellularNetwork_4G      = 6 //移动网络-4G
)

//****************** Table 7.6 Data Asset Types *********************
// 以下列出了在发布时原生广告的常见资产元素类型编写此规范。 此清单并非详尽无遗，打算由买家进行扩展和卖方随着格式的发展而变化。
// 实施中的交易所可能不支持所有资产变体或引入新的资产变体该系统独有。
const (
	_                 = iota
	Native_Sponsored  = 1  //消息赞助，回应中应包含品牌 保荐人名称。推荐最大25长度
	Native_Desc       = 2  //相关的描述性文字宣传的产品或服务。回复中的文字更长可能会被截断或省略交换。
	Native_Rating     = 3  //所提供产品的评级给用户。 例如，某个应用的在应用商店中的评分为0-5。
	Native_Likes      = 4  //社会评价或“赞”次数
	Native_Downloads  = 5  //此产品的下载/安装数量
	Native_Price      = 6  //产品/应用/应用内价格采购。 值应包括本地化格式的货币符号
	Native_Saleprice  = 7  //可以一起使用的销售价格价格表明打折价格与正常价格相比。价值应包括货币本地化格式的符号。
	Native_Phone      = 8  //Phone number
	Native_Address    = 9  //Address
	Native_Desc2      = 10 //关联的其他描述性文字
	Native_Displayurl = 11 //文字广告的显示网址。
	Native_Ctatext    = 12 //CTA说明-描述性文字。推荐最大15长度
)
