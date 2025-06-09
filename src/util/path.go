package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// # CopyFile
//
// CopyFile copies the contents of the source file to the destination file.
// If the destination file does not exist, it will be created.
func CopyFile(srcFile, dstFile string) error {
	// Open the source file
	source, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer source.Close()

	// Create the destination file
	destination, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer destination.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(destination, source)
	return err
}

// # CopyDir
//
// CopyDir copies the contents of the source directory to the destination directory.
// If the destination directory does not exist, it will be created.
func CopyDir(src string, dst string) error {
	// Read all files in the source directory
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// Create the destination directory if it does not exist
	err = os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return err
	}

	// Iterate over all files in the source directory
	for _, file := range files {
		srcFile := filepath.Join(src, file.Name())
		dstFile := filepath.Join(dst, file.Name())

		if file.IsDir() {
			// If the file is a directory, recursively call CopyDir()
			err := CopyDir(srcFile, dstFile)
			if err != nil {
				return err
			}
		} else {
			// If the file is not a directory, call CopyFile()
			err := CopyFile(srcFile, dstFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// # DeleteDir
//
// DeleteDir deletes the specified directory and all its contents.
// It recursively removes all files and subdirectories within the given directory.
func DeleteDir(dir string) error {
	// Read the directory's contents
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// Iterate over each file in the directory
	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			// If the file is a directory, recursively delete its contents
			err := DeleteDir(filePath)
			if err != nil {
				return err
			}
		} else {
			// If the file is not a directory, remove it
			err := os.Remove(filePath)
			if err != nil {
				return err
			}
		}
	}

	// Remove the now-empty directory
	return os.Remove(dir)
}

// # MoveDir
//
// MoveDir moves a directory from oldPath to newPath by first copying
// the directory and its contents to the new location, and then deleting
// the original directory.
//
// It returns an error if the copy or delete operations fail.
func MoveDir(oldPath, newPath string) error {
	// Copy the contents of the old directory to the new directory
	err := CopyDir(oldPath, newPath)
	if err != nil {
		fmt.Println("Error copying folder:", err)
		return err
	}

	// Delete the original directory after successful copying
	err = DeleteDir(oldPath)
	if err != nil {
		fmt.Println("Error deleting original folder:", err)
		return err
	}

	// Inform the user that the directory has been moved successfully
	fmt.Println("Folder moved successfully.")
	return nil
}

// # CreateDir
//
// CreateDir creates a directory and all necessary parent directories
// at the specified path. If the directory already exists, no error
// is returned. Logs a fatal error if the directory cannot be created.
func CreateDir(path string) {
	// Attempt to create the directory with the specified permissions
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("unable to create dir: %s, error: %v", path, err)
	}
}
