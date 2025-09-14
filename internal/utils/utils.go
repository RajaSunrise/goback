// internal/utils/utils.go

package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// RunCommand executes a shell command and returns an error if it fails
func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// RunCommandWithOutput executes a shell command and returns its output
func RunCommandWithOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// RunCommandInDir executes a shell command in a specific directory
func RunCommandInDir(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DirExists checks if a directory exists
func DirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// CreateDir creates a directory with all parent directories
func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// CreateFile creates a file with the given content
func CreateFile(path, content string) error {
	dir := filepath.Dir(path)
	if err := CreateDir(dir); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write content to file %s: %w", path, err)
	}

	return nil
}

// ReadFile reads the entire content of a file
func ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}
	return string(content), nil
}

// AppendToFile appends content to a file
func AppendToFile(path, content string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %s for appending: %w", path, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to append content to file %s: %w", path, err)
	}

	return nil
}

// CopyFile copies a file from src to dst
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", src, err)
	}
	defer srcFile.Close()

	dstDir := filepath.Dir(dst)
	if err2 := CreateDir(dstDir); err2 != nil {
		return fmt.Errorf("failed to create destination directory %s: %w", dstDir, err2)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", dst, err)
	}
	defer dstFile.Close()

	_, err = srcFile.WriteTo(dstFile)
	if err != nil {
		return fmt.Errorf("failed to copy content from %s to %s: %w", src, dst, err)
	}

	return nil
}

// DeleteFile deletes a file
func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file %s: %w", path, err)
	}
	return nil
}

// DeleteDir deletes a directory and all its contents
func DeleteDir(path string) error {
	err := os.RemoveAll(path)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete directory %s: %w", path, err)
	}
	return nil
}

// ListFiles lists all files in a directory
func ListFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list files in directory %s: %w", dir, err)
	}

	return files, nil
}

// ListDirs lists all directories in a directory
func ListDirs(dir string) ([]string, error) {
	var dirs []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", dir, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, filepath.Join(dir, entry.Name()))
		}
	}

	return dirs, nil
}

// GetCurrentDir returns the current working directory
func GetCurrentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}
	return dir, nil
}

// ChangeDir changes the current working directory
func ChangeDir(dir string) error {
	err := os.Chdir(dir)
	if err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", dir, err)
	}
	return nil
}

// GetFileSize returns the size of a file in bytes
func GetFileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, fmt.Errorf("failed to get file info for %s: %w", path, err)
	}
	return info.Size(), nil
}

// IsEmpty checks if a directory is empty
func IsEmpty(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, fmt.Errorf("failed to read directory %s: %w", dir, err)
	}
	return len(entries) == 0, nil
}

// MakeExecutable makes a file executable
func MakeExecutable(path string) error {
	err := os.Chmod(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to make file %s executable: %w", path, err)
	}
	return nil
}

// ReplaceInFile replaces all occurrences of old with new in a file
func ReplaceInFile(path, old, newVal string) error {
	content, err := ReadFile(path)
	if err != nil {
		return err
	}

	newContent := strings.ReplaceAll(content, old, newVal)

	err = CreateFile(path, newContent)
	if err != nil {
		return fmt.Errorf("failed to write updated content to file %s: %w", path, err)
	}

	return nil
}

// ReadLines reads all lines from a file
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file %s: %w", path, err)
	}

	return lines, nil
}

// WriteLines writes lines to a file
func WriteLines(path string, lines []string) error {
	content := strings.Join(lines, "\n")
	return CreateFile(path, content)
}

// GetRelativePath returns the relative path from base to target
func GetRelativePath(base, target string) (string, error) {
	relPath, err := filepath.Rel(base, target)
	if err != nil {
		return "", fmt.Errorf("failed to get relative path from %s to %s: %w", base, target, err)
	}
	return relPath, nil
}

// GetAbsolutePath returns the absolute path of a file
func GetAbsolutePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path of %s: %w", path, err)
	}
	return absPath, nil
}

// EnsureDir ensures that a directory exists, creating it if necessary
func EnsureDir(dir string) error {
	if !DirExists(dir) {
		return CreateDir(dir)
	}
	return nil
}

// CleanPath cleans a file path, resolving any relative components
func CleanPath(path string) string {
	return filepath.Clean(path)
}

// JoinPath joins path elements into a single path
func JoinPath(elements ...string) string {
	return filepath.Join(elements...)
}

// GetFileName returns the filename component of a path
func GetFileName(path string) string {
	return filepath.Base(path)
}

// GetFileExt returns the file extension
func GetFileExt(path string) string {
	return filepath.Ext(path)
}

// GetDirName returns the directory component of a path
func GetDirName(path string) string {
	return filepath.Dir(path)
}

// SplitPath splits a path into directory and filename
func SplitPath(path string) (string, string) {
	return filepath.Split(path)
}

// ValidateProjectName validates a project name
func ValidateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	if strings.Contains(name, " ") {
		return fmt.Errorf("project name cannot contain spaces")
	}

	if strings.Contains(name, "/") || strings.Contains(name, "\\") {
		return fmt.Errorf("project name cannot contain path separators")
	}

	return nil
}

// ValidateModulePath validates a Go module path
func ValidateModulePath(path string) error {
	if path == "" {
		return fmt.Errorf("module path cannot be empty")
	}

	// Basic validation - should contain at least one slash
	if !strings.Contains(path, "/") {
		return fmt.Errorf("module path should be in format 'domain.com/username/project'")
	}

	return nil
}

// SanitizeFileName sanitizes a filename by removing invalid characters
func SanitizeFileName(name string) string {
	// Replace invalid characters with underscores
	invalid := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range invalid {
		name = strings.ReplaceAll(name, char, "_")
	}
	return name
}

// GetHomeDir returns the user's home directory
func GetHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	return homeDir, nil
}

// IsExecutable checks if a file is executable
func IsExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode()&0111 != 0
}
