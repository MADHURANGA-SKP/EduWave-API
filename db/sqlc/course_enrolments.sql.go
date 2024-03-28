// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: course_enrolments.sql

package db

import (
	"context"
	"database/sql"
)

const listEnrolments = `-- name: ListEnrolments :many
SELECT enrolment_id, course_id, request_id, student_id FROM course_enrolments
WHERE student_id = $1 AND course_id = $2
ORDER BY enrolment_id
LIMIT $3
OFFSET $4
`

type ListEnrolmentsParams struct {
	StudentID sql.NullInt64 `json:"student_id"`
	CourseID  sql.NullInt64 `json:"course_id"`
	Limit     int32         `json:"limit"`
	Offset    int32         `json:"offset"`
}

func (q *Queries) ListEnrolments(ctx context.Context, arg ListEnrolmentsParams) ([]CourseEnrolment, error) {
	rows, err := q.db.QueryContext(ctx, listEnrolments,
		arg.StudentID,
		arg.CourseID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CourseEnrolment{}
	for rows.Next() {
		var i CourseEnrolment
		if err := rows.Scan(
			&i.EnrolmentID,
			&i.CourseID,
			&i.RequestID,
			&i.StudentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}