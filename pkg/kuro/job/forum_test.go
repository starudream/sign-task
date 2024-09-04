package job

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/pkg/kuro/api"
)

func TestSignForumRecords_String(t *testing.T) {
	tests := []SignForumRecords{
		{
			Records: []SignForumRecord{
				{
					GameId: api.GameIdPNS,
					Error:  fmt.Errorf("list post error"),
				},
				{
					GameId:    api.GameIdMC,
					HasSigned: true,
					PostView:  3,
					PostLike:  5,
					PostShare: 1,
					LoopCount: 1,
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
