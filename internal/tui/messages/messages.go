package messages

import "github.com/Tkdefender88/booky/internal/bookmarks"

type ErrMsg struct {
	Err error
}

func (e ErrMsg) Error() string {
	return e.Err.Error()
}

func NewErrMsg(err error) ErrMsg {
	return ErrMsg{Err: err}
}

type BookmarkAddedMsg struct{}

type DbConnectedMsg struct {
	Manager *bookmarks.BookmarkManager
	Close   func() error
}

type BookmarksFetchedMsg struct {
	Bookmarks []bookmarks.Bookmark
	Tags      []string
}

type FormClosedStatus int

const (
	FormClosedSuccess FormClosedStatus = iota
	FormClosedAborted
)

type FormClosedMsg struct {
	Status FormClosedStatus
	Name   string
	Url    string
	Desc   string
	Tags   []string
}

type ListFocusTarget string

const (
	BookmarkFocus ListFocusTarget = "bookmarks"
	TagFocus      ListFocusTarget = "tags"
	FormFocus     ListFocusTarget = "form"
)

type ChangeListFocusMsg struct {
	Target ListFocusTarget
}
