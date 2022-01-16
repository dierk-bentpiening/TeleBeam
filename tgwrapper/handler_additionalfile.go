/*
 * handler_additionalfile.go of  TeleBeam from modul TeleBeam
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
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetAdditionalFileInformation(FileID string) JSONResponse {
	var decodeErr error
	var TGURL string = fmt.Sprintf("%s/bot%s/getFile?file_id=%s", libs.CFG.Telegram.APIBaseURL, libs.CFG.Telegram.APIToken, FileID)
	RawResponse, retrievingErr := http.Get(TGURL)
	if retrievingErr != nil {
		libs.LogError(fmt.Sprintf("Fatal Error: attempt to retrieve AdditionalFileData from Telegram API!\nCause: %s\n", retrievingErr.Error()))
	} else {
		libs.LogInfo(fmt.Sprintf("Successfully retrieved AdditionalFile information for FileID: %s\n", FileID))
	}
	decoder := json.NewDecoder(RawResponse.Body)
	var AdditionalFileData JSONResponse
	decodeErr = decoder.Decode(&AdditionalFileData)
	if decodeErr != nil {
		libs.LogError(fmt.Sprintf("Fatal Error at Retreiving AdditionalFileData from Telegram API!\nCause: %s\n", decodeErr.Error()))
		os.Exit(120)
	}
	return AdditionalFileData
}
