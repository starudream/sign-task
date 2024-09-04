package job

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/pkg/douyu/api"
)

func TestRenewalRecords_String(t *testing.T) {
	tests := []RenewalRecord{
		{
			Error: fmt.Errorf("refresh error"),
		},
		{
			Badges1: []*api.Badge{
				{
					Name:     "aaa",
					Level:    1,
					Intimacy: 100,
				},
				{
					Name:     "bbb",
					Level:    20,
					Intimacy: 7890,
				},
			},
			Gifts1: []*api.Gift{
				{
					Id:    268,
					Name:  "粉丝荧光棒",
					Count: 100,
				},
			},
			Badges2: []*api.Badge{
				{
					Name:     "aaa",
					Level:    2,
					Intimacy: 200,
				},
				{
					Name:     "bbb",
					Level:    20,
					Intimacy: 8890,
				},
			},
			Gifts2: []*api.Gift{},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testutil.Log(t, "\n"+tt.String())
		})
	}
}
