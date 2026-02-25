package repositories

import (
	"context"
	"database/sql"
	"open-music-go/model/domain"
)

type PlaylistCollabRepositoryImpl struct {
}

func NewPlaylistCollabRepositoryImpl() *PlaylistCollabRepositoryImpl {
	return &PlaylistCollabRepositoryImpl{}
}

func (pcr *PlaylistCollabRepositoryImpl) Add(ctx context.Context, tx *sql.Tx, collab domain.PlaylistCollab) error {
	SQL := "INSERT INTO playlist_collaborators(playlist_id, user_id) VALUES (?, ?)"
	_, err := tx.ExecContext(ctx, SQL, collab.PlaylistId, collab.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (pcr *PlaylistCollabRepositoryImpl) Remove(ctx context.Context, tx *sql.Tx, collab domain.PlaylistCollab) error {
	SQL := "DELETE FROM playlist_collaborators WHERE playlist_id = ? AND user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, collab.PlaylistId, collab.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (pcr *PlaylistCollabRepositoryImpl) IsCollab(ctx context.Context, tx *sql.Tx, collab domain.PlaylistCollab) (bool, error) {
	SQL := "SELECT 1 FROM playlist_collaborators WHERE playlist_id = ? AND user_id = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, SQL, collab.PlaylistId, collab.UserId)

	var dummy int
	err := row.Scan(&dummy)
	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *PlaylistCollabRepositoryImpl) FindByPlaylistId(ctx context.Context, tx *sql.Tx, playlistId int) ([]int, error) {
	SQL := "SELECT user_id FROM playlist_collaborators WHERE playlist_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, playlistId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []int{}
	for rows.Next() {

		var userId int

		err := rows.Scan(&userId)
		if err != nil {
			return nil, err
		}

		users = append(users, userId)
	}

	return users, nil
}
