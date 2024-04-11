package main

type Root struct {
}

func (r *Root) GetName() string {
	return "root"
}

func (r *Root) IsAdmin() bool {
	return true
}

func (r *Root) SetActive(active bool) {
}
