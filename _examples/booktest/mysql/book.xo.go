package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// Book represents a row from 'booktest.books'.
type Book struct {
	BookID      int       `json:"book_id"`     // book_id
	AuthorID    int       `json:"author_id"`   // author_id
	Isbn        string    `json:"isbn"`        // isbn
	BookType    BookType  `json:"book_type"`   // book_type
	Title       string    `json:"title"`       // title
	Year        int       `json:"year"`        // year
	Available   time.Time `json:"available"`   // available
	Description string    `json:"description"` // description
	Tags        string    `json:"tags"`        // tags
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Book exists in the database.
func (b *Book) Exists() bool {
	return b._exists
}

// Deleted returns true when the Book has been marked for deletion from
// the database.
func (b *Book) Deleted() bool {
	return b._deleted
}

// Insert inserts the Book to the database.
func (b *Book) Insert(ctx context.Context, db DB) error {
	switch {
	case b._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case b._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO booktest.books (` +
		`author_id, isbn, book_type, title, year, available, description, tags` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, b.AuthorID, b.Isbn, b.BookType, b.Title, b.Year, b.Available, b.Description, b.Tags)
	res, err := db.ExecContext(ctx, sqlstr, b.AuthorID, b.Isbn, b.BookType, b.Title, b.Year, b.Available, b.Description, b.Tags)
	if err != nil {
		return err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	} // set primary key
	b.BookID = int(id)
	// set exists
	b._exists = true
	return nil
}

// Update updates a Book in the database.
func (b *Book) Update(ctx context.Context, db DB) error {
	switch {
	case !b._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case b._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE booktest.books SET ` +
		`author_id = ?, isbn = ?, book_type = ?, title = ?, year = ?, available = ?, description = ?, tags = ? ` +
		`WHERE book_id = ?`
	// run
	logf(sqlstr, b.AuthorID, b.Isbn, b.BookType, b.Title, b.Year, b.Available, b.Description, b.Tags, b.BookID)
	if _, err := db.ExecContext(ctx, sqlstr, b.AuthorID, b.Isbn, b.BookType, b.Title, b.Year, b.Available, b.Description, b.Tags, b.BookID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Book to the database.
func (b *Book) Save(ctx context.Context, db DB) error {
	if b.Exists() {
		return b.Update(ctx, db)
	}
	return b.Insert(ctx, db)
}

// Upsert performs an upsert for Book.
func (b *Book) Upsert(ctx context.Context, db DB) error {
	switch {
	case b._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO booktest.books (` +
		`book_id, author_id, isbn, book_type, title, year, available, description, tags` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`author_id = VALUES(author_id), isbn = VALUES(isbn), book_type = VALUES(book_type), title = VALUES(title), year = VALUES(year), available = VALUES(available), description = VALUES(description), tags = VALUES(tags)`
	// run
	logf(sqlstr, b.BookID, b.AuthorID, b.Isbn, b.BookType, b.Title, b.Year, b.Available, b.Description, b.Tags)
	if _, err := db.ExecContext(ctx, sqlstr, b.BookID, b.AuthorID, b.Isbn, b.BookType, b.Title, b.Year, b.Available, b.Description, b.Tags); err != nil {
		return err
	}
	// set exists
	b._exists = true
	return nil
}

// Delete deletes the Book from the database.
func (b *Book) Delete(ctx context.Context, db DB) error {
	switch {
	case !b._exists: // doesn't exist
		return nil
	case b._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM booktest.books ` +
		`WHERE book_id = ?`
	// run
	logf(sqlstr, b.BookID)
	if _, err := db.ExecContext(ctx, sqlstr, b.BookID); err != nil {
		return logerror(err)
	}
	// set deleted
	b._deleted = true
	return nil
}

// BooksByAuthorID retrieves a row from 'booktest.books' as a Book.
//
// Generated from index 'author_id'.
func BooksByAuthorID(ctx context.Context, db DB, authorID int) ([]*Book, error) {
	// query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, book_type, title, year, available, description, tags ` +
		`FROM booktest.books ` +
		`WHERE author_id = ?`
	// run
	logf(sqlstr, authorID)
	rows, err := db.QueryContext(ctx, sqlstr, authorID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Book
	for rows.Next() {
		b := Book{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.BookType, &b.Title, &b.Year, &b.Available, &b.Description, &b.Tags); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &b)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// BookByBookID retrieves a row from 'booktest.books' as a Book.
//
// Generated from index 'books_book_id_pkey'.
func BookByBookID(ctx context.Context, db DB, bookID int) (*Book, error) {
	// query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, book_type, title, year, available, description, tags ` +
		`FROM booktest.books ` +
		`WHERE book_id = ?`
	// run
	logf(sqlstr, bookID)
	b := Book{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, bookID).Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.BookType, &b.Title, &b.Year, &b.Available, &b.Description, &b.Tags); err != nil {
		return nil, logerror(err)
	}
	return &b, nil
}

// BooksByTitleYear retrieves a row from 'booktest.books' as a Book.
//
// Generated from index 'books_title_idx'.
func BooksByTitleYear(ctx context.Context, db DB, title string, year int) ([]*Book, error) {
	// query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, book_type, title, year, available, description, tags ` +
		`FROM booktest.books ` +
		`WHERE title = ? AND year = ?`
	// run
	logf(sqlstr, title, year)
	rows, err := db.QueryContext(ctx, sqlstr, title, year)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Book
	for rows.Next() {
		b := Book{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.BookType, &b.Title, &b.Year, &b.Available, &b.Description, &b.Tags); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &b)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// BookByIsbn retrieves a row from 'booktest.books' as a Book.
//
// Generated from index 'isbn'.
func BookByIsbn(ctx context.Context, db DB, isbn string) (*Book, error) {
	// query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, book_type, title, year, available, description, tags ` +
		`FROM booktest.books ` +
		`WHERE isbn = ?`
	// run
	logf(sqlstr, isbn)
	b := Book{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, isbn).Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.BookType, &b.Title, &b.Year, &b.Available, &b.Description, &b.Tags); err != nil {
		return nil, logerror(err)
	}
	return &b, nil
}

// Author returns the Author associated with the Book's (AuthorID).
//
// Generated from foreign key 'books_ibfk_1'.
func (b *Book) Author(ctx context.Context, db DB) (*Author, error) {
	return AuthorByAuthorID(ctx, db, b.AuthorID)
}
