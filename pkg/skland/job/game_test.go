package job

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestSignGameRecords_String(t *testing.T) {
	tests := []SignGameRecords{
		{
			Error: fmt.Errorf("list player error"),
		},
		{
			Records: []SignGameRecord{
				{
					GameId:        "1",
					GameName:      "明日方舟",
					PlayerName:    "森空岛#1234",
					PlayerChannel: "官服",
					IsSuccess:     true,
					Award:         "合成玉*8888",
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testutil.Log(t, "\n"+tt.String())
		})
	}
}
