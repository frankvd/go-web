Go web package(WIP)
===============

Example
--------
	package main

	import (
		"vrenc/web"
	)

	func main() {
		app := web.CreateServer()
		router := app.Router()

		router.Get("/abc", func(res *web.Response, req *web.Request) {
			res.Send("abcdefghijklmnopqrstuvwxyz")
		})

		app.Run(8080)
	}