// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"github.com/ugorji/go/codec"
	"net/http"
)

type MsgPack struct {
	Data interface{}
}

var msgpackContentType = []string{"application/msgpack; charset=utf-8"}

func (r MsgPack) Render(w http.ResponseWriter) error {
	writeContentType(w, msgpackContentType)
	var h codec.Handle = new(codec.MsgpackHandle)
	return codec.NewEncoder(w, h).Encode(r.Data)
}
