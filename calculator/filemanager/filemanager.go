package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
	"time"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func NewFileManager(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}

func (fm *FileManager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.InputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func (fm *FileManager) WriteJSON(data interface{}) error {

	time.Sleep(3 * time.Second)
	//save a file as json
	file, err := os.Create(fm.OutputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
