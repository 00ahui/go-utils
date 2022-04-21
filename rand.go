// Copyright (c) 2022 Yaohui Wang (yaohuiwang@outlook.com)
// utils is licensed under Mulan PubL v2.
// You can use this software according to the terms and conditions of the Mulan PubL v2.
// You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PubL v2 for more details.

package utils

import (
	"math/rand"
	"time"
	"fmt"
)

func RandNum(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(max + 1)
}

func RandomString(len int) string {
        rand.Seed(time.Now().UnixNano())
        b := make([]byte, len)
        rand.Read(b)
        return fmt.Sprintf("%x", b)[:len]
}
