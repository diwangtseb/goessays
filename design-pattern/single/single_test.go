package single

import (
	"reflect"
	"testing"
)

func TestNewSingleTone(t *testing.T) {
	tests := []struct {
		name string
		want *SingleTon
	}{
		// TODO: Add test cases.
		{
			name: "",
			want: &SingleTon{},
		},
		{
			name: "",
			want: &SingleTon{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingleTone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingleTone() = %v, want %v", got, tt.want)
			}
		})
	}
}
