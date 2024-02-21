module main

go 1.21

replace (
	arritn => ../arritn
	workspace/user/repo/somepackage => ../somemodule/somepackage
	ypmodule => ../ypmodule
)

require (
	arritn v0.0.0-00010101000000-000000000000
	github.com/figurecode/mymath v0.0.0-20240220131431-a649138275f8
	workspace/user/repo/somepackage v0.0.0-00010101000000-000000000000
	ypmodule v0.0.0-00010101000000-000000000000
)
