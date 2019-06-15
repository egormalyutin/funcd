# funcd
funcd is daemon for functional keys (works without X11). It can:
- Decrease brightness
- Increase brightness
- Toggle brightness
- Toggle touchpad
- Mute volume
- Decrease volume
- Increase volume

### Status of project
All basic functional is already done, but you can feel free to open an issue with feature request or/and pull request :)

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
