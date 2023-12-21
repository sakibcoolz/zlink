package domain

import (
	"sync"
	"testing"
	"zlink/model"

	"go.uber.org/zap"
)

func TestStore_GetCounter(t *testing.T) {
	cntStore := NewCountStore(0, new(sync.Mutex))

	type fields struct {
		log         *zap.Logger
		ms          *model.MemoryStore
		sc          *model.CountStore
		mr          *model.MappingRev
		collections *model.URLCountCollections
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Pass1",
			fields: fields{
				sc: cntStore,
			},
			want: 1,
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
			if got := s.GetCounter(); got != tt.want {
				t.Errorf("Store.GetCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}
