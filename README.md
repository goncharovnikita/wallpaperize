# Wallpaperize
Util for setting wallpapers from most famous sources

## Supported platforms

- Ubuntu 
- - Gnome
- Mac OS X

## Installing

```go
  go get github.com/goncharovnikita/wallpaperize
  cd ~/go/src/github.com/goncharovnikita/wallpaperize
  go install
```

## Usage

```go
  // Show all usage
  wallpaperize --help

  // Set daily image from bing as wallpaper
  wallpaperize daily

  // Set random high-quality photo from random source as wallpaper
  wallpaperize random

  // Get info about disk usage
  wallpaperize info

  // If you bored or angry, set your previous wallpaper
  wallpaperize restore
```

## TODO
- [x] Save initial image and allow to rollback
- [x] Support other platforms
- [ ] Add more sources
- [ ] Support as many platforms as possible
- [ ] Provide useful API
- [ ] Daemonize
- [ ] Add UI