// Code generated by qtc from "metrics_find_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/graphite/metrics_find_response.qtpl:1
package graphite

//line app/vmselect/graphite/metrics_find_response.qtpl:1
import (
	"strings"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
)

// MetricsFindResponse generates response for /metrics/find .See https://graphite-api.readthedocs.io/en/latest/api.html#metrics-find

//line app/vmselect/graphite/metrics_find_response.qtpl:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/graphite/metrics_find_response.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/graphite/metrics_find_response.qtpl:11
func StreamMetricsFindResponse(qw422016 *qt422016.Writer, paths []string, delimiter, format string, addWildcards bool, jsonp string) {
//line app/vmselect/graphite/metrics_find_response.qtpl:12
	if jsonp != "" {
//line app/vmselect/graphite/metrics_find_response.qtpl:12
		qw422016.N().S(jsonp)
//line app/vmselect/graphite/metrics_find_response.qtpl:12
		qw422016.N().S(`(`)
//line app/vmselect/graphite/metrics_find_response.qtpl:12
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:13
	switch format {
//line app/vmselect/graphite/metrics_find_response.qtpl:14
	case "completer":
//line app/vmselect/graphite/metrics_find_response.qtpl:15
		streammetricsFindResponseCompleter(qw422016, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:16
	case "treejson":
//line app/vmselect/graphite/metrics_find_response.qtpl:17
		streammetricsFindResponseTreeJSON(qw422016, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:18
	default:
//line app/vmselect/graphite/metrics_find_response.qtpl:19
		logger.Panicf("BUG: unexpected format=%q", format)

//line app/vmselect/graphite/metrics_find_response.qtpl:20
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:21
	if jsonp != "" {
//line app/vmselect/graphite/metrics_find_response.qtpl:21
		qw422016.N().S(`)`)
//line app/vmselect/graphite/metrics_find_response.qtpl:21
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:22
}

//line app/vmselect/graphite/metrics_find_response.qtpl:22
func WriteMetricsFindResponse(qq422016 qtio422016.Writer, paths []string, delimiter, format string, addWildcards bool, jsonp string) {
//line app/vmselect/graphite/metrics_find_response.qtpl:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:22
	StreamMetricsFindResponse(qw422016, paths, delimiter, format, addWildcards, jsonp)
//line app/vmselect/graphite/metrics_find_response.qtpl:22
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:22
}

//line app/vmselect/graphite/metrics_find_response.qtpl:22
func MetricsFindResponse(paths []string, delimiter, format string, addWildcards bool, jsonp string) string {
//line app/vmselect/graphite/metrics_find_response.qtpl:22
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/graphite/metrics_find_response.qtpl:22
	WriteMetricsFindResponse(qb422016, paths, delimiter, format, addWildcards, jsonp)
//line app/vmselect/graphite/metrics_find_response.qtpl:22
	qs422016 := string(qb422016.B)
//line app/vmselect/graphite/metrics_find_response.qtpl:22
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:22
	return qs422016
//line app/vmselect/graphite/metrics_find_response.qtpl:22
}

//line app/vmselect/graphite/metrics_find_response.qtpl:24
func streammetricsFindResponseCompleter(qw422016 *qt422016.Writer, paths []string, delimiter string, addWildcards bool) {
//line app/vmselect/graphite/metrics_find_response.qtpl:24
	qw422016.N().S(`{"metrics":[`)
//line app/vmselect/graphite/metrics_find_response.qtpl:27
	for i, path := range paths {
//line app/vmselect/graphite/metrics_find_response.qtpl:27
		qw422016.N().S(`{"path":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:29
		qw422016.N().Q(path)
//line app/vmselect/graphite/metrics_find_response.qtpl:29
		qw422016.N().S(`,"name":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:30
		streammetricPathName(qw422016, path, delimiter)
//line app/vmselect/graphite/metrics_find_response.qtpl:30
		qw422016.N().S(`,"is_leaf":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:31
		if strings.HasSuffix(path, delimiter) {
//line app/vmselect/graphite/metrics_find_response.qtpl:31
			qw422016.N().S(`0`)
//line app/vmselect/graphite/metrics_find_response.qtpl:31
		} else {
//line app/vmselect/graphite/metrics_find_response.qtpl:31
			qw422016.N().S(`1`)
//line app/vmselect/graphite/metrics_find_response.qtpl:31
		}
//line app/vmselect/graphite/metrics_find_response.qtpl:31
		qw422016.N().S(`}`)
//line app/vmselect/graphite/metrics_find_response.qtpl:33
		if i+1 < len(paths) {
//line app/vmselect/graphite/metrics_find_response.qtpl:33
			qw422016.N().S(`,`)
//line app/vmselect/graphite/metrics_find_response.qtpl:33
		}
//line app/vmselect/graphite/metrics_find_response.qtpl:34
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:35
	if addWildcards && len(paths) > 1 {
//line app/vmselect/graphite/metrics_find_response.qtpl:35
		qw422016.N().S(`,{"name": "*"}`)
//line app/vmselect/graphite/metrics_find_response.qtpl:39
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:39
	qw422016.N().S(`]}`)
//line app/vmselect/graphite/metrics_find_response.qtpl:42
}

//line app/vmselect/graphite/metrics_find_response.qtpl:42
func writemetricsFindResponseCompleter(qq422016 qtio422016.Writer, paths []string, delimiter string, addWildcards bool) {
//line app/vmselect/graphite/metrics_find_response.qtpl:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:42
	streammetricsFindResponseCompleter(qw422016, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:42
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:42
}

//line app/vmselect/graphite/metrics_find_response.qtpl:42
func metricsFindResponseCompleter(paths []string, delimiter string, addWildcards bool) string {
//line app/vmselect/graphite/metrics_find_response.qtpl:42
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/graphite/metrics_find_response.qtpl:42
	writemetricsFindResponseCompleter(qb422016, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:42
	qs422016 := string(qb422016.B)
//line app/vmselect/graphite/metrics_find_response.qtpl:42
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:42
	return qs422016
//line app/vmselect/graphite/metrics_find_response.qtpl:42
}

//line app/vmselect/graphite/metrics_find_response.qtpl:44
func streammetricsFindResponseTreeJSON(qw422016 *qt422016.Writer, paths []string, delimiter string, addWildcards bool) {
//line app/vmselect/graphite/metrics_find_response.qtpl:44
	qw422016.N().S(`[`)
//line app/vmselect/graphite/metrics_find_response.qtpl:46
	for i, path := range paths {
//line app/vmselect/graphite/metrics_find_response.qtpl:46
		qw422016.N().S(`{`)
//line app/vmselect/graphite/metrics_find_response.qtpl:49
		allowChildren := "0"
		isLeaf := "1"
		if strings.HasSuffix(path, delimiter) {
			allowChildren = "1"
			isLeaf = "0"
		}

//line app/vmselect/graphite/metrics_find_response.qtpl:55
		qw422016.N().S(`"id":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:56
		qw422016.N().Q(path)
//line app/vmselect/graphite/metrics_find_response.qtpl:56
		qw422016.N().S(`,"text":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:57
		streammetricPathName(qw422016, path, delimiter)
//line app/vmselect/graphite/metrics_find_response.qtpl:57
		qw422016.N().S(`,"allowChildren":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:58
		qw422016.N().S(allowChildren)
//line app/vmselect/graphite/metrics_find_response.qtpl:58
		qw422016.N().S(`,"expandable":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:59
		qw422016.N().S(allowChildren)
//line app/vmselect/graphite/metrics_find_response.qtpl:59
		qw422016.N().S(`,"leaf":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:60
		qw422016.N().S(isLeaf)
//line app/vmselect/graphite/metrics_find_response.qtpl:60
		qw422016.N().S(`}`)
//line app/vmselect/graphite/metrics_find_response.qtpl:62
		if i+1 < len(paths) {
//line app/vmselect/graphite/metrics_find_response.qtpl:62
			qw422016.N().S(`,`)
//line app/vmselect/graphite/metrics_find_response.qtpl:62
		}
//line app/vmselect/graphite/metrics_find_response.qtpl:63
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:64
	if addWildcards && len(paths) > 1 {
//line app/vmselect/graphite/metrics_find_response.qtpl:64
		qw422016.N().S(`,{`)
//line app/vmselect/graphite/metrics_find_response.qtpl:67
		path := paths[0]
		for strings.HasSuffix(path, delimiter) {
			path = path[:len(path)-1]
		}
		id := ""
		if n := strings.LastIndexByte(path, delimiter[0]); n >= 0 {
			id = path[:n+1]
		}
		id += "*"

		allowChildren := "0"
		isLeaf := "1"
		for _, path := range paths {
			if strings.HasSuffix(path, delimiter) {
				allowChildren = "1"
				isLeaf = "0"
				break
			}
		}

//line app/vmselect/graphite/metrics_find_response.qtpl:86
		qw422016.N().S(`"id":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:87
		qw422016.N().Q(id)
//line app/vmselect/graphite/metrics_find_response.qtpl:87
		qw422016.N().S(`,"text": "*","allowChildren":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:89
		qw422016.N().S(allowChildren)
//line app/vmselect/graphite/metrics_find_response.qtpl:89
		qw422016.N().S(`,"expandable":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:90
		qw422016.N().S(allowChildren)
//line app/vmselect/graphite/metrics_find_response.qtpl:90
		qw422016.N().S(`,"leaf":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:91
		qw422016.N().S(isLeaf)
//line app/vmselect/graphite/metrics_find_response.qtpl:91
		qw422016.N().S(`}`)
//line app/vmselect/graphite/metrics_find_response.qtpl:93
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:93
	qw422016.N().S(`]`)
//line app/vmselect/graphite/metrics_find_response.qtpl:95
}

//line app/vmselect/graphite/metrics_find_response.qtpl:95
func writemetricsFindResponseTreeJSON(qq422016 qtio422016.Writer, paths []string, delimiter string, addWildcards bool) {
//line app/vmselect/graphite/metrics_find_response.qtpl:95
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:95
	streammetricsFindResponseTreeJSON(qw422016, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:95
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:95
}

//line app/vmselect/graphite/metrics_find_response.qtpl:95
func metricsFindResponseTreeJSON(paths []string, delimiter string, addWildcards bool) string {
//line app/vmselect/graphite/metrics_find_response.qtpl:95
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/graphite/metrics_find_response.qtpl:95
	writemetricsFindResponseTreeJSON(qb422016, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:95
	qs422016 := string(qb422016.B)
//line app/vmselect/graphite/metrics_find_response.qtpl:95
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:95
	return qs422016
//line app/vmselect/graphite/metrics_find_response.qtpl:95
}

//line app/vmselect/graphite/metrics_find_response.qtpl:97
func streammetricPathName(qw422016 *qt422016.Writer, path, delimiter string) {
//line app/vmselect/graphite/metrics_find_response.qtpl:99
	name := path
	for strings.HasSuffix(name, delimiter) {
		name = name[:len(name)-1]
	}
	if n := strings.LastIndexByte(name, delimiter[0]); n >= 0 {
		name = name[n+1:]
	}

//line app/vmselect/graphite/metrics_find_response.qtpl:107
	qw422016.N().Q(name)
//line app/vmselect/graphite/metrics_find_response.qtpl:108
}

//line app/vmselect/graphite/metrics_find_response.qtpl:108
func writemetricPathName(qq422016 qtio422016.Writer, path, delimiter string) {
//line app/vmselect/graphite/metrics_find_response.qtpl:108
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:108
	streammetricPathName(qw422016, path, delimiter)
//line app/vmselect/graphite/metrics_find_response.qtpl:108
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:108
}

//line app/vmselect/graphite/metrics_find_response.qtpl:108
func metricPathName(path, delimiter string) string {
//line app/vmselect/graphite/metrics_find_response.qtpl:108
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/graphite/metrics_find_response.qtpl:108
	writemetricPathName(qb422016, path, delimiter)
//line app/vmselect/graphite/metrics_find_response.qtpl:108
	qs422016 := string(qb422016.B)
//line app/vmselect/graphite/metrics_find_response.qtpl:108
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:108
	return qs422016
//line app/vmselect/graphite/metrics_find_response.qtpl:108
}