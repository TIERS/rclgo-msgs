#!/bin/bash

if [ $# != 2 ]; then
    echo "usage: $0 <ros-package-path> <dest-dir>"
    exit 1
fi

ros_package_path="$(realpath "$1")"
dest_dir="$(realpath "$2")"

ros_package="$(basename "$ros_package_path")"
tools_dir="$(dirname "$(realpath "$BASH_SOURCE")")"

if [ -z "$GOPATH" ]; then
    export GOPATH="$HOME/go"
fi
export PATH="$PATH:$GOPATH/bin"

cd "$ros_package_path"

author_name="$(git log --format='%an' HEAD^!)" || exit 1
author_email="$(git log --format='%ae' HEAD^!)" || exit 1

cd "$tools_dir"

rclgo_version="$(go run getrclgoversion.go -mod-file "$dest_dir/go.mod")" || exit 1

cd "$dest_dir"

go install "github.com/TIERS/rclgo/cmd/rclgo-gen@$rclgo_version" || exit 1

retry_delay=3s
for i in {1..5}; do
    if [ $i -gt 1 ]; then
        echo "Trying again after $retry_delay"
        sleep "$retry_delay"
    fi
    git reset --hard "origin/$(git branch --show-current)"
    git clean -dxf
    git pull -X theirs || continue
    rm -rf "$ros_package"
    rclgo-gen generate --root-path "$ros_package_path" || continue
    git add -A || continue
    commit_output="$(git -c "user.name=$author_name" -c "user.email=$author_email" commit -m "Update bindings for $ros_package")"
    ret_code=$?
    echo "$commit_output"
    if [[ $ret_code != 0 ]]; then
        if grep -iq "nothing to commit, working tree clean" <<< "$commit_output"; then
            exit 0
        else
            continue
        fi
    fi
    git push || continue
    exit 0
done
exit 1
