CREATE TABLE song(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title TEXT NOT NULL,
    year TEXT NOT NULL,
    genre TEXT NOT NULL,
    performer TEXT NOT NULL,
    duration INT NOT NULL,
    album_id INT NOT NULL,
    CONSTRAINT fk_album_id FOREIGN KEY (album_id) REFERENCES album(id) ON DELETE CASCADE
) ENGINE = InnoDB;