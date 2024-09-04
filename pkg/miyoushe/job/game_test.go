package job

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/pkg/miyoushe/api"
)

func TestSignGameRecords_String(t *testing.T) {
	tests := []SignGameRecords{
		{
			Error: fmt.Errorf("list game role error"),
		},
		{
			Records: []SignGameRecord{
				{
					GameId:   "2",
					RoleName: "米哈游",
					RoleUid:  "123456",
					Error:    fmt.Errorf("get home error"),
				},
				{
					GameId:    "2",
					RoleName:  "米哈游",
					RoleUid:   "123456",
					IsSuccess: true,
					Verify:    1,
					Award:     "原石*9999",
				},
				{
					GameId:    "6",
					RoleName:  "米哈游",
					RoleUid:   "123456",
					IsSuccess: true,
					Verify:    1,
					Award:     "星琼*9999",
				},
			},
		},
	}
	for i, tt := range tests {
		for j := range tt.Records {
			tt.Records[j].GameName = api.GameCNNameById[tt.Records[j].GameId]
		}
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testutil.Log(t, "\n"+tt.String())
		})
	}
}
