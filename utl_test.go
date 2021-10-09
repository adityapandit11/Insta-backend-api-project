package nakama

import (
	"reflect"
	"testing"
)

func Test_collectTags(t *testing.T) {
	tt := []struct {
		name  string
		given string
		want  []string
	}{
		{
			given: "#tag",
			want:  []string{"tag"},
		},
		{
			given: "foo #tag bar",
			want:  []string{"tag"},
		},
		{
			given: "#unique #unique",
			want:  []string{"unique"},
		},
		{
			given: "#tág",
			want:  []string{"tág"},
		},
		{
			given: "#123",
			want:  []string{"123"},
		},
		{
			given: "#世界",
			want:  []string{"世界"},
		},
		{
			given: "##",
			want:  nil,
		},
		{
			given: "###",
			want:  nil,
		},
		{
			given: "#nope#",
			want:  nil,
		},
		{
			given: "nope",
			want:  nil,
		},
		{
			given: "https://example.org#nope",
			want:  nil,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := collectTags(tc.given)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("collectTags(%q) = %+v, want %+v", tc.given, got, tc.want)
			}
		})
	}
}
