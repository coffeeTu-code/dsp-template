package base

type Template struct {
	ID     string          `json:"id"`     //模板id
	Parent map[string]bool `json:"parent"` //父模版ID，当存在父模版时，优先使用父模版

	// element list
	Elements                    []string            `json:"element"`                  //base element
	OptionalElements            []string            `json:"optional_elements"`        //可选的展示element
	CreativeCombination         map[string][]string `json:"creative_combination"`     //强制要求element，有一个即可，模板有素材组合随机选逻辑
	ElementsRecalledPriorityArr []map[string]int    `json:"elements_recall_priority"` //需要按照优先级进行召回的元素集合,key 元素名，val 为 按优先级从大到小排列的slice

	// html
	FrontID      string            `json:"front,omitempty"`       //前端代码id， html或js等
	Placeholders map[string]string `json:"placeholder,omitempty"` //替换宏， key=召回的元素，val=前段占位符

	// attr
	HasVideoCreative   bool    `json:"has_video_creative"` //包含video素材，认为是视频模版，有特殊的过滤条件
	TemplateType       int32   `json:"template_type"`
	ApiFramework       int32   `json:"api_framework"`
	Orientation        int32   `json:"orientation"`
	TemplateScreenSize int32   `json:"template_screen_size"`
	ClickArea          []int32 `json:"click_area"`
	CrType             string  `json:"crtype"`

	// target
	CampaignTarget map[string]bool `json:"campaigns"`
}
