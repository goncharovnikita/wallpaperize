build:
	cd app && go build && cd ..

install:
	cd app && go build -o $(GOPATH)/bin/wallpaperize && cd ..

run:
	make install && wallpaperize