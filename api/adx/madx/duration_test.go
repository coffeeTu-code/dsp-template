package madx

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDurationMarshaler(t *testing.T) {
	Convey("序列化", t, func() {
		b, err := Duration(0).MarshalText()
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, "00:00:00")
		b, err = Duration(2 * time.Millisecond).MarshalText()
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, "00:00:00.002")

		b, err = Duration(2 * time.Second).MarshalText()
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, "00:00:02")

		b, err = Duration(2 * time.Minute).MarshalText()
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, "00:02:00")

		b, err = Duration(2 * time.Hour).MarshalText()
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, "02:00:00")

	})

}

func TestDurationUnmarshal(t *testing.T) {
	Convey("反序列化", t, func() {
		var d Duration
		So(d.UnmarshalText([]byte("00:00:00")), ShouldBeNil)
		So(d, ShouldEqual, Duration(0))

		d = 0
		So(d.UnmarshalText([]byte("00:00:02")), ShouldBeNil)
		So(d, ShouldEqual, Duration(2*time.Second))

		d = 0
		So(d.UnmarshalText([]byte(" 00:00:02 ")), ShouldBeNil)
		So(d, ShouldEqual, Duration(2*time.Second))

		d = 0
		So(d.UnmarshalText([]byte("00:02:00")), ShouldBeNil)
		So(d, ShouldEqual, Duration(2*time.Minute))

		d = 0
		So(d.UnmarshalText([]byte("02:00:00")), ShouldBeNil)
		So(d, ShouldEqual, Duration(2*time.Hour))

		d = 0
		So(d.UnmarshalText([]byte("00:00:00.123")), ShouldBeNil)
		So(d, ShouldEqual, Duration(123*time.Millisecond))

		d = 0
		So(d.UnmarshalText([]byte("undefined")), ShouldBeNil)
		So(d, ShouldEqual, Duration(0))

		d = 0
		So(d.UnmarshalText([]byte("")), ShouldBeNil)
		So(d, ShouldEqual, Duration(0))

		//assert.EqualError(t, d.UnmarshalText([]byte("00:00:60")), "invalid duration: 00:00:60")
		//assert.EqualError(t, d.UnmarshalText([]byte("00:60:00")), "invalid duration: 00:60:00")
		//assert.EqualError(t, d.UnmarshalText([]byte("00:00:00.-1")), "invalid duration: 00:00:00.-1")
		//assert.EqualError(t, d.UnmarshalText([]byte("00:00:00.1000")), "invalid duration: 00:00:00.1000")
		//assert.EqualError(t, d.UnmarshalText([]byte("00h01m")), "invalid duration: 00h01m")
	})

}
