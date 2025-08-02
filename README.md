# Telegram Search CLI Tool 🔍

A powerful command-line interface tool built with Go and Cobra for searching Telegram content on Lyzem.com. Search for channels, groups, bots, and messages with ease, and export results to multiple formats.

[![Go Version](https://img.shields.io/badge/Go-1.23+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![CLI](https://img.shields.io/badge/CLI-Cobra-orange.svg)](https://github.com/spf13/cobra)

## ✨ Features

- 🔍 **Multi-type Search**: Search for channels, groups, bots, and messages
- 📄 **Smart Pagination**: Automatic page calculation based on your limit
- 💾 **Export Support**: Save results to JSON, CSV, or TXT formats
- ⚡ **Fast & Efficient**: Built with Go for high performance
- 🎯 **Flexible Limits**: Set custom result limits with auto-pagination
- 🛠️ **Easy to Use**: Simple and intuitive command-line interface
- 🌐 **Web Scraping**: Scrapes data from Lyzem.com search engine
- 📁 **Auto-filenames**: Automatic timestamped filename generation

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

### Export Options

| Flag | Short | Description |
|------|-------|-------------|
| `--export` | `-e` | Export format: json, csv, txt |
| `--output` | `-o` | Output filename (auto-generated if not specified) |

## 🎯 Examples

### Channel Search

```bash
# Search for channels with default settings (10 results)
telegram-search-cli search --channel "pakistan"

# Search for channels with custom limit (auto-calculates pages)
telegram-search-cli search --channel "pakistan" --limit 50

# Search for channels with manual pagination
telegram-search-cli search --channel "pakistan" --pages 3 --per-page 15 --limit 45

# Export channels to JSON
telegram-search-cli search --channel "pakistan" --limit 50 --export json

# Export channels to CSV with custom filename
telegram-search-cli search --channel "pakistan" --limit 50 --export csv --output pakistan_channels.csv
```

### Group Search

```bash
# Search for groups
telegram-search-cli search --group "programming"

# Search for groups with high limit
telegram-search-cli search --group "programming" --limit 100

# Export groups to JSON
telegram-search-cli search --group "programming" --limit 100 --export json

# Export groups to TXT
telegram-search-cli search --group "programming" --limit 100 --export txt
```

### Bot Search

```bash
# Search for bots
telegram-search-cli search --bot "sports"

# Search for bots with custom limit
telegram-search-cli search --bot "sports" --limit 75

# Export bots to CSV
telegram-search-cli search --bot "sports" --limit 75 --export csv

# Export bots to JSON with custom filename
telegram-search-cli search --bot "sports" --limit 75 --export json --output sports_bots.json
```

### Message Search

```bash
# Search for messages (typically returns more results)
telegram-search-cli search --message "pakistan"

# Search for messages with high limit
telegram-search-cli search --message "pakistan" --limit 200

# Export messages to JSON
telegram-search-cli search --message "pakistan" --limit 200 --export json

# Export messages to CSV with custom filename
telegram-search-cli search --message "pakistan" --limit 200 --export csv --output pakistan_messages.csv
```

### Advanced Examples

```bash
# Search for channels with specific pagination and export
telegram-search-cli search --channel "tech news" --pages 5 --per-page 20 --limit 100 --export json

# Search for groups with high results per page and export to CSV
telegram-search-cli search --group "education" --per-page 25 --limit 150 --export csv

# Quick search for bots with export (auto-calculates 10 pages for 100 results)
telegram-search-cli search --bot "music" --limit 100 --export txt

# Export all search types to different formats
telegram-search-cli search --channel "pakistan" --limit 50 --export json --output channels.json
telegram-search-cli search --group "pakistan" --limit 50 --export csv --output groups.csv
telegram-search-cli search --bot "pakistan" --limit 50 --export txt --output bots.txt
telegram-search-cli search --message "pakistan" --limit 100 --export json --output messages.json
```

## 📁 Export Formats

### JSON Format
Perfect for data processing and APIs:
```json
{
  "query": "pakistan",
  "type": "channels",
  "timestamp": "2024-01-15T10:30:00Z",
  "count": 50,
  "results": [
    {
      "Title": "Channel Name",
      "URL": "https://t.me/channel",
      "Description": "Channel description"
    }
  ]
}
```

### CSV Format
Ideal for spreadsheet applications:
```csv
Title,URL,Description
Channel Name,https://t.me/channel,Channel description
Another Channel,https://t.me/another,Another description
```

### TXT Format
Human-readable plain text:
```
Search Results
==============
Query: pakistan
Type: channels
Timestamp: 2024-01-15T10:30:00Z

1. Channel Name
   URL: https://t.me/channel
   Description: Channel description
----------------------------------------
```

### Auto-Generated Filenames

When you don't specify `--output`, the CLI automatically generates timestamped filenames:
- `channels_pakistan_2024-01-15_10-30-05.json`
- `groups_programming_2024-01-15_10-30-05.csv`
- `bots_sports_2024-01-15_10-30-05.txt`
- `messages_pakistan_2024-01-15_10-30-05.json`

## 🏗️ Project Structure

```
Telegram-Search-CLI-tool-golang-osint/
├── main.go                    # Application entry point
├── cmd/
│   ├── root.go               # Root command configuration
│   ├── version.go            # Version command
│   └── search.go             # Search command with all search types and export
├── scraper/
│   ├── channel_search.go     # Channel search functionality
│   ├── group_search.go       # Group search functionality
│   ├── bot_search.go         # Bot search functionality
│   ├── message_search.go     # Message search functionality
│   └── convert.go            # Type conversion helpers for export
├── utils/
│   └── export.go             # Export functionality (JSON, CSV, TXT)
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
6. **Export Data**: Saves results to file if export is requested
7. **Format Output**: Displays results in a clean, readable format

### Export Process

1. **Format Selection**: Choose JSON, CSV, or TXT format
2. **Data Conversion**: Convert search results to export format
3. **Filename Generation**: Auto-generate or use custom filename
4. **File Creation**: Write data to file with proper formatting
5. **Metadata Addition**: Include query, type, timestamp, and count

## 🛠️ Development

### Adding New Search Types

1. Create a new file in `scraper/` directory
2. Implement the search function following the existing pattern
3. Add the flag in `cmd/search.go`
4. Update the search logic to handle the new type
5. Add conversion function in `scraper/convert.go`

### Adding New Export Formats

1. Add new format case in `utils/export.go`
2. Implement the export function for the new format
3. Update the format validation in `ExportToFile`

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
- **encoding/csv** - CSV export functionality
- **encoding/json** - JSON export functionality

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

**"Export error"**
- Check if you have write permissions in the current directory
- Ensure the export format is supported (json, csv, txt)
- Verify the filename doesn't contain invalid characters

**"Unsupported format"**
- Supported formats: json, csv, txt
- Use lowercase format names
- Check the format spelling

## 📞 Support

If you encounter any issues or have questions:

1. Check the troubleshooting section above
2. Review the help documentation: `telegram-search-cli --help`
3. Open an issue on GitHub
4. Check the project documentation

---

**Happy Searching and Exporting! 🔍💾✨** 