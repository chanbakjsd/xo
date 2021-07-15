package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// BooktestTag represents a row from 'django.booktest_tag'.
type BooktestTag struct {
	TagID int64  `json:"tag_id"` // tag_id
	Tag   string `json:"tag"`    // tag
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the BooktestTag exists in the database.
func (bt *BooktestTag) Exists() bool {
	return bt._exists
}

// Deleted returns true when the BooktestTag has been marked for deletion from
// the database.
func (bt *BooktestTag) Deleted() bool {
	return bt._deleted
}

// Insert inserts the BooktestTag to the database.
func (bt *BooktestTag) Insert(ctx context.Context, db DB) error {
	switch {
	case bt._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case bt._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.booktest_tag (` +
		`tag` +
		`) VALUES (` +
		`?` +
		`)`
	// run
	logf(sqlstr, bt.Tag)
	res, err := db.ExecContext(ctx, sqlstr, bt.Tag)
	if err != nil {
		return err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	} // set primary key
	bt.TagID = int64(id)
	// set exists
	bt._exists = true
	return nil
}

// Update updates a BooktestTag in the database.
func (bt *BooktestTag) Update(ctx context.Context, db DB) error {
	switch {
	case !bt._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case bt._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.booktest_tag SET ` +
		`tag = ? ` +
		`WHERE tag_id = ?`
	// run
	logf(sqlstr, bt.Tag, bt.TagID)
	if _, err := db.ExecContext(ctx, sqlstr, bt.Tag, bt.TagID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the BooktestTag to the database.
func (bt *BooktestTag) Save(ctx context.Context, db DB) error {
	if bt.Exists() {
		return bt.Update(ctx, db)
	}
	return bt.Insert(ctx, db)
}

// Upsert performs an upsert for BooktestTag.
func (bt *BooktestTag) Upsert(ctx context.Context, db DB) error {
	switch {
	case bt._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO django.booktest_tag (` +
		`tag_id, tag` +
		`) VALUES (` +
		`?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`tag = VALUES(tag)`
	// run
	logf(sqlstr, bt.TagID, bt.Tag)
	if _, err := db.ExecContext(ctx, sqlstr, bt.TagID, bt.Tag); err != nil {
		return err
	}
	// set exists
	bt._exists = true
	return nil
}

// Delete deletes the BooktestTag from the database.
func (bt *BooktestTag) Delete(ctx context.Context, db DB) error {
	switch {
	case !bt._exists: // doesn't exist
		return nil
	case bt._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.booktest_tag ` +
		`WHERE tag_id = ?`
	// run
	logf(sqlstr, bt.TagID)
	if _, err := db.ExecContext(ctx, sqlstr, bt.TagID); err != nil {
		return logerror(err)
	}
	// set deleted
	bt._deleted = true
	return nil
}

// BooktestTagByTagID retrieves a row from 'django.booktest_tag' as a BooktestTag.
//
// Generated from index 'booktest_tag_tag_id_pkey'.
func BooktestTagByTagID(ctx context.Context, db DB, tagID int64) (*BooktestTag, error) {
	// query
	const sqlstr = `SELECT ` +
		`tag_id, tag ` +
		`FROM django.booktest_tag ` +
		`WHERE tag_id = ?`
	// run
	logf(sqlstr, tagID)
	bt := BooktestTag{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, tagID).Scan(&bt.TagID, &bt.Tag); err != nil {
		return nil, logerror(err)
	}
	return &bt, nil
}
