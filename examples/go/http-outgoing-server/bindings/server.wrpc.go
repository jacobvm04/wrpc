// Generated by `wit-bindgen-wrpc-go` 0.1.0. DO NOT EDIT!
// server package contains wRPC bindings for `server` world
package server

import (
	exports__wrpc__http__outgoing_handler "github.com/wrpc/wrpc/examples/go/http-outgoing-server/bindings/exports/wrpc/http/outgoing_handler"
	wrpc "github.com/wrpc/wrpc/go"
)

func Serve(c wrpc.Client, h interface {
	exports__wrpc__http__outgoing_handler.Handler
}) (stop func() error, err error) {
	stops := make([]func() error, 0, 1)
	stop = func() error {
		for _, stop := range stops {
			if err := stop(); err != nil {
				return err
			}
		}
		return nil
	}
	stop0, err := exports__wrpc__http__outgoing_handler.ServeInterface(c, h)
	if err != nil {
		return
	}
	stops = append(stops, stop0)
	stop = func() error {
		if err := stop0(); err != nil {
			return err
		}
		return nil
	}
	return
}
