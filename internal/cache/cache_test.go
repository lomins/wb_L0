package cache

import (
	"reflect"
	"testing"
)

func TestCache(t *testing.T) {
	type args struct {
		key string
		val []byte
	}
	c := New()
	tests := []struct {
		name      string
		c         *Cache
		args      args
		wantData  []byte
		wantFound bool
	}{
		{"Should pass", c,
			args{"some_key", []byte("value")},
			[]byte("value"), true},
		{"empty key error", c,
			args{"", []byte("value")},
			[]byte(nil), false},
		{"key already exists error", c,
			args{"some_key", []byte("new value")},
			[]byte("value"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Add(tt.args.key, tt.args.val)
			gotData, gotFound := tt.c.Get(tt.args.key)
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Cache.Get() gotData = %v, want %v", gotData, tt.wantData)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Cache.Get() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
