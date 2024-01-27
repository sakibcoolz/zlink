package domain

import (
	"reflect"
	"sync"
	"testing"
	"zlink/log"
	"zlink/model"
)

func TestStore_SetStack(t *testing.T) {
	memStore := NewMemoryStore(map[string]string{"xyz": "xyz"}, new(sync.Mutex))
	mapRevStore := NewMappingRev(make(map[string]string), new(sync.Mutex))
	cntStore := NewCountStore(0, new(sync.Mutex))
	collectCount := NewUrlCollectionCount(model.Collections{URLs: []string{}, Counts: []int{}}, new(sync.Mutex))
	lg := log.New()

	type fields struct {
		log         *log.Log
		ms          *model.MemoryStore
		sc          *model.CountStore
		mr          *model.MappingRev
		collections *model.URLCountCollections
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Pass1",
			fields: fields{
				log:         lg,
				ms:          memStore,
				sc:          cntStore,
				mr:          mapRevStore,
				collections: collectCount,
			},
			args: args{
				path: "xyz",
			},
		},
		{
			name: "Pass2",
			fields: fields{
				log:         lg,
				ms:          memStore,
				sc:          cntStore,
				mr:          mapRevStore,
				collections: collectCount,
			},
			args: args{
				path: "xyz",
			},
		},
		{
			name: "failed1",
			fields: fields{
				log:         lg,
				ms:          memStore,
				sc:          cntStore,
				mr:          mapRevStore,
				collections: collectCount,
			},
			args: args{
				path: "yz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				log:         tt.fields.log,
				ms:          tt.fields.ms,
				sc:          tt.fields.sc,
				mr:          tt.fields.mr,
				collections: tt.fields.collections,
			}
			s.SetStack(tt.args.path)
		})
	}
}

func TestStore_GetMostUrl(t *testing.T) {
	memStore := NewMemoryStore(map[string]string{"xyz": "xyz"}, new(sync.Mutex))
	mapRevStore := NewMappingRev(make(map[string]string), new(sync.Mutex))
	cntStore := NewCountStore(0, new(sync.Mutex))
	collectCount := NewUrlCollectionCount(model.Collections{URLs: []string{"xyz", "abc", "sfd"}, Counts: []int{12, 34, 56}}, new(sync.Mutex))
	lg := log.New()
	type fields struct {
		log         *log.Log
		ms          *model.MemoryStore
		sc          *model.CountStore
		mr          *model.MappingRev
		collections *model.URLCountCollections
	}
	type args struct {
		top int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]int
	}{
		{
			name: "Pass",
			fields: fields{
				log:         lg,
				ms:          memStore,
				sc:          cntStore,
				mr:          mapRevStore,
				collections: collectCount,
			},
			args: args{
				top: 2,
			},
			want: map[string]int{
				"sfd": 56,
				"abc": 34,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				log:         tt.fields.log,
				ms:          tt.fields.ms,
				sc:          tt.fields.sc,
				mr:          tt.fields.mr,
				collections: tt.fields.collections,
			}
			if got := s.GetMostUrl(tt.args.top); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.GetMostUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
