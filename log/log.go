// Copyright (c) 2022 Yaohui Wang (yaohuiwang@outlook.com)
// go-utils is licensed under Mulan PubL v2.
// You can use this software according to the terms and conditions of the Mulan PubL v2.
// You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PubL v2 for more details.

package log

import (
	"log"
	"fmt"
)

const (
	LOG_DEBUG int = iota
	LOG_INFO
	LOG_WARN
	LOG_ERROR
)

var (
	Logger   *log.Logger = log.Default()
	LogLevel int         = LOG_DEBUG
	LogLabel []string    = []string{"DEBUG", "INFO", "WARN", "ERROR"}
)

func Printf(level int, module string, format string, v ...interface{}) {
        if level < LogLevel {
                return
        }
        Logger.Printf("[%s] %s: %s", LogLabel[level], module, fmt.Sprintf(format, v...))
}
