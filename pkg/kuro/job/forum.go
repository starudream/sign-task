package job

import (
	"bytes"
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"text/template"
	"time"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/sign-task/pkg/kuro/api"
	"github.com/starudream/sign-task/pkg/kuro/config"
	"github.com/starudream/sign-task/util"
)

const (
	postView  = 3
	postLike  = 5
	postShare = 1
	postLoop  = 3
)

type SignForumRecord struct {
	GameId    int
	GameName  string
	HasSigned bool
	IsSuccess bool
	PostView  int
	PostLike  int
	PostShare int
	LoopCount int
	Error     error
}

type SignForumRecords struct {
	Records []SignForumRecord
}

//go:embed forum.tmpl
var forumTplRaw string

var forumTpl = template.Must(template.New("kuro-sign-forum").Parse(forumTplRaw))

func (rs SignForumRecords) String() string {
	val := util.ToMap[any](rs)
	val["TotalPostView"] = postView
	val["TotalPostLike"] = postLike
	val["TotalPostShare"] = postShare
	buf := &bytes.Buffer{}
	_ = forumTpl.Execute(buf, val)
	return strings.TrimSpace(buf.String())
}

func SignForum(account config.Account) (records SignForumRecords) {
	c := api.NewClient(account)

	for id, name := range api.GameNameById {
		record := SignForumGame(c, id)
		record.GameId = id
		record.GameName = name
		records.Records = append(records.Records, record)
	}

	slices.SortFunc(records.Records, func(a, b SignForumRecord) int {
		return cmp.Compare(a.GameId, b.GameId)
	})

	return
}

func SignForumGame(c *api.Client, gid int) (record SignForumRecord) {
	today, err := c.GetSignForum(gid)
	if err != nil {
		record.Error = fmt.Errorf("get sign forum error: %w", err)
		return
	}

	record.HasSigned = today.HasSignIn

	if record.HasSigned {
		goto post
	}

	_, err = c.SignForum(gid)
	if err != nil {
		if api.IsCode(err, api.CodeHasSigned) {
			record.HasSigned = true
		} else {
			record.Error = fmt.Errorf("sign forum error: %w", err)
			return
		}
	} else {
		record.IsSuccess = true
	}

post:

	record.LoopCount++

	fid := api.ForumIdByGameId[gid]

	posts, err := c.ListPost(gid, fid, record.LoopCount)
	if err != nil {
		record.Error = fmt.Errorf("list post error: %w", err)
		return
	}

	for i := 0; i < len(posts.PostList); i++ {
		p := posts.PostList[i]
		if record.PostView < postView {
			_, e := c.GetPost(p.PostId)
			if e != nil {
				slog.Error("get post error: %v", e)
				continue
			}
			record.PostView++
			time.Sleep(500 * time.Millisecond)
		}
		if record.PostLike < postLike && p.IsLike == 0 {
			e := c.LikePost(gid, fid, p.PostId, p.UserId, false)
			if e != nil {
				slog.Error("like post error: %v", e)
				continue
			}
			time.Sleep(500 * time.Millisecond)
			_ = c.LikePost(gid, fid, p.PostId, p.UserId, true)
			record.PostLike++
			time.Sleep(500 * time.Millisecond)
		}
		if record.PostShare < postShare {
			e := c.SharePost(gid)
			if e != nil {
				slog.Error("share post error: %v", e)
				continue
			}
			record.PostShare++
			time.Sleep(500 * time.Millisecond)
		}
		time.Sleep(500 * time.Millisecond)
	}

	if record.LoopCount < postLoop && (record.PostView < postView || record.PostLike < postLike || record.PostShare < postShare) {
		goto post
	}

	return
}
