// Copyright (c) 2014-2020 Canonical Ltd
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License version 3 as
// published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package daemon

import (
	"net/http"
	"time"

	"github.com/canonical/pebble/internal/overlord/state"
	"github.com/canonical/pebble/internal/osutil/sys"
)

type Resp = resp
type ErrorResult = errorResult

func FakeMuxVars(f func(*http.Request) map[string]string) (restore func()) {
	old := muxVars
	muxVars = f
	return func() {
		muxVars = old
	}
}

func FakeStateEnsureBefore(f func(st *state.State, d time.Duration)) (restore func()) {
	old := stateEnsureBefore
	stateEnsureBefore = f
	return func() {
		stateEnsureBefore = old
	}
}

func FakeGetuid(f func() sys.UserID) (restore func()) {
	old := sysGetuid
	sysGetuid = f
	return func() {
		sysGetuid = old
	}
}
