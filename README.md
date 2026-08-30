<div align="center">
  <h1>🧹 DevSweeper</h1>
  
  **A lightning-fast, interactive TUI tool to scan and purge system junk, app caches, and developer artifacts directly from your terminal.**

  <p>
    <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
    <img src="https://img.shields.io/badge/TUI-Bubble_Tea-FF5F87?style=for-the-badge" />
    <img src="https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-24292e?style=for-the-badge" />
  </p>
</div>

---

## ✨ Features

- ⚡ **High-Speed Scanning:** Concurrently crawls system and user directories to discover bloated caches in milliseconds.
- 📊 **Dynamic Size Sorting:** Automatically ranks targets by calculated disk usage as scans resolve.
- 🎨 **Terminal UI:** Smooth keyboard navigation, size-based color-coding, pagination, and real-time deletion progress bar.
- 🛡️ **Built-in Safety:** Automatically skips `.git` repositories and gracefully handles locked or protected system files.
- 🗂️ **Batch Deletion:** Select multiple individual directories or select all at once for single-keystroke cleaning.

---

## 🎯 What DevSweeper Cleans

| Category | Targeted Directories & Patterns |
| :--- | :--- |
| **System Temporary Files** | Windows `Temp`, `tmp`, user local temporary caches |
| **Application Caches** | `Cache`, `.cache`, `CachedData`, `GPUCache`, browser & app cache stores |
| **Development Artifacts** | `__pycache__`, `.pytest_cache`, `.vs`, language runtime build caches |
| **Crash Reports & Logs** | `Crashpad`, `CrashReports`, application `.log` dumps |

---

## ⌨️ Controls & Keybindings

| Keybinding | Action |
| :--- | :--- |
| `↑ / ↓` or `k / j` | Navigate up/down the directory list |
| `Space` | Toggle selection of highlighted folder |
| `A` | Select / Deselect **All** discovered folders |
| `Enter` | Execute safe batch deletion |
| `Q` / `Ctrl+C` | Exit application |

---

## 📦 Installation & Usage

### 1. Download Pre-compiled Binary
Grab the standalone `DevSweeper.exe` directly from the [Releases](https://github.com/amiraliT-88/DevSweeper/releases) page. No installation or runtime required!

### 2. Build from Source
Make sure you have [Go](https://go.dev/) (1.20+) installed:

```bash
# Clone the repository
git clone https://github.com/amiraliT-88/DevSweeper.git
cd DevSweeper

# Build executable
go build -o DevSweeper.exe .

# Run
./DevSweeper.exe
```

---

## 🤝 Contributing
Issues, feature requests, and Pull Requests are welcome! Feel free to open an issue to suggest new junk folder patterns.
