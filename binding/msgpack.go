// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/ugorji/go/codec"

	"net/http"
)

type msgpackBinding struct{}

func (msgpackBinding) Name() string {
	return "msgpack"
}

func (msgpackBinding) Bind(req *http.Request, obj interface{}) error {

	decoder = codec.NewDecoder(req.Body, &codec.MsgpackHandle)
	if err := decoder.Decode(&obj); err != nil {
		return err
	}
	return validate(obj)

}