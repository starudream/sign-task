package job

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/pkg/douyu/api"
)

func TestRefreshRecord_String(t *testing.T) {
	tests := []RefreshRecord{
		{
			Error: fmt.Errorf("refresh error"),
		},
		{
			Gifts1: []*api.Gift{
				{
					Id:    268,
					Name:  "粉丝荧光棒",
					Count: 100,
				},
				{
					Id:    999,
					Name:  "AAA",
					Count: 1000,
				},
			},
			Gifts2: []*api.Gift{
				{
					Id:    268,
					Name:  "粉丝荧光棒",
					Count: 200,
				},
				{
					Id:    999,
					Name:  "AAA",
					Count: 2000,
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
