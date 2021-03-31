package html

import "testing"

func TestRemoveStyle(t *testing.T) {
	content := `<img src=\"/i/eg_tulip.jpg\"  alt=\"上海鲜花港 - 郁金香\"/> 内同`
	tagArray := []string{"img"}

	s := RemoveStyle(content,tagArray)
	t.Log(s)
	s = FetchText(s,tagArray)
	t.Log(s)
}
