PocketDAV

Based on rbastic/webdav, a fork of gogits/webdav, for core WebDAV
implementation.  Uses facebookgo/grace for graceful in-place restarts without
interrupting user experience.

This was initially intended to be a practice ground for integrating
facebookgo/grace -- lucky us, we've now got WebDAV.

Quick start:

go build
... install all the dependencies, sorry
./pocketdav

cURL examples:

$ curl -v -X PUT http://localhost:8081/webdav/test.txt --data-ascii "ohnoeszzt"
$ curl -v -X GET http://localhost:8081/webdav/test.txt

TODO:
- add godep
