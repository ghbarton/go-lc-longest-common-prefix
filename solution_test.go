package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name string
		data []string
		exp  string
	}{
		{
			name: "returns empty string",
			data: []string{},
			exp:  "",
		},
		{
			name: "returns empty string",
			data: []string{"dog", "racecar", "car"},
			exp:  "",
		},
		{
			name: "returns fl string",
			data: []string{"flower", "flow", "flight"},
			exp:  "fl",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			audit := solution(tc.data)
			assert.Equal(t, tc.exp, audit)
		})
	}
}
