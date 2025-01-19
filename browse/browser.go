package browse

import (
	"os/exec"
)

type Browser interface {
	OpenURL(URL string) (*exec.Cmd, error)
}

type WebBrowser struct{}

func (browser WebBrowser) OpenURL(url string) (*exec.Cmd, error) {
	cmd := exec.Command("open", url)
	error := cmd.Start()
	return cmd, error
}

func NewBrowser() Browser {
	return WebBrowser{}
}
