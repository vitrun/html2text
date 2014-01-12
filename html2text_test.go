/**
 * Copyright ©2014-01-12 Alex <zhirun.yu@duitang.com>
 */
package html2text

import (
    "testing"
    "strings"
    )

func TestText(t *testing.T) {
    raw_html := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/
                bar/baz">BarBaz</a></ul>`
    text := Text(strings.NewReader(raw_html))
    right := "\nLinks:\nFoo\nBarBaz"
    if text != right {
        t.Error("Expected " + right + ", got " + text)
    }
}
