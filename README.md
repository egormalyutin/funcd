# funcd
funcd is deamon for functional keys (works without X11). It can:
- Decrease brightness
- Increase brightness
- Toggle brightness
- Toggle touchpad
- Mute volume
- Decrease volume
- Increase volume

## Installation

### Dependencies
- amixer
- xinput

### Building from source
```
go get github.com/malyutinegor/funcd
cd $GOPATH/src/github.com/malyutinegor/funcd
make
sudo make install
```
