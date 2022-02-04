# Wallpaperize

Util for setting wallpapers from most famous sources

## Installation

### Mac OS

Amd64:

```bash
  curl -L https://github.com/goncharovnikita/wallpaperize/releases/latest/download/wallpaperize-darwin-amd64.tar.gz > ./wallpaperize.tar.gz
  tar -xvzf ./wallpaperize.tar.gz
  mv ./wallpaperize /usr/local/bin/wallpaperize
  rm ./wallpaperize.tar.gz
```

Arm64:

```bash
  curl -L https://github.com/goncharovnikita/wallpaperize/releases/latest/download/wallpaperize-darwin-arm64.tar.gz > ./wallpaperize.tar.gz
  tar -xvzf ./wallpaperize.tar.gz
  mv ./wallpaperize /usr/local/bin/wallpaperize
  rm ./wallpaperize.tar.gz
```

## Usage

```bash
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
