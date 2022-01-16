/*
 * configloader.go of  TeleBeam from modul TeleBeam
 * Created at 16.1.2022
 * Created from: dpiening
 * Last modified: 16.01.22, 15:30
 * Copyright (C) 2021 - 2022 Dierk-Bent Piening & the TeleBeam Team.
 *
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package libs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Logging struct {
		FolderName string `yaml:"foldername"`
	} `yaml:"logging"`
	Database struct {
		Hostname string `yaml:"hostname"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
	} `yaml:"database"`
	Telegram struct {
		APIToken   string `yaml:"apitoken"`
		APIBaseURL string `yaml:"apibaseurl"`
	} `yaml:"telegram"`
	Download struct {
		TempPath string `yaml:"temppath"`
	} `yaml:"download"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

var userConfig Config

func GetConfigValues() Config {
	file, err := os.Open("conf/config.yaml")
	if err != nil {
		processError(err)
	} else {
		defer file.Close()

		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&userConfig)
		if err != nil {
			processError(err)
		} else {
			return userConfig
		}
	}
	return userConfig
}
