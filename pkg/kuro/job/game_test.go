package job

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/pkg/kuro/api"
)

func TestSignGameRecords_String(t *testing.T) {
	tests := []SignGameRecords{
		{
			Error: fmt.Errorf("list game role error"),
		},
		{
			Records: []SignGameRecord{
				{
					GameId:     api.GameIdMC,
					ServerName: "鸣潮",
					RoleName:   "鸣潮",
					IsSuccess:  true,
					Award:      "星声*9999",
				},
			},
		},
	}
	for i, tt := range tests {
		for j := range tt.Records {
			tt.Records[j].GameName = api.GameNameById[tt.Records[j].GameId]
		}
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testutil.Log(t, "\n"+tt.String())
		})
	}
}
