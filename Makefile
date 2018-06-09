build:
	cd app && go build && cd ..

install:
	cd app && go install && cd ..

run:
	make install && wallpaperize