package mgotables

type App struct {
	AppId             int64                   `bson:"appId,omitempty" json:"appId"`
	PublisherId       int64                   `bson:"publisherId,omitempty" json:"publisherId"`
	App               AppInfo                 `bson:"app,omitempty" json:"app"`
	Setting           AppSettingDetail        `bson:"setting,omitempty" json:"setting"`
	DirectOfferConfig DirectOfferConfigDetail `bson:"directOfferConfig,omitempty" json:"directOfferConfig"`
	BlackPackageList  []string                `bson:"blackPackageList,omitempty" json:"blackPackageList"`
	BlackCategoryList []int64                 `bson:"blackCategoryList,omitempty" json:"blackCategoryList"`
	Updated           int64                   `bson:"updated,omitempty" json:"updated"`
}

type AppInfo struct {
	Grade               int64               `bson:"grade,omitempty" json:"grade"`
	Name                string              `bson:"name,omitempty" json:"name"`
	DirectMarket        int64               `bson:"directMarket,omitempty" json:"directMarket"`
	CampaignType        int64               `bson:"campaignType,omitempty" json:"campaignType"`
	Status              int8                `bson:"status,omitempty" json:"status"`
	IsIncent            int64               `bson:"isIncent,omitempty" json:"isIncent"`
	BtClass             int64               `bson:"btClass,omitempty" json:"btClass"`
	OfferPreference     []int64             `bson:"offerPreference,omitempty" json:"offerPreference"`
	ContentPreference   int64               `bson:"contentPreference,omitempty" json:"contentPreference"`
	TrafficType         int64               `bson:"trafficType,omitempty" json:"trafficType"`
	IabCategoryV2       map[string][]string `bson:"iabCategoryV2,omitempty" json:"iabCategoryV2"`
	SubCategoryV2       []string            `bson:"subCategoryV2,omitempty" json:"subCategoryV2"`
	Coppa               int64               `bson:"coppa,omitempty" json:"coppa"`
	SiteIntegrationType int64               `bson:"siteIntegrationType,omitempty" json:"siteIntegrationType"`
	Alac                int64               `bson:"alac,omitempty" json:"alac"`
}

type AppSettingDetail struct {
	StoreKit int64 `bson:"storekit,omitempty" json:"storekit"`
}

type DirectOfferConfigDetail struct {
	Status  int8    `bson:"status,omitempty" json:"status"`
	Type    int64   `bson:"type,omitempty" json:"type"`
	OfferId []int64 `bson:"offerId,omitempty" json:"offerId"`
}



// ----------------------------------------------- Parser --------------------------------------------------------------

func (this *App) Parser() {
	return
}