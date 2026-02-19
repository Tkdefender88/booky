package bookmarklist

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/Tkdefender88/booky/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func switchToTags() tea.Msg {
	return messages.ChangeListFocusMsg{Target: messages.TagFocus}
}

func openBookmark(b bookmark) tea.Cmd {
	return func() tea.Msg {
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "linux":
			cmd = exec.Command("xdg-open", b.Url())
		case "darwin":
			cmd = exec.Command("open", b.Url())
		default:
			return messages.NewErrMsg(fmt.Errorf("unsupported platform"))
		}

		if err := cmd.Start(); err != nil {
			return messages.NewErrMsg(fmt.Errorf("failed to open browser: %w", err))
		}
		return nil
	}
}

