CREATE TABLE students_subjects
(
    id         INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    id_student INT UNSIGNED NOT NULL,
    id_subject INT UNSIGNED NOT NULL,
    frequency  FLOAT UNSIGNED NULL,
    status     ENUM('in_progress', 'approved', 'reproved') NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_student) REFERENCES students(id) ON DELETE CASCADE,
    FOREIGN KEY (id_subject) REFERENCES subjects(id) ON DELETE CASCADE
);
