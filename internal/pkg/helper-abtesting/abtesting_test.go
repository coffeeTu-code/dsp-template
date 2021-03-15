package abtesting

import (
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"hash/crc32"
	"strconv"
	"testing"
	"time"
	"unsafe"
)

const abtestInfo = `[{
  "flow_control" : {
    "region" : [],
    "ad_type" : [],
    "adx" : [],
    "device" : ["null", "notNull"],
    "country" : [],
    "start_time" : "2020-05-16 12:00:00"
  },
  "experiment_type" : "exclusive",
  "layer" : "layer1",
  "experiment" : {
    "name" : "test1",
    "content" : {
        "control": 1000, 
        "expr1": 3000,
        "expr2": 5000
    }
  }
},{
  "flow_control" : {
    "region" : [],
    "ad_type" : [],
    "adx" : [],
    "device" : ["null", "notNull"],
    "country" : [],
    "start_time" : "2020-05-16 12:00:00"
  },
  "experiment_type" : "",
  "layer" : "layer2",
  "experiment" : {
    "name" : "test2",
    "content" : {
        "control" : 1000, 
        "expr1": 1000,
        "expr2": 4000
    }
  }
},{
  "flow_control" : {
    "region" : [],
    "ad_type" : [],
    "adx" : [],
    "device" : ["null", "notNull"],
    "country" : [],
    "start_time" : "2020-05-16 12:00:00"
  },
  "experiment_type" : "",
  "layer" : "layer3",
  "experiment" : {
    "name" : "test12",
    "content" : {
        "control" : 1, 
        "expr1": 1000,
        "expr2": 2000
    }
  }
},{
  "flow_control" : {
    "region" : [],
    "ad_type" : [],
    "adx" : [],
    "device" : ["null", "notNull"],
    "country" : [],
    "start_time" : "2020-05-16 12:00:00"
  },
  "experiment_type" : "",
  "layer" : "layer4",
  "experiment" : {
    "name" : "test22",
    "content" : {
        "control": 40000,
        "expr1": 100, 
        "expr2": 20
    }
  }
}]`

func TestAbTesting_ToString(t *testing.T) {
	m := map[string]string{
		"abc": "123",
		"cde": "345",
	}

	Convey("TestAbTesting_ToString", t, func() {
		fmt.Println(ToString(m))
		ShouldEqual("abc=123;cde=345", ToString(m))
	})

}

func TestConsul(t *testing.T) {
	ab, err := NewAbTesting("127.0.0.1:8500", "abtest/abtest")
	e := ab.Init(context.Background())
	fmt.Println(e)
	Convey("TestAbTesting_ToString", t, func() {
		ShouldBeNil(err)
		kv := ab.client.KV()
		So(kv, ShouldNotBeNil)
		So(err, ShouldBeNil)
		fmt.Println(ExperimentsInfo)
		for i := 0; i < 100; i++ {
			m := ab.GetAbTesting(&FlowInfo{
				HashKey: "1bcddsfaxsxxs",
				AdType:  "b",
			})
			fmt.Println(ToString(m))
		}
	})
}

func TestAbTesting_GetAbTesting(t *testing.T) {
	ab, err := NewAbTesting("18.197.217.185:8500", "abtest/abtest")
	Convey("TestAbTesting_ToString", t, func() {
		So(err, ShouldBeNil)
		err = ab.Init(context.Background())
		So(err, ShouldBeNil)
		flow := &FlowInfo{
			Device:  "null",
			Region:  "vg",
			Adx:     "1",
			AdType:  "296",
			Country: "US",
		}
		a, b, c := 0, 0, 0
		a1, b1, c1 := 0, 0, 0
		a2, b2, c2 := 0, 0, 0
		a3, b3, c3 := 0, 0, 0
		for i := 0; i < 100000; i++ {
			v := ab.GetAbTesting(flow)
			if v["test1"] == "expr1" {
				a++
			}
			if v["test1"] == "expr2" {
				b++
			}
			if v["test1"] == "control" {
				c++
			}
			if v["test2"] == "expr1" {
				a1++
			}
			if v["test2"] == "expr2" {
				b1++
			}

			if v["test2"] == "control" {
				c1++
			}

			if v["test12"] == "expr1" {
				a2++
			}
			if v["test12"] == "expr2" {
				b2++
			}
			if v["test12"] == "control" {
				c2++
			}

			if v["test22"] == "expr1" {
				a3++
			}
			if v["test22"] == "expr2" {
				b3++
			}
			if v["test22"] == "control" {
				c3++
			}
		}
		fmt.Println(a, b, c, "expect: 3:5:1")
		fmt.Println(a1, b1, c1)
		fmt.Println(a2, b2, c2)
		fmt.Println(a3, b3, c3)
		fmt.Println(a + b + a1 + b1 + a2 + b2 + a3 + b3)
	})
}

func BenchmarkAbTesting_GetAbTesting(b *testing.B) {
	ab, _ := NewAbTesting("18.197.217.185:8500", "abtest/abtest")
	_ = ab.Init(context.Background())
	flow := &FlowInfo{
		Device:  "null",
		Region:  "vg",
		Adx:     "1",
		AdType:  "296",
		Country: "US",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100000; j++ {
			_ = ab.GetAbTesting(flow)
		}
	}
	b.StopTimer()
}

func TestRate(t *testing.T) {
	Arr := make([]int, 10000)
	ExperimentsInfo = &Experiments{
		Exclusive:  make([]Info, 0),
		MultiLayer: make(map[string][]Info, 0),
	}

	ExperimentsInfo.MultiLayer["abc"] = []Info{
		{
			Name: "abc",
			Content: []ContentInfo{
				{"a", 10},
				{"b", 10},
			},
		},
	}

	ab := AbTesting{}
	num := 1000 * 1000
	a, b := 0, 0
	for i := 0; i < num; i++ {
		m := ab.GetAbTesting(&FlowInfo{
			AdType: "b",
		})
		for _, v := range m {
			if v == "a" {
				a++
			} else if v == "b" {
				b++
			}
		}
	}
	println(a, b, float64(a)/float64(num)*100, float64(b)/float64(num)*100)
	//println(calNum, calNumA, calNumB)
	for i, v := range Arr {
		println(i, v)
	}
}

func TestRandom(t *testing.T) {
	num := 10000 * 10000
	arr := make([]int, 10000)
	a, b, c := 0, 0, 0
	for i := 0; i < num; i++ {
		// hashSalt := ExclusiveFlow + strconv.Itoa(int(time.Now().UnixNano()))
		hashSalt := ExclusiveFlow + strconv.Itoa(int(time.Now().UnixNano()))
		// hashNum := int(crc32.ChecksumIEEE(*(*[]byte)(unsafe.Pointer(&hashSalt))) % uint32(10000))
		hashNum := int(crc32.ChecksumIEEE(*(*[]byte)(unsafe.Pointer(&hashSalt))) % uint32(10000))
		arr[hashNum]++
		//if hashNum < 10 {
		//	a++
		//} else if hashNum < 20 {
		//	b++
		//}
	}
	m := 0
	for _, v := range arr {
		//println(v)
		if v >= 9900 && v <= 10100 {
			m++
		}
		if v < 9800 || v > 10200 {
			a++
		} else if v < 9900 || v > 10100 {
			b++
		} else {
			c++
		}
	}
	println(m, a, b, c)
}

const (
	BIG_M = 0xc6a4a7935bd1e995
	BIG_R = 47
	SEED  = 0x1234ABCD
)

func MurmurHash64A(data []byte) (h uint64) {
	var k uint64
	h = SEED ^ uint64(len(data))*BIG_M

	var ubigm uint64 = BIG_M
	var ibigm = ubigm
	for l := len(data); l >= 8; l -= 8 {
		k = uint64(int64(data[0]) | int64(data[1])<<8 | int64(data[2])<<16 | int64(data[3])<<24 |
			int64(data[4])<<32 | int64(data[5])<<40 | int64(data[6])<<48 | int64(data[7])<<56)

		k := k * ibigm
		k ^= k >> BIG_R
		k = k * ibigm

		h = h ^ k
		h = h * ibigm
		data = data[8:]
	}

	switch len(data) {
	case 7:
		h ^= uint64(data[6]) << 48
		fallthrough
	case 6:
		h ^= uint64(data[5]) << 40
		fallthrough
	case 5:
		h ^= uint64(data[4]) << 32
		fallthrough
	case 4:
		h ^= uint64(data[3]) << 24
		fallthrough
	case 3:
		h ^= uint64(data[2]) << 16
		fallthrough
	case 2:
		h ^= uint64(data[1]) << 8
		fallthrough
	case 1:
		h ^= uint64(data[0])
		h *= ibigm
	}

	h ^= h >> BIG_R
	h *= ibigm
	h ^= h >> BIG_R
	return
}
