// Copyright (c) 2022 Yaohui Wang (yaohuiwang@outlook.com)
// utils is licensed under Mulan PubL v2.
// You can use this software according to the terms and conditions of the Mulan PubL v2.
// You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PubL v2 for more details.

package net

import (
	"bytes"
	"fmt"
	"net"
)

// cidr contai
func IpInNetwork(ipstr string, cidr string) bool {
	ip := net.ParseIP(ipstr)
	if _, ipnet, err := net.ParseCIDR(cidr); err != nil {
		return false
	} else {
		return ipnet.Contains(ip)
	}
}

// allocate ip from cidr
func AllocIPFromCidr(cidr string, exist []string) (*string, error) {
	cidrIp, cidrNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	for ip := cidrIp.Mask(cidrNet.Mask); cidrNet.Contains(ip); inc(ip) {
		ipstr := ip.String()
		if !ListContains(exist, ipstr) {
			return &ipstr, nil
		}
	}
	return nil, fmt.Errorf("cidr is full")
}

// allocate ip from range
func AllocIPFromRange(start string, end string, exist []string) (*string, error) {
	startIp := net.ParseIP(start)
	endIp := net.ParseIP(end)

	for ip := startIp; bytes.Compare(ip, endIp) < 0; inc(ip) {
		ipstr := ip.String()
		if !ListContains(exist, ipstr) {
			return &ipstr, nil
		}
	}
	return nil, fmt.Errorf("ip range is full")
}

// inc ip
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
