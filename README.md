# ffmpeg-cli

A simple command-line tool built in Go for converting media files using ffmpeg-go, a Go binding for FFmpeg.

## Description

This tool provides a convenient way to convert media files using FFmpeg functionality through a Go interface. Place your files in the `Input` folder, and the converted files will appear in the `Output` folder.

## Prerequisites

- Go 1.16 or higher
- FFmpeg installed on your system
  - macOS: `brew install ffmpeg`
  - Debian/Ubuntu: `sudo apt install ffmpeg`
  - Other systems: Download from [FFmpeg's official site](https://ffmpeg.org/download.html)

## Installation

```bash
git clone https://github.com/colinchambachan/ffmpeg-cli.git
cd ffmpeg-cli
go build
```

## Usage

Place your files in the `Input` folder, then run:

```bash
go run . -t <desired_output_type>
```

Example:

```bash
go run . -t mp3
```

### Options

- `-t`: Desired output format (e.g., mp3, mp4, wav)

Note: All files must be in the `Input` folder. Converted files will be saved to the `Output` folder.
