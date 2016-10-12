package GopyByEachEnv

import (
	"log"
	"os"
	"time"
)

// Modifier ...
type Modifier struct {
	beforeTime    time.Time
	checkFilePath string
}

func (m *Modifier) isModify() (bool, error) {
	fp, err := os.Stat(m.checkFilePath)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if m.beforeTime != fp.ModTime() {
		m.beforeTime = fp.ModTime()
		return true, nil
	}
	return false, nil
}
