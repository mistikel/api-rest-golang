CREATE TABLE records (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    marks JSON NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX(created_at)
);

INSERT INTO mezink_db.records (name,marks,created_at) VALUES
	 ('First','[1, 2, 3, 4]','2024-08-29 04:12:38'),
	 ('Second','[4, 5, 6, 7]','2024-08-27 08:18:27'),
	 ('Third','[8, 9, 10, 11]','2024-08-28 08:18:26'),
	 ('Fourth','[10, 11, 12, 13, 14]','2024-08-26 08:18:26');
