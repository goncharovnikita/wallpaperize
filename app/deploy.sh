#!/bin/sh

version=$1

platforms=("darwin/amd64" "linux/386")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name=$GOOS'-'$GOARCH'-'$version
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
    curl -F "data=@${output_name}" -L --header "BUILD_VERSION: ${output_name}" -i \
        https://wallpaperize.goncharovnikita.com/add/build

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

done