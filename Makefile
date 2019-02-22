default: build

build:
	go build

build-cross:
	gox -os="linux"

install:
	echo Hi!
