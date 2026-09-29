package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

const AppName = "chronarch"

func New() (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("toml")

	dev := os.Getenv("CHRONARCH_DEV") == "1"
	configDir := ".chronarch"
	if !dev {
		var err error
		configDir, err = userConfigDir(v)
		if err != nil {
			return nil, err
		}
	}

	v.AddConfigPath(configDir)

	if err := setDefaults(v, dev); err != nil {
		return nil, err
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := os.MkdirAll(configDir, 0700); err != nil {
				return nil, err
			}
			if err := v.SafeWriteConfig(); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// TODO: activate watcher when GUI or TUI components added; unnecessary for CLI
	// viper.WatchConfig()

	return v, nil
}

func setDefaults(v *viper.Viper, dev bool) error {
	if dev {
		v.SetDefault("storage.path", filepath.Join(".chronarch", AppName+".db"))
		return nil
	}

	userDataDir, err := userDataDir()
	if err != nil {
		return err
	}

	v.SetDefault("storage.path", filepath.Join(userDataDir, AppName, AppName+".db"))

	return nil
}

func userDataDir() (string, error) {
	switch runtime.GOOS {
	case "darwin":
		return os.UserConfigDir() // ~/Library/Application Support
	case "windows":
		dir := os.Getenv("LOCALAPPDATA")
		if dir == "" {
			return "", fmt.Errorf("LOCALAPPDATA is not set")
		}

		return dir, nil
	default:
		if dir := os.Getenv("XDG_DATA_HOME"); dir != "" {
			if !filepath.IsAbs(dir) {
				return "", fmt.Errorf("XDG_DATA_HOME must be absolute")
			}
			return dir, nil
		}

		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		return filepath.Join(home, ".local", "share"), nil
	}
}

func userConfigDir(v *viper.Viper) (string, error) {
	var configDir string
	var err error

	v.SetEnvPrefix(AppName)
	v.BindEnv("conf_dir")

	switch runtime.GOOS {
	case "windows":
		configDir, err = os.UserConfigDir()
	default:
		configDir, err = os.UserHomeDir()
	}
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, AppName), nil
}
