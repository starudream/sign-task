package job

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/sign-task/pkg/miyoushe/api"
)

func TestSignForumRecord_String(t *testing.T) {
	tests := []SignForumRecord{
		{
			Error: fmt.Errorf("no businesses"),
		},
		{
			GameId:     "2",
			Verify:     1,
			Points:     50,
			PostView:   3,
			PostUpvote: 10,
			PostShare:  1,
			SignErr:    fmt.Errorf("verify max retry and give up"),
		},
		{
			GameId:     "6",
			IsSuccess:  true,
			Verify:     1,
			Points:     50,
			PostView:   3,
			PostUpvote: 10,
			PostShare:  1,
		},
	}
	for i, tt := range tests {
		tt.GameName = api.GameCNNameById[tt.GameId]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testutil.Log(t, "\n"+tt.String())
		})
	}
}
