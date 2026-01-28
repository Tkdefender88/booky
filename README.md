# booky

A terminal-based bookmark manager that's agnostic of which browser you use. Easily organize and access your bookmarks across different browsers without the hassle of constant import/export cycles.

<div align="center">
  <img src="examples/booky.gif" alt="booky demo">
</div>

## What is booky?

booky solves the problem of managing bookmarks when you frequently switch between browsers. Instead of manually exporting and importing bookmarks between Chrome, Firefox, Safari, or any other browser, booky provides:

- **Cross-browser bookmark storage**: Keep all your bookmarks in one place, independent of your browser choice
- **Custom tagging system**: Organize bookmarks with custom tags for easy categorization and discovery
- **Terminal interface**: A fast, keyboard-driven TUI (Terminal User Interface) for browsing and opening bookmarks
- **One-click opening**: Press Enter on any bookmark to open it in your default browser

## How to Use

### Launch the application

```bash
booky
```

This opens the interactive terminal interface.

### Navigation

- **Tab** - Switch between the tags view (left) and bookmarks view (right)
- **Up/Down arrows** - Navigate through tags or bookmarks
- **Enter** - Open the selected bookmark in your default browser
- **q** - Quit the application

### Add bookmarks

```bash
booky add https://example.com -n "My Bookmark" -t tag1,tag2 -d "Optional description"
```

Options:
- `-n, --name` - Name/title of the bookmark (default: "bookmark")
- `-t, --tags` - Comma-separated tags (optional)
- `-d, --description` - Description of the bookmark (optional)

### Filtering

By default, the "All" tag is selected, showing all bookmarks. Select any other tag from the left panel to see only bookmarks associated with that tag.

## Installation

### Using `go install`

```bash
go install github.com/Tkdefender88/booky@latest
```

Make sure your `$GOPATH/bin` directory is in your `$PATH`.

### Manual building

1. Clone the repository
2. Build the binary
3. Move it to a directory in your `$PATH`

## Development

### Prerequisites

- Go 1.25+
- [mise](https://mise.jdx.dev/) for tool management and task running

### Clone and setup

```bash
git clone https://github.com/Tkdefender88/booky.git
cd booky
```

### Install dependencies and build

```bash
mise run build
```

This runs the default mise tasks which generates code and builds the binary to `./bin/booky`.

### Available tasks

```bash
mise run build      # Generate code and build the binary
mise run generate   # Run sqlc to generate database code
mise run migrate    # Run database migrations
mise run clean      # Remove build artifacts
```

### Running during development

After building:

```bash
./bin/booky
```

Or build and run in one command:

```bash
mise run && ./bin/booky
```

## Database

booky stores bookmarks in a SQLite database located at:

```
~/.local/share/booky/bookmarks.db
```

The database is automatically created on first run. The schema is managed through migrations in `internal/repo/migrations/`.

### Manual database access

If you need to inspect or manipulate the database directly:

```bash
sqlite3 ~/.local/share/booky/bookmarks.db
```

## Architecture

- **internal/bookmarks** - Core bookmark and tag management logic
- **internal/repo** - Database access and migrations
- **internal/tui** - Terminal user interface built with [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- **cmd** - CLI command definitions

## Credits

This project was built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) by [Charm](https://charm.land/), a fantastic framework for building terminal user interfaces in Go. Thanks to the Charm team for the excellent tools and documentation.

## Contributing

This is a personal project, but feel free to open issues or submit PRs if you have ideas or bug reports.

## License

MIT - Do whatever you want with it! See [LICENSE](LICENSE) for details.
