# Telegram Search CLI Tool

A simple command-line interface tool built with Go and Cobra for Telegram search operations.

## Installation

```bash
go mod tidy
go build -o telegram-search-cli
```

## Usage

### Basic Commands

```bash
# Run the CLI (shows welcome message)
./telegram-search-cli

# Show help
./telegram-search-cli --help

# Show version
./telegram-search-cli version

# Search for content
./telegram-search-cli search "your search query"
```

### Examples

```bash
# Search for programming content
./telegram-search-cli search "golang programming"

# Search for channels
./telegram-search-cli search "tech news"
```

## Project Structure

```
.
├── main.go          # Entry point
├── cmd/
│   ├── root.go      # Root command
│   ├── version.go   # Version command
│   └── search.go    # Search command
├── go.mod           # Go module file
└── README.md        # This file
```

## Development

This is a basic CLI structure using Cobra. You can extend it by:

1. Adding more commands in the `cmd/` directory
2. Implementing actual Telegram API integration
3. Adding configuration management
4. Adding more flags and options

## Dependencies

- `github.com/spf13/cobra` - CLI framework 