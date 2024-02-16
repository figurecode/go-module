module main

go 1.21

replace (
	workspace/user/repo/somepackage => ../somemodule/somepackage
	ypmodule => ../ypmodule
)

require (
	workspace/user/repo/somepackage v0.0.0-00010101000000-000000000000
	ypmodule v0.0.0-00010101000000-000000000000
)
