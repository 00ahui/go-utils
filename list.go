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

func ListContains(slice []string, obj string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == obj {
			return true
		}
	}
	return false
}

func ListDelete(slice []string, obj string) (result []string) {
	for i := 0; i < len(slice); i++ {
		if slice[i] == obj {
			continue
		}
		result = append(result, slice[i])
	}
	return
}

