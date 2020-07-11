package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestTicket(t *testing.T) {
	e := httptest.New(t, newApp())
	e.GET("/").Expect().Status(httptest.StatusOK)

	e.GET("/prize").Expect().Status(httptest.StatusOK)
}
