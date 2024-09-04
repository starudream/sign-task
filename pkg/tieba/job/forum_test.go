package job

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestSignForumRecord_String(t *testing.T) {
	tests := []SignForumRecord{
		{
			Error: fmt.Errorf("list forum error"),
		},
		{
			Success: []string{"大主宰", "黑神话", "王者荣耀"},
			SignErr: map[string]error{
				"AAA": fmt.Errorf("错误信息111"),
				"BBB": fmt.Errorf("错误信息222"),
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testutil.Log(t, "\n"+tt.String())
		})
	}
}
