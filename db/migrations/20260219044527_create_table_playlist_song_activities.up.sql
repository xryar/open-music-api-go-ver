CREATE TABLE playlist_song_activities(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    playlist_id INT NOT NULL,
    song_id INT NOT NULL,
    user_id INT NOT NULL,
    action VARCHAR(255) NOT NULL,
    time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_playlist_activities_id FOREIGN KEY(playlist_id) REFERENCES playlists(id) ON DELETE CASCADE
) ENGINE = InnoDB;