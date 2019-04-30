package store

import (
	"bytes"
	"testing"
	"time"
)

func TestMemorySetGet(t *testing.T) {
	s := NewMemoryStore(time.Second, time.Second*2)
	id := "captcha id"
	d := []byte("123456")
	s.Set(id, d)
	d2 := s.Get(id, false)

	if d2 == nil || !bytes.Equal(d, d2) {
		t.Errorf("saved %v, getDigits returned got %v", d, d2)
	}
}

func TestMemoryGetClear(t *testing.T) {
	s := NewMemoryStore(time.Second, time.Second*2)
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

func TestMemoryGC(t *testing.T) {
	s := NewMemoryStore(time.Millisecond*10, time.Millisecond*100)
	id := "captcha id"
	d := []byte("123456")
	s.Set(id, d)

	time.Sleep(time.Millisecond * 200)
	d2 := s.Get(id, false)

	if d2 != nil {
		t.Errorf("gc didn't clear (%q=%v)", id, d2)
	}
}
