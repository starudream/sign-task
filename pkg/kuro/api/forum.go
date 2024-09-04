package api

import (
	"strconv"

	"github.com/starudream/go-lib/core/v2/gh"
)

type ListPostData struct {
	HasNext  int         `json:"hasNext"`
	PostList []*PostInfo `json:"postList"`
}

type PostInfo struct {
	GameId          int       `json:"gameId"`
	GameName        string    `json:"gameName"`
	GameForumId     int       `json:"gameForumId"`
	PostId          string    `json:"postId"`
	PostType        int       `json:"postType"`
	PostUserId      string    `json:"postUserId"` // detail
	PostTitle       string    `json:"postTitle"`
	PostContent     any       `json:"postContent"`
	UserId          string    `json:"userId"` // list
	UserName        string    `json:"userName"`
	BrowseCount     string    `json:"browseCount"`
	CommentCount    int       `json:"commentCount"`
	LikeCount       int       `json:"likeCount"`
	IsFollow        int       `json:"isFollow"`
	IsLike          int       `json:"isLike"`
	IsLock          int       `json:"isLock"`
	IsPublisher     int       `json:"isPublisher"`
	CreateTimestamp Timestamp `json:"createTimestamp"`
}

func (c *Client) ListPost(gid, fid, page int) (*ListPostData, error) {
	// searchType 1最新发布 2最新回复 3默认
	req := c.R().SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "forumId": strconv.Itoa(fid), "searchType": "1", "pageIndex": strconv.Itoa(page), "pageSize": "20"})
	return Exec[*ListPostData](req, "POST", "/forum/list")
}

func (c *Client) GetPost(id string) (*GetPostData, error) {
	req := c.R().SetFormData(gh.MS{"postId": id})
	return Exec[*GetPostData](req, "POST", "/forum/getPostDetail")
}

type GetPostData struct {
	PostDetail *PostInfo `json:"postDetail"`
	IsCollect  int       `json:"isCollect"`
	IsFollow   int       `json:"isFollow"`
	IsLike     int       `json:"isLike"`
}

func (c *Client) LikePost(gid, fid int, pid, uid string, cancel bool) error {
	req := c.R().SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "forumId": strconv.Itoa(fid), "likeType": "1", "operateType": gh.Ternary(cancel, "2", "1"), "postId": pid, "postType": "1", "toUserId": uid})
	_, err := Exec[bool](req, "POST", "/forum/like")
	return err
}

func (c *Client) SharePost(gid int) error {
	req := c.R().SetFormData(gh.MS{"gameId": strconv.Itoa(gid)})
	_, err := Exec[*GetPostData](req, "POST", "/encourage/level/shareTask")
	return err
}

type SignForumData struct {
	GainVoList   []*GainVo `json:"gainVoList"`
	ContinueDays int       `json:"continueDays"`
}

type GainVo struct {
	GainTyp   int `json:"gainTyp"`
	GainValue int `json:"gainValue"`
}

func (c *Client) SignForum(gid int) (*SignForumData, error) {
	req := c.R().SetFormData(gh.MS{"gameId": strconv.Itoa(gid)})
	return Exec[*SignForumData](req, "POST", "/user/signIn")
}

type GetSignForumData struct {
	ContinueDays int  `json:"continueDays"`
	HasSignIn    bool `json:"hasSignIn"`
}

func (c *Client) GetSignForum(gid int) (*GetSignForumData, error) {
	req := c.R().SetFormData(gh.MS{"gameId": strconv.Itoa(gid)})
	return Exec[*GetSignForumData](req, "POST", "/user/signIn/info")
}
