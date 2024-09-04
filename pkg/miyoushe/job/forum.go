package job

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/sign-task/pkg/geetest"
	"github.com/starudream/sign-task/pkg/miyoushe/api"
	"github.com/starudream/sign-task/pkg/miyoushe/config"
	"github.com/starudream/sign-task/util"
)

const (
	verifyRetry = 10

	postView   = 3
	postUpvote = 10
	postShare  = 1
	postLoop   = 10
)

type SignForumRecord struct {
	GameId     string
	GameName   string
	HasSigned  bool
	IsRisky    bool
	IsSuccess  bool
	Verify     int
	Points     int
	PostView   int
	PostUpvote int
	PostShare  int
	LoopCount  int
	Error      error
	SignErr    error
}

//go:embed forum.tmpl
var forumTplRaw string

var forumTpl = template.Must(template.New("miyoushe-sign-forum").Parse(forumTplRaw))

func (r SignForumRecord) String() string {
	val := util.ToMap[any](r)
	val["TotalPostView"] = postView
	val["TotalPostUpvote"] = postUpvote
	val["TotalPostShare"] = postShare
	buf := &bytes.Buffer{}
	_ = forumTpl.Execute(buf, val)
	return strings.TrimSpace(buf.String())
}

func SignForum(account config.Account) (record SignForumRecord) {
	c := api.NewClient(account)

	businesses, err := c.GetBusinesses()
	if err != nil {
		record.Error = fmt.Errorf("get bussinesses error: %w", err)
		return
	}

	if len(businesses.Businesses) == 0 {
		record.Error = fmt.Errorf("no businesses")
		return
	}

	record.GameId = businesses.Businesses[0]
	record.GameName = api.GameCNNameById[record.GameId]

	today, err := c.GetSignForum(record.GameId)
	if err != nil {
		record.Error = fmt.Errorf("get sign forum error: %w", err)
		return
	}

	var (
		gt  *geetest.V3Data
		sfd *api.SignForumData
	)

	if today.IsSigned {
		record.HasSigned = true
		goto post
	}

sign:

	sfd, err = c.SignForum(record.GameId, gt)
	if err != nil {
		if api.IsRetCode(err, api.RetCodeForumHasSigned) {
			record.HasSigned = true
		} else if api.IsRetCode(err, api.RetCodeForumNeedGeetest) {
			record.IsRisky = true
		verify:
			record.Verify++
			gt, err = verify(c)
			if err == nil {
				goto sign
			} else {
				slog.Error("verify error: %v", err)
				if record.Verify < verifyRetry {
					slog.Info("retry verify, count: %d", record.Verify)
					goto verify
				} else {
					record.SignErr = fmt.Errorf("verify max retry and give up")
					goto post
				}
			}
		} else {
			record.SignErr = fmt.Errorf("sign forum error: %w", err)
			goto post
		}
	} else {
		record.IsSuccess = true
	}

	if se := record.SignErr; se != nil {
		slog.Error(se.Error())
	}

	if sfd != nil {
		record.Points = sfd.Points
	}

post:

	record.LoopCount++

	posts, err := c.ListFeedPost(record.GameId)
	if err != nil {
		record.Error = fmt.Errorf("list feed post error: %w", err)
		return
	}

	for i := 0; i < len(posts.List); i++ {
		p := posts.List[i]
		pid := p.Post.PostId
		if record.PostView < postView {
			_, e := c.GetPost(pid)
			if e != nil {
				slog.Error("get post error: %v", e)
				continue
			}
			record.PostView++
			time.Sleep(500 * time.Millisecond)
		}
		if record.PostUpvote < postUpvote && !p.IsUpvote() {
			e := c.UpvotePost(pid, false)
			if e != nil {
				slog.Error("upvote post error: %v", e)
				continue
			}
			time.Sleep(500 * time.Millisecond)
			_ = c.UpvotePost(pid, true)
			record.PostUpvote++
			time.Sleep(500 * time.Millisecond)
		}
		if record.PostShare < postShare {
			_, e := c.SharePost(pid)
			if e != nil {
				slog.Error("share post error: %v", e)
				continue
			}
			record.PostShare++
			time.Sleep(500 * time.Millisecond)
		}
		time.Sleep(500 * time.Millisecond)
	}

	if record.LoopCount < postLoop && (record.PostView < postView || record.PostUpvote < postUpvote || record.PostShare < postShare) {
		goto post
	}

	return
}
