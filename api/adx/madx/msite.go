package madx

// 如果广告载体是一个网站时应该包含这个对象，如果是非浏览器应用时则不需要。
// 一个竞价请求一定不能同时包含Site对象和App对象。
// 提供一个站点标识或者页面地址是很有用的， 但是不是严格必须的
type MSite struct {
	ID            string      `json:"id,omitempty"`            // ID on the exchange
	Name          string      `json:"name,omitempty"`          // Site name (may be aliased at the publisher’s request).
	Domain        string      `json:"domain,omitempty"`        // Domain of the site (e.g., “mysite.foo.com”).
	Cat           []string    `json:"cat,omitempty"`           // Array of IAB content categories
	SectionCat    []string    `json:"sectioncat,omitempty"`    // Array of IAB content categories for subsection
	PageCat       []string    `json:"pagecat,omitempty"`       // Array of IAB content categories for page
	PrivacyPolicy int         `json:"privacypolicy,omitempty"` // Default: 1 ("1": has a privacy policy)
	Publisher     *MPublisher `json:"publisher,omitempty"`     // Details about the Publisher
	Content       *MContent   `json:"content,omitempty"`       // Details about the Content
	Keywords      string      `json:"keywords,omitempty"`      // Comma separated list of keywords about the site.
	Page          string      `json:"page,omitempty"`          // 展示广告将要被展示的页面地址
	Ref           string      `json:"ref,omitempty"`           // 引导到当前页面的referrer地址
	Search        string      `json:"search,omitempty"`        // 引导到当前页面的搜索字符串
	Mobile        int         `json:"mobile,omitempty"`        // 移动优化标志， 0表示否，1表示是
}
