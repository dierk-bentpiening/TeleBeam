/*
 * dbhandler.go of  TeleBeam from modul TeleBeam
 * Created at 16.1.2022
 * Created from: dpiening
 * Last modified: 16.01.22, 17:23
 * Copyright (C) 2021 - 2022 Dierk-Bent Piening & the TeleBeam Team.
 *
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package dbhandler

import (
	"TeleBeam/libs"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var (
	DB *gorm.DB
)

//TODO: Change Initialization of DB to be Function with just setting a Pointer to DB Variable
func init() {
	var err error
	var cfg = libs.GetConfigValues()
	var DBConfig string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Database.Hostname, cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName, strconv.Itoa(cfg.Database.Port))
	//TODO: Implement bedder Error Catching with custom Error types defined over an Interface.
	DB, err = gorm.Open(postgres.Open(DBConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	//TODO: Work on bedder Error Catching. Fixing Problem instead.
	if err != nil {
		libs.LogError(fmt.Sprintf("Fatal Error: Could not etablish Connection to database: %s\nCause: %s!\nTerminating Application, please check connection to Database Server", cfg.Database.Hostname, err.Error()))
		os.Exit(128)
	} else {
		libs.LogInfo(fmt.Sprintf("Successfully Connected to database server: %s:%s\n", cfg.Database.Hostname, strconv.Itoa(cfg.Database.Port)))
	}

}
