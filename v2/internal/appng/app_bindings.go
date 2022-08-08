//go:build bindings
// +build bindings

package appng

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/leaanthony/gosod"
	"github.com/wailsapp/wails/v2/internal/binding"
	"github.com/wailsapp/wails/v2/internal/frontend/runtime/wrapper"
	"github.com/wailsapp/wails/v2/internal/fs"
	"github.com/wailsapp/wails/v2/internal/logger"
	"github.com/wailsapp/wails/v2/internal/project"
	"github.com/wailsapp/wails/v2/pkg/options"
)

// App defines a Wails application structure
type App struct {
	logger     *logger.Logger
	appoptions *options.App
}

func (a *App) Run() error {

	// Create binding exemptions - Ugly hack. There must be a better way
	bindingExemptions := []interface{}{a.appoptions.OnStartup, a.appoptions.OnShutdown, a.appoptions.OnDomReady}
	appBindings := binding.NewBindings(a.logger, a.appoptions.Bind, bindingExemptions)

	var configDir string
	flag.StringVar(&configDir, "config", "", "Path to Wails config file")
	flag.Parse()

	err := generateBindings(configDir, appBindings)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Shutdown() {}

// CreateApp creates the app!
func CreateApp(appoptions *options.App) (*App, error) {
	// Set up logger
	myLogger := logger.New(appoptions.Logger)
	myLogger.SetLogLevel(appoptions.LogLevel)

	result := &App{
		logger:     myLogger,
		appoptions: appoptions,
	}

	return result, nil

}

func generateBindings(configDir string, bindings *binding.Bindings) error {

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	if configDir == "" {
		configDir = cwd
	}

	projectConfig, err := project.Load(configDir)
	if err != nil {
		return err
	}

	if projectConfig.WailsJSDir == "" {
		projectConfig.WailsJSDir = filepath.Join(cwd, "frontend")
	}
	wrapperDir := filepath.Join(projectConfig.WailsJSDir, "wailsjs", "runtime")
	_ = os.RemoveAll(wrapperDir)
	extractor := gosod.New(wrapper.RuntimeWrapper)
	err = extractor.Extract(wrapperDir, nil)
	if err != nil {
		return err
	}

	targetDir := filepath.Join(projectConfig.WailsJSDir, "wailsjs", "go")
	err = os.RemoveAll(targetDir)
	if err != nil {
		return err
	}
	_ = fs.MkDirs(targetDir)

	err = bindings.GenerateGoBindings(targetDir)
	if err != nil {
		return err
	}

	return nil

}
