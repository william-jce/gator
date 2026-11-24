package main

import (
	"context"
	"fmt"
	"time"

	"github.com/william-jce/gator/internal/database"
)

func handlerBookmark(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <post ID>", cmd.Name)
	}

	ctx := context.Background()
	idPrefix := cmd.Args[0]

	post, err := findFollowedPost(ctx, s, user.ID, idPrefix)
	if err != nil {
		return err
	}

	_, err = s.db.CreateBookmark(ctx, database.CreateBookmarkParams{
		UserID:    user.ID,
		PostID:    post.ID,
		CreatedAt: time.Now(),
	})
	if isUniqueViolation(err) {
		return fmt.Errorf("already bookmarked: [%s] %s", shortID(post.ID), post.Title)
	}
	if err != nil {
		return fmt.Errorf("creating bookmark: %w", err)
	}

	fmt.Printf("Bookmarked: [%s] %s\n", shortID(post.ID), post.Title)
	return nil
}

func handlerUnbookmark(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <post ID>", cmd.Name)
	}

	ctx := context.Background()
	idPrefix := cmd.Args[0]

	post, err := findBookmarkedPost(ctx, s, user.ID, idPrefix)
	if err != nil {
		return err
	}

	rows, err := s.db.DeleteBookmark(ctx, database.DeleteBookmarkParams{
		UserID: user.ID,
		PostID: post.ID,
	})
	if err != nil {
		return fmt.Errorf("deleting bookmark: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("not bookmarked: [%s] %s", shortID(post.ID), post.Title)
	}

	fmt.Printf("Removed bookmark: [%s] %s\n", shortID(post.ID), post.Title)

	return nil
}

func handlerBookmarks(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	bookmarks, err := s.db.GetBookmarksForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("getting bookmarks: %w", err)
	}

	if len(bookmarks) == 0 {
		fmt.Println("No bookmarks found.")
		return nil
	}
	fmt.Printf("%s's Bookmarks:\n", user.Name)
	for _, bookmark := range bookmarks {
		fmt.Printf("[%s] %s: %s\n", shortID(bookmark.ID), bookmark.Title, bookmark.Url)
	}

	return nil
}
