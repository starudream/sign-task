package util

import (
	"strings"

	"golang.org/x/net/html"
)

func NodeSearch(node *html.Node, fn func(node *html.Node) bool) *html.Node {
	if node == nil {
		return nil
	}
	if fn(node) {
		return node
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if n := NodeSearch(child, fn); n != nil {
			return n
		}
	}
	return nil
}

func NodeAttrSearch(node *html.Node, fn func(attr html.Attribute) bool) string {
	if node == nil {
		return ""
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		for _, attr := range child.Attr {
			if fn(attr) {
				return attr.Val
			}
		}
		if n := NodeAttrSearch(child, fn); n != "" {
			return n
		}
	}
	return ""
}

func NodeAttrExists(node *html.Node, fn func(attr html.Attribute) bool) bool {
	if node == nil {
		return false
	}
	for _, attr := range node.Attr {
		if fn(attr) {
			return true
		}
	}
	return false
}

func NodeChildren(node *html.Node, childTags ...string) (nodes []*html.Node) {
	childTag := func() string {
		if len(childTags) > 0 && childTags[0] != "" {
			return childTags[0]
		}
		return ""
	}()
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		data := strings.TrimSpace(child.Data)
		if data != "" && (childTag == "" || data == childTag) {
			nodes = append(nodes, child)
		}
	}
	return
}

func NodeTitle(node *html.Node) string {
	title := NodeSearch(node, func(node *html.Node) bool {
		return node.Type == html.ElementNode && strings.TrimSpace(node.Data) == "title"
	})
	return strings.TrimSpace(title.FirstChild.Data)
}
