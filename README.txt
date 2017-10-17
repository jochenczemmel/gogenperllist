package gogenperllist provides perl list functions to go
using go generate

$> go install
generates gogenperllist

develop/
	literal methods for type int for developing and testing

	$> go test

	create file ../template.go:

	$> go build
	$> ./develop

usage/
	examples for using gogenperllist 

	local type: Player
	//go:generate gogenperllist Player

	type: big.Int
	//go:generate gogenperllist -import "math/big" Int

	$> go generate
	$> go build
	$> ./usage

