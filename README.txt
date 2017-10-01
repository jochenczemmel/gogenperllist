package gogenperllist provides perl list functions to go
using go generate

$> go install
generates gogenperllist

develop/
	literal methods for type int for developing and testing

	$> go test

usage/
	example for using gogenperllist (type: Player)
	//go:generate gogenperllist Player

	$> go generate
	$> go build
	$> ./usage

