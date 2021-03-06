/*
 * cons.go of  TeleBeam from modul TeleBeam
 * Created at 16.1.2022
 * Created from: dpiening
 * Last modified: 16.01.22, 15:44
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
	"time"
)

const ApplicationName string = "TeleBeam - Beam your Files"
const ApplicationLicense string = "MIT License"
const Version float32 = 1.0

var Year int
var Month time.Month
var Day int

var ApplicationCopyright string
var CFG = GetConfigValues()

func init() {
	Year = time.Now().Year()
	Month = time.Now().Month()
	Day = time.Now().Day()

	ApplicationCopyright = fmt.Sprintf("%s Version %f \n(C) 2021 - %d the TeleBeam Team, Dierk-Bent Piening\nLicense: %s", ApplicationName, Version, Year, ApplicationLicense)
}
