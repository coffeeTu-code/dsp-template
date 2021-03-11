package mortb

import (
	"encoding/xml"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMediaFile_UnmarshalXML(t *testing.T) {
	Convey("MediaFile_UmarshalXML", t, func() {
		mf := MediaFile{
			ID:                  "id",
			Delivery:            "d",
			Type:                "t",
			Codec:               "cc",
			Bitrate:             10,
			MinBitrate:          12,
			MaxBitrate:          30,
			Width:               40,
			Height:              50,
			Scalable:            true,
			MaintainAspectRatio: true,
			APIFramework:        "a",
			URI:                 "http://",
			Size:                200,
		}
		bb, _ := xml.Marshal(&mf)
		t.Log(string(bb))
		mRet := new(MediaFile)
		err := xml.Unmarshal(bb, mRet)
		So(err, ShouldBeNil)
		So(*mRet, ShouldResemble, mf)
		t.Logf("ret:%+v\n", mRet)
	})
}

func TestVideoClick_UnmarshalXML(t *testing.T) {
	Convey("VideoClick_UnmarshalXML", t, func() {
		v := VideoClick{
			ID:  "aaa",
			URI: "http://",
		}
		bb, _ := xml.Marshal(&v)
		mRet := new(VideoClick)
		err := xml.Unmarshal(bb, mRet)
		So(err, ShouldBeNil)
		So(*mRet, ShouldResemble, v)
	})
}

func TestVideoClicks_UnmarshalXML(t *testing.T) {
	Convey("VideoClicks_UnmarshalXML", t, func() {
		v := VideoClicks{
			ClickThroughs: []VideoClick{{
				ID:  "ct1",
				URI: "http:ct1",
			}, {
				ID:  "ct2",
				URI: "http://ct2",
			}},
			ClickTrackings: []VideoClick{
				{
					ID:  "ctk1",
					URI: "http://ctk1",
				},
			},
			CustomClicks: []VideoClick{
				{
					ID:  "cc1",
					URI: "http://cc1",
				},
				{
					ID:  "cc2",
					URI: "http://cc2.com?a=b",
				},
			},
		}
		bb, _ := xml.Marshal(&v)
		mRet := new(VideoClicks)
		err := xml.Unmarshal(bb, mRet)
		So(err, ShouldBeNil)
		So(*mRet, ShouldResemble, v)
	})
}

func TestTracking_UnmarshalXML(t *testing.T) {
	Convey("Tracking_UnmarshalXML", t, func() {
		duration := Duration(1000000)
		tk := Tracking{
			Event: "E",
			Offset: &Offset{
				Duration: &duration,
				Percent:  0,
			},
			URI: "http://tk",
		}
		bb, _ := xml.Marshal(&tk)
		mRet := new(Tracking)
		err := xml.Unmarshal(bb, mRet)
		So(err, ShouldBeNil)
		So(*mRet, ShouldResemble, tk)
	})
}

func TestCompanion_UnmarshalXML(t *testing.T) {
	Convey("Companion_UnmarshalXML", t, func() {
		duration1 := Duration(10000000)
		duration2 := Duration(20000000)
		cn := Companion{
			ID:                     "a",
			Width:                  20,
			Height:                 30,
			AssetWidth:             40,
			AssetHeight:            50,
			ExpandedWidth:          60,
			ExpandedHeight:         70,
			APIFramework:           "af",
			AdSlotID:               "asi",
			CompanionClickThrough:  CDATAString{CDATA: "ccth"},
			CompanionClickTracking: []CDATAString{{"cct1"}, {"cct2"}},
			AltText:                "altt",
			TrackingEvents: []Tracking{
				{
					Event:  "te1",
					Offset: &Offset{Duration: &duration1},
					URI:    "http://te1",
				},
				{
					Event:  "te2",
					Offset: &Offset{Duration: &duration2},
					URI:    "http://te2",
				},
			},
			AdParameters: &AdParameters{
				XMLEncoded: true,
				Parameters: "param",
			},
			StaticResource: &StaticResource{
				CreativeType: "ct",
				URI:          "http://src",
			},
			IFrameResource: CDATAString{"IFR"},
			HTMLResource: &HTMLResource{
				XMLEncoded: true,
				HTML:       "html",
			},
		}
		bb, _ := xml.Marshal(&cn)
		mRet := new(Companion)
		err := xml.Unmarshal(bb, mRet)
		So(err, ShouldBeNil)
		So(*mRet, ShouldResemble, cn)
	})
}

func TestImpression_UnmarshalXML(t *testing.T) {
	Convey("Impresssion_UnmarshalXML", t, func() {
		imp := Impression{
			ID:  "id",
			URI: "http://uri",
		}
		bb, _ := xml.Marshal(&imp)
		mRet := new(Impression)
		err := xml.Unmarshal(bb, mRet)
		So(err, ShouldBeNil)
		So(*mRet, ShouldResemble, imp)
		//impxml := `<Impression>
		//        <![CDATA[http://usimpdsp.mobvista.com/smaato_imp?price=${AUCTION_PRICE}&region=us&reqid=5c3dc3f700b2de6d4c63e598&campaignid=266185733&exchange=mintegral&restype=video&ctype=cpi&devid=C98B99BB-1B2F-434E-BE53-FD651B03FB6A&day=20190115&reqinfo=CS0JY3BpCW1pbnRlZ3JhbAkxNTQ3NTUxNzM1Njk3CS9taW50ZWdyYWxfcmVxdWVzdAktCTVjM2RjM2Y3MDBiMmRlNmQ0YzYzZTU5OAkwLjAwMDAwMDAwMDAJMC4wMDAwMDAwMDAwCS0Jbm9ybWFsCXJhbmtfbW9kZWxfY2FtcAl7ImFiIjoib2EiLCJhZG5tZGwiOiJpcGhvbmUgNnMiLCJhdCI6IjIiLCJjaXZyIjoiMC40NDczMzQzIiwiY3JlYXRpdmUiOiJkNDFkOGNkOThmMDBiMjA0ZTk4MDA5OThlY2Y4NDI3ZSMxIyMxMjYiLCJjdHIiOiIwLjAwMDAwMDAiLCJkaXZyIjoiMC4wMDAwMDAwIiwiZG1wX21zIjoiNCIsImlwIjoiMzQuMTk3Ljg1LjMwIiwiaXZyIjoiMC40NDczMzQzIiwicmFuayI6IjE1IiwicmVjZWl2ZWQiOiIxfDEuNCIsInJ2ciI6IjAuNDQ3MzM0MyIsInNlYXJjaCI6IjEyIiwic3JjIjoiNTQwIiwidGVtcGxhdGUiOiIxMjYiLCJ2ZXJudW0iOiIxMTUxOCIsInZlcnNpb24iOiJoZWFkcy96aCJ9CTMJLQkyCTMyZjk0MjI4LTUyZDMtNGJjZS1iZTg4LTVkZDFkNzk5YWU0ZgkxCTUwOTUJY29tLmFiYy50ZXN0CS0JMjUyNTkJSUFCMS0xJklBQjEtMyZJQUIxLTcmSUFCMiZJQUIzLTEJMAktCTE3Ny4xMDIuMjA0LjEwCWFwcGxlCWlwaG9uZTgsMQlpb3MJMTEuMS4yCTQJMglVUwlDOThCOTlCQi0xQjJGLTQzNEUtQkU1My1GRDY1MUIwM0ZCNkEJLQktCS0JLQkwCS0JTW96aWxsYS81LjAlMjAoaVBob25lJTNCJTIwQ1BVJTIwaVBob25lJTIwT1MlMjAxMF8yXzElMjBsaWtlJTIwTWFjJTIwT1MlMjBYKSUyMEFwcGxlV2ViS2l0LzYwMi40LjYlMjAoS0hUTUwlMkMlMjBsaWtlJTIwR2Vja28pJTIwTW9iaWxlLzE0RDI3JTIwQ2FtZXJhMzYwLzguMy40CTEzNDEJLG1pcyxEMTYsVGwscGEsemIsSTEsb2EJMC4wMDAwMDAwMDAwCXsiY2FycmllciI6IjQ0MDUwIiwiY3R5cGUiOiIxIiwiZGVDbGsiOiIxMDE2NyIsImRlSW1wIjoiMCIsInBhcnNlIjoiMSIsInJlcXR5cGUiOiJ2cmUiLCJza2lwIjoiZmFsc2UiLCJ2dGEiOiIxIn0JMjY2MTg1NzMzCTAuMDAwMDAwMDAwMAktCS0JMAkw]]>
		//    </Impression>`
		//impression := new(Impression)
		//xml.Unmarshal([]byte(impxml), impression)
		//t.Log(impression)
	})
}
