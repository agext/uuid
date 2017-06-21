// Copyright 2015 ALRUX Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package uuid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var (
	zero       = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	uuid       = []byte{0xf2, 0x54, 0xdf, 0x4a, 0x18, 0x4c, 0x10, 0x19, 0x80, 0xa4, 0xc6, 0x1c, 0xd0, 0x0a, 0x68, 0x99}
	uuidString = "f254df4a-184c-1019-80a4-c61cd00a6899"
)

func TestNodeId(t *testing.T) {
	nodeId := uint32(rand.Int31n(0x40000000))
	err := SetNodeId(nodeId)
	if err != nil {
		t.Error("SetNodeId:", err)
	}
	act := NodeId()
	if act != nodeId {
		t.Errorf("NodeId: expecting % x, got % x", nodeId, act)
	}
	uuid := New()
	act = uuid.NodeId()
	if act != nodeId {
		t.Errorf("New().NodeId: expecting % x, got % x", nodeId, act)
	}
	for i := 0; i < randBufCap; i++ {
		uuid = NewCrypto()
	}
	act = uuid.NodeId()
	if act != nodeId {
		t.Errorf("NewCrypto().NodeId: expecting % x, got % x", nodeId, act)
	}
	err = SetNodeId(nodeId | 0x40000000)
	if err == nil {
		t.Error("SetNodeId(30-bit overflow): expecting error, got nil")
	}
	act = NodeId()
	if act != nodeId {
		t.Errorf("NodeId after SetNodeId(30-bit overflow): expecting % x, got % x", nodeId, act)
	}
}

func TestNewFromBytes(t *testing.T) {
	_, err := NewFromBytes(zero)
	if err != nil {
		t.Error("TestNewFromBytes:", err)
	}
	_, err = NewFromBytes(zero[1:])
	if err == nil {
		t.Error("TestNewFromBytes(too short): expecting error, got nil")
	}
	_, err = NewFromBytes(append([]byte{0x00}, zero...))
	if err == nil {
		t.Error("TestNewFromBytes(too long): expecting error, got nil")
	}
}

func TestNewFromString(t *testing.T) {
	uuid1, err := NewFromString(uuidString)
	if err != nil {
		t.Error("TestNewFromString:", err)
	}

	if bytes.Compare(uuid, []byte(uuid1)) != 0 {
		t.Errorf("TestNewFromString: Expecting % x, got % x", uuid, uuid1)
	}

	uuid2, err := NewFromString(strings.Replace(uuidString, "-", "", -1))
	if err != nil {
		t.Error("TestNewFromString:", err)
	}

	if uuid2.String() != uuid1.String() {
		t.Error("TestNewFromString: Stripping dashes should not affect string parsing", uuid1, uuid2)
	}

	_, err = NewFromString("0000")
	if err == nil {
		t.Error("TestNewFromString: Should fail on short UUID")
	}

	_, err = NewFromString("0000------------------------0000")
	if err == nil {
		t.Error("TestNewFromString: Should fail on short UUID, ignoring dashes")
	}

	_, err = NewFromString("00000000000000000000000000000000000000000")
	if err == nil {
		t.Error("TestNewFromString: Should fail on long UUID")
	}

	_, err = NewFromString("f254df4a-184c-1z19-80a4-c61cd00a6899")
	if err == nil {
		t.Error("TestNewFromString: Should fail on invalid hex digit(s)")
	}

	_, err = NewFromString("-0--000-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0-0--0--")
	if err != nil {
		t.Error("TestNewFromString: Should ignore extra dashes")
	}
}

func TestString(t *testing.T) {
	uuid1, err := NewFromBytes(uuid)

	if err != nil {
		t.Error("TestString:", err)
	}

	if uuid1.String() != uuidString {
		t.Errorf("TestString: Expecting %s, got %s", uuidString, uuid1.String())
	}
}

func TestUnixNano(t *testing.T) {
	now := int64(time.Now().UTC().UnixNano())
	act := toUnixNano(fromUnixNano(now))
	now = (now / 100) * 100
	if act != now {
		t.Errorf("TestUnixNano: Expecting % x, got % x", now, act)
	}
}

func TestTime(t *testing.T) {
	now := time.Now()
	timeNow = func() time.Time {
		return now
	}
	uuid1 := New()
	ts := toUnixNano(fromUnixNano(int64(now.UTC().UnixNano())))

	if act := uuid1.Time(); act.UnixNano() != ts {
		t.Errorf("TestTime: Expecting %d, got %d", ts, act.UnixNano())
	}
}

func TestVersion(t *testing.T) {
	uuid1, err := NewFromBytes(uuid)
	if err != nil {
		t.Error("TestVersion:", err)
	}

	if uuid1.Version() != 1 {
		t.Errorf("TestVersion: Expecting %d, got %d", 1, uuid1.Version())
	}
}

func TestVariant(t *testing.T) {
	uuid1, err := NewFromBytes(uuid)
	if err != nil {
		t.Error("TestVariant:", err)
	}

	if uuid1.Variant() != 4 {
		t.Errorf("TestVariant: Expecting %d, got %d", 4, uuid1.Variant())
	}
}

func TestNew(t *testing.T) {
	uuid1 := New()

	if uuid1.Version() != 1 {
		t.Errorf("TestNew: Expecting version %d, got %d", 1, uuid1.Version())
	}

	if uuid1.Variant() != 1 {
		t.Errorf("TestNew: Expecting variant %d, got %d", 1, uuid1.Variant())
	}
}

func TestUnmarshalJSON(t *testing.T) {
	s := fmt.Sprintf(`{"uuid":"%s"}`, uuidString)
	d := new(struct{ Uuid UUID })

	if err := json.Unmarshal([]byte(s), d); err != nil {
		t.Error("TestUnmarshalJSON:", err)
	}

	got := d.Uuid.String()

	if got != uuidString {
		t.Errorf("TestUnmarshalJSON: Expecting %s, got %s", uuidString, got)
	}

	if err := json.Unmarshal([]byte(s[:len(s)-2]), d); err == nil {
		t.Error("TestUnmarshalJSON: Should fail on invalid JSON")
	}

	if err := json.Unmarshal([]byte(`{"uuid":"f254df4a-184c-1z19-80a4-c61cd00a6899"}`), d); err == nil {
		t.Error("TestUnmarshalJSON: Should fail on invalid UUID")
	}
}

func TestMarshalJSON(t *testing.T) {
	uuid1, err := NewFromString(uuidString)
	if err != nil {
		t.Error(err)
	}

	b, err := json.Marshal(struct{ Uuid UUID }{uuid1})
	if err != nil {
		t.Error(err)
	}

	got := string(b)
	want := fmt.Sprintf(`{"Uuid":"%s"}`, uuidString)

	if got != want {
		t.Errorf("TestMarshalJSON: Expecting %s, got %s", want, got)
	}
}
