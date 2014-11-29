PocketDAV

Based on rbastic/webdav, a fork of gogits/webdav, for core WebDAV
implementation.  Uses facebookgo/grace for graceful in-place restarts without
interrupting user experience, and github.com/golang/glog for leveled execution
logs.

This was initially intended to be a practice ground for integrating
facebookgo/grace -- lucky us, we've now got WebDAV.

Quick start:

go build
... install all the dependencies, sorry
./pocketdav

cURL examples:

$ curl -v -X PUT http://localhost:8081/webdav/test.txt --data-ascii "ohnoeszzt"

$ curl -v -X GET http://localhost:8081/webdav/test.txt

--

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

