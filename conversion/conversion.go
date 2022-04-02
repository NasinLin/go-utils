package conversion

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"

	"log"
	"strconv"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

func ByteToDateString(src []byte) (string, error) {
	y := hex.EncodeToString(src[0:2])
	m := hex.EncodeToString(src[2:3])
	d := hex.EncodeToString(src[3:4])

	num, err := strconv.ParseInt(y, 16, 64)
	if err != nil {
		return "", errors.New("parse year failed")
	}

	num1, err := strconv.ParseInt(m, 16, 64)
	if err != nil {
		return "", errors.New("parse month failed")
	}

	num2, err := strconv.ParseInt(d, 16, 64)
	if err != nil {
		return "", errors.New("parse day failed")
	}

	out := fmt.Sprintf("%4d-%02d-%02d", num, num1, num2)

	return out, nil
}

func ByteToTimeString(src []byte) (string, error) {

	h := hex.EncodeToString(src[0:1])
	m := hex.EncodeToString(src[1:2])
	s := hex.EncodeToString(src[2:3])

	num, err := strconv.ParseInt(h, 16, 64)
	if err != nil {
		return "", errors.New("parse hour failed")
	}

	num1, err := strconv.ParseInt(m, 16, 64)
	if err != nil {
		return "", errors.New("parse minute failed")
	}

	num2, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return "", errors.New("parse second failed")
	}

	out := fmt.Sprintf("%02d:%02d:%02d", num, num1, num2)

	return out, nil

}

func ByteToDateTimeString(src []byte) (string, error) {
	type dateDetail struct {
		year   int64
		month  int64
		day    int64
		hour   int64
		minute int64
		second int64
	}
	var detail dateDetail
	y := hex.EncodeToString(src[0:2])
	m := hex.EncodeToString(src[2:3])
	d := hex.EncodeToString(src[3:4])
	hh := hex.EncodeToString(src[5:6])
	mm := hex.EncodeToString(src[6:7])
	ss := hex.EncodeToString(src[7:8])

	if num, err := strconv.ParseInt(y, 16, 64); err != nil {
		log.Println(err.Error())
		return "", errors.New("parse year failed")
	} else {
		detail.year = num
	}

	if detail.year == 65535 {
		return hex.EncodeToString(src), nil
	}

	if num, err := strconv.ParseInt(m, 16, 64); err != nil {
		log.Println(err.Error())
		return "", errors.New("parse month failed")
	} else {
		detail.month = num
	}

	if num, err := strconv.ParseInt(d, 16, 64); err != nil {
		log.Println(err.Error())
		return "", errors.New("parse day failed")
	} else {
		detail.day = num
	}

	if num, err := strconv.ParseInt(hh, 16, 64); err != nil {
		log.Println(err.Error())
		return "", errors.New("parse hour failed")
	} else {
		detail.hour = num
	}

	if num, err := strconv.ParseInt(mm, 16, 64); err != nil {
		log.Println(err.Error())
		return "", errors.New("parse minute failed")
	} else {
		detail.minute = num
	}

	if num, err := strconv.ParseInt(ss, 16, 64); err != nil {
		log.Println(err.Error())
		return "", errors.New("parse second failed")
	} else {
		detail.second = num
	}

	time := time.Date(int(detail.year), time.Month(detail.month), int(detail.day), int(detail.hour), int(detail.minute), int(detail.second), 0, time.UTC)
	timeString, _ := time.MarshalText()
	return string(timeString), nil
}

func ByteToNumberString(src []byte) (string, error) {
	x := hex.EncodeToString(src)
	num, err := strconv.ParseInt(x, 16, 64)
	if err != nil {
		log.Println(err.Error())
		errMsg := fmt.Sprintf("parse byte to number string failed, %+v", err.Error())
		return "", errors.New(errMsg)
	}

	x1 := int(num)
	s := strconv.Itoa(x1)
	return s, nil
}

func ByteToInt64(src []byte) (int64, error) {
	x := hex.EncodeToString(src)
	num, err := strconv.ParseInt(x, 16, 64)
	if err != nil {
		log.Println(err.Error())
		errMsg := fmt.Sprintf("parse byte to int64 failed, %+v", err.Error())
		return -1, errors.New(errMsg)
	}
	return num, nil

}

func ByteToInt32(src []byte) (int32, error) {
	x := hex.EncodeToString(src)
	num, err := strconv.ParseInt(x, 16, 32)
	if err != nil {
		log.Println(err.Error())
		errMsg := fmt.Sprintf("parse byte to int failed, %+v", err.Error())
		return -1, errors.New(errMsg)
	}
	return int32(num), nil
}

func ByteToBinary8BitString(src []byte) (string, error) {
	buf := bytes.NewBuffer([]byte{})
	for _, v := range src {
		buf.WriteString(fmt.Sprintf("%08b", v))
	}
	return buf.String(), nil
}
