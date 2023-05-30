package tail

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

const (
	MAXTXTLEN = 10000 // 最长的文本长度，为了避免发太大的邮件，最多只会发这么多字符的内容。
)

// 读取指定文件 fileName ，看是否有新的内容， 如果有则把数据读取到 buffer 中。
// 有新的时 ， hasNewInfo 为 true
// 是否有新的， 是基于 旧的 oldFileSize 来做比较的。 同时返回新的文件尺寸 newFileSize。
func Tail(fileName string, oldFileSize int64, buffer *bytes.Buffer) (hasNewInfo bool, newFileSize int64, err error) {
	hasNewInfo = false
	newFileSize = oldFileSize

	// 检查文件是否发生变化，
	// 通过比较之前记录的文件尺寸，来判断文件是否被新增内容。
	// 注意，文件被修改，或者删除部分内容，是判断不出来的。
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Println("err:", err)
		return
	}
	newFileSize = fileInfo.Size()

	if newFileSize <= 0 {
		// 文件被清理了。不用发送内容。
		hasNewInfo = false
		err = nil
		return
	}

	tailLen := newFileSize - oldFileSize

	if tailLen != 0 {
		hasNewInfo = true

		hasCut := false // 发送的内容是否被截断了。
		// 避免去太长的文本
		if tailLen > MAXTXTLEN || tailLen < 0 {
			// 太长，截断之， 文件被删除内容，意味着需要重新读取，读取最后的 10000 字符
			tailLen = MAXTXTLEN
			hasCut = true
		}

		tailLen = tailLen + 100 // 从文件最后往前读的位移量
		// 如果合并后的 tailLen 大于文件的长度，则仍然返回文件头。

		tailPos := 0 - tailLen

		err = fileReader(fileName, buffer, tailPos, hasCut)
	}
	return
}

// 从文件中读取最后 tailLen 长度的内容。
// 第一个不满的空行排除
func fileReader(fileName string, buffer *bytes.Buffer, tailPos int64, hasCut bool) (err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	log.Println("tailPos：", tailPos)
	file.Seek(tailPos, os.SEEK_END)
	scanner := bufio.NewScanner(file)

	var line int
	line = 0
	for scanner.Scan() {
		line++
		if hasCut && line <= 1 {
			// 如果需要发送的内容被截断了，
			// 这时候，我们假设 第一行不是完整的一行，跳过它。
			continue
		}
		buffer.WriteString(scanner.Text())
		buffer.WriteString("\r\n")

		//log.Println(line)
		// 获取从文件中读取的内容
		//fmt.Println(scanner.Text())
	}

	return
}
