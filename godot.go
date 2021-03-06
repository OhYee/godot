package godot

import (
	"fmt"
	"io"
	"os/exec"
)

func Dot(src []byte, args ...string) (dist []byte, err error) {
	var dotPath string
	var stdin io.WriteCloser

	dotPath, err = exec.LookPath("dot")
	if err != nil {
		return
	}

	params := append([]string{"-Tsvg"}, args...)
	cmd := exec.Command(dotPath, params...)
	stdin, err = cmd.StdinPipe()
	if err != nil {
		return
	}

	_, err = stdin.Write(src)
	if err != nil {
		return
	}

	stdin.Close()
	if err != nil {
		return
	}

	dist, err = cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf(string(dist))
		dist = []byte{}
	}

	return
}
