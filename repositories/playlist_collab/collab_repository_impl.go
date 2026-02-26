package repositories

import (
	"context"
	"database/sql"
	"open-music-go/model/domain"
)

type PlaylistCollabRepositoryImpl struct {
}

func NewPlaylistCollabRepository() *PlaylistCollabRepositoryImpl {
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

func (r *PlaylistCollabRepositoryImpl) GetCollaborators(ctx context.Context, tx *sql.Tx, playlistId int) ([]domain.User, error) {
	SQL := `
	SELECT u.id, u.username
	JOIN users u ON pc.user_id = u.id
	FROM playlist_collaborators pc 
	WHERE pc.playlist_id = ?
	`
	rows, err := tx.QueryContext(ctx, SQL, playlistId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, user.Username)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
