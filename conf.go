package main

import (
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

type configFile struct {
	Debug bool `toml:"debug"`

	HTTP struct {
		Addr     string `toml:"addr"`
		XHeaders bool   `toml:"xheaders"`
	} `toml:"http_server"`

	HTTPS struct {
		Addr     string `toml:"addr"`
		CertFile string `toml:"cert_file"`
		KeyFile  string `toml:"key_file"`
	} `toml:"https_server"`

	S3 struct {
		Region      string `toml:"region"`
		Credentials string `toml:"credentials"`
		EndPoint    string `toml:"endpoint"`
		LogLevel    uint   `toml:"logLevel"`
		BucketName  string `toml:"bucket_name"`
		DisableSSL  bool   `toml:"disableSSL"`
		LogHTTPBody bool   `toml:"logHttpBody"`
	} `toml:"s3"`
}

// loadConfig reads and parses the configuration file.
func loadConfig(filename string) (*configFile, error) {
	c := &configFile{}
	if _, err := toml.DecodeFile(filename, c); err != nil {
		return nil, err
	}

	// Make files' path relative to the config file's directory.
	basedir := filepath.Dir(filename)
	relativePath(basedir, &c.HTTPS.CertFile)
	relativePath(basedir, &c.HTTPS.KeyFile)

	return c, nil
}

func relativePath(basedir string, path *string) {
	p := *path
	if p != "" && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}
