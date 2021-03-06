package gquery_test

import (
	"github.com/wusuluren/gquery"
	"testing"
)

func printMarkdownNodeList(t *testing.T, nodeList []*gquery.MarkdownNode) {
	for _, node := range nodeList {
		t.Log(node)
		printMarkdownNodeList(t, node.Children(gquery.MdAll))
	}
}

func TestParseMarkdown(t *testing.T) {
	testData := `
# Title
This is title
- baidu
http://www.baidu.com
	- baidu
	http://www.baidu.com
- google
http://www.google.com
`
	gq := gquery.NewMarkdown(testData)
	children := gq.Gquery(gquery.MdAll)
	t.Log(len(children))
	printMarkdownNodeList(t, children)

	t.Log("test search")
	t.Log(gq.Gquery(gquery.MdTitle)[0])
	t.Log(gq.Gquery(gquery.MdUnorderList)[0].First(gquery.MdUnorderList) ==
		gq.Gquery(gquery.MdUnorderList)[0].Last(gquery.MdUnorderList))

	node := gquery.NewMarkdownNode(map[string]interface{}{
		"type":  gquery.MdUnorderList,
		"value": "test",
		"text":  "test",
		"html":  "test",
	})
	t.Log(node)
	gq.Gquery(gquery.MdUnorderList)[0].Append(node)
	t.Log(gq.Gquery(gquery.MdUnorderList)[0].Last(gquery.MdUnorderList) == node)
}
