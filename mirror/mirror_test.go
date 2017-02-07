//: ----------------------------------------------------------------------------
//: Copyright (C) 2017 Verizon.  All Rights Reserved.
//: All Rights Reserved
//:
//: file:    mirror_test.go
//: details: TODO
//: author:  Mehrdad Arshad Rad
//: date:    02/01/2017
//:
//: Licensed under the Apache License, Version 2.0 (the "License");
//: you may not use this file except in compliance with the License.
//: You may obtain a copy of the License at
//:
//:     http://www.apache.org/licenses/LICENSE-2.0
//:
//: Unless required by applicable law or agreed to in writing, software
//: distributed under the License is distributed on an "AS IS" BASIS,
//: WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//: See the License for the specific language governing permissions and
//: limitations under the License.
//: ----------------------------------------------------------------------------
package mirror

import (
	"net"
	"testing"

	"golang.org/x/net/ipv4"
)

func TestNewRawConn(t *testing.T) {
	ip := net.ParseIP("127.0.0.1")
	c, err := NewRawConn(ip)
	if err != nil {
		t.Error("unexpected error", err)
	}
	if c.family != 2 {
		t.Error("expected family# 2, got", c.family)
	}
}

func TestIPv4Header(t *testing.T) {
	ipv4RawHeader := NewIPv4HeaderTpl(17)
	b := ipv4RawHeader.Marshal()
	h, err := ipv4.ParseHeader(b)
	if err != nil {
		t.Error("unexpected error", err)
	}
	if h.Version != 4 {
		t.Error("expect version: 4, got", h.Version)
	}
	if h.Protocol != 17 {
		t.Error("expect protocol: 17, got", h.Protocol)
	}
	if h.TTL != 64 {
		t.Error("expect TTL: 64, got", h.TTL)
	}
	if h.Len != 20 {
		t.Error("expect Len: 20, got", h.Len)
	}
	if h.Checksum != 0 {
		t.Error("expect Checksum: 0, got", h.Checksum)
	}
}