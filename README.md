# go-file-converter

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

### Option 1: Build from source

```bash
# Clone the repository
git clone https://github.com/colinchambachan/go-file-converter.git
cd go-file-converter

# For macOS/Linux
go mod download
go build -o converter
chmod +x converter

# For Windows
go mod download
go build -o converter.exe
```

The build process will create an executable in your current directory:

- On macOS/Linux: `converter`
- On Windows: `converter.exe`

### Option 2: Download pre-built binaries

You can download pre-built binaries from the [Releases](https://github.com/colinchambachan/go-file-converter/releases) page.

## Usage

1. Create an `Input` folder in the same directory as the executable
2. Place your media files in the `Input` folder
3. Run the converter:

### macOS/Linux:

```bash
./converter -t jpg    # Convert to JPG
./converter -t mp3    # Convert to MP3
./converter -t mp4    # Convert to MP4
```

### Windows:

```bash
converter.exe -t jpg  # Convert to JPG
converter.exe -t mp3  # Convert to MP3
converter.exe -t mp4  # Convert to MP4
```

### Options

- `-t`: Desired output format (e.g., mp3, mp4, wav, jpg)

Note: Converted files will be saved to the `Output` folder automatically.
