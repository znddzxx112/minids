all:	install

install:
	go build minids.go
	go build minidsd.go

clean:
	rm -f minids minidsd

