package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// DjangoAdminLog represents a row from 'django.django_admin_log'.
type DjangoAdminLog struct {
	ID            int            `json:"id"`              // id
	ActionTime    time.Time      `json:"action_time"`     // action_time
	ObjectID      sql.NullString `json:"object_id"`       // object_id
	ObjectRepr    string         `json:"object_repr"`     // object_repr
	ActionFlag    Smallint5      `json:"action_flag"`     // action_flag
	ChangeMessage string         `json:"change_message"`  // change_message
	ContentTypeID sql.NullInt64  `json:"content_type_id"` // content_type_id
	UserID        int            `json:"user_id"`         // user_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the DjangoAdminLog exists in the database.
func (dal *DjangoAdminLog) Exists() bool {
	return dal._exists
}

// Deleted returns true when the DjangoAdminLog has been marked for deletion from
// the database.
func (dal *DjangoAdminLog) Deleted() bool {
	return dal._deleted
}

// Insert inserts the DjangoAdminLog to the database.
func (dal *DjangoAdminLog) Insert(ctx context.Context, db DB) error {
	switch {
	case dal._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dal._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.django_admin_log (` +
		`action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID)
	res, err := db.ExecContext(ctx, sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID)
	if err != nil {
		return err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	} // set primary key
	dal.ID = int(id)
	// set exists
	dal._exists = true
	return nil
}

// Update updates a DjangoAdminLog in the database.
func (dal *DjangoAdminLog) Update(ctx context.Context, db DB) error {
	switch {
	case !dal._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dal._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.django_admin_log SET ` +
		`action_time = ?, object_id = ?, object_repr = ?, action_flag = ?, change_message = ?, content_type_id = ?, user_id = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, dal.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, dal.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjangoAdminLog to the database.
func (dal *DjangoAdminLog) Save(ctx context.Context, db DB) error {
	if dal.Exists() {
		return dal.Update(ctx, db)
	}
	return dal.Insert(ctx, db)
}

// Upsert performs an upsert for DjangoAdminLog.
func (dal *DjangoAdminLog) Upsert(ctx context.Context, db DB) error {
	switch {
	case dal._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO django.django_admin_log (` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`action_time = VALUES(action_time), object_id = VALUES(object_id), object_repr = VALUES(object_repr), action_flag = VALUES(action_flag), change_message = VALUES(change_message), content_type_id = VALUES(content_type_id), user_id = VALUES(user_id)`
	// run
	logf(sqlstr, dal.ID, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID)
	if _, err := db.ExecContext(ctx, sqlstr, dal.ID, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID); err != nil {
		return err
	}
	// set exists
	dal._exists = true
	return nil
}

// Delete deletes the DjangoAdminLog from the database.
func (dal *DjangoAdminLog) Delete(ctx context.Context, db DB) error {
	switch {
	case !dal._exists: // doesn't exist
		return nil
	case dal._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.django_admin_log ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, dal.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dal.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dal._deleted = true
	return nil
}

// DjangoAdminLogByContentTypeID retrieves a row from 'django.django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'django_admin_log_content_type_id_c4bce8eb_fk_django_co'.
func DjangoAdminLogByContentTypeID(ctx context.Context, db DB, contentTypeID sql.NullInt64) ([]*DjangoAdminLog, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id ` +
		`FROM django.django_admin_log ` +
		`WHERE content_type_id = ?`
	// run
	logf(sqlstr, contentTypeID)
	rows, err := db.QueryContext(ctx, sqlstr, contentTypeID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjangoAdminLog
	for rows.Next() {
		dal := DjangoAdminLog{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dal.ID, &dal.ActionTime, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjangoAdminLogByID retrieves a row from 'django.django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'django_admin_log_id_pkey'.
func DjangoAdminLogByID(ctx context.Context, db DB, id int) (*DjangoAdminLog, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id ` +
		`FROM django.django_admin_log ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	dal := DjangoAdminLog{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dal.ID, &dal.ActionTime, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID); err != nil {
		return nil, logerror(err)
	}
	return &dal, nil
}

// DjangoAdminLogByUserID retrieves a row from 'django.django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'django_admin_log_user_id_c564eba6_fk_auth_user_id'.
func DjangoAdminLogByUserID(ctx context.Context, db DB, userID int) ([]*DjangoAdminLog, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id ` +
		`FROM django.django_admin_log ` +
		`WHERE user_id = ?`
	// run
	logf(sqlstr, userID)
	rows, err := db.QueryContext(ctx, sqlstr, userID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjangoAdminLog
	for rows.Next() {
		dal := DjangoAdminLog{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dal.ID, &dal.ActionTime, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjangoContentType returns the DjangoContentType associated with the DjangoAdminLog's (ContentTypeID).
//
// Generated from foreign key 'django_admin_log_content_type_id_c4bce8eb_fk_django_co'.
func (dal *DjangoAdminLog) DjangoContentType(ctx context.Context, db DB) (*DjangoContentType, error) {
	return DjangoContentTypeByID(ctx, db, int(dal.ContentTypeID.Int64))
}

// AuthUser returns the AuthUser associated with the DjangoAdminLog's (UserID).
//
// Generated from foreign key 'django_admin_log_user_id_c564eba6_fk_auth_user_id'.
func (dal *DjangoAdminLog) AuthUser(ctx context.Context, db DB) (*AuthUser, error) {
	return AuthUserByID(ctx, db, dal.UserID)
}
