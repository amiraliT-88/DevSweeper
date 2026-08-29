# DevSweeper

A fast, lightweight TUI tool to clean up system junk, app caches, and temp folders directly from your terminal.

Built with **Go** and **Bubble Tea**.

---

## Features
- **Fast Scanning:** Quickly discovers junk folders and caches across your system.
- **Dynamic Sorting:** Automatically ranks folders by calculated disk usage.
- **Terminal UI:** Smooth keyboard navigation, size-based color coding, and deletion progress bar.
- **Safe:** Automatically protects `.git` repositories and safely handles locked files.
- **Batch Operations:** Select multiple directories (or press `A` to select all) and clean in one keystroke.

## Targets Cleaned:
- Windows `Temp` & `tmp` directories
- Application caches (`Cache`, `.cache`, `CachedData`, `GPUCache`)
- Development artifacts (`__pycache__`, `.pytest_cache`, `.vs`)
- Crash logs & reports (`Crashpad`, `CrashReports`, `logs`)

---

## Installation & Usage

### Pre-compiled Binary
Download `DevSweeper.exe` directly from the [Releases](https://github.com/amiraliT-88/DevSweeper/releases) page.

### Build from Source
Make sure you have [Go](https://golang.org/) installed:
```bash
git clone https://github.com/amiraliT-88/DevSweeper.git
cd DevSweeper
go build -o DevSweeper.exe .
./DevSweeper.exe
```

## Controls
- `↑ / ↓` or `k / j`: Navigate list
- `Space`: Toggle folder selection
- `A`: Select / Deselect All
- `Enter`: Delete selected items
- `Q`: Quit

## Contributing
Issues and Pull Requests are welcome!
