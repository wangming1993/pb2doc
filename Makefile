move:
	cp -r output/*.html  ~/Desktop/www

run:
	rm -rf $GOPATH/bin/pb2doc
	go get github.com/wangming1993/pb2doc
	pb2doc build pb from ./protos --dist=html

clean:
	rm -r html
