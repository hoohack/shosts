main:
	go build cmd/main.go
install:
	mv main /usr/bin/shosts
clean:
	rm -rf main test
