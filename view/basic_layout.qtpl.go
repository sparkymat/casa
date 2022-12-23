// Code generated by qtc from "basic_layout.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/basic_layout.qtpl:1
package view

//line view/basic_layout.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/basic_layout.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/basic_layout.qtpl:1
func StreamBasicLayout(qw422016 *qt422016.Writer, title string, content string) {
//line view/basic_layout.qtpl:1
	qw422016.N().S(`
  <!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <meta name="viewport" content="width=device-width,initial-scale=1">
      <title>`)
//line view/basic_layout.qtpl:7
	qw422016.E().S(title)
//line view/basic_layout.qtpl:7
	qw422016.N().S(`</title>
      <link rel="stylesheet" type="text/css" href="/css/uikit.min.css">
    </head>
    <body>
      `)
//line view/basic_layout.qtpl:11
	qw422016.N().S(content)
//line view/basic_layout.qtpl:11
	qw422016.N().S(`
      <script src="/js/uikit.min.js"></script>
      <script src="/js/uikit-icons.min.js"></script>
    </body>
  </html>
`)
//line view/basic_layout.qtpl:16
}

//line view/basic_layout.qtpl:16
func WriteBasicLayout(qq422016 qtio422016.Writer, title string, content string) {
//line view/basic_layout.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/basic_layout.qtpl:16
	StreamBasicLayout(qw422016, title, content)
//line view/basic_layout.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line view/basic_layout.qtpl:16
}

//line view/basic_layout.qtpl:16
func BasicLayout(title string, content string) string {
//line view/basic_layout.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
//line view/basic_layout.qtpl:16
	WriteBasicLayout(qb422016, title, content)
//line view/basic_layout.qtpl:16
	qs422016 := string(qb422016.B)
//line view/basic_layout.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
//line view/basic_layout.qtpl:16
	return qs422016
//line view/basic_layout.qtpl:16
}
