package store

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"
)

const (
	addr = "localhost:6379"
	db   = 15
)

var (
	logger = log.New(os.Stderr, "", log.LstdFlags)
)

func TestRedisSetGet(t *testing.T) {
	s := NewRedisStore(&RedisOptions{
		Addr: addr,
		DB:   db,
	}, time.Second, logger)
	id := "captcha id"
	d := []byte("123456")
	s.Set(id, d)
	d2 := s.Get(id, false)

	if d2 == nil || !bytes.Equal(d, d2) {
		t.Errorf("saved %v, getDigits returned got %v", d, d2)
	}
}

func TestRedisGetClear(t *testing.T) {
	s := NewRedisStore(&RedisOptions{
		Addr: addr,
		DB:   db,
	}, time.Second, logger)
	id := "captcha id"
	d := []byte("123456")
	s.Set(id, d)
	d2 := s.Get(id, true)
	if d2 == nil || !bytes.Equal(d, d2) {
		t.Errorf("saved %v, getDigitsClear returned got %v", d, d2)
	}

	d2 = s.Get(id, false)
	if d2 != nil {
		t.Errorf("getDigitClear didn't clear (%q=%v)", id, d2)
	}
}

func TestRedisGC(t *testing.T) {
	s := NewRedisStore(&RedisOptions{
		Addr: addr,
		DB:   db,
	}, time.Millisecond*10, logger)
	id := "captcha id"
	d := []byte("123456")
	s.Set(id, d)

	time.Sleep(time.Millisecond * 200)
	d2 := s.Get(id, false)

	if d2 != nil {
		t.Errorf("gc didn't clear (%q=%v)", id, d2)
	}
}
