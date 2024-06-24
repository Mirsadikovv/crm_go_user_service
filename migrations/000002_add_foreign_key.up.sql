ALTER TABLE students ADD FOREIGN KEY (group_id) REFERENCES groups (id);

ALTER TABLE teachers ADD FOREIGN KEY (branch_id) REFERENCES branches (id);

ALTER TABLE support_teachers ADD FOREIGN KEY (branch_id) REFERENCES branches (id);

ALTER TABLE managers ADD FOREIGN KEY (branch_id) REFERENCES branches (id);

ALTER TABLE administrations ADD FOREIGN KEY (branch_id) REFERENCES branches (id);

ALTER TABLE groups ADD FOREIGN KEY (branch_id) REFERENCES branches (id);

ALTER TABLE groups ADD FOREIGN KEY (teacher_id) REFERENCES teachers (id);

ALTER TABLE groups ADD FOREIGN KEY (support_teacher_id) REFERENCES support_teachers (id);

ALTER TABLE schedules ADD FOREIGN KEY (group_id) REFERENCES groups (id);

ALTER TABLE schedules ADD FOREIGN KEY (lesson_id) REFERENCES lessons (id);

ALTER TABLE journals ADD FOREIGN KEY (schedule_id) REFERENCES schedules (id);

ALTER TABLE events ADD FOREIGN KEY (branch_id) REFERENCES branches (id);

ALTER TABLE events ADD FOREIGN KEY (student_id) REFERENCES students (id);

ALTER TABLE student_performance ADD FOREIGN KEY (student_id) REFERENCES students (id);

ALTER TABLE student_performance ADD FOREIGN KEY (schedule_id) REFERENCES schedules (id);
