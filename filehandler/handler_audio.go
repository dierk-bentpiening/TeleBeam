/*
 * handler_audio.go of  TeleBeam from modul TeleBeam
 * Created at 16.1.2022
 * Created from: dpiening
 * Last modified: 16.01.22, 23:10
 * Copyright (C) 2021 - 2022 Dierk-Bent Piening & the TeleBeam Team.
 *
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package filehandler

import (
	"TeleBeam/dbhandler"
	"TeleBeam/dbschema"
	"TeleBeam/libs"
	"TeleBeam/mimetools"
	"TeleBeam/tgwrapper"
	"bytes"
	"fmt"
	"github.com/dhowden/tag"
	"github.com/google/uuid"
	tb "gopkg.in/tucnak/telebot.v2"

	"os"
	"path"
)

func GetTitleByInterpreter(m *tb.Message, b *tb.Bot) {
	var titles []dbschema.AudioEntitiy
	var searchInterpret string = "%" + m.Payload + "%"
	dbhandler.DB.Where("Interpret LIKE ?", searchInterpret).Find(&titles)
	var message string
	var anmountTitles int = len(titles)
	if anmountTitles == 0 {
		b.Send(m.Sender, fmt.Sprintf("<b>Sorry %s,</b>\ni do <u>not</u> have any Titles for Interpret %s in the Database", m.Sender.FirstName, m.Payload), tb.ModeHTML)

	} else {
		var counter int = 1
		var messageBuffer string
		for _, title := range titles {
			if counter == 1 {
				messageBuffer = fmt.Sprintf("<b>Found %d Titles of Interpret %s</b>\n\n<b>Title %d of %d</b>\n<b>Title:</b> %s\n<b>Interpret:</b> %s\n<b>Album:</b> %s\n<b>Composer:</b> %s\n<b>Year:</b> %d\n<b>Genre:</b> %s\n<b><u>FileID: %s</u></b>\n\n", anmountTitles, m.Payload, counter, anmountTitles, title.Title, title.Interpret, title.Album, title.Composer, title.Year, title.Genre, title.ID)
				message = message + messageBuffer
			} else {
				messageBuffer = fmt.Sprintf("<b>Title %d of %d\n</b><b>Title:</b> %s\n<b>Interpret:</b> %s\n<b>Album:</b> %s\n<b>Composer:</b> %s\n<b>Year:</b> %d\n<b>Genre:</b> %s\n<b><u>FileID: %s</u></b>\n\n", counter, anmountTitles, title.Title, title.Interpret, title.Album, title.Composer, title.Year, title.Genre, title.ID)
				message = message + messageBuffer
			}

			counter++
		}
		b.Send(m.Sender, message, tb.ModeHTML)
	}

}
func GetByTitle(m *tb.Message, b *tb.Bot) {
	var titles []dbschema.AudioEntitiy
	var searchInterpret string = "%" + m.Payload + "%"
	dbhandler.DB.Where("Title LIKE ?", searchInterpret).Find(&titles)
	var message string
	var anmountTitles int = len(titles)
	if anmountTitles == 0 {
		b.Send(m.Sender, fmt.Sprintf("<b>Sorry %s,</b>\ni do <u>not</u> have any Titles for Title %s in the Database", m.Sender.FirstName, m.Payload), tb.ModeHTML)

	} else {
		var counter int = 1
		var messageBuffer string
		for _, title := range titles {
			if counter == 1 {
				messageBuffer = fmt.Sprintf("<b>Found %d Titles matching Title %s</b>\n\n<b>Title %d of %d</b>\n<b>Title:</b> %s\n<b>Interpret:</b> %s\n<b>Album:</b> %s\n<b>Composer:</b> %s\n<b>Year:</b> %d\n<b>Genre:</b> %s\n<b><u>FileID: %s</u></b>\n\n", anmountTitles, m.Payload, counter, anmountTitles, title.Title, title.Interpret, title.Album, title.Composer, title.Year, title.Genre, title.ID)
				message = message + messageBuffer
			} else {
				messageBuffer = fmt.Sprintf("<b>Title %d of %d\n</b><b>Title:</b> %s\n<b>Interpret:</b> %s\n<b>Album:</b> %s\n<b>Composer:</b> %s\n<b>Year:</b> %d\n<b>Genre:</b> %s\n<b><u>FileID: %s</u></b>\n\n", counter, anmountTitles, title.Title, title.Interpret, title.Album, title.Composer, title.Year, title.Genre, title.ID)
				message = message + messageBuffer
			}

			counter++
		}
		b.Send(m.Sender, message, tb.ModeHTML)
	}
}
func GetByGenre(m *tb.Message, b *tb.Bot) {
	var titles []dbschema.AudioEntitiy
	var searchInterpret string = "%" + m.Payload + "%"
	dbhandler.DB.Where("genre LIKE ?", searchInterpret).Find(&titles)
	var message string
	var anmountTitles int = len(titles)
	if anmountTitles == 0 {
		b.Send(m.Sender, fmt.Sprintf("<b>Sorry %s,</b>\ni do <u>not</u> have any Titles for Genre %s in the Database", m.Sender.FirstName, m.Payload), tb.ModeHTML)

	} else {
		var counter int = 1
		var messageBuffer string
		for _, title := range titles {
			if counter == 1 {
				messageBuffer = fmt.Sprintf("<b>Found %d Titles of Genre %s</b>\n\n<b>Title %d of %d</b>\n<b>Title:</b> %s\n<b>Interpret:</b> %s\n<b>Album:</b> %s\n<b>Composer:</b> %s\n<b>Year:</b> %d\n<b>Genre:</b> %s\n<b><u>FileID: %s</u></b>\n\n", anmountTitles, m.Payload, counter, anmountTitles, title.Title, title.Interpret, title.Album, title.Composer, title.Year, title.Genre, title.ID)
				message = message + messageBuffer
			} else {
				messageBuffer = fmt.Sprintf("<b>Title %d of %d\n</b><b>Title:</b> %s\n<b>Interpret:</b> %s\n<b>Album:</b> %s\n<b>Composer:</b> %s\n<b>Year:</b> %d\n<b>Genre:</b> %s\n<b><u>FileID: %s</u></b>\n\n", counter, anmountTitles, title.Title, title.Interpret, title.Album, title.Composer, title.Year, title.Genre, title.ID)
				message = message + messageBuffer
			}

			counter++
		}
		b.Send(m.Sender, message, tb.ModeHTML)
	}
}
func GetByYear(m *tb.Message, b *tb.Bot) {
	var titles []dbschema.AudioEntitiy
	var searchInterpret string = "%" + m.Payload + "%"
	dbhandler.DB.Where("year LIKE ?", searchInterpret).Find(&titles)
	var message string
	var anmountTitles int = len(titles)
	if anmountTitles == 0 {
		b.Send(m.Sender, fmt.Sprintf("<b>Sorry %s,</b>\ni do <u>not</u> have any Titles for Year %s in the Database", m.Sender.FirstName, m.Payload), tb.ModeHTML)

	} else {
		var counter int = 1
		var messageBuffer string
		for _, title := range titles {
			if counter == 1 {
				messageBuffer = fmt.Sprintf("<b>Found %d Titles of Year %s</b>\n\n<b>Title %d of %d</b>\n<b>Title:</b> %s\n<b>Interpret:</b> %s\n<b>Album:</b> %s\n<b>Composer:</b> %s\n<b>Year:</b> %d\n<b>Genre:</b> %s\n<b><u>FileID: %s</u></b>\n\n", anmountTitles, m.Payload, counter, anmountTitles, title.Title, title.Interpret, title.Album, title.Composer, title.Year, title.Genre, title.ID)
				message = message + messageBuffer
			} else {
				messageBuffer = fmt.Sprintf("<b>Title %d of %d\n</b><b>Title:</b> %s\n<b>Interpret:</b> %s\n<b>Album:</b> %s\n<b>Composer:</b> %s\n<b>Year:</b> %d\n<b>Genre:</b> %s\n<b><u>FileID: %s</u></b>\n\n", counter, anmountTitles, title.Title, title.Interpret, title.Album, title.Composer, title.Year, title.Genre, title.ID)
				message = message + messageBuffer
			}

			counter++
		}
		b.Send(m.Sender, message, tb.ModeHTML)
	}
}

func GetByAlbum(m *tb.Message, b *tb.Bot) {
	var titles []dbschema.AudioEntitiy
	var searchInterpret string = "%" + m.Payload + "%"
	dbhandler.DB.Where("album LIKE ?", searchInterpret).Find(&titles)
	var message string
	var anmountTitles int = len(titles)
	if anmountTitles == 0 {
		b.Send(m.Sender, fmt.Sprintf("<b>Sorry %s,</b>\ni do <u>not</u> have any Titles for Album %s in the Database", m.Sender.FirstName, m.Payload), tb.ModeHTML)

	} else {
		var counter int = 1
		var messageBuffer string
		for _, title := range titles {
			if counter == 1 {
				messageBuffer = fmt.Sprintf("<b>Found %d Titles of Album %s</b>\n\n<b>Title %d of %d</b>\n<b>Title:</b> %s\n<b>Interpret:</b> %s\n<b>Album:</b> %s\n<b>Composer:</b> %s\n<b>Year:</b> %d\n<b>Genre:</b> %s\n<b><u>FileID: %s</u></b>\n\n", anmountTitles, m.Payload, counter, anmountTitles, title.Title, title.Interpret, title.Album, title.Composer, title.Year, title.Genre, title.ID)
				message = message + messageBuffer
			} else {
				messageBuffer = fmt.Sprintf("<b>Title %d of %d\n</b><b>Title:</b> %s\n<b>Interpret:</b> %s\n<b>Album:</b> %s\n<b>Composer:</b> %s\n<b>Year:</b> %d\n<b>Genre:</b> %s\n<b><u>FileID: %s</u></b>\n\n", counter, anmountTitles, title.Title, title.Interpret, title.Album, title.Composer, title.Year, title.Genre, title.ID)
				message = message + messageBuffer
			}

			counter++
		}
		b.Send(m.Sender, message, tb.ModeHTML)
	}
}
func RecievedAudio(m *tb.Message, b *tb.Bot) {
	libs.LogInfo(fmt.Sprintf("Recieved AudioFile from User %s;\nUserID: %d\nLanguageCode: %s;\nFirstname: %s;\nLastname: %s", m.Sender.Username, m.Sender.ID, m.Sender.LanguageCode, m.Sender.FirstName, m.Sender.LastName))

	Fileextension, FileAllowed := mimetools.GetFileExtensionByMIME(m.Audio.MIME)
	var AudioTitle dbschema.AudioEntitiy
	dbhandler.DB.Where("Title = ? AND Interpret = ?", m.Audio.Title, m.Audio.Performer).Find(&AudioTitle)
	if (dbschema.AudioEntitiy{} == AudioTitle) {
		if FileAllowed == true {
			var FileGuid string = uuid.New().String()
			TelegramFile := &tb.Audio{File: m.Audio.File}
			var AdditionalFileData = tgwrapper.GetAdditionalFileInformation(TelegramFile.FileID)
			filedownloaded, err := tgwrapper.DownloadFile(AdditionalFileData)
			if err != nil {
				libs.LogError(fmt.Sprintf("Fatal: Error %s\n", err.Error()))
			}
			if filedownloaded == true {
				libs.LogInfo("File downloaded successfully...·\nReading Meta Data....\n")
				file, errFileRead := os.Open(path.Join(libs.CFG.Download.TempPath, AdditionalFileData.FileData.FileUniqueID))
				if errFileRead != nil {
					libs.LogError(fmt.Sprintf("Fatal Error: Attemp to read File %s failed\nCause: %s\n", AdditionalFileData.FileData.FileUniqueID, errFileRead.Error()))
					b.Send(m.Sender, fmt.Sprintf("Sorry %s, a Internal Server Error occurred.\nPlease try again later...", m.Sender.FirstName))
				} else {
					defer file.Close()
					fileInfo, _ := file.Stat()
					var size int64 = fileInfo.Size()
					buffer := make([]byte, size)
					_, err := file.Read(buffer)
					if err != nil {
						libs.LogError(fmt.Sprintf("Fatal Error: Attemp to read File %s failed\nCause: %s\n", AdditionalFileData.FileData.FileUniqueID, err.Error()))
						b.Send(m.Sender, fmt.Sprintf("Sorry %s, a Internal Server Error occurred.\nPlease try again later...", m.Sender.FirstName))
					} else {
						fileBytes := bytes.NewReader(buffer)
						fileMetaData, errReadTags := tag.ReadFrom(fileBytes)
						if errReadTags != nil {
							libs.LogError(fmt.Sprintf("Fatal Error: Attemp to read tags from file %s failed\nCause: %s\n", AdditionalFileData.FileData.FileUniqueID, errReadTags.Error()))
							b.Send(m.Sender, fmt.Sprintf("Sorry %s, a Internal Server Error occurred.\nPlease try again later...", m.Sender.FirstName))
						} else {
							result := dbhandler.DB.Create(&dbschema.AudioEntitiy{
								ID:          FileGuid,
								FileEntryID: FileGuid,
								FileEntry: dbschema.FileEntry{
									FileEntryID:       FileGuid,
									GUID:              FileGuid,
									IsAudio:           true,
									IsFile:            false,
									IsMovie:           false,
									IsImage:           false,
									IsOtherFile:       false,
									IsDownloadBlocked: false,
									FileID:            TelegramFile.FileID,
									FilePath:          AdditionalFileData.FileData.FilePath,
									UniqueID:          AdditionalFileData.FileData.FileUniqueID,
								},

								Title:       m.Audio.Title,
								Format:      Fileextension,
								MimeType:    m.Audio.MIME,
								Interpret:   m.Audio.Performer,
								Album:       fileMetaData.Album(),
								Year:        fileMetaData.Year(),
								Description: m.Audio.Caption,
								Genre:       fileMetaData.Genre(),
								Composer:    fileMetaData.Composer(),
								Lyrics:      fileMetaData.Lyrics(),
							})

							if result.Error != nil {

								b.Send(m.Sender, fmt.Sprintf("I am so Sorry %s, Could not add %s to the Database...\nThis issue will be reported!", m.Sender.FirstName, m.Audio.Title))
								libs.LogError(fmt.Sprintf("Fatal Error: Could not send  answer to request from %s\nCause: %s\n", m.Sender.ID, result.Error))
							} else {

								b.Send(m.Sender, fmt.Sprintf("Thank you %s, you Successfully added Title %s to the System\nFileID: %s", m.Sender.FirstName, m.Audio.Title, FileGuid))
							}
						}
					}

				}

			}

		} else {
			_, err := b.Send(m.Sender,
				fmt.Sprintf("Sorry %s, Files of Type %s are not supported or allowed. Try something different!",
					m.Sender.FirstName,
					m.Audio.MIME))
			if err != nil {
				libs.LogError(fmt.Sprintf("Fatal Error: Could not send  answer to request from %s\nCause: %s\n", m.Sender.ID, err.Error()))
			}
		}
	} else {
		b.Send(m.Sender, fmt.Sprintf("Sorry %s, The Title %s from Interpret %s is allready present and has the FileID %s\nPlease upload another Title...", m.Sender.FirstName, AudioTitle.Title, AudioTitle.Interpret, AudioTitle.ID))
		libs.LogWarning(fmt.Sprintf("User %d <%s> tried to add the Title %s from Interpret on Chat %s, but is allready present and has the FileID %s\n", m.Sender.ID, m.Sender.Username, AudioTitle.Title, AudioTitle.Interpret, m.Chat, AudioTitle.ID))
	}
}
