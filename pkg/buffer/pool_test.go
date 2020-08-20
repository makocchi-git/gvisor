// Copyright 2020 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package buffer

import (
	"testing"
)

func TestGetDefaultBufferSize(t *testing.T) {
	var p pool
	for i := 0; i < embeddedCount*2; i++ {
		buf := p.Get()
		if got, want := len(buf.data), defaultBufferSize; got != want {
			t.Errorf("#%d len(buf.data) = %d, want %d", i, got, want)
		}
	}
}

func TestGetCustomBufferSize(t *testing.T) {
	const size = 100

	var p pool
	p.setBufferSize(size)
	for i := 0; i < embeddedCount*2; i++ {
		buf := p.Get()
		if got, want := len(buf.data), size; got != want {
			t.Errorf("#%d len(buf.data) = %d, want %d", i, got, want)
		}
	}
}

func TestPut(t *testing.T) {
	var p pool
	buf := p.Get()
	p.Put(buf)
	if buf.data != nil {
		t.Errorf("buf.data = %x, want nil", buf.data)
	}
}
