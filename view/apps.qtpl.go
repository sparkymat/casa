// Code generated by qtc from "apps.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/apps.qtpl:1
package view

//line view/apps.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/apps.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/apps.qtpl:1
func StreamApps(qw422016 *qt422016.Writer) {
//line view/apps.qtpl:1
	qw422016.N().S(`
  <h3>apps</h3>
`)
//line view/apps.qtpl:3
}

//line view/apps.qtpl:3
func WriteApps(qq422016 qtio422016.Writer) {
//line view/apps.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/apps.qtpl:3
	StreamApps(qw422016)
//line view/apps.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line view/apps.qtpl:3
}

//line view/apps.qtpl:3
func Apps() string {
//line view/apps.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line view/apps.qtpl:3
	WriteApps(qb422016)
//line view/apps.qtpl:3
	qs422016 := string(qb422016.B)
//line view/apps.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line view/apps.qtpl:3
	return qs422016
//line view/apps.qtpl:3
}