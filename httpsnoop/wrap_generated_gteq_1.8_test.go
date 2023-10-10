//go:build go1.8
// +build go1.8

// Code generated by "httpsnoop/codegen"; DO NOT EDIT

package httpsnoop

import (
	http "gitee.com/zhaochuninhefei/gmgo/gmhttp"
	"io"
	"testing"
)

func TestWrap(t *testing.T) {
	// combination 1/32
	{
		t.Log("http.ResponseWriter")
		inner := struct {
			http.ResponseWriter
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 2/32
	{
		t.Log("http.ResponseWriter, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 3/32
	{
		t.Log("http.ResponseWriter, io.ReaderFrom")
		inner := struct {
			http.ResponseWriter
			io.ReaderFrom
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 4/32
	{
		t.Log("http.ResponseWriter, io.ReaderFrom, http.Pusher")
		inner := struct {
			http.ResponseWriter
			io.ReaderFrom
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 5/32
	{
		t.Log("http.ResponseWriter, http.Hijacker")
		inner := struct {
			http.ResponseWriter
			http.Hijacker
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 6/32
	{
		t.Log("http.ResponseWriter, http.Hijacker, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.Hijacker
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 7/32
	{
		t.Log("http.ResponseWriter, http.Hijacker, io.ReaderFrom")
		inner := struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 8/32
	{
		t.Log("http.ResponseWriter, http.Hijacker, io.ReaderFrom, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 9/32
	{
		t.Log("http.ResponseWriter, http.CloseNotifier")
		inner := struct {
			http.ResponseWriter
			http.CloseNotifier
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 10/32
	{
		t.Log("http.ResponseWriter, http.CloseNotifier, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 11/32
	{
		t.Log("http.ResponseWriter, http.CloseNotifier, io.ReaderFrom")
		inner := struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 12/32
	{
		t.Log("http.ResponseWriter, http.CloseNotifier, io.ReaderFrom, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 13/32
	{
		t.Log("http.ResponseWriter, http.CloseNotifier, http.Hijacker")
		inner := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 14/32
	{
		t.Log("http.ResponseWriter, http.CloseNotifier, http.Hijacker, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 15/32
	{
		t.Log("http.ResponseWriter, http.CloseNotifier, http.Hijacker, io.ReaderFrom")
		inner := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 16/32
	{
		t.Log("http.ResponseWriter, http.CloseNotifier, http.Hijacker, io.ReaderFrom, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 17/32
	{
		t.Log("http.ResponseWriter, http.Flusher")
		inner := struct {
			http.ResponseWriter
			http.Flusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 18/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 19/32
	{
		t.Log("http.ResponseWriter, http.Flusher, io.ReaderFrom")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 20/32
	{
		t.Log("http.ResponseWriter, http.Flusher, io.ReaderFrom, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 21/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.Hijacker")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 22/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.Hijacker, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 23/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.Hijacker, io.ReaderFrom")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 24/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.Hijacker, io.ReaderFrom, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 25/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.CloseNotifier")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 26/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.CloseNotifier, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 27/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.CloseNotifier, io.ReaderFrom")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 28/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.CloseNotifier, io.ReaderFrom, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 29/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.CloseNotifier, http.Hijacker")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 30/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.CloseNotifier, http.Hijacker, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 31/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.CloseNotifier, http.Hijacker, io.ReaderFrom")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

	// combination 32/32
	{
		t.Log("http.ResponseWriter, http.Flusher, http.CloseNotifier, http.Hijacker, io.ReaderFrom, http.Pusher")
		inner := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{}
		w := Wrap(inner, Hooks{})
		if _, ok := w.(http.ResponseWriter); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Flusher); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Error("unexpected interface")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Error("unexpected interface")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != inner {
				t.Error("w.Unwrap() failed")
			}
		} else {
			t.Error("Unwrapper interface not implemented")
		}
	}

}
