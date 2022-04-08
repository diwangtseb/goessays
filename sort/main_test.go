package main

import (
	"reflect"
	"testing"
)

func TestQuicklySort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "quicklysort",
			args: args{in: []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}},
			want: []int{1, 2, 5, 8, 9, 10, 12, 30, 45, 63, 234},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuicklySort(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuicklySort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQuicklySort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuicklySort([]int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12})
	}
}

func TestBubbleSort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "bubbelsort",
			args: args{in: []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}},
			want: []int{1, 2, 5, 8, 9, 10, 12, 30, 45, 63, 234},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort([]int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12})
	}
}
