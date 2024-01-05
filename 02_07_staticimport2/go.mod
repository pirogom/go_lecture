module github.com/pirogom/go_lecture/02_07_staticimport2

go 1.21.5

require (
	firstpkg v0.0.0-00010101000000-000000000000
	secondpkg v0.0.0-00010101000000-000000000000
)

replace (
	firstpkg => ./firstpkg
	secondpkg => ./secondpkg
)
