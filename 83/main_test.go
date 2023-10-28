// 83. Remove Duplicates from Sorted List

package main

import (
	"testing"

	"github.com/dkandr/leetcode/list"
	"github.com/stretchr/testify/assert"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "empty",
			args: nil,
			want: nil,
		},
		{
			name: "test_1",
			args: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "test_2",
			args: []int{1, 1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "test_2",
			args: []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, deleteDuplicates(list.SliceToList(tt.args)), list.SliceToList(tt.want))
		})
	}
}
