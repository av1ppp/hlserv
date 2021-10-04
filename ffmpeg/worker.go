package ffmpeg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/thanhpk/randstr"
)

type FFmpegWorker struct {
	cmd       *exec.Cmd
	cmdBackup exec.Cmd
	cmdString string

	murderPlanned bool
}

// Start starts the specified worker but does not wait for it to complete.
func (w *FFmpegWorker) Start() error {
	return w.cmd.Start()
}

// Run starts the specified worker and waits for it to complete.
func (w *FFmpegWorker) Run() error {
	return w.cmd.Run()
}

func (w *FFmpegWorker) Stop() error {
	w.murderPlanned = true

	if err := w.cmd.Process.Signal(os.Kill); err != nil {
		return err
	}
	w.cmd.Wait()
	return nil
}

// Генерация уникального имени для воркера.
func (f *FFmpeg) genWorkerKey() string {
	var workername string
	for {
		_workername := "ONCE_" + randstr.String(4)

		if _, exists := f.Worker(_workername); !exists {
			workername = _workername
			break
		}
	}
	return workername
}

func (f *FFmpeg) RunOnceWorker(files ...OptionIO) error {
	key := f.genWorkerKey()

	w, err := f.NewWorker(key, files...)
	if err != nil {
		return err
	}
	defer f.RmWorker(key)

	var stderr bytes.Buffer
	w.cmd.Stderr = &stderr

	err = w.Run()
	if err != nil {
		return fmt.Errorf("%s\nstderr:\n%s", err, stderr.String())
	}

	return nil
}

func (w *FFmpegWorker) Activated() bool {
	return w.cmd.ProcessState == nil
}

// FFmpeg process checker..
func (w *FFmpegWorker) Cron(timeout time.Duration) {
	for {
		w.cmd.Process.Wait()

		if w.murderPlanned {
			break
		}

		fmt.Fprintf(os.Stderr, "[WARN] the worker stopped the process unexpectedly")

		w.cmd = &(w.cmdBackup)
		if err := w.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "error when trying to restart the worker: %s\n", err)
		}

		time.Sleep(timeout)
	}
}

// Format files to string.
func (ff *FFmpeg) ftos(files []OptionIO) string {
	cmd := filepath.FromSlash(ff.BinPath) + " -loglevel error"

	for _, f := range files {
		cmd += " " + f.String()
	}

	return cmd
}

// Format string to exec.Cmd.
func (ff *FFmpeg) stoc(s string) *exec.Cmd {
	splits := strings.Split(s, " ")
	cmd := exec.Command(splits[0], splits[1:]...)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	if ff.Report != (report{}) {
		_ = os.MkdirAll(filepath.Dir(ff.Report.File), 0755)
		cmd.Env = append(cmd.Env, fmt.Sprintf("FFREPORT=file=%s:level=%d", ff.Report.File, ff.Report.LogLevel))
	}
	return cmd
}

func (ff *FFmpeg) NewWorker(key string, files ...OptionIO) (*FFmpegWorker, error) {
	if _, find := ff.Worker(key); find {
		return nil, ErrWorkerAlreadyExists
	}

	ff.Lock()
	defer ff.Unlock()

	s := ff.ftos(files)
	// fmt.Println(s)
	cmd := ff.stoc(s)

	w := FFmpegWorker{cmd: cmd, cmdBackup: *cmd, cmdString: s, murderPlanned: false}

	ff.workers[key] = &w

	return &w, nil
}

func (ff *FFmpeg) Worker(key string) (*FFmpegWorker, bool) {
	ff.RLock()
	defer ff.RUnlock()

	w, find := ff.workers[key]
	return w, find
}

func (ff *FFmpeg) RmWorker(key string) {
	w, find := ff.Worker(key)
	if !find {
		return
	}

	ff.Lock()
	defer ff.Unlock()

	if w.Activated() {
		w.Stop()
	}

	delete(ff.workers, key)
}
