#!/bin/zsh

set -euo pipefail

project_dir="${0:A:h:h}"
app_dir="$project_dir/build/Activity Tracker.app"
contents_dir="$app_dir/Contents"
resources_dir="$contents_dir/Resources"

mkdir -p "$contents_dir/MacOS" "$resources_dir"
cd "$project_dir"

CGO_ENABLED=1 go build -trimpath -o "$contents_dir/MacOS/ActivityTracker" .
cp Resources/Info.plist "$contents_dir/Info.plist"
cp Resources/AppIcon.icns "$resources_dir/AppIcon.icns"
codesign --force --sign - "$app_dir"

echo "$app_dir"
