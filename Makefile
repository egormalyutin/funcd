default: build

build:
	go build

run: build
	su -c "./funcd" # Exactly su -c, NOT sudo

install:
	echo Hi!
