package logger

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"
)

// 日志文件句柄
var logfile *os.File

// 执行结果是否输出到文件上，如果 true 表示输出到文件，否则输出到屏幕上。
var outPutLogFile = false

// 如果要输出到文件，准备好文件
// out2File 是否要输出到文件， dir 如果要输出到文件，这个文件放在那个目录下。
func InitLogFile(out2File bool, dir string) {
	outPutLogFile = out2File
	if outPutLogFile {
		logfilename := getLogFileName(dir)
		logfile, err := os.OpenFile(logfilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0700)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		log.SetOutput(logfile)
	}
}

// 应用退出时 关闭文件句柄
func LoggerFinish() {
	if outPutLogFile {
		logfile.Close()
	}
}

// 获取当前文件名
func getLogFileName(dir string) (filename string) {
	file, _ := exec.LookPath(os.Args[0])
	fileName := filepath.Base(file)
	logfileName := fmt.Sprintf("%s.%s.log", fileName, time.Now().Format("2006-01-02"))

	dirPath := path.Join(dir, "log")

	if !isDirExists(dirPath) {
		// 参考 http://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permission-value
		os.Mkdir(dirPath, 0700)
	}

	logfileName = path.Join(dirPath, logfileName)
	fmt.Println("log文件：", logfileName)

	return logfileName
}

// 判断目录是否存在
func isDirExists(dir string) bool {
	fi, err := os.Stat(dir)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
	panic("not reached")
}
