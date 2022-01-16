/*
 * main.go of  TeleBeam from modul TeleBeam
 * Created at 16.1.2022
 * Created from: dpiening
 * Last modified: 16.01.22, 22:23
 * Copyright (C) 2021 - 2022 Dierk-Bent Piening & the TeleBeam Team.
 *
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"TeleBeam/dbhandler"
	"TeleBeam/dbschema"
	"TeleBeam/filehandler"
	"TeleBeam/libs"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func main() {
	fmt.Println(libs.ApplicationCopyright)
	b, err := tb.NewBot(tb.Settings{
		Token:  libs.CFG.Telegram.APIToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	b.Handle("/start", func(m *tb.Message) {
		//TODO: Put LogInfo Logic for Recieved Command inside of an own function.
		go libs.LogInfo(fmt.Sprintf("Recieved start Command from User %s\nUserID: %d\nFirstname: %s\nLastname: %s\n", m.Sender.Username, m.Sender.ID, m.Sender.FirstName, m.Sender.LastName))
		b.Send(m.Sender, fmt.Sprintf("Hello %s, i am TeleBeam.\nYour Contact Bot to get all the Files you need.", m.Sender.FirstName))
	})
	b.Handle(tb.OnAudio, func(m *tb.Message) {
		go filehandler.RecievedAudio(m, b)
	})
	b.Handle("/download", func(m *tb.Message) {
		var FileEntry dbschema.FileEntry
		dbhandler.DB.Where("GUID = ?", m.Payload).Find(&FileEntry)
		SendFile := &tb.Audio{File: tb.File{FileID: FileEntry.FileID, FilePath: FileEntry.FilePath, UniqueID: FileEntry.UniqueID}}
		go b.Send(m.Sender, SendFile)
	})
	b.Handle("/searchbyinterpret", func(m *tb.Message) {
		go libs.LogInfo(fmt.Sprintf("Recieved /searchbyinterpret Command from User %s\nUserID: %d\nFirstname: %s\nLastname: %s\n", m.Sender.Username, m.Sender.ID, m.Sender.FirstName, m.Sender.LastName))

		go filehandler.GetTitleByInterpreter(m, b)
	})
	b.Handle("/searchbygenre", func(m *tb.Message) {
		go libs.LogInfo(fmt.Sprintf("Recieved /searchbygenre Command from User %s\nUserID: %d\nFirstname: %s\nLastname: %s\n", m.Sender.Username, m.Sender.ID, m.Sender.FirstName, m.Sender.LastName))

		go filehandler.GetByGenre(m, b)
	})
	b.Handle("/searchbytitle", func(m *tb.Message) {
		go libs.LogInfo(fmt.Sprintf("Recieved /searchbytitle Command from User %s\nUserID: %d\nFirstname: %s\nLastname: %s\n", m.Sender.Username, m.Sender.ID, m.Sender.FirstName, m.Sender.LastName))

		go filehandler.GetByTitle(m, b)
	})
	b.Handle("/searchbyyear", func(m *tb.Message) {
		go libs.LogInfo(fmt.Sprintf("Recieved /searchbyyear Command from User %s\nUserID: %d\nFirstname: %s\nLastname: %s\n", m.Sender.Username, m.Sender.ID, m.Sender.FirstName, m.Sender.LastName))

		go filehandler.GetByYear(m, b)
	})
	b.Handle("/searchbyalbum", func(m *tb.Message) {
		go libs.LogInfo(fmt.Sprintf("Recieved /searchbyalbum Command from User %s\nUserID: %d\nFirstname: %s\nLastname: %s\n", m.Sender.Username, m.Sender.ID, m.Sender.FirstName, m.Sender.LastName))

		go filehandler.GetByAlbum(m, b)
	})

	b.Start()
}
