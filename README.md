PocketDAV - A tiny 'webdav' server written in Go.

 * It only supports GET, HEAD, PUT, and DELETE.
 * It supports in-place restarts. To upgrade to a new binary, you can
   'sudo mv pocketdav /usr/bin/pocketdav' and run kill -SIGUSR2 `/sbin/pidof
   pocketdav` to restart it.

    There are some potential CAVEATS about the current in-place restart approach.

 * It has extensive logging capabilities since it uses 'glog'. This is not yet
   being fully utilized, but a lot of logging code has been put in place already.
 * PUT /foobar/test.txt will automatically create /foobar/ if needed. This works
   for infinite levels of nested directories (e.g. whatever is allowed by the
   filesystem; os.MkdirAll() is called AKA the equivalent of mkdir -p in Go).
 * No caching takes place anywhere, so a DELETE always immediately takes effect.
 * Deleting entire folders/subfolders is not possible. This is considered to be
   a feature.

This was initially intended to be a small practice ground for integrating
facebookgo/grace. It has since become a bit more, but not much.

Quick start:

```
go build
./pocketdav
```

cURL examples:

```
$ curl -v -X PUT http://localhost:8081/webdav/test.txt --data-ascii "ohnoeszzt"
$ curl -v -X GET http://localhost:8081/webdav/test.txt
$ curl -v -X HEAD http://localhost:8081/webdav/test.txt
$ curl -v -X DELETE http://localhost:8081/webdav/test.txt
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
