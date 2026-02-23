package repositories

import (
	"context"
	"database/sql"
	"open-music-go/model/domain"
)

type PlaylistActivityRepositoryImpl struct {
}

func NewPlaylistActivityRepository() *PlaylistActivityRepositoryImpl {
	return &PlaylistActivityRepositoryImpl{}
}

func (ar *PlaylistActivityRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, activity domain.PlaylistActivity) error {
	SQL := "INSERT INTO playlist_song_activities (playlist_id, song_id, user_id, action) VALUES (?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, activity.PlaylistId, activity.SongId, activity.UserId, activity.Action)
	if err != nil {
		return err
	}

	return nil
}

func (ar *PlaylistActivityRepositoryImpl) FindPlaylistById(ctx context.Context, tx *sql.Tx, playlistId int) ([]domain.PlaylistActivityJoin, error) {
	SQL := `
	SELECT
		u.username,
		s.title,
		a.action,
		a.time
	FROM playlist_song_activities a
	JOIN users u ON u.id = a.user_id
	JOIN song s ON s.id = a.song_id
	WHERE a.playlist_id = ?
	ORDER BY a.time DESC
	`

	rows, err := tx.QueryContext(ctx, SQL, playlistId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	activities := []domain.PlaylistActivityJoin{}

	for rows.Next() {
		activity := domain.PlaylistActivityJoin{}
		err := rows.Scan(&activity.Username, &activity.SongTitle, &activity.Action, &activity.Time)
		if err != nil {
			return nil, err
		}

		activities = append(activities, activity)
	}

	return activities, nil
}
