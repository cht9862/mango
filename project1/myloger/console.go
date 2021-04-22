package myloger

import "fmt"
import "time"

// 向终端输出
type Loger struct{
	DebugBool bool
}

	
func NewLog() Loger {
	return Loger{}
}

func (l *Loger) Debug(msg string) {
	if !l.DebugBool {
		return
	} 
	now := time.Now()
	fmt.Printf("\x1b[%dm[%s] %s->  %s \x1b[0m\n", 33,"Debug",now.Format("2006-01-02 03:04:05 UTC"),msg)
}
func (l *Loger) Trace(msg string) {
	if !l.DebugBool {
		return
	} 
	now := time.Now()
	fmt.Printf("\x1b[%dm[%s] %s->  %s \x1b[0m\n", 33,"Trace",now.Format("2006-01-02 03:04:05 UTC"),msg)
}
func (l *Loger) Info(msg string) {
	now := time.Now()
	fmt.Printf("\x1b[%dm[%s] %s->  %s \x1b[0m\n", 33,"Info",now.Format("2006-01-02 03:04:05 UTC"),msg)
}
func (l *Loger) Warning(msg string) {
	now := time.Now()
	fmt.Printf("\x1b[%dm[%s] %s->  %s \x1b[0m\n", 31,"Warning",now.Format("2006-01-02 03:04:05 UTC"),msg)
}
func (l *Loger) Error(msg string) {
	now := time.Now()
	fmt.Printf("\x1b[%dm[%s] %s->  %s \x1b[0m\n", 31,"Error",now.Format("2006-01-02 03:04:05 UTC"),msg)
}
func (l *Loger) Fatal(msg string) {
	now := time.Now()
	fmt.Printf("\x1b[%dm[%s] %s->  %s \x1b[0m\n", 31,"Fatal",now.Format("2006-01-02 03:04:05 UTC"),msg)
}