package ascendex

import "fmt"

type Stock int

const (
	Ascendex = iota
)

type URLCreator interface {
	Create() string
}

type AscendexURL struct {
	host         string
	accountGroup string
	path         string
}

func NewAscendexURL() *AscendexURL {
	const (
		defaultHost         = "wss://ascendex.com"
		defaultAccountGroup = "0"
		defaultPath         = "api/pro/v1/stream"
	)

	return &AscendexURL{
		host:         defaultHost,
		accountGroup: defaultAccountGroup,
		path:         defaultPath,
	}
}

func (a *AscendexURL) SetHost(host string) *AscendexURL {
	if host != "" {
		a.host = host
	}

	return a
}

func (a *AscendexURL) SetAccountGroup(accountGroup string) *AscendexURL {
	if accountGroup != "" {
		a.accountGroup = accountGroup
	}

	return a
}

func (a *AscendexURL) Create() string {
	return fmt.Sprintf("%s/%s/%s", a.host, a.accountGroup, a.path)
}

func (a *AscendexURL) SetPath(path string) *AscendexURL {
	if path != "" {
		a.path = path
	}

	return a
}
