package main

type predefined map[identifier]float64

type expr interface {
	eval(p predefined) float64
}

type additive struct {
	op string // + -
	x, y expr
}

type multiplicative struct {
	op string // * / %
	x, y expr
}

type unary struct {
	op string // - +
	x expr
}

type postfix struct {
	op string // ()
	x, y expr
}

type argument struct {
	x, y expr
}

type primary struct {
	x expr
}

type identifier string

type constant float64
