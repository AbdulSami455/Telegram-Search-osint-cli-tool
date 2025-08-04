# Telegram Search CLI

A powerful command-line interface tool for searching and discovering Telegram channels, groups, bots, and messages. Built with modern CLI design principles for the best user experience.

## ✨ Features

- **🔍 Multi-type Search**: Search for channels, groups, bots, and messages
- **📊 Progress Indicators**: Visual progress bars and spinners for long-running operations
- **📁 Export Support**: Export results in JSON, CSV, and TXT formats
- **🎯 Human-readable Errors**: Clear error messages with actionable suggestions
- **⚡ Real-time Feedback**: Immediate visual feedback for all user actions
- **📈 Pagination Control**: Configurable pages and results per page
- **🎨 Modern UI**: Beautiful, intuitive interface with emojis and clear formatting

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd Telegram-Search-CLI-tool-golang-osint

# Build the application
go build -o telegram-search-cli

# Make it executable (Linux/macOS)
chmod +x telegram-search-cli
```

### Basic Usage

```bash
# Search for channels
./telegram-search-cli search --channel "technology" --limit 10

# Search for groups with export
./telegram-search-cli search --group "programming" --export json --output results.json

# Search for bots with multiple pages
./telegram-search-cli search --bot "news" --pages 3 --per-page 20

# Search for messages
./telegram-search-cli search --message "golang" --limit 50
```

## 📖 Command Reference

### Main Commands

```bash
telegram-search-cli [command]

Available Commands:
  search      Search for content on Telegram
  version     Print the version number
  help        Help about any command
```

### Search Command

```bash
telegram-search-cli search [query] [flags]

Flags:
  -c, --channel     Search for Telegram channels
  -g, --group       Search for Telegram groups
  -b, --bot         Search for Telegram bots
  -m, --message     Search for Telegram messages
  -l, --limit       Maximum number of results (default: no limit)
  -p, --pages       Number of pages to fetch (auto-calculated if limit specified)
  -n, --per-page    Results per page (default: 10)
  -e, --export      Export format: json, csv, txt
  -o, --output      Output filename (auto-generated if not specified)
```

## 🎯 CLI Design Principles Implemented

### 1. Visual Progress Indicators

The CLI now provides real-time visual feedback for all operations:

- **Progress Bars**: For multi-page searches showing completion percentage and ETA
- **Spinners**: For export operations and data processing
- **Status Messages**: Clear indication of what's happening at each step

**Example Output:**
```
ℹ️  Searching for channels with keyword: 'technology'
ℹ️  Will fetch 3 pages with 10 results per page
Fetching channels [████████████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░] 2/3 (66.7%) ETA: 2s
ℹ️  Page 1: Found 8 channels
ℹ️  Page 2: Found 10 channels
✅ Completed! Found 18 channels total
```

### 2. Human-readable Error Messages

All errors now provide clear, actionable information:

- **Contextual Messages**: Errors explain what went wrong in user terms
- **Actionable Suggestions**: Specific steps to resolve the issue
- **Error Categorization**: Different types of errors (network, validation, etc.)

**Example Error Output:**
```
❌ Failed to fetch page 2
   Details: connection timeout
   Suggestions:
   1. Check your internet connection
   2. Verify the URL is accessible
   3. Try again in a few moments
```

### 3. Reactions for Every Action

Every user action now provides immediate feedback:

- **Success Confirmations**: Clear success messages with result counts
- **Progress Updates**: Real-time status during operations
- **Completion Notifications**: Summary of what was accomplished

**Example Success Output:**
```
✅ Channel search completed successfully! Found 18 results
✅ Results exported to 'channels_technology_2024-01-15_14-30-25.json' successfully
```

## 📊 Export Formats

### JSON Export
```json
{
  "query": "technology",
  "type": "channels",
  "timestamp": "2024-01-15T14:30:25Z",
  "count": 18,
  "results": [
    {
      "Title": "Tech News Channel",
      "URL": "https://t.me/technews",
      "Description": "Latest technology updates"
    }
  ]
}
```

### CSV Export
```csv
Title,URL,Description
Tech News Channel,https://t.me/technews,Latest technology updates
Programming Hub,https://t.me/programming,Programming tutorials and tips
```

### TXT Export
```
Search Results
==============
Query: technology
Type: channels
Timestamp: 2024-01-15T14:30:25Z

1. Tech News Channel
   URL: https://t.me/technews
   Description: Latest technology updates
----------------------------------------
```

## 🛠️ Advanced Usage

### Multi-page Search with Progress
```bash
# Search with visual progress for large result sets
./telegram-search-cli search --channel "news" --pages 5 --per-page 20
```

### Export with Custom Filename
```bash
# Export to specific file with custom name
./telegram-search-cli search --group "programming" --export json --output my_results.json
```

### Limited Results with Auto-pagination
```bash
# Automatically calculate pages needed for 50 results
./telegram-search-cli search --bot "helper" --limit 50
```

## 🎨 User Interface Features

### Visual Elements
- **Emojis**: Intuitive icons for different types of information
- **Headers**: Clear section separators with decorative borders
- **Progress Bars**: Visual representation of completion status
- **Color-coded Messages**: Different styles for info, success, warning, and error

### Interactive Feedback
- **Real-time Updates**: Live progress during operations
- **Status Indicators**: Clear indication of current operation
- **Completion Summaries**: Detailed results with statistics

## 🔧 Error Handling

The CLI implements comprehensive error handling with:

### Error Types
- **Network Errors**: Connection issues, timeouts, DNS failures
- **HTTP Errors**: Status codes with specific handling (404, 429, 500, etc.)
- **Validation Errors**: Invalid input parameters
- **Permission Errors**: Access denied scenarios
- **Rate Limit Errors**: Too many requests handling

### Error Recovery
- **Automatic Retries**: For transient network issues
- **Graceful Degradation**: Continue operation when possible
- **User Guidance**: Clear next steps for resolution

## 📈 Performance Features

- **Smart Pagination**: Automatic calculation of required pages
- **Progress Tracking**: Real-time feedback for long operations
- **Efficient Parsing**: Optimized HTML parsing for large result sets
- **Memory Management**: Efficient handling of large datasets

## 🤝 Contributing

This project follows modern CLI design principles. When contributing:

1. **User Experience First**: Always consider the user's perspective
2. **Clear Feedback**: Provide immediate and clear feedback for all actions
3. **Error Handling**: Implement comprehensive error handling with helpful messages
4. **Progress Indicators**: Show progress for operations that take time
5. **Consistent Interface**: Maintain consistent UI patterns throughout

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For issues, questions, or contributions:

1. Check the error messages for specific guidance
2. Review the help documentation: `./telegram-search-cli --help`
3. Try the examples in this README
4. Open an issue with detailed error information

---

**Built with ❤️ and modern CLI design principles for the best user experience.** 