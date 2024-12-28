package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyFile 复制文件
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", src, err)
	}
	defer sourceFile.Close()

	// 创建目标文件
	destinationFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", dst, err)
	}
	defer destinationFile.Close()

	// 复制文件内容
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file %s to %s: %w", src, dst, err)
	}

	// 获取源文件的权限，并设置到目标文件
	sourceInfo, err := sourceFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat source file: %w", err)
	}

	err = os.Chmod(dst, sourceInfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to set permissions on destination file: %w", err)
	}

	return nil
}

// CopyDir 递归复制文件夹 src 资源路径 dst 目标路径
func CopyDir(src, dst string) error {
	// 获取源目录的信息
	sourceInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat source directory %s: %w", src, err)
	}

	// 如果源路径是文件，直接复制文件
	if !sourceInfo.IsDir() {
		return CopyFile(src, dst)
	}

	// 创建目标文件夹
	err = os.MkdirAll(dst, sourceInfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to create destination directory %s: %w", dst, err)
	}

	// 打开源目录
	dir, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source directory %s: %w", src, err)
	}
	defer dir.Close()

	// 读取目录中的内容
	files, err := dir.Readdir(-1) // -1 表示读取所有文件
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", src, err)
	}

	// 递归复制目录中的每个文件和子目录
	for _, file := range files {
		srcPath := filepath.Join(src, file.Name())
		dstPath := filepath.Join(dst, file.Name())

		if file.IsDir() {
			// 递归复制子目录
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return fmt.Errorf("failed to copy subdirectory %s to %s: %w", srcPath, dstPath, err)
			}
		} else {
			// 复制文件
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return fmt.Errorf("failed to copy file %s to %s: %w", srcPath, dstPath, err)
			}
		}
	}

	return nil
}
