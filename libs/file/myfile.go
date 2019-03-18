//文件处理
package file

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

/**
 * 文件操作:
 * 1. os package
 *    os/file
 * 2. buffer package
 * 3. io package
 * 4. Path package
 * 5.
 *
 */
type MyFile struct {
	Path     string
	FileName string
}

func (f *MyFile) FilePath() string {
	return f.Path + string(filepath.Separator) + f.FileName
}

func (myFile *MyFile) getFilePath() string {
	return myFile.Path + string(filepath.Separator) + myFile.FileName
}

//创建文件夹,若不存在则创建,反之不处理
func CreateDir(file MyFile) error {
	var err error
	//文件夹处理
	if err = readDirectrory(file); os.IsNotExist(err) {
		log.Println(file.Path, "不存在,", "将创建")
		if err = createDirectory(file); err != nil {
			log.Fatal(file.Path, " 文件夹创建错误:", err)
		}
		log.Println(file.Path, "创建成功")
	} else if isNotDirectory(err) {
		log.Fatal(file.Path, "已存在，但非文件夹:,", err)
	}
	return err
}

//创建文件
func CreateFile(file MyFile) error {
	return errors.New("not implement this methdo")
}

/**
 * 删除文件夹文件测试入口
 */
func RemoveDirFile(file MyFile) {
	log.Println("--------------开始删除文件-------------")
	if err := removeFile(file.getFilePath()); err != nil {
		log.Fatal("文件 ", file.getFilePath(), " 删除失败 :", err)
	} else {
		log.Print("文件 ", file.getFilePath(), " 已删除")
	}
	log.Println("--------------开始删除文件夹-------------")
	if err := removeFile(file.Path); err != nil {
		log.Fatal("文件夹 ", file.Path, " 删除失败 :", err)
	} else {
		log.Print("文件夹 ", file.Path, " 已删除")
	}
}

/**
 * 文件操作读写测试入口
 */
func fileHandleTest(file MyFile) {
	log.Println("--------------开始文件数据写入-------------")
	var file_ *os.File
	var err error
	if file_, err = readWriteFile(file.getFilePath()); err != nil {
		log.Fatal("文件 ", file.getFilePath(), " 打开失败 :", err)
	}
	defer (*file_).Close()
	//prefix := []byte{0xCA,0xFE,0xBA,0xBE}
	//4 字节前缀 TODO
	prefix := []byte{0xBA, 0xBE, 0xFA, 0xCE}
	//1字节主版本本号
	version_major := 999
	//1字节次版本号
	version_minor := 999
	prefix = append(prefix, byte(version_major), byte(version_minor))
	if _, err = (*file_).Write(prefix); err != nil {
		log.Fatal("文件 ", file.getFilePath(), " 写入失败 :", err)
	}
	if _, err = (*file_).Read(prefix); err != nil && io.EOF != err {
		log.Fatal("文件 ", file.getFilePath(), " 读取失败 :", err)
	}
	log.Println("文件 ", file.getFilePath(), " 写入数据为 :", string(prefix))

	fileContent := `这是一个普通文件
			你说呢，"【普通】"
			吗？`
	if _, err = (*file_).WriteString(fileContent); err != nil {
		log.Fatal("文件 ", file.getFilePath(), " 写入失败 :", err)
	}
	log.Println("文件 ", file.getFilePath(), " 写入数据为 :", fileContent)

}

/**
 * 修改文件大小
 */
func truncFile(file string, size int64) error {
	return os.Truncate(file, size)
}

/**
 * 创建文件夹
 */
func createDirectory(file MyFile) error {
	return os.MkdirAll(file.Path, os.ModeDir)
}

/**
 * 获取文件夹
 */
func readDirectrory(file MyFile) error {
	folder_, err := openFile(file.Path)
	defer folder_.Close()
	if err == nil {
		if fileInfo_, err := folder_.Stat(); err != nil && !fileInfo_.IsDir() {
			return FILE_NOT_DIRECTORY
		}

	}
	return err
}

/**
 * 创建文件
 */
func createFile(name string) (*os.File, error) {
	return os.Create(name)
}

/**
 * 打开文件或文件夹
 * Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。
 */
func openFile(Path string) (*os.File, error) {
	return os.Open(Path)
}

/**
 * 打开文件
 * Open打开一个文件用于读写
 */
func readWriteFile(Path string) (*os.File, error) {
	return os.OpenFile(Path, os.O_RDWR|os.O_APPEND, os.ModeType)
}

/**
 * 读取文件
 * OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
 */
func readFile(name string) (os.FileInfo, error) {
	var err error
	var file *os.File
	if file, err = os.OpenFile(name, os.O_RDWR, 0); err == nil {
		defer file.Close()
		if fileInfo_, err := file.Stat(); err == nil {
			if fileInfo_.IsDir() {
				return nil, FILE_IS_DIRECTORY
			}
			return fileInfo_, nil
		} else {
			return fileInfo_, nil
		}
	} else {
		//log.Fatal("readFile err: ", err)
		return nil, err
	}
}

/**
 * 写入文件
 */
func writeFile() {

}

/**
 * 压缩文件
 */
func zipFile() {

}

/**
 * 删除文件
 */
func removeFile(file string) error {
	return os.RemoveAll(file)
}

var (
	FILE_NOT_DIRECTORY = errors.New("FILE IS NOT DIRECTORY")
	FILE_IS_DIRECTORY  = errors.New("FILE IS DIRECTORY")
)

/**
 * 非文件夹
 */
func isNotDirectory(err error) bool {
	return err == FILE_NOT_DIRECTORY
}

/**
 * 文件夹
 */
func isDirectory(err error) bool {
	return err == FILE_IS_DIRECTORY
}
