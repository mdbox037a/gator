package main

import ()

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

}
