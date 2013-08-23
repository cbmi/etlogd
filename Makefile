all:
	go get -d -v github.com/bruth/assert
	go get -d -v labix.org/v2/mgo
	go get -d code.google.com/p/gcfg
	go build -v
