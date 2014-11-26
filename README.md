PocketDAV


Quick start:

go build
...
./pocketdav

cURL examples:

$ curl -v -X PUT http://localhost:8081/webdav/test.txt --data-ascii "ohnoeszzt"
$ curl -v -X GET http://localhost:8081/webdav/test.txt

