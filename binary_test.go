package dlt645_test

import (
	"reflect"
	"testing"

	"github.com/LaoYutang/go-dlt645"
)

func TestString2BcdBytes(t *testing.T) {
	tests := []struct {
		input    string
		expected []byte
	}{
		{"12", []byte{0x12}},
		{"1234", []byte{0x12, 0x34}},
		{"1", []byte{0x01}},
		{"", []byte{}},
		{"987654", []byte{0x98, 0x76, 0x54}},
		{"0001", []byte{0x00, 0x01}},
		{"a2", []byte{0xa2}},
		{"9b", []byte{0x9b}},
		{"1!", []byte{0x10}},
	}

	for _, tt := range tests {
		got := dlt645.String2BcdBytes(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("String2BcdBytes(%q) = %#v, want %#v", tt.input, got, tt.expected)
		}
	}
}

func TestBcdBytes2String(t *testing.T) {
	tests := []struct {
		input    []byte
		expected string
	}{
		{[]byte{}, ""},
		{[]byte{0x12}, "12"},
		{[]byte{0x00}, "00"},
		{[]byte{0x98, 0x76, 0x54}, "987654"},
		{[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, "0123456789abcdef"},
		{[]byte{0xff}, "ff"},
		{[]byte{0x09}, "09"},
		{[]byte{0x10, 0x20}, "1020"},
	}

	for _, tt := range tests {
		got := dlt645.BcdBytes2String(tt.input)
		if got != tt.expected {
			t.Errorf("BcdBytes2String2(%#v) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func BenchmarkBcdBytes2String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dlt645.BcdBytes2String([]byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0})
	}
}

func TestBytesSub(t *testing.T) {
	tests := []struct {
		input    []byte
		sub      byte
		expected []byte
	}{
		{[]byte{}, 1, []byte{}},
		{[]byte{0x10, 0x20, 0x30}, 0x10, []byte{0x00, 0x10, 0x20}},
		{[]byte{0x00, 0x01, 0x02}, 0x01, []byte{0xff, 0x00, 0x01}},
		{[]byte{0xff, 0xfe, 0xfd}, 0x01, []byte{0xfe, 0xfd, 0xfc}},
		{[]byte{0x05}, 0x05, []byte{0x00}},
		{[]byte{0x00}, 0x00, []byte{0x00}},
	}

	for _, tt := range tests {
		got := dlt645.BytesSub(tt.input, tt.sub)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("BytesSub(%#v, %#x) = %#v, want %#v", tt.input, tt.sub, got, tt.expected)
		}
	}
}

func TestBytesAdd(t *testing.T) {
	tests := []struct {
		input    []byte
		add      byte
		expected []byte
	}{
		{[]byte{}, 1, []byte{}},
		{[]byte{0x10, 0x20, 0x30}, 0x10, []byte{0x20, 0x30, 0x40}},
		{[]byte{0x00, 0x01, 0x02}, 0x01, []byte{0x01, 0x02, 0x03}},
		{[]byte{0xff, 0xfe, 0xfd}, 0x01, []byte{0x00, 0xff, 0xfe}},
		{[]byte{0x05}, 0x05, []byte{0x0a}},
		{[]byte{0x00}, 0x00, []byte{0x00}},
	}

	for _, tt := range tests {
		got := dlt645.BytesAdd(tt.input, tt.add)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("BytesAdd(%#v, %#x) = %#v, want %#v", tt.input, tt.add, got, tt.expected)
		}
	}
}

func TestBytesReverse(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte{0x01}, []byte{0x01}},
		{[]byte{0x01, 0x02}, []byte{0x02, 0x01}},
		{[]byte{0x01, 0x02, 0x03}, []byte{0x03, 0x02, 0x01}},
		{[]byte{0xff, 0x00, 0x55, 0xaa}, []byte{0xaa, 0x55, 0x00, 0xff}},
	}

	for _, tt := range tests {
		got := dlt645.BytesReverse(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("BytesReverse(%#v) = %#v, want %#v", tt.input, got, tt.expected)
		}
	}
}
