package models

import (
    "errors"
    "time"
    "unicode"
)

type Post struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Author    string    `json:"author"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

var (
    ErrInvalidID       = errors.New("invalid ID")
    ErrInvalidName     = errors.New("invalid name")
    ErrEmptyContent    = errors.New("content should not be empty")
    ErrEmptyAuthor     = errors.New("author should not be empty")
)

func (p *Post) ValidatePost() error {
    if err := p.validName(); err != nil {
        return err
    }
    if err := p.validContent(); err != nil {
        return err
    }
    return nil
}

func (p *Post) validName() error {
    // Check if the name is empty
    if len(p.Author) == 0 {
        return ErrEmptyAuthor
    }

    // Check if the name contains any numbers
    for _, char := range p.Author {
        if unicode.IsDigit(char) {
            return ErrInvalidName
        }
    }

    return nil
}

func (p *Post) validContent() error {
    if len(p.Content) == 0 {
        return ErrEmptyContent
    }
    return nil
}