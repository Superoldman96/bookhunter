package flags

import (
	"os"
	"runtime"
	"strings"

	"github.com/bookstairs/bookhunter/internal/client"
	"github.com/bookstairs/bookhunter/internal/fetcher"
	"github.com/bookstairs/bookhunter/internal/file"
)

var (
	// Flags for talebook registering.

	Username = ""
	Password = ""
	Email    = ""

	// Common flags.

	Website    = ""
	Proxy      = ""
	ConfigRoot = ""
	Keywords   []string
	Retry      = 3
	SkipError  = true

	// Common download flags.

	Formats = []string{
		string(file.EPUB),
		string(file.AZW3),
		string(file.MOBI),
		string(file.PDF),
		string(file.ZIP),
	}
	Extract         = false
	DownloadPath, _ = os.Getwd()
	InitialBookID   = int64(1)
	Rename          = false
	Thread          = runtime.NumCPU()
	RateLimit       = 30

	// Telegram configurations.

	ChannelID = ""
	Mobile    = ""
	ReLogin   = false
	AppID     = int64(0)
	AppHash   = ""

	// SoBooks configurations.

	SoBooksCode = "881120"
)

func NewClientConfig() (*client.Config, error) {
	return client.NewConfig(Website, Proxy, ConfigRoot)
}

// NewFetcher will create the fetcher by the command line arguments.
func NewFetcher(category fetcher.Category, properties map[string]string) (fetcher.Fetcher, error) {
	cc, err := NewClientConfig()
	if err != nil {
		return nil, err
	}

	fs, err := fetcher.ParseFormats(Formats)
	if err != nil {
		return nil, err
	}

	return fetcher.New(&fetcher.Config{
		Config:        cc,
		Category:      category,
		Formats:       fs,
		Keywords:      Keywords,
		Extract:       Extract,
		DownloadPath:  DownloadPath,
		InitialBookID: InitialBookID,
		Rename:        Rename,
		Thread:        Thread,
		RateLimit:     RateLimit,
		Properties:    properties,
		Retry:         Retry,
		SkipError:     SkipError,
	})
}

// HideSensitive will replace the sensitive content with star but keep the original length.
func HideSensitive(content string) string {
	if content == "" {
		return ""
	}

	// Preserve only the prefix and suffix, replace others with *
	s := []rune(content)
	c := len(s)

	// Determine the visible length of the prefix and suffix.
	l := 1
	if c >= 9 {
		l = 3
	} else if c >= 6 {
		l = 2
	}

	return string(s[0:l]) + strings.Repeat("*", c-l*2) + string(s[c-l:c])
}
