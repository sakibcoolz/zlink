package domain

import (
	"reflect"
	"sync"
	"testing"
	"zlink/log"
	"zlink/model"
)

func TestNewStore(t *testing.T) {
	memStore := NewMemoryStore(map[string]string{"xyz": "xyz"}, new(sync.Mutex))
	mapRevStore := NewMappingRev(make(map[string]string), new(sync.Mutex))
	cntStore := NewCountStore(0, new(sync.Mutex))
	collectCount := NewUrlCollectionCount(model.Collections{URLs: []string{}, Counts: []int{}}, new(sync.Mutex))
	lg := log.New()

	type args struct {
		logger      *log.Log
		ms          *model.MemoryStore
		sc          *model.CountStore
		mr          *model.MappingRev
		collections *model.URLCountCollections
	}
	tests := []struct {
		name string
		args args
		want *Store
	}{
		{
			name: "Pass",
			args: args{
				logger:      lg,
				ms:          memStore,
				sc:          cntStore,
				mr:          mapRevStore,
				collections: collectCount,
			},
			want: &Store{
				log:         lg,
				ms:          memStore,
				sc:          cntStore,
				mr:          mapRevStore,
				collections: collectCount,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStore(tt.args.logger, tt.args.ms, tt.args.sc, tt.args.mr, tt.args.collections); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStore() = %v, want %v", got, tt.want)
			}
		})
	}
}
