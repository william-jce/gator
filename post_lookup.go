package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/william-jce/gator/internal/database"
)

const shortIDLen = 8

func shortID(id uuid.UUID) string {
	return id.String()[:shortIDLen]
}

func findFollowedPost(ctx context.Context, s *state, userID uuid.UUID, prefix string) (database.Post, error) {
	prefix = strings.ToLower(strings.TrimSpace(prefix))
	if !isValidIDPrefix(prefix) {
		return database.Post{}, fmt.Errorf("invalid post ID: %q", prefix)
	}

	posts, err := s.db.GetPostByIDPrefixForUser(ctx, database.GetPostByIDPrefixForUserParams{
		UserID:   userID,
		IDPrefix: prefix,
	})
	if err != nil {
		return database.Post{}, fmt.Errorf("looking up posts: %w", err)
	}

	return choosePost(prefix, posts)
}

func findBookmarkedPost(ctx context.Context, s *state, userID uuid.UUID, prefix string) (database.Post, error) {
	prefix = strings.ToLower(strings.TrimSpace(prefix))
	if !isValidIDPrefix(prefix) {
		return database.Post{}, fmt.Errorf("invalid post ID: %q", prefix)
	}

	posts, err := s.db.GetBookmarkedPostByIDPrefix(ctx, database.GetBookmarkedPostByIDPrefixParams{
		UserID:   userID,
		IDPrefix: prefix,
	})
	if err != nil {
		return database.Post{}, fmt.Errorf("looking up bookmark: %w", err)
	}

	return choosePost(prefix, posts)
}

func choosePost(prefix string, posts []database.Post) (database.Post, error) {
	switch len(posts) {
	case 0:
		return database.Post{}, fmt.Errorf("no post found with ID %q", prefix)
	case 1:
		return posts[0], nil
	default:
		var matches strings.Builder
		fmt.Fprintf(&matches, "post ID %q matches multiple posts:", prefix)
		for _, post := range posts {
			fmt.Fprintf(&matches, "\n\t[%s] %s", post.ID.String()[:12], post.Title)
		}
		return database.Post{}, errors.New(matches.String())
	}
}

func isValidIDPrefix(prefix string) bool {
	if prefix == "" {
		return false
	}

	for _, r := range prefix {
		if !strings.ContainsRune("0123456789abcdef-", r) {
			return false
		}
	}

	return true
}
