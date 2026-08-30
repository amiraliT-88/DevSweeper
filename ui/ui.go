package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"devsweeper/scanner"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle    = lipgloss.NewStyle().Margin(1, 2)
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFFDF5")).Background(lipgloss.Color("#FF5F87")).Padding(0, 2).MarginBottom(1)
	footerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#A8CC8C")).MarginTop(1)
	listStyle   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#874BFD")).Padding(0, 1)

	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF79C6")).Bold(true)
	markedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5555")).Strikethrough(true)
)

type model struct {
	root         string
	items        []scanner.FoundDir
	cursor       int
	marked       map[string]bool
	scanning         bool
	spinner          spinner.Model
	finished         bool
	deletedCount     int
	totalSize        int64
	selectedSize     int64
	startIndex       int
	pageSize         int
	sizesToCalculate int
	
	deleting         bool
	itemsToDelete    []string
	deleteTotal      int
	deleteProgress   int
}

type scanResultMsg []scanner.FoundDir

type deletedItemMsg struct {
	path string
}

func deleteItemCmd(path string) tea.Cmd {
	return func() tea.Msg {
		scanner.DeleteDir(path)
		return deletedItemMsg{path}
	}
}

type sizeMsg struct {
	path string
	size int64
}

func calculateSize(path string) tea.Cmd {
	return func() tea.Msg {
		var size int64
		filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() {
				size += info.Size()
			}
			return nil
		})
		return sizeMsg{path, size}
	}
}

func InitialModel(root string) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		root:         root,
		items:        []scanner.FoundDir{},
		marked:       make(map[string]bool),
		scanning:     true,
		spinner:      s,
		pageSize:     10,
	}
}

func runScanner(root string) tea.Cmd {
	return func() tea.Msg {
		res := scanner.Scan(root)
		return scanResultMsg(res)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		runScanner(m.root),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				if m.cursor < m.startIndex {
					m.startIndex = m.cursor
				}
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
				if m.cursor >= m.startIndex+m.pageSize {
					m.startIndex = m.cursor - m.pageSize + 1
				}
			}
		case "a", "A":
			if !m.scanning && len(m.items) > 0 {
				allSelected := true
				for _, item := range m.items {
					if !m.marked[item.Path] {
						allSelected = false
						break
					}
				}
				for _, item := range m.items {
					m.marked[item.Path] = !allSelected
				}
				m.selectedSize = 0
				for _, item := range m.items {
					if m.marked[item.Path] {
						m.selectedSize += item.Size
					}
				}
			}
		case " ":
			if !m.scanning && len(m.items) > 0 {
				itemPath := m.items[m.cursor].Path
				m.marked[itemPath] = !m.marked[itemPath]
				if m.marked[itemPath] {
					m.selectedSize += m.items[m.cursor].Size
				} else {
					m.selectedSize -= m.items[m.cursor].Size
				}
			}
		case "enter":
			if !m.scanning && !m.deleting {
				if m.finished {
					return m, tea.Quit
				}
				var toDelete []string
				for _, item := range m.items {
					if m.marked[item.Path] {
						toDelete = append(toDelete, item.Path)
					}
				}
				if len(toDelete) == 0 {
					return m, nil
				}
				
				m.deleting = true
				m.itemsToDelete = toDelete
				m.deleteTotal = len(toDelete)
				m.deleteProgress = 0
				
				return m, deleteItemCmd(m.itemsToDelete[0])
			}
		}
		if m.finished {
			return m, tea.Quit
		}
		return m, nil
	}

	if dMsg, ok := msg.(deletedItemMsg); ok {
		_ = dMsg
		m.deleteProgress++
		if m.deleteProgress >= m.deleteTotal {
			m.deleting = false
			m.finished = true
			m.deletedCount = m.deleteTotal
			return m, nil
		}
		return m, deleteItemCmd(m.itemsToDelete[m.deleteProgress])
	}

	if resMsg, ok := msg.(scanResultMsg); ok {
		m.items = resMsg
		m.scanning = false
		m.sizesToCalculate = len(m.items)
		var cmds []tea.Cmd
		for _, item := range m.items {
			cmds = append(cmds, calculateSize(item.Path))
		}
		
		if len(m.items) == 0 {
			m.sizesToCalculate = 0
		}
		
		return m, tea.Batch(cmds...)
	}

	if sMsg, ok := msg.(sizeMsg); ok {
		for i, item := range m.items {
			if item.Path == sMsg.path {
				m.items[i].Size = sMsg.size
				m.totalSize += sMsg.size
				if m.marked[item.Path] {
					m.selectedSize += sMsg.size
				}
				break
			}
		}
		
		m.sizesToCalculate--
		if m.sizesToCalculate <= 0 {
			sort.SliceStable(m.items, func(i, j int) bool {
				return m.items[i].Size > m.items[j].Size
			})
		}
		return m, nil
	}

	if wMsg, ok := msg.(tea.WindowSizeMsg); ok {
		m.pageSize = wMsg.Height - 8
		if m.pageSize < 5 {
			m.pageSize = 5
		}
		return m, nil
	}

	if tickMsg, ok := msg.(spinner.TickMsg); ok {
		if m.scanning {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(tickMsg)
			return m, cmd
		}
	}

	return m, nil
}

func formatSize(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func colorForSize(size int64) lipgloss.Style {
	mb := size / (1024 * 1024)
	if mb > 1024 { // > 1 GB
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5555")).Bold(true) // Red
	} else if mb > 100 {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB86C")) // Orange
	} else if mb > 0 {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#50FA7B")) // Green
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")) // Gray
}

func (m model) View() string {
	if m.finished {
		successStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#50FA7B")).
			Border(lipgloss.DoubleBorder(), true).
			BorderForeground(lipgloss.Color("#50FA7B")).
			Padding(2, 4).
			Margin(2, 4)

		msg := fmt.Sprintf("✨ MISSION ACCOMPLISHED ✨\n\nSuccessfully incinerated %d junk folders!\nFreed up %s of space.\n\nPress any key to exit.", m.deletedCount, formatSize(m.selectedSize))
		return successStyle.Render(msg)
	}

	if m.deleting {
		percentage := 0.0
		if m.deleteTotal > 0 {
			percentage = float64(m.deleteProgress) / float64(m.deleteTotal)
		}
		barWidth := 40
		filled := int(percentage * float64(barWidth))
		empty := barWidth - filled
		if filled < 0 { filled = 0 }
		if empty < 0 { empty = 0 }
		
		bar := strings.Repeat("█", filled) + strings.Repeat("░", empty)
		
		s := fmt.Sprintf("\n\n 🗑️  DELETING JUNK... Please wait.\n\n")
		s += fmt.Sprintf(" [%s] %d%%\n\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5555")).Render(bar), int(percentage*100))
		s += fmt.Sprintf(" %d / %d items deleted.\n", m.deleteProgress, m.deleteTotal)
		
		boxStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF5555")).
			Padding(1, 4).
			Margin(2, 4)
			
		return boxStyle.Render(s)
	}

	s := headerStyle.Render(" 🚀 DevSweeper ") + "\n"

	if !m.scanning && len(m.items) > 0 {
		s += fmt.Sprintf(" 💾 Total Size: %s\n 🗑️  Selected to Delete: %s\n\n",
			lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00")).Render(formatSize(m.totalSize)),
			lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5555")).Render(formatSize(m.selectedSize)))
	}

	if m.scanning {
		s += fmt.Sprintf("\n %s Scanning %s ...\n", m.spinner.View(), m.root)
		return docStyle.Render(s)
	}

	if len(m.items) == 0 {
		s += "\nNo targets found. Your system is super clean! ✨\n"
		return docStyle.Render(s)
	}

	var listContent string
	endIndex := m.startIndex + m.pageSize
	if endIndex > len(m.items) {
		endIndex = len(m.items)
	}

	for i := m.startIndex; i < endIndex; i++ {
		item := m.items[i]
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("> ")
		}

		check := "[ ]"
		if m.marked[item.Path] {
			check = "[x]"
		}

		sizeStr := "..."
		if item.Size > 0 {
			sizeStr = formatSize(item.Size)
		}
		
		sizeColored := colorForSize(item.Size).Render(sizeStr)
		line := fmt.Sprintf("%s%s %s (%s)", cursor, check, item.Path, sizeColored)

		if m.marked[item.Path] {
			listContent += markedStyle.Render(line) + "\n"
		} else if m.cursor == i {
			listContent += lipgloss.NewStyle().Bold(true).Render(line) + "\n"
		} else {
			listContent += line + "\n"
		}
	}

	s += listStyle.Render(listContent)
	s += footerStyle.Render("\n[Space: Select 1]  [A: Select All]  [Enter: Delete]  [Q: Quit]")

	return docStyle.Render(s)
}
