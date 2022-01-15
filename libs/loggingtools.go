/*
 * loggingtools.go of  TeleBeam from modul TeleBeam
 * Created at 15.1.2022
 * Created from: dpiening
 * Last modified: 15.01.22, 14:48
 * Copyright (C) 2021 - 2022 Dierk-Bent Piening & the TeleBeam Team.
 *
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

// Package libs /*
package libs

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

var logpath string = CFG.Logging.FolderName

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	var acdate string = time.Now().Format("01-02-2006")

	var logfilename string = ApplicationName + "-" + acdate + "-debuglog.log"
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		err := os.Mkdir(logpath, 0755)
		if err != nil {
			log.Println(fmt.Sprintf(`FATAL ERROR!\nCould not Create Log Directory with the Name %s\nFatal Error Cause: %s \n`, logpath, err.Error()))
		} else {
			fmt.Println(fmt.Sprintf(`Successfully Created Log Directory with the Name %s \n`, logpath))
		}
	}
	var file, fileerr = os.OpenFile(path.Join(logpath, logfilename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileerr != nil {
		log.Fatal(fileerr)
	} else {
		log.Println(fmt.Sprintf(`Successfully Created Logfile %s in Directory %s`, logfilename, logpath))
	}

	//Define Loggers
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
}

func LogError(logerrortext string) {
	fmt.Println(logerrortext)
	ErrorLogger.Println("\n" + logerrortext + "\n")
}

func LogWarning(logwarningtext string) {
	fmt.Println(logwarningtext)
	WarningLogger.Println("\n" + logwarningtext + "\n")
}

func LogInfo(loginfotext string) {
	fmt.Println(loginfotext)
	InfoLogger.Println("\n" + loginfotext + "\n")
}
