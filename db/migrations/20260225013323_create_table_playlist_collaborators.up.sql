CREATE TABLE playlist_collaborators(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    playlist_id INT NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT fk_collab_playlist_id FOREIGN KEY(playlist_id) REFERENCES playlists(id) ON DELETE CASCADE,
    CONSTRAINT fk_collab_user_id FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE KEY unique_collab (playlist_id, user_id)
) ENGINE = InnoDB;