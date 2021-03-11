package mortb

import (
	"encoding/xml"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	extensionCustomTracking = []byte(`<Extension type="testCustomTracking"><Asset type="CTA">Install</Asset><Asset type="Icon"><![CDATA[http://icon.com/icon.jpg]]></Asset><CustomTracking><Tracking event="event.1"><![CDATA[http://event.1]]></Tracking><Tracking event="event.2"><![CDATA[http://event.2]]></Tracking></CustomTracking></Extension>`)
	extensionData           = []byte(`<Extension type="testCustomTracking"><SkippableAdType>Generic</SkippableAdType></Extension>`)
)

func TestExtensionCustomTrackingMarshal(t *testing.T) {
	Convey("ExtensionCustomTrackingMarshal", t, func() {
		e := Extension{
			Asset: []Asset{
				{
					AssetType: "CTA",
					Value:     "Install",
				},
				{
					AssetType: "Icon",
					Cdata:     "http://icon.com/icon.jpg",
				},
			},
			Type: "testCustomTracking",
			CustomTracking: []Tracking{
				{
					Event: "event.1",
					URI:   "http://event.1",
				},
				{
					Event: "event.2",
					URI:   "http://event.2",
				},
			},
		}

		// marshal the extension
		xmlExtensionOutput, err := xml.Marshal(e)
		So(err, ShouldBeNil)

		// assert the resulting marshaled extension
		So(string(extensionCustomTracking), ShouldEqual, string(xmlExtensionOutput))
	})

}

func TestExtensionCustomTracking(t *testing.T) {
	Convey("ExtensionCustomTracking", t, func() {
		// unmarshal the Extension
		var e Extension
		So(xml.Unmarshal(extensionCustomTracking, &e), ShouldBeNil)

		// assert the resulting extension
		So("testCustomTracking", ShouldEqual, e.Type)
		So(string(e.Data), ShouldBeEmpty)
		So(len(e.CustomTracking), ShouldEqual, 2)
		// first event
		So("event.1", ShouldEqual, e.CustomTracking[0].Event)
		So("http://event.1", ShouldEqual, e.CustomTracking[0].URI)
		// second event
		So("event.2", ShouldEqual, e.CustomTracking[1].Event)
		So("http://event.2", ShouldEqual, e.CustomTracking[1].URI)

		//// marshal the extension
		//xmlExtensionOutput, err := xml.Marshal(e)
		//So(err, ShouldBeNil)
		//// assert the resulting marshaled extension
		//So(string(extensionCustomTracking), ShouldEqual, string(xmlExtensionOutput))
	})

}

func TestExtensionGeneric(t *testing.T) {
	Convey("ExtensionGeneric", t, func() {
		// unmarshal the Extension
		var e Extension
		So(xml.Unmarshal(extensionData, &e), ShouldBeNil)

		// assert the resulting extension
		So("testCustomTracking", ShouldEqual, e.Type)
		So("<SkippableAdType>Generic</SkippableAdType>", ShouldEqual, string(e.Data))
		So(e.CustomTracking, ShouldBeEmpty)

		// marshal the extension
		xmlExtensionOutput, err := xml.Marshal(e)
		So(err, ShouldBeNil)

		// assert the resulting marshaled extension
		So(string(extensionData), ShouldEqual, string(xmlExtensionOutput))
	})
}

func TestExtensionTemplate(t *testing.T) {
	Convey("ExtensionTemplate", t, func() {
		e := Extension{
			Type: "t",
			Asset: []Asset{
				{
					AssetType:    "at",
					Height:       "h",
					Width:        "w",
					CreativeType: "cr",
					Value:        "cd",
				},
			},
			Templates: []Templates{
				{
					Id:   "i",
					Type: "t",
					Asset: []Asset{
						{
							AssetType:    "aat",
							Height:       "h",
							Width:        "w",
							CreativeType: "cre",
							Value:        "cd",
						},
					},
					Cdata: "",
				},
			},
			Template: &Template{
				Id: "id",
				Asset: []Asset{
					{
						AssetType:    "aat",
						Height:       "h",
						Width:        "w",
						CreativeType: "cre",
						Value:        "cd",
					},
				},
			},
			VideoEndType:            &VideoEndType{"vet"},
			PlayableAdsWithoutVideo: &PlayableAdsWithoutVideo{"pawv"},
			Orientation:             &Orientation{Value: "o"},
			Fcb:                     &Fcb{Value: "fcb"},
			Verification: &Verification{
				Vendor: "v",
			},
		}
		b, _ := xml.Marshal(&e)
		newExtension := new(Extension)
		err := xml.Unmarshal(b, newExtension)
		So(err, ShouldBeNil)
		So(newExtension.Asset, ShouldResemble, e.Asset)
		So(newExtension.Template, ShouldResemble, e.Template)
		So(newExtension.Templates, ShouldResemble, e.Templates)
		So(newExtension.Orientation, ShouldResemble, e.Orientation)
		So(newExtension.PlayableAdsWithoutVideo, ShouldResemble, e.PlayableAdsWithoutVideo)
		So(newExtension.Fcb, ShouldResemble, e.Fcb)
		So(newExtension.Verification, ShouldResemble, e.Verification)
	})
}
