CREATE TABLE playlist_songs(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    playlist_id INT NOT NULL,
    song_id INT NOT NULL,
    CONSTRAINT fk_playlist_id FOREIGN KEY(playlist_id) REFERENCES playlists(id) ON DELETE CASCADE,
    CONSTRAINT fk_song_id FOREIGN KEY(song_id) REFERENCES song(id) ON DELETE CASCADE
) ENGINE = InnoDB;