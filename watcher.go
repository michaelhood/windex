package windex

import (
	"github.com/howeyc/fsnotify"
)

type Watcher struct {
	watcher *fsnotify.Watcher
}

func NewWatcher() (watcher *Watcher, err error) {
	return &Watcher{
		watcher: &fsnotify.Watcher,
	}
}

func (log *LogFile) Watch() (err error) {
	err = log.watchable()
	if err != nil {
		return
	}

	log.Watcher.Watch(log.FileName)

	return
}

func (log *LogFile) watchable() (err error) {
	if log.File == nil {
		err = ErrNoFile
	}

	if log.FileName == "" {
		err = ErrNoFileName
	}

	if log.FileSize < 0 {
		err = ErrInvalidFileSize
	}

	return
}

/*
go func() {
  for {
	select {
	case ev := <-watcher.Event:
		if ev != nil && ev.IsModify() && ev.Name == filename {
			log.moveAndFlush()
		}
	case err := <-watcher.Error:
		if err != nil {
		}
	}
  }
}()
*/
/*

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	logfile := &LogFile{}

	if err = log.updateFileSize(); err != nil {
		return nil, err
	}

*/
