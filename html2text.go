/**
 * Copyright ©2014-01-12 Alex <zhirun.yu@duitang.com>
 */
package html2text

import (
    "code.google.com/p/go.net/html"
    "log"
    "bytes"
    "io"
    "strings"
)

func extract(node *html.Node, buff *bytes.Buffer){
    if node.Type == html.TextNode {
        data := strings.Trim(node.Data, "\r\n ")
        if data != "" {
            buff.WriteString("\n")
            buff.WriteString(data)
        }
    }
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        extract(c, buff)
    }
}

func Text(reader io.Reader) string {
    var buffer bytes.Buffer
    doc, err := html.Parse(reader)
    if err != nil {
        log.Fatal(err)
    }
    extract(doc, &buffer)
    return buffer.String()
}

