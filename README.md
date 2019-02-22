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

#### Install on Arch Linux
Just install AUR package [funcd-git](https://aur.archlinux.org/packages/funcd-git/) and enable systemd service:
```bash
sudo systemctl enable funcd.service
```

#### Building from source
```bash
go get github.com/malyutinegor/funcd
cd $GOPATH/src/github.com/malyutinegor/funcd
go build
sudo cp ./funcd /usr/bin/
sudo cp ./funcd.service /etc/systemd/system/
sudo systemctl enable funcd.service
```
