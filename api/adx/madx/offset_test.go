package madx

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOffsetMarshaler(t *testing.T) {
	Convey("OffsetMarshaler", t, func() {
		b, err := Offset{}.MarshalText()
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, "0%")

		b, err = Offset{Percent: .1}.MarshalText()
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, "10%")

		d := Duration(0)
		b, err = Offset{Duration: &d}.MarshalText()
		So(err, ShouldBeNil)
		So("00:00:00", ShouldEqual, string(b))
	})

}

func TestOffsetUnmarshaler(t *testing.T) {
	Convey("OffsetUnmarshaler", t, func() {
		var o Offset
		So(o.UnmarshalText([]byte("0%")), ShouldBeNil)
		So(o.Duration, ShouldBeNil)
		So(float32(0.0), ShouldEqual, o.Percent)

		o = Offset{}
		So(o.UnmarshalText([]byte("10%")), ShouldBeNil)
		So(o.Duration, ShouldBeNil)
		So(float32(0.1), ShouldEqual, o.Percent)

		o = Offset{}
		So(o.UnmarshalText([]byte("00:00:00")), ShouldBeNil)
		So(o.Duration, ShouldNotBeNil)
		So(Duration(0), ShouldEqual, *o.Duration)
		So(float32(0), ShouldEqual, o.Percent)

		o = Offset{}
		So(o.UnmarshalText([]byte("abc%")), ShouldBeError, "invalid offset: abc%")
	})
}
