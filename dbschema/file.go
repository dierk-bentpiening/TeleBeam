/*
 * file.go of  TeleBeam from modul TeleBeam
 * Created at 16.1.2022
 * Created from: dpiening
 * Last modified: 16.01.22, 22:54
 * Copyright (C) 2021 - 2022 Dierk-Bent Piening & the TeleBeam Team.
 *
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package dbschema

import "gorm.io/gorm"

type FileEntry struct {
	gorm.Model
	FileEntryID       string `gorm:"primaryKey"`
	GUID              string `gorm:"unique"`
	IsAudio           bool   `gorm:"dafault:false"`
	IsFile            bool   `gorm:"dafault:false"`
	IsMovie           bool   `gorm:"dafault:false"`
	IsImage           bool   `gorm:"dafault:false"`
	IsOtherFile       bool   `gorm:"dafault:false"`
	IsDownloadBlocked bool   `gorm:"dafault:false"`
	FilePath          string `gorm:"unique"`
	FileID            string `gorm:"unique"`
	UniqueID          string `gorm:"unique"`
}
