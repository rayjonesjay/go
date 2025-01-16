package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
)

// defines an interface for reading files
type FileReader interface {
	ReadFile(filename string) error
}

// reads .txt files
type TXTreader struct{}

func (r *TXTreader) ReadFile(filename string) error {
	// check if filename has a .txt extension using the filepath function
	if filepath.Ext(filename) != ".txt" {
		return fmt.Errorf("unsupported file type %s", filepath.Ext(filename))
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	fmt.Printf("read .txt file: %s\nContent:\n%s\n", filename, content)
	return nil
}

type PluginManager struct {
	reader FileReader
}

func (pm *PluginManager) LoadPlugin(pluginPath string, reply *string) error {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return fmt.Errorf("failed to load plugin %v", err)
	}

	symbol, err := p.Lookup("NewReader")

	if err != nil {
		return fmt.Errorf("failed to find 'NewReader' symbol: %v",err)
	}

	newReader , ok := symbol.(func() FileReader)
	if !ok {
		return fmt.Errorf("invalid plugin: new reader must return FileReader")
	}

	pm.reader = newReader()
	*reply = "plugin loaded successfully"
	return nil
}

func (pm *PluginManager) ReadFile(filename string, reply *string) error {
	if pm.reader == nil {
		return fmt.Errorf("no file reader available")
	}

	err := pm.reader.ReadFile(filename)

	if err != nil {
		return err
	}

	*reply = "file read successfully"
	return nil
}

