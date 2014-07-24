package services

import (
	"errors"
	"inspeqtor/core"
	"inspeqtor/util"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Upstart struct {
	Path        string
	dummyOutput string
}

var (
	pidScanner *regexp.Regexp = regexp.MustCompile(" (?:start|stop)\\/(?:running|waiting)(?:, process (\\d+))?")
)

func DetectUpstart(path string) (*Upstart, error) {
	result, err := util.FileExists(path)
	if err != nil {
		return nil, err
	}

	if !result {
		util.Debug("upstart not detected, no " + path)
		return nil, nil
	}

	matches, err := filepath.Glob(path + "/*.conf")
	if err != nil {
		return nil, err
	}

	if len(matches) > 0 {
		util.Debug("Detected upstart in " + path)
		return &Upstart{path, ""}, nil
	}

	util.Debug("upstart not detected, empty " + path)
	return nil, nil
}

func (u *Upstart) LookupService(serviceName string) (core.ProcessId, core.ServiceStatus, error) {
	matches, err := filepath.Glob(u.Path + "/" + serviceName + ".conf")
	if err != nil {
		return 0, core.Unknown, err
	}

	if len(matches) == 0 {
		return -1, core.Unknown, nil
	}

	var sout []byte
	if len(u.dummyOutput) != 0 {
		sout = []byte(u.dummyOutput)
	} else {
		cmd := exec.Command("status", serviceName)
		sout, err = cmd.CombinedOutput()
		if err != nil {
			return 0, core.Unknown, err
		}
	}

	lines, err := util.ReadLines(sout)
	if len(lines) != 1 {
		return 0, core.Unknown, errors.New("Unexpected output: " + strings.Join(lines, "\n"))
	}

	// mysql start/running, process 14190
	// sshdgenkeys stop/waiting
	line := lines[0]
	results := pidScanner.FindStringSubmatch(line)

	if len(results) > 1 && len(results[1]) > 0 {
		pid, err := strconv.ParseInt(results[1], 10, 32)
		if err != nil {
			return 0, core.Unknown, err
		}
		return core.ProcessId(pid), core.Up, nil
	}
	if len(results) > 1 {
		return -1, core.Unknown, nil
	}

	return 0, core.Unknown, errors.New("Unknown upstart output: " + line)
}
