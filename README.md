# Telegram Search CLI Tool 🔍

A powerful command-line interface tool built with Go and Cobra for searching Telegram content on Lyzem.com. Search for channels, groups, bots, and messages with ease.

[![Go Version](https://img.shields.io/badge/Go-1.23+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![CLI](https://img.shields.io/badge/CLI-Cobra-orange.svg)](https://github.com/spf13/cobra)

## ✨ Features

- 🔍 **Multi-type Search**: Search for channels, groups, bots, and messages
- 📄 **Smart Pagination**: Automatic page calculation based on your limit
- ⚡ **Fast & Efficient**: Built with Go for high performance
- 🎯 **Flexible Limits**: Set custom result limits with auto-pagination
- 🛠️ **Easy to Use**: Simple and intuitive command-line interface
- 🌐 **Web Scraping**: Scrapes data from Lyzem.com search engine

## 🚀 Installation

### Prerequisites
- Go 1.23 or higher
- Git

### Quick Install

```bash
# Clone the repository
git clone <your-repo-url>
cd Telegram-Search-CLI-tool-golang-osint

# Install dependencies
go mod tidy

# Build the CLI tool
go build -o telegram-search-cli

# Install globally (optional)
cp telegram-search-cli ~/.local/bin/
# or
sudo cp telegram-search-cli /usr/local/bin/
```

### Using Go Install

```bash
# Install directly from your project directory
go install .

# Make sure ~/go/bin is in your PATH
export PATH="$PATH:$(go env GOPATH)/bin"
```

## 📖 Usage

### Basic Commands

```bash
# Show help
telegram-search-cli --help

# Show version
telegram-search-cli version

# Show search command help
telegram-search-cli search --help
```

### Search Types

The CLI supports four types of Telegram content search:

| Type | Flag | Short | Description |
|------|------|-------|-------------|
| **Channels** | `--channel` | `-c` | Search for Telegram channels |
| **Groups** | `--group` | `-g` | Search for Telegram groups |
| **Bots** | `--bot` | `-b` | Search for Telegram bots |
| **Messages** | `--message` | `-m` | Search for Telegram messages |

### Search Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--limit` | `-l` | 0 (no limit) | Maximum number of results to return |
| `--pages` | `-p` | Auto-calculated | Number of pages to fetch |
| `--per-page` | `-n` | 10 | Results per page |

## 🎯 Examples

### Channel Search

```bash
# Search for channels with default settings (10 results)
telegram-search-cli search --channel "pakistan"

# Search for channels with custom limit (auto-calculates pages)
telegram-search-cli search --channel "pakistan" --limit 50

# Search for channels with manual pagination
telegram-search-cli search --channel "pakistan" --pages 3 --per-page 15 --limit 45
```

### Group Search

```bash
# Search for groups
telegram-search-cli search --group "programming"

# Search for groups with high limit
telegram-search-cli search --group "programming" --limit 100
```

### Bot Search

```bash
# Search for bots
telegram-search-cli search --bot "sports"

# Search for bots with custom limit
telegram-search-cli search --bot "sports" --limit 75
```

### Message Search

```bash
# Search for messages (typically returns more results)
telegram-search-cli search --message "pakistan"

# Search for messages with high limit
telegram-search-cli search --message "pakistan" --limit 200
```

### Advanced Examples

```bash
# Search for channels with specific pagination
telegram-search-cli search --channel "tech news" --pages 5 --per-page 20 --limit 100

# Search for groups with high results per page
telegram-search-cli search --group "education" --per-page 25 --limit 150

# Quick search for bots (auto-calculates 10 pages for 100 results)
telegram-search-cli search --bot "music" --limit 100
```

## 🏗️ Project Structure

```
Telegram-Search-CLI-tool-golang-osint/
├── main.go                    # Application entry point
├── cmd/
│   ├── root.go               # Root command configuration
│   ├── version.go            # Version command
│   └── search.go             # Search command with all search types
├── scraper/
│   ├── channel_search.go     # Channel search functionality
│   ├── group_search.go       # Group search functionality
│   ├── bot_search.go         # Bot search functionality
│   └── message_search.go     # Message search functionality
├── go.mod                    # Go module definition
├── go.sum                    # Dependency checksums
├── README.md                 # This documentation
└── telegram-search-cli       # Compiled binary
```

## 🔧 How It Works

### Auto-Pagination Logic

When you specify `--limit`, the CLI automatically calculates the required pages:

- **Limit 100** → Fetches 10 pages (100 ÷ 10 = 10 pages)
- **Limit 50** → Fetches 5 pages (50 ÷ 10 = 5 pages)
- **Limit 75** → Fetches 8 pages (75 ÷ 10 = 7.5, rounded up)

### Search Process

1. **Parse Query**: Validates and processes your search keyword
2. **Calculate Pages**: Determines how many pages to fetch
3. **Scrape Data**: Fetches results from Lyzem.com
4. **Parse Results**: Extracts title, URL, and description
5. **Apply Limits**: Returns results up to your specified limit
6. **Format Output**: Displays results in a clean, readable format

## 🛠️ Development

### Adding New Search Types

1. Create a new file in `scraper/` directory
2. Implement the search function following the existing pattern
3. Add the flag in `cmd/search.go`
4. Update the search logic to handle the new type

### Building from Source

```bash
# Clone and navigate to project
git clone <your-repo-url>
cd Telegram-Search-CLI-tool-golang-osint

# Install dependencies
go mod tidy

# Build
go build -o telegram-search-cli

# Run tests (if available)
go test ./...
```

## 📋 Dependencies

- **github.com/spf13/cobra** - CLI framework
- **github.com/PuerkitoBio/goquery** - HTML parsing
- **golang.org/x/net** - HTTP client utilities

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⚠️ Disclaimer

This tool is for educational and research purposes. Please respect the terms of service of Lyzem.com and use responsibly. The tool includes rate limiting and respectful scraping practices.

## 🆘 Troubleshooting

### Common Issues

**"Command not found"**
```bash
# Make sure the binary is in your PATH
export PATH="$PATH:$(go env GOPATH)/bin"
# or
export PATH="$PATH:$HOME/.local/bin"
```

**"No results found"**
- Check your internet connection
- Verify the search keyword
- Try different search terms
- Check if Lyzem.com is accessible

**"Too many requests"**
- The tool includes built-in delays
- Wait a few minutes before trying again
- Consider reducing the number of pages

## 📞 Support

If you encounter any issues or have questions:

1. Check the troubleshooting section above
2. Review the help documentation: `telegram-search-cli --help`
3. Open an issue on GitHub
4. Check the project documentation

---

**Happy Searching! 🔍✨** 