/*
 * downloads.go of  TeleBeam from modul TeleBeam
 * Created at 16.1.2022
 * Created from: dpiening
 * Last modified: 16.01.22, 18:59
 * Copyright (C) 2021 - 2022 Dierk-Bent Piening & the TeleBeam Team.
 *
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package tgwrapper

import (
	"TeleBeam/libs"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

var url string

func DownloadFile(fileinfo JSONResponse) (bool, error) {
	url = fmt.Sprintf("%s/file/bot%s/%s", libs.CFG.Telegram.APIBaseURL, libs.CFG.Telegram.APIToken, fileinfo.FileData.FilePath)

	rawresponse, err := http.Get(url)
	if err != nil {
		libs.LogError(fmt.Sprintf("Fatal Error: Attemp to Download File: %s <%s>\nCause: %s\n", fileinfo.FileData.FileUniqueID, fileinfo.FileData.FilePath, err.Error()))
		return false, err
	} else {
		if rawresponse.StatusCode == 200 {
			if _, err := os.Stat(libs.CFG.Download.TempPath); os.IsNotExist(err) {
				err := os.Mkdir(libs.CFG.Download.TempPath, 0755)
				if err != nil {
					libs.LogError(fmt.Sprintf(`FATAL ERROR!\nCould not Create Temp Directory with the Name %s\nFatal Error Cause: %s \n`, libs.CFG.Download.TempPath, err.Error()))
					return false, err

				} else {
					libs.LogInfo(fmt.Sprintf("Successfully Created Temp Directory with the Name %s\n", libs.CFG.Download.TempPath))
				}
			}
			file, errFileCreate := os.Create(path.Join(libs.CFG.Download.TempPath, fileinfo.FileData.FileUniqueID))
			if errFileCreate != nil {
				libs.LogError(fmt.Sprintf("Error creating file %s!\nCause: %s\n", fileinfo.FileData.FileUniqueID, errFileCreate.Error()))
				return false, errFileCreate
			} else {
				defer file.Close()
				_, errWriteFile := io.Copy(file, rawresponse.Body)
				if errWriteFile != nil {
					libs.LogError(fmt.Sprintf("Fatal Error: Attemp to write file %s failed!\nCause: %s\n", fileinfo.FileData.FileUniqueID, errWriteFile.Error()))
					return false, errWriteFile
				} else {
					libs.LogInfo(fmt.Sprintf("Successfully wrote file %s to Temp Directory %s\n", fileinfo.FileData.FileUniqueID, libs.CFG.Download.TempPath))
					defer rawresponse.Body.Close()
					return true, errWriteFile
				}
			}
		} else {
			var errMessage string = fmt.Sprintf("Fatal Error: Attemp to download file %s failes with status code %d", fileinfo.FileData.FileUniqueID, rawresponse.StatusCode)
			var statusCodeError error = errors.New(errMessage)
			libs.LogError(errMessage)
			return false, statusCodeError
		}
	}
}
