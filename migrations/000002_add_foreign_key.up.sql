ALTER TABLE students ADD FOREIGN KEY (group_id) REFERENCES groups (id);

ALTER TABLE teachers ADD FOREIGN KEY (branch_id) REFERENCES branches (id);

ALTER TABLE support_teachers ADD FOREIGN KEY (branch_id) REFERENCES branches (id);

ALTER TABLE managers ADD FOREIGN KEY (branch_id) REFERENCES branches (id);

ALTER TABLE administrators ADD FOREIGN KEY (branch_id) REFERENCES branches (id);

ALTER TABLE groups ADD FOREIGN KEY (branch_id) REFERENCES branches (id);

ALTER TABLE groups ADD FOREIGN KEY (teacher_id) REFERENCES teachers (id);

ALTER TABLE groups ADD FOREIGN KEY (support_teacher_id) REFERENCES support_teachers (id);

ALTER TABLE event_registrate ADD FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE;

ALTER TABLE event_registrate ADD FOREIGN KEY (student_id) REFERENCES students (id) ON DELETE CASCADE;
