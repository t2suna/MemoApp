package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/t2suna/memo_server/model"
)

type MemoService struct {
	db *sql.DB
}

func (s *MemoService) CreateMemo(ctx context.Context, subject, description string) (*model.Memo, error) {
	const (
		insert  = `INSERT INTO memos(subject,description) VALUES(?,?)`
		confirm = `SELECT subject,description,created_at,updated_at FROM memos WHERE id =?`
	)
	result, err := s.db.ExecContext(ctx, insert, subject, description)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	stmt, err := s.db.PrepareContext(ctx, confirm)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var Memo model.Memo
	Memo.ID = id
	err = stmt.QueryRowContext(ctx, id).Scan(&Memo.Subject, &Memo.Description, &Memo.CreatedAt, &Memo.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err

	}
	return &Memo, nil
}

func (s *MemoService) ReadMemo(ctx context.Context, prevID, size int64) ([]*model.Memo, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM Memos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM Memos WHERE id < ? ORDER BY id DESC LIMIT ?`
		readAll    = `SELECT id, subject, description, created_at, updated_at FROM Memos ORDER BY id DESC`
	)
	Memos := []*model.Memo{}

	if prevID == 0 {
		if size == 0 {
			rows, err := s.db.Query(readAll)
			switch {
			case err == sql.ErrNoRows:
				return Memos, err
			case err != nil:
				return Memos, err
			}

			for rows.Next() {
				var Memo model.Memo
				err = rows.Scan(&Memo.ID, &Memo.Subject, &Memo.Description, &Memo.CreatedAt, &Memo.UpdatedAt)
				if err != nil {
					break
				}
				Memos = append(Memos, &Memo)
			}

			return Memos, nil

		}
		stmt, err := s.db.PrepareContext(ctx, read)
		if err != nil {
			return Memos, err
		}
		defer stmt.Close()

		rows, err := stmt.QueryContext(ctx, size)
		switch {
		case err == sql.ErrNoRows:
			return Memos, err
		case err != nil:
			return Memos, err
		}

		for rows.Next() {
			var Memo model.Memo
			err = rows.Scan(&Memo.ID, &Memo.Subject, &Memo.Description, &Memo.CreatedAt, &Memo.UpdatedAt)
			if err != nil {
				break
			}
			Memos = append(Memos, &Memo)
		}

		return Memos, nil
	} else {
		stmt, err := s.db.PrepareContext(ctx, readWithID)
		if err != nil {
			return Memos, err
		}
		defer stmt.Close()

		rows, err := stmt.QueryContext(ctx, prevID, size)
		switch {
		case err == sql.ErrNoRows:
			return Memos, err
		case err != nil:
			return Memos, err
		}

		for rows.Next() {
			var Memo model.Memo
			err = rows.Scan(&Memo.ID, &Memo.Subject, &Memo.Description, &Memo.CreatedAt, &Memo.UpdatedAt)
			if err != nil {
				break
			}
			Memos = append(Memos, &Memo)
		}

		return Memos, nil

	}
}

// UpdateMemo updates the Memo on DB.
func (s *MemoService) UpdateMemo(ctx context.Context, id int64, subject, description string) (*model.Memo, error) {
	const (
		update  = `UPDATE Memos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM Memos WHERE id = ?`
	)

	result, err := s.db.ExecContext(ctx, update, subject, description, id)
	if err != nil {
		return nil, err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return nil, &model.ErrNotFound{}
	}

	stmt, err := s.db.PrepareContext(ctx, confirm)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var Memo model.Memo
	Memo.ID = id
	err = stmt.QueryRowContext(ctx, id).Scan(&Memo.Subject, &Memo.Description, &Memo.CreatedAt, &Memo.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err

	}
	return &Memo, nil
}

// DeleteMemo deletes Memos on DB by ids.
func (s *MemoService) DeleteMemo(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM Memos WHERE id IN (?%s)`

	stmt, err := s.db.PrepareContext(ctx, fmt.Sprintf(deleteFmt, strings.Repeat(`,?`, len(ids)-1)))
	if err != nil {
		return err
	}
	defer stmt.Close()

	idsArr := make([]interface{}, len(ids))
	for i, v := range ids {
		idsArr[i] = v
	}

	result, err := stmt.ExecContext(ctx, idsArr...)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return &model.ErrNotFound{}
	}

	return nil
}
