/*
 * handler.go of  TeleBeam from modul TeleBeam
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

package audiohandler

import (
	"TeleBeam/dbhandler"
	"TeleBeam/dbschema"
	"TeleBeam/libs"
	"TeleBeam/mimetools"
	"TeleBeam/tgwrapper"
	"fmt"
	"github.com/google/uuid"
	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
	"time"
)

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
			result := dbhandler.DB.Create(&dbschema.AudioEntitiy{
				Model:       gorm.Model{},
				GUID:        FileGuid,
				Title:       m.Audio.Title,
				Format:      Fileextension,
				MimeType:    m.Audio.MIME,
				UploaderID:  m.Sender.ID,
				UploadDate:  time.Now(),
				Deleted:     false,
				Public:      true,
				Interpret:   m.Audio.Performer,
				Album:       "",
				Year:        "2022",
				Description: m.Audio.Caption,
				Genre:       m.Audio.Title,
				FileID:      TelegramFile.FileID,
				FilePath:    AdditionalFileData.FileData.FilePath,
				UniqueID:    AdditionalFileData.FileData.FileUniqueID,
			})
			if result.Error != nil {

				b.Send(m.Sender, fmt.Sprintf("I am so Sorry %s, Could not add %s to the Database...\nThis issue will be reported!", m.Sender.FirstName, m.Audio.Title))
				libs.LogError(fmt.Sprintf("Fatal Error: Could not send  answer to request from %s\nCause: %s\n", m.Sender.ID, result.Error))
			} else {
				b.Send(m.Sender, fmt.Sprintf("Thank you %s, you Successfully added Title %s to the System\nFileID: %s", m.Sender.FirstName, m.Audio.Title, FileGuid))
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
		b.Send(m.Sender, fmt.Sprintf("Sorry %s, The Title %s from Interpret %s is allready present and has the FileID %s\nPlease upload another Title...", m.Sender.FirstName, AudioTitle.Title, AudioTitle.Interpret, AudioTitle.GUID))
		libs.LogWarning(fmt.Sprintf("User %d <%s> tried to add the Title %s from Interpret on Chat %s, but is allready present and has the FileID %s\n", m.Sender.ID, m.Sender.Username, AudioTitle.Title, AudioTitle.Interpret, m.Chat, AudioTitle.GUID))
	}
}
