package madx

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMOrtbRequest_FromJson(t *testing.T) {
	Convey("FromJson", t, func() {
		jsonBidRequest := &MOrtbRequest{}
		data := `
				{
				  "id":"ed24882da38ee500ee941fd9f2988e1d5f28f64f",
				  "imp":[
					  {
						  "id":"ed24882da38ee500ee941fd9f2988e1d5f28f64f",
						  "banner":{
							  "w":480,
							  "h":320,
							  "wmax":480,
							  "hmax":320,
							  "btype":[
								  4
							  ],
							  "battr":[
								  5,
								  10,
								  14,
								  6,
								  9,
								  12,
								  7,
								  3,
								  11,
								  8,
								  4
							  ],
							  "pos":7,
							  "topframe":1,
							  "api":[
								  3,
								  5
							  ],
							  "ext":{
								  "placementtype":"rewarded",
								  "playableonly":true,
								  "allowscustomclosebutton":false
							  }
						  },
						  "video":{
							  "mimes":[
								  "video/mp4"
							  ],
							  "minduration":6,
							  "maxduration":30,
							  "protocols":[
								  5,
								  1,
								  6,
								  2,
								  3
							  ],
							  "w":480,
							  "h":320,
							  "placement":5,
							  "linearity":1,
							  "skip":0,
							  "battr":[
								  5,
								  14,
								  9,
								  12,
								  3,
								  11,
								  8,
								  4
							  ],
							  "delivery":[
								  2
							  ],
							  "pos":7,
							  "companiontype":[
								  1,
								  2
							  ],
							  "companionad":[
								  {
									  "btype":[
										  4
									  ],
									  "battr":[
										  1,
										  2,
										  6,
										  7
									  ],
									  "pos":7,
									  "h":320,
									  "w":480,
									  "id":"1"
								  }
							  ],
							  "ext":{
								  "placementtype":"rewarded"
							  }
						  },
						  "instl":1,
						  "tagid":"Default",
						  "displaymanager":"Chartboost-Android-SDK",
						  "displaymanagerver":"7.3.0",
						  "bidfloor":0.25,
						  "bidfloorcur":"USD",
						  "secure":1
					  }
				  ],
				  "app":{
					  "id":"530321159ddc356afb1d558c",
					  "name":"Online Soccer Manager (OSM)",
					  "bundle":"com.gamebasics.osm",
					  "storeurl":"market://details?id=com.gamebasics.osm",
					  "cat":[
						  "IAB9-30"
					  ],
					  "publisher":{
						  "id":"51949d9617ba475226000002",
						  "name":"Gamebasics BV"
					  },
					  "ext":{
						  "itunesid":"com.gamebasics.osm",
						  "premium":false
					  }
				  },
				  "device":{
					  "ua":"Mozilla/5.0 (Linux; Android 7.1.2; Redmi 4X Build/N2G47H; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36",
					  "geo":{
						  "lat":51.339698791503906,
						  "lon":12.373100280761719,
						  "country":"DEU",
						  "type":2,
						  "region":"Saxony",
						  "city":"Leipzig",
						  "zip":"04109"
					  },
					  "lmt":0,
					  "ip":"93.242.41.42",
					  "devicetype":4,
					  "model":"Redmi 4X",
					  "os":"Android",
					  "osv":"7.1.2",
					  "h":720,
					  "w":1280,
					  "connectiontype":2,
					  "ifa":"76057bcd-12e2-4d54-abf0-33e7d544d292",
					  "language":"de",
					  "carrier":"WIFI",
					  "js":1,
					  "pxratio":2
				  },
				  "user":{
					  "id":"57516c9b-2552-11e8-8fb5-155b99d6dc70",
					  "geo":{
						  "lat":51.339698791503906,
						  "lon":12.373100280761719,
						  "country":"DEU",
						  "type":2,
						  "region":"Saxony",
						  "city":"Leipzig",
						  "zip":"04109"
					  },
					  "ext":{
						  "consent":0
					  }
				  },
				  "test":0,
				  "at":2,
				  "tmax":400,
				  "cur":[
					  "USD"
				  ],
				  "bcat":[
					  "IAB7-39",
					  "IAB7-42",
					  "IAB8-5",
					  "IAB9-9",
					  "IAB14-1",
					  "IAB19-3",
					  "IAB25-1",
					  "IAB25-2",
					  "IAB25-3",
					  "IAB25-4",
					  "IAB25-5",
					  "IAB25-6",
					  "IAB25-7",
					  "IAB26-1",
					  "IAB26-2",
					  "IAB26-3",
					  "IAB26-4"
				  ],
				  "bapp":[
					  "eu.nordeus.topeleven.android",
					  "com.telsis.android.drivesafe",
					  "eu.nordeus.TopEleven"
				  ],
				  "regs":{
					  "coppa":0,
					  "ext":{
						  "gdpr":1
					  }
				  }
				}
				`
		jsonBidRequest.FromJson([]byte(data))
		So(jsonBidRequest.Imp[0].Video.Ext.PlacementType, ShouldEqual, "rewarded")
	})
}

func TestMOrtbRequest_String(t *testing.T) {
	Convey("To String", t, func() {
		mines := []string{"mp3", "mp4"}
		imp := MImp{
			Video: &MVideo{
				Mimes: mines,
			},
		}

		jsonBidRequest := &MOrtbRequest{
			Imp: []*MImp{&imp},
		}

		t.Log(jsonBidRequest.String())
		So(jsonBidRequest.String(), ShouldContainSubstring, "mp4")
	})
}
