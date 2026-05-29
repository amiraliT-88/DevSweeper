# DevSweeper

A fast TUI tool to clean up system junk, app caches, and temp folders directly from your terminal.

Built with **Go** and **Bubble Tea**.

---

## Features
- **Fast Scanning:** Finds junk folders across your system quickly.
- **Dynamic Sorting:** Automatically sorts folders by size once calculated.
- **Terminal UI:** Uses pagination, size-based color coding, and a deletion progress bar.
- **Safe:** Skips `.git` repositories and respects locked files.
- **Batch Deletion:** Select multiple targets (or press `A` to select all) and hit `Enter` to delete.

## What it targets:
- Windows `Temp` & `tmp` directories
- App caches (`Cache`, `.cache`, `CachedData`, `GPUCache`)
- Python & Visual Studio caches (`__pycache__`, `.pytest_cache`, `.vs`)
- Crash reports (`Crashpad`, `CrashReports`, `logs`)

---

## Installation & Usage

### The Easy Way
You can download the pre-compiled `DevSweeper.exe` directly from the Releases page, double-click it, and start cleaning.

### Build from source
Make sure you have [Go](https://golang.org/) installed.
```bash
git clone https://github.com/yourusername/DevSweeper.git
cd DevSweeper
go build -o DevSweeper.exe .
./DevSweeper.exe
```

## Controls
- `Up/Down` or `k/j`: Navigate the list
- `Space`: Select/Deselect a folder
- `A`: Select/Deselect All
- `Enter`: Delete selected folders
- `Q`: Quit

## Contributing
Pull requests and issues are welcome.
