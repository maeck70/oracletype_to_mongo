package main

type lowhigh_t struct {
	low  int
	high int
}

type num_t struct {
	scale     lowhigh_t
	precision lowhigh_t
}

type xlat_t struct {
	ot   string
	otnp num_t
	mt   string
}
