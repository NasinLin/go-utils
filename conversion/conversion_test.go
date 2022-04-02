package conversion

import (
	"log"
	"testing"
)

func TestByteToDateTimeString(t *testing.T) {

	testItems := [...]struct {
		v1 []byte
		v2 string
	}{
		{v1: []byte{7, 225, 1, 1, 255, 1, 0, 0, 255, 128, 0, 0}, v2: "2017-01-01T01:00:00Z"},
		{v1: []byte{7, 225, 1, 1, 255, 0, 45, 0, 255, 128, 0, 0}, v2: "2017-01-01T00:45:00Z"},
		{v1: []byte{7, 225, 1, 1, 255, 0, 30, 0, 255, 128, 0, 0}, v2: "2017-01-01T00:30:00Z"},
		{v1: []byte{7, 225, 1, 1, 255, 0, 15, 0, 255, 128, 0, 0}, v2: "2017-01-01T00:15:00Z"},
		{v1: []byte{7, 225, 1, 1, 255, 0, 0, 0, 255, 128, 0, 0}, v2: "2017-01-01T00:00:00Z"},
	}

	for _, i := range testItems {
		out, err := ByteToDateTimeString(i.v1)
		if err != nil {
			t.Errorf(err.Error())
		}
		if out != i.v2 {
			t.Errorf("failed, get: %s, should: %s", out, i.v2)
		}
	}

}

func TestByteToNumberString(t *testing.T) {
	testItems := [...]struct {
		v1 []byte
		v2 string
	}{
		{v1: []byte{255}, v2: "255"},
		{v1: []byte{0}, v2: "0"},
		{v1: []byte{1}, v2: "1"},
		{v1: []byte{255, 255}, v2: "65535"},
		{v1: []byte{255, 254}, v2: "65534"},
	}

	for _, i := range testItems {
		out, err := ByteToNumberString(i.v1)
		if err != nil {
			t.Errorf(err.Error())
		}
		if out != i.v2 {
			t.Errorf("failed, get: %s, should: %s", out, i.v2)
		}
	}
}

func TestByteToInt64(t *testing.T) {
	testItems := [...]struct {
		v1 []byte
		v2 int64
	}{
		{v1: []byte{255}, v2: 255},
		{v1: []byte{0}, v2: 0},
		{v1: []byte{1}, v2: 1},
		{v1: []byte{255, 255}, v2: 65535},
		{v1: []byte{255, 254}, v2: 65534},
	}

	for _, i := range testItems {
		out, err := ByteToInt64(i.v1)
		if err != nil {
			t.Errorf(err.Error())
		}
		if out != i.v2 {
			t.Errorf("failed, get: %d, should: %d", out, i.v2)
		}
	}
}

func TestByteToInt32(t *testing.T) {
	testItems := [...]struct {
		v1 []byte
		v2 int32
	}{
		{v1: []byte{255}, v2: 255},
		{v1: []byte{0}, v2: 0},
		{v1: []byte{1}, v2: 1},
		{v1: []byte{255, 255}, v2: 65535},
		{v1: []byte{255, 254}, v2: 65534},
	}

	for _, i := range testItems {
		out, err := ByteToInt32(i.v1)
		if err != nil {
			t.Errorf(err.Error())
		}
		if out != i.v2 {
			t.Errorf("failed, get: %d, should: %d", out, i.v2)
		}
	}
}

func TestByteToBinary8BitString(t *testing.T) {
	testItems := [...]struct {
		v1 []byte
		v2 string
	}{
		{v1: []byte{0}, v2: "00000000"},
		{v1: []byte{1}, v2: "00000001"},
		{v1: []byte{10}, v2: "00001010"},
		{v1: []byte{200}, v2: "11001000"},
		{v1: []byte{255}, v2: "11111111"},
	}

	for _, i := range testItems {
		out, err := ByteToBinary8BitString(i.v1)
		if err != nil {
			t.Errorf(err.Error())
		}
		if out != i.v2 {
			t.Errorf("failed, get: %s, should: %s", out, i.v2)
		}
	}

}

func TestString(t *testing.T) {
	str := "ABC"
	asciiByte := []byte(str)
	log.Printf("A : %+v", asciiByte)
}
