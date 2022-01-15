/*
 * main.go of  TeleBeam from modul TeleBeam
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

package main

import (
	"TeleBeam/audiohandler"
	"TeleBeam/dbhandler"
	"TeleBeam/dbschema"
	"TeleBeam/libs"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
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

	var errMigrate = dbhandler.DB.AutoMigrate(
		&dbschema.AudioEntitiy{},
	)
	if errMigrate != nil {
		libs.LogError(err.Error())
		os.Exit(128)
	}
	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, fmt.Sprintf("Hello %s, i am TeleTransferNG.\nYour Contact Bot to get all the Files you need.", m.Sender.FirstName))
	})
	b.Handle(tb.OnAudio, func(m *tb.Message) {
		audiohandler.RecievedAudio(m, b)
	})
	b.Handle("/downloadaudio", func(m *tb.Message) {
		fmt.Println(m.Payload)
		var AudioFileD dbschema.AudioEntitiy
		dbhandler.DB.Where("GUID=?", m.Payload).Find(&AudioFileD)
		SendFile := &tb.Audio{File: tb.File{FileID: AudioFileD.FileID, FilePath: AudioFileD.FilePath, UniqueID: AudioFileD.UniqueID}}
		fmt.Println(SendFile)
		b.Send(m.Sender, SendFile)
	})

	b.Start()
}
