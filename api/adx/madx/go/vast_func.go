package mortb

import (
	"bytes"
	"encoding/xml"
	"io"
	"strconv"
)

func (i *Impression) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			i.ID = attr.Value
		}
	}
	chardata := bytes.Buffer{}
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		switch tt := t.(type) {
		case xml.CharData:
			chardata.Write(tt)
		}
	}
	if chardata.Len() > 0 {
		i.URI = chardata.String()
	}
	return nil
}
func (v *VideoClicks) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		switch tt := t.(type) {
		case xml.StartElement:
			var vc VideoClick
			d.DecodeElement(&vc, &tt)
			switch tt.Name.Local {
			case "ClickThrough":
				v.ClickThroughs = append(v.ClickThroughs, vc)
			case "ClickTracking":
				v.ClickTrackings = append(v.ClickTrackings, vc)
			case "CustomClick":
				v.CustomClicks = append(v.CustomClicks, vc)
			}
			//case xml.CharData:
			//case xml.EndElement:
		}

	}
	return nil
}

func (v *VideoClick) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			v.ID = attr.Value
		}
	}
	chardata := bytes.Buffer{}
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		switch tt := t.(type) {
		case xml.CharData:
			chardata.Write(tt)
		}
	}
	if chardata.Len() > 0 {
		v.URI = chardata.String()
	}
	return nil
}

func (f *MediaFile) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			f.ID = attr.Value
		case "delivery":
			f.Delivery = attr.Value
		case "type":
			f.Type = attr.Value
		case "codec":
			f.Codec = attr.Value
		case "bitrate":
			f.Bitrate, _ = strconv.Atoi(attr.Value)
		case "minBitrate":
			f.MinBitrate, _ = strconv.Atoi(attr.Value)
		case "maxBitrate":
			f.MaxBitrate, _ = strconv.Atoi(attr.Value)
		case "width":
			f.Width, _ = strconv.Atoi(attr.Value)
		case "height":
			f.Height, _ = strconv.Atoi(attr.Value)
		case "scalable":
			f.Scalable, _ = strconv.ParseBool(attr.Value)
		case "maintainAspectRatio":
			f.MaintainAspectRatio, _ = strconv.ParseBool(attr.Value)
		case "apiFramework":
			f.APIFramework = attr.Value
		case "size":
			f.Size, _ = strconv.Atoi(attr.Value)
		}
	}
	chardata := bytes.Buffer{}
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		switch tt := t.(type) {
		case xml.CharData:
			chardata.Write(tt)
		}
	}
	if chardata.Len() > 0 {
		f.URI = chardata.String()
	}

	return nil
}

func (c *Companion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			c.ID = attr.Value
		case "height":
			c.Height, _ = strconv.Atoi(attr.Value)
		case "width":
			c.Width, _ = strconv.Atoi(attr.Value)
		case "assetHeight":
			c.AssetHeight, _ = strconv.Atoi(attr.Value)
		case "assetWidth":
			c.AssetWidth, _ = strconv.Atoi(attr.Value)
		case "expandedHeight":
			c.ExpandedHeight, _ = strconv.Atoi(attr.Value)
		case "expandedWidth":
			c.ExpandedWidth, _ = strconv.Atoi(attr.Value)
		case "apiFramework":
			c.APIFramework = attr.Value
		case "adSlotId":
			c.AdSlotID = attr.Value
		}
	}
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		switch tt := t.(type) {
		//case xml.CharData:
		//	f.URI = string(tt)
		case xml.StartElement:
			switch tt.Name.Local {
			case "CompanionClickThrough":
				var s CDATAString
				d.DecodeElement(&s, &tt)
				c.CompanionClickThrough = s
			case "CompanionClickTracking":
				var s CDATAString
				d.DecodeElement(&s, &tt)
				c.CompanionClickTracking = append(c.CompanionClickTracking, s)
			case "AltText":
				var s string
				d.DecodeElement(&s, &tt)
				c.AltText = s
			case "Tracking":
				var s Tracking
				d.DecodeElement(&s, &tt)
				c.TrackingEvents = append(c.TrackingEvents, s)
			case "StaticResource":
				var s StaticResource
				d.DecodeElement(&s, &tt)
				c.StaticResource = &s
			case "AdParameters":
				var s AdParameters
				d.DecodeElement(&s, &tt)
				c.AdParameters = &s
			case "IFrameResource":
				var s CDATAString
				d.DecodeElement(&s, &tt)
				c.IFrameResource = s
			case "HTMLResource":
				var s HTMLResource
				d.DecodeElement(&s, &tt)
				c.HTMLResource = &s
			}
		}
	}
	return nil
}
func (t *Tracking) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "event":
			t.Event = attr.Value
		case "offset":
			//fmt.Println("offset:", attr.Value)
			t.Offset = new(Offset)
			t.Offset.UnmarshalText([]byte(attr.Value))
		}
	}
	chardata := bytes.Buffer{}
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		switch tt := t.(type) {
		case xml.CharData:
			chardata.Write(tt)
		}
	}
	if chardata.Len() > 0 {
		t.URI = chardata.String()
	}
	return nil
}
