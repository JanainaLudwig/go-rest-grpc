CREATE DATABASE go_api;
USE go_api;
#
# CREATE TABLE students (
#                           id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
#                           first_name VARCHAR(255) NOT NULL,
#                           last_name VARCHAR(255) NOT NULL,
#                           identifier VARCHAR(255) NOT NULL UNIQUE,
#                           created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
#                           updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
# );
#
# CREATE INDEX idx_students_identifier ON students (identifier);
# CREATE FULLTEXT INDEX idx_students_full_name ON students (first_name, last_name);
#
# INSERT INTO students (first_name, last_name, identifier) VALUES
#                                                              ('Alisa', 'Roberts',	'01c6c4bb-7f9e-4622-b8ae-9b425e12d066'),
#                                                              ('Olivia', 'Rogers',	'2b89c8b0-100f-40da-be26-44b9087735a9'),
#                                                              ('Rubie', 'Hawkins',	'0b346cf9-3066-45c0-a1e5-f83ff6080d02');
