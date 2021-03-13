package madx

// 如果广告载体是非浏览器应用（通常是移动设备）时应该包含该对象， 网站则不需要包含。
// 一个竞价请求一定不能同时包含Site对象和App对象。 提供一个App标识或者bundle是很有用的， 但是不是严格必须的。
type MApp struct {
	ID         string      `json:"id,omitempty"`         // ID on the exchange
	Name       string      `json:"name,omitempty"`       // Site name (may be aliased at the publisher’s request).
	Domain     string      `json:"domain,omitempty"`     // Domain of the site (e.g., “mysite.foo.com”).
	Cat        []string    `json:"cat,omitempty"`        // Array of IAB content categories
	SectionCat []string    `json:"sectioncat,omitempty"` // Array of IAB content categories for subsection
	PageCat    []string    `json:"pagecat,omitempty"`    // Array of IAB content categories for page
	Publisher  *MPublisher `json:"publisher,omitempty"`  // Details about the Publisher
	Content    *MContent   `json:"content,omitempty"`    // Details about the Content
	Keywords   string      `json:"keywords,omitempty"`   // Comma separated list of keywords about the site.
	Bundle     string      `json:"bundle,omitempty"`     // 应用信息或者包名（例如， com.foo.mygame); 需要是在整个交易过程中唯一的标识
	StoreURL   string      `json:"storeurl,omitempty"`   // 应用的商店地址， 遵循AQG 1.5
	Ver        string      `json:"ver,omitempty"`        // 应用版本号
	Paid       int         `json:"paid,omitempty"`       // 应用是否需要付费， 0表示免费， 1表示付费
	Ext        *MAppExt    `json:"ext,omitempty"`
}

type MAppExt struct {
	DevUserId string `json:"devuserid,omitempty"` // inneractive
}
