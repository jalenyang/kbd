package main

type KbdAction interface {
	Install() error
	Uninstall() error
	Upgrade() error
	Rollback() error
}

type K8SClient interface {
	List() error
}
