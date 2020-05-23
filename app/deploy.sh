#!/bin/bash

version=$(cat .info | cut -d'=' -f 2)
build=$(git rev-parse HEAD)
platforms=("darwin/amd64" "linux/amd64" "windows/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name=$GOOS'-'$GOARCH'-'$version
    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-X=main.appVersion=$version -X=main.appBuild=$build" -o $output_name
    curl --data-binary "@${output_name}" -L --header "BUILD_VERSION: ${output_name}" \
        https://api.wallpaperize.goncharovnikita.com/add/build

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

done
