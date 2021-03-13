package madx

import "encoding/xml"

// Extension represent arbitrary XML provided by the platform to extend the
// VAST response or by custom trackers.
type Extension struct {
	Type           string     `xml:"type,attr,omitempty" json:",omitempty"`
	Asset          []Asset    `xml:"Asset,omitempty" json:",omitempty"`
	CustomTracking []Tracking `xml:"CustomTracking>Tracking,omitempty" json:",omitempty"`
	Data           []byte     `xml:",innerxml" json:",omitempty"`

	Template                *Template                `xml:"Template,omitempty" json:",omitempty"` // Template 里是旧逻辑， Templates+Asset是新逻辑
	Templates               []Templates              `xml:"Templates,omitempty" json:",omitempty"`
	VideoEndType            *VideoEndType            `xml:"VideoEndType,omitempty" json:",omitempty"`
	PlayableAdsWithoutVideo *PlayableAdsWithoutVideo `xml:"PlayableAdsWithoutVideo,omitempty" json:",omitempty"`
	Orientation             *Orientation             `xml:"Orientation,omitempty" json:",omitempty"`
	Fcb                     *Fcb                     `xml:"Fcb,omitempty" json:",omitempty"`
	Verification            *Verification            `xml:"AdVerifications>Verification,omitempty" json:",omitempty"`
}

// the extension type as a middleware in the encoding process.
type extension Extension

type extensionNoCT struct {
	Type  string  `xml:"type,attr,omitempty" json:",omitempty"`
	Asset []Asset `xml:"Asset,omitempty" json:",omitempty"`
	Data  []byte  `xml:",innerxml" json:",omitempty"`

	Template                *Template                `xml:"Template,omitempty" json:",omitempty"` // Template 里是旧逻辑， Templates+Asset是新逻辑
	Templates               []Templates              `xml:"Templates,omitempty" json:",omitempty"`
	VideoEndType            *VideoEndType            `xml:"VideoEndType,omitempty" json:",omitempty"`
	PlayableAdsWithoutVideo *PlayableAdsWithoutVideo `xml:"PlayableAdsWithoutVideo,omitempty" json:",omitempty"`
	Orientation             *Orientation             `xml:"Orientation,omitempty" json:",omitempty"`
	Fcb                     *Fcb                     `xml:"Fcb,omitempty" json:",omitempty"`
	Verification            *Verification            `xml:"AdVerifications>Verification,omitempty" json:",omitempty"`
}

type Asset struct {
	AssetType    string `xml:"type,attr,omitempty" json:",omitempty"` // 取值： starrating, numberrating, CTA, icon
	Height       string `xml:"height,attr,omitempty" json:",omitempty"`
	Width        string `xml:"width,attr,omitempty" json:",omitempty"`
	CreativeType string `xml:"creativeType,attr,omitempty" json:",omitempty"`
	Value        string `xml:",chardata" json:",omitempty"`
	Cdata        string `xml:",cdata" json:",omitempty"`
}

// MarshalXML implements xml.Marshaler interface.
func (e Extension) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	// create a temporary element from a wrapper Extension, copy what we need to
	// it and return it's encoding.
	var e2 interface{}
	// if we have custom trackers, we should ignore the data, if not, then we
	// should consider only the data.
	if len(e.CustomTracking) > 0 {
		e2 = extension{
			Type: e.Type, CustomTracking: e.CustomTracking, Asset: e.Asset,
			Templates: e.Templates, Template: e.Template, VideoEndType: e.VideoEndType,
			PlayableAdsWithoutVideo: e.PlayableAdsWithoutVideo, Orientation: e.Orientation,
			Fcb: e.Fcb, Verification: e.Verification,
		}
	} else {
		e2 = extensionNoCT{Type: e.Type, Data: e.Data, Asset: e.Asset,
			Templates: e.Templates, Template: e.Template, VideoEndType: e.VideoEndType,
			PlayableAdsWithoutVideo: e.PlayableAdsWithoutVideo, Orientation: e.Orientation,
			Fcb: e.Fcb, Verification: e.Verification,
		}
	}

	return enc.EncodeElement(e2, start)
}

// UnmarshalXML implements xml.Unmarshaler interface.
func (e *Extension) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	// decode the extension into a temporary element from a wrapper Extension,
	// copy what we need over.
	var e2 extension
	if err := dec.DecodeElement(&e2, &start); err != nil {
		return err
	}
	// copy the type and the customTracking
	e.Type = e2.Type
	e.CustomTracking = e2.CustomTracking
	e.Asset = e2.Asset
	// copy the data only of customTracking is empty
	if len(e.CustomTracking) == 0 {
		e.Data = e2.Data
	}

	e.Template = e2.Template
	e.Templates = e2.Templates
	e.VideoEndType = e2.VideoEndType
	e.PlayableAdsWithoutVideo = e2.PlayableAdsWithoutVideo
	e.Orientation = e2.Orientation
	e.Fcb = e2.Fcb
	e.Verification = e2.Verification
	return nil
}

type Template struct {
	Id    string  `xml:"id"`
	Asset []Asset `xml:"Asset,omitempty" json:",omitempty"`
}

type Templates struct {
	Id    string  `xml:"id,attr,omitempty" json:",omitempty"`
	Type  string  `xml:"type,attr,omitempty" json:",omitempty"`
	Asset []Asset `xml:"Asset,omitempty" json:",omitempty"`
	Cdata string  `xml:",cdata" json:",omitempty"`
}

type VideoEndType struct {
	Value string `xml:",chardata" json:",omitempty"`
}
type PlayableAdsWithoutVideo struct {
	Value string `xml:",chardata"`
}
type Orientation struct {
	Value string `xml:",chardata"`
}
type Fcb struct {
	Value string `xml:",chardata"`
}
type Verification struct {
	Vendor                 string              `xml:"vendor,attr,omitempty"`
	JavaScriptResource     *JavaScriptResource `xml:"JavaScriptResource,omitempty"`
	VerificationParameters *VerificationParameters
	// pokkt 的vast中还有tracking Event,用不上，暂不实现
}

type JavaScriptResource struct {
	ApiFramework string `xml:"apiFramework,attr,omitempty"`
	URI          string `xml:",cdata"`
}

type VerificationParameters struct {
	Cdata string `xml:",cdata"`
}
