package domain

import (
	"net/http"
	"sync"
	"testing"
	"zlink/log"
	"zlink/model"

	"github.com/gin-gonic/gin"
)

func TestStore_GetCounter(t *testing.T) {
	cntStore := NewCountStore(0, new(sync.Mutex))

	type fields struct {
		log         *log.Log
		ms          *model.MemoryStore
		sc          *model.CountStore
		mr          *model.MappingRev
		collections *model.URLCountCollections
	}

	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Pass1",
			fields: fields{
				sc: cntStore,
			},
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						Header: http.Header{
							"skm": []string{""},
						},
					},
				},
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
			if got := s.GetCounter(tt.args.ctx); got != tt.want {
				t.Errorf("Store.GetCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}
