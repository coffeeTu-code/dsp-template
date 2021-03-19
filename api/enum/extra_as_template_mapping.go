package enum

var VideoTemplateTypeToVideoTemplateId map[int][]string
var EndcardTemplateTypeToEndcardTemplateId map[int][]string
var GroupMapping map[int][]int

func init() {
	VideoTemplateTypeToVideoTemplateId = make(map[int][]string)
	VideoTemplateTypeToVideoTemplateId[-1] = []string{"0"}
	VideoTemplateTypeToVideoTemplateId[1] = []string{"5002001__Low", "5002001__High"}
	VideoTemplateTypeToVideoTemplateId[2] = []string{"5002002__Low", "5002002__High"}
	VideoTemplateTypeToVideoTemplateId[3] = []string{"5002003__Low", "5002003__High"}
	VideoTemplateTypeToVideoTemplateId[4] = []string{"5002004__Low", "5002004__High"}
	VideoTemplateTypeToVideoTemplateId[7] = []string{"5002005__Low", "5002005__High"}
	VideoTemplateTypeToVideoTemplateId[8] = []string{"5002006__Low", "5002006__High"}
	VideoTemplateTypeToVideoTemplateId[9] = []string{"5002007__Low", "5002007__High"}
	VideoTemplateTypeToVideoTemplateId[11] = []string{"5002008__Low", "5002008__High"}

	EndcardTemplateTypeToEndcardTemplateId = make(map[int][]string)
	EndcardTemplateTypeToEndcardTemplateId[-2] = []string{"6004001"}
	EndcardTemplateTypeToEndcardTemplateId[-3] = []string{"6003001"}
	EndcardTemplateTypeToEndcardTemplateId[10] = []string{"6002004"}
	EndcardTemplateTypeToEndcardTemplateId[4] = []string{"6002001"}
	EndcardTemplateTypeToEndcardTemplateId[7] = []string{"6002002"}
	EndcardTemplateTypeToEndcardTemplateId[14] = []string{"6002007"}
	EndcardTemplateTypeToEndcardTemplateId[8] = []string{"6002005"}
	EndcardTemplateTypeToEndcardTemplateId[9] = []string{"6002006"}
	EndcardTemplateTypeToEndcardTemplateId[12] = []string{"6002003"}

	GroupMapping = make(map[int][]int)
	GroupMapping[1] = []int{18}
	GroupMapping[2] = []int{19}
	GroupMapping[3] = []int{20}
	GroupMapping[4] = []int{21}
	GroupMapping[5] = []int{22}
	GroupMapping[6] = []int{28}
	GroupMapping[7] = []int{29}
	GroupMapping[8] = []int{23}
	GroupMapping[9] = []int{24}
	GroupMapping[10] = []int{25}
	GroupMapping[11] = []int{27}
	GroupMapping[12] = []int{26}

}

var AsCampaignDefaultVideoTemplateSet = map[string]bool{
	"5002001__Low": true, "5002001__High": true,
	"5002002__Low": true, "5002002__High": true,
	"5002003__Low": true, "5002003__High": true,
	"5002004__Low": true, "5002004__High": true,
	"5002007__Low": true, "5002007__High": true,
	"5002009__Low": true, "5002009__High": true,
	"5002008__Low": true, "5002008__High": true,
}

var AsCampaignDefaultEndcardTemplateSet = map[string]bool{
	"6004001": true, "6003001": true,
	"6002001": true, "6002002": true,
	"6002003": true, "6002007": true,
}

var AsDefaultTemplateGroupMap = map[int]bool{
	1: true, 2: true, 3: true,
	6: true, 7: true, 18: true,
	19: true, 20: true, 21: true,
}
