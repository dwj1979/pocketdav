PocketDAV

Based on rbastic/webdav, a fork of gogits/webdav, for core WebDAV implementation.
Uses facebookgo/grace for graceful in-place restarts without interrupting user experience.

Quick start:

go build
...
./pocketdav

cURL examples:

$ curl -v -X PUT http://localhost:8081/webdav/test.txt --data-ascii "ohnoeszzt"
$ curl -v -X GET http://localhost:8081/webdav/test.txt

