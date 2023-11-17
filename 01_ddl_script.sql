CREATE TABLE IF NOT EXISTS users5 (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL
);

INSERT INTO users5 (username, email) VALUES
('john_doe', 'john.doe@example.com'),
('jane_doe', 'jane.doe@example.com');