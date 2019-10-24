// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !js !wasm

package http

// RoundTrip implements the RoundTripper interface.
//
// For higher-level HTTP client support (such as handling of cookies
// and redirects), see Get, Post, and the Client type.
//
// Like the RoundTripper interface, the error types returned
// by RoundTrip are unspecified.
//
// RoundTrip 实现 RoundTripper 接口。
//
// 有关更高级别的 HTTP 客户端支持（例如 cookie 和重定向的处理），请参阅 Get，Post 和 Client 类型。
//
// 与 RoundTripper 接口类似，RoundTrip 返回的错误类型未指定。
func (t *Transport) RoundTrip(req *Request) (*Response, error) {
	return t.roundTrip(req)
}
