package job

import (
	"bytes"
	_ "embed"
	"fmt"
	"math/rand/v2"
	"strings"
	"text/template"
	"time"

	"github.com/starudream/sign-task/pkg/tieba/api"
	"github.com/starudream/sign-task/pkg/tieba/config"
	"github.com/starudream/sign-task/util"
)

type SignForumRecord struct {
	Success []string
	SignErr map[string]error
	Error   error
}

//go:embed forum.tmpl
var forumTplRaw string

var forumTpl = template.Must(template.New("tieba-sign-forum").Parse(forumTplRaw))

func (r SignForumRecord) String() string {
	val := util.ToMap[any](r)
	buf := &bytes.Buffer{}
	_ = forumTpl.Execute(buf, val)
	return strings.TrimSpace(buf.String())
}

func SignForum(account config.Account) (record SignForumRecord) {
	c := api.NewClient(account)

	forums, err := c.ListForum()
	if err != nil {
		record.Error = fmt.Errorf("list forum error: %w", err)
		return
	}

	if len(forums) == 0 {
		record.Error = fmt.Errorf("no liked forum")
		return
	}

	record.SignErr = map[string]error{}

	for _, forum := range forums {
		_, err = c.SignForum(forum.Name)
		if err != nil && !api.IsCode(err, api.CodeHasSigned) {
			record.SignErr[forum.Name] = err
		} else {
			record.Success = append(record.Success, forum.Name)
		}
		time.Sleep(2*time.Second + time.Duration(rand.Float64()*float64(time.Second)))
	}

	return
}
