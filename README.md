PocketDAV - A tiny 'webdav' server written in Go.

This was initially intended to be a small practice ground for integrating
facebookgo/grace -- I've since added glog and cleaned up a few 'golint'
complaints in the original webdav server at https://github.com/rbastic/webdav
and am hoping to get everything merged upstream.

Quick start:

```
go build
./pocketdav
```

cURL examples:

```
$ curl -v -X PUT http://localhost:8081/webdav/test.txt --data-ascii "ohnoeszzt"

$ curl -v -X GET http://localhost:8081/webdav/test.txt
```

--

```
ml36806:pocketdav rbastic$ ./pocketdav -h
Usage of ./pocketdav:
  -alsologtostderr=false: log to standard error as well as files
  -docroot="www": document root folder to serve from
  -gracehttp.log=true: Enable logging.
  -httpport="8080": port to listen on
  -log_backtrace_at=:0: when logging hits line file:N, emit a stack trace
  -log_dir="": If non-empty, write log files in this directory
  -logtostderr=false: log to standard error instead of files
  -stderrthreshold=0: logs at or above this threshold go to stderr
  -v=0: log level for V logs
  -vmodule=: comma-separated list of pattern=N settings for file-filtered logging
```
