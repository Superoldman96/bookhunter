package argument

import (
	"os"
	"strings"

	"golang.org/x/exp/utf8string"

	"github.com/bookstairs/bookhunter/internal/client"
	"github.com/bookstairs/bookhunter/internal/fetcher"
)

var (
	// Flags for talebook registering.

	Username = ""
	Password = ""
	Email    = ""

	// Common flags.

	Website    = ""
	UserAgent  = client.DefaultUserAgent
	Proxy      = ""
	ConfigRoot = ""

	// Common download flags.

	Formats = []string{
		string(fetcher.EPUB),
		string(fetcher.AZW3),
		string(fetcher.MOBI),
		string(fetcher.PDF),
		string(fetcher.ZIP),
	}
	Extract         = false
	DownloadPath, _ = os.Getwd()
	InitialBookID   = int64(1)
	Rename          = false
	Thread          = 1
	RateLimit       = 30

	// Drive ISP configurations.

	RefreshToken = "" // Aliyun Drive

	TelecomUsername = "" // Telecom Cloud.
	TelecomPassword = ""

	// Telegram configurations.

	ChannelID = ""
	Mobile    = ""
	ReLogin   = false
	AppID     = int64(0)
	AppHash   = ""
)

// NewFetcher will create the fetcher by the command line arguments.
func NewFetcher(category fetcher.Category, properties map[string]string) (fetcher.Fetcher, error) {
	cc, err := client.NewConfig(Website, UserAgent, Proxy, ConfigRoot)
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
		Extract:       Extract,
		DownloadPath:  DownloadPath,
		InitialBookID: InitialBookID,
		Rename:        Rename,
		Thread:        Thread,
		RateLimit:     RateLimit,
		Properties:    properties,
	})
}

func HideSensitive(content string) string {
	if content == "" {
		return ""
	}

	// Preserve only the prefix and suffix, replace others with *
	s := utf8string.NewString(content)
	c := s.RuneCount()

	// Determine the visible length of the prefix and suffix.
	l := 1
	if c >= 9 {
		l = 3
	} else if c >= 6 {
		l = 2
	}

	return s.Slice(0, l) + strings.Repeat("*", c-l*2) + s.Slice(c-l, c)
}
