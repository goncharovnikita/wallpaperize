# Wallpaperize
Util for setting wallpapers from most famous sources

## Supported platforms

- Ubuntu 
- - Gnome

## Installing

```go
  go get github.com/goncharovnikita/wallpaperize
  cd ~/go/src/github.com/goncharovnikita/wallpaperize
  go install
```

## Usage

```go
  // Set Bing's photo of the day as wallpaper
  wallpaperize

  // Set random high-quality photo as wallpaper
  wallpaperize -r
```

## TODO
- [ ] Save initial image and allow to rollback
- [ ] Support other platforms
- [ ] Add more sources