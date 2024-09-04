package api

import (
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/sign-task/pkg/geetest"
)

type ListPostData struct {
	IsLast   bool        `json:"is_last"`
	IsOrigin bool        `json:"is_origin"`
	LastId   string      `json:"last_id"`
	List     []*PostData `json:"list"`
}

type PostData struct {
	Post          *PostInfo          `json:"post"`
	Stat          *PostStat          `json:"stat"`
	User          *PostUser          `json:"user"`
	SelfOperation *PostSelfOperation `json:"self_operation"`
}

func (p *PostData) IsUpvote() bool {
	return p != nil && p.SelfOperation != nil && p.SelfOperation.Attitude == 1
}

func (p *PostData) IsCollected() bool {
	return p != nil && p.SelfOperation != nil && p.SelfOperation.IsCollected
}

type PostInfo struct {
	PostId    string `json:"post_id"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
	MaxFloor  int    `json:"max_floor"`
	ReplyTime string `json:"reply_time"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type PostStat struct {
	BookmarkNum int `json:"bookmark_num"`
	ForwardNum  int `json:"forward_num"`
	LikeNum     int `json:"like_num"`
	ReplyNum    int `json:"reply_num"`
	ViewNum     int `json:"view_num"`
}

type PostUser struct {
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
}

type PostSelfOperation struct {
	Attitude    int  `json:"attitude"`
	IsCollected bool `json:"is_collected"`
	UpvoteType  int  `json:"upvote_type"`
}

func (c *Client) ListPost(forumId, lastId string) (*ListPostData, error) {
	query := gh.MS{"forum_id": forumId, "is_good": "false", "is_hot": "false", "sort_type": "1", "last_id": lastId, "page_size": "10"}
	req := c.R().SetCookies(c.sToken()).SetQueryParams(query)
	return Exec[*ListPostData](req, "GET", AddrTakumi+"/post/api/getForumPostList")
}

func (c *Client) ListFeedPost(gameId string) (*ListPostData, error) {
	query := gh.MS{"gids": gameId, "fresh_action": "1", "is_first_initialize": "true"}
	req := c.R().SetCookies(c.sToken()).SetQueryParams(query)
	return Exec[*ListPostData](req, "GET", AddrBBS+"/post/api/feeds/posts")
}

type GetPostData struct {
	Post *PostData `json:"post"`
}

func (c *Client) GetPost(postId string) (*GetPostData, error) {
	req := c.R().SetCookies(c.sToken()).SetQueryParam("post_id", postId)
	return Exec[*GetPostData](req, "GET", AddrBBS+"/post/api/getPostFull")
}

func (c *Client) UpvotePost(postId string, cancel bool) error {
	req := c.R().SetCookies(c.sToken()).SetBody(gh.M{"post_id": postId, "upvote_type": gh.Ternary(cancel, "0", "1"), "is_cancel": cancel})
	_, err := Exec[any](req, "POST", AddrBBS+"/post/api/post/upvote")
	return err
}

func (c *Client) CollectPost(postId string, cancel bool) error {
	req := c.R().SetCookies(c.sToken()).SetBody(gh.M{"post_id": postId, "is_cancel": cancel})
	_, err := Exec[any](req, "POST", AddrBBS+"/post/api/collectPost")
	return err
}

type SharePostData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Icon    string `json:"icon"`
	Url     string `json:"url"`
}

func (c *Client) SharePost(postId string) (*SharePostData, error) {
	req := c.R().SetCookies(c.sToken()).SetQueryParam("entity_type", "1").SetQueryParam("entity_id", postId)
	return Exec[*SharePostData](req, "GET", AddrBBS+"/apihub/api/getShareConf")
}

type SignForumData struct {
	Points int `json:"points"`
}

func (c *Client) SignForum(gameId string, gt *geetest.V3Data) (*SignForumData, error) {
	req := c.R().SetCookies(c.sToken()).SetBody(gh.M{"gids": gameId})
	return Exec[*SignForumData](req, "POST", AddrBBS+"/apihub/app/api/signIn", SignDS2, gt)
}

type GetSignForumData struct {
	IsSigned bool `json:"is_signed"`
}

func (c *Client) GetSignForum(gameId string) (*GetSignForumData, error) {
	req := c.R().SetCookies(c.sToken()).SetQueryParam("gids", gameId)
	return Exec[*GetSignForumData](req, "GET", AddrBBS+"/apihub/sapi/querySignInStatus")
}
