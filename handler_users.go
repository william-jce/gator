package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/william-jce/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("cmd args empty")
	}
	userName := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err == nil {
		return errors.New("user already exists")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}

	fmt.Println("user was created")
	log.Printf("user: %v", user)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("cmd args empty")
	}
	userName := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		fmt.Println("user doesn't exist")
		os.Exit(1)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set to %s\n", userName)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}

	return nil
}
