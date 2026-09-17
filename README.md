<p align="center">
  <img src="Resources/AppIcon.png" width="112" alt="Activity Tracker app icon" />
</p>

<h1 align="center">Activity Tracker</h1>

<p align="center">
  A small macOS menu bar app that keeps today's active time visible at a glance.
</p>

<p align="center">
  <a href="https://support.apple.com/macos"><img src="https://img.shields.io/badge/macOS-13%2B-black?logo=apple" alt="macOS 13+" /></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.22.2%2B-00ADD8?logo=go&amp;logoColor=white" alt="Go 1.22.2+" /></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT" /></a>
</p>

Activity Tracker counts time while you use your Mac and pauses after one minute of inactivity. The total appears in the menu bar and resets when a new day starts. It stores the count locally; no account is needed.

<p align="center">
  <img src="images/screenshot.png" width="256" alt="Activity Tracker showing today's active time in the macOS menu bar" />
</p>

## Features

- Today's active time in the menu bar, shown in minutes or hours and minutes
- Automatic pause after one minute of inactivity
- Local persistence across app restarts
- No Dock icon

Inspired by [timeow](https://github.com/f-person/timeow-mac).

## Requirements

- macOS 13 Ventura or newer
- Go 1.22.2 or newer to build from source

## Install

Download the DMG from the [latest GitHub release](https://github.com/mguellsegarra/activity-tracker-mac/releases/latest), open it, and drag **Activity Tracker.app** to **Applications**. Launch it from Applications; the counter appears in the menu bar.

The downloadable app is ad-hoc signed but not Apple-notarized. macOS may ask you to confirm the first launch of a downloaded copy.

To start Activity Tracker when you log in, add the installed app in **System Settings → General → Login Items & Extensions**. If you previously installed the Launch Agent from older instructions, unload or disable it first to avoid running two copies.

## Build from source

```sh
git clone https://github.com/mguellsegarra/activity-tracker-mac.git
cd activity-tracker-mac
./scripts/build-app.sh
ditto "build/Activity Tracker.app" "/Applications/Activity Tracker.app"
open "/Applications/Activity Tracker.app"
```

Quit any running copy before replacing it. On macOS 27, running the app from `/Applications` helps menu bar managers such as Bartender recognize it; `go run .` is still available for development.

The source icon is `Resources/AppIcon.png`, and the bundled icon is `Resources/AppIcon.icns`. To regenerate the bundled icon, install ImageMagick and run `./scripts/build-icon.sh`. To build the drag-to-Applications DMG, run `./scripts/build-dmg.sh`; it writes `build/Activity-Tracker.dmg`.

## Data and privacy

Activity Tracker stores its count in `~/.config/activity_tracker/active_time`. This file remains in place when you replace the app with a newer version. The app does not require an account or cloud service.

## Contributing

Contributions and bug reports are welcome through GitHub issues and pull requests.

## License

[MIT](LICENSE) © Marc Güell Segarra. More about the author at [Ondori.dev](https://ondori.dev). If you find the app useful, you can [buy me a coffee](https://buymeacoffee.com/mguellsegarra).
