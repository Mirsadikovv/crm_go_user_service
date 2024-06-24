-- Create enum types
CREATE TYPE gender AS ENUM ('male', 'female', 'other');
CREATE TYPE level AS ENUM ('beginner', 'elementary', 'intermediate', 'ielts');

CREATE SEQUENCE student_external_id_seq START WITH 1;
CREATE SEQUENCE teacher_external_id_seq START WITH 1;
CREATE SEQUENCE support_teacher_external_id_seq START WITH 1;
CREATE SEQUENCE manager_external_id_seq START WITH 1;
CREATE SEQUENCE administration_external_id_seq START WITH 1;
CREATE SEQUENCE superadmin_external_id_seq START WITH 1;

CREATE TABLE IF NOT EXISTS students (
    id UUID PRIMARY KEY,
    group_id UUID NOT NULL,
    user_login VARCHAR(20),
    birthday VARCHAR(20),
    gender gender,
    fullname VARCHAR(35),
    email VARCHAR(35),
    phone VARCHAR(20),
    user_password VARCHAR(20),
    paid_sum DECIMAL,
    started_at DATE DEFAULT NOW(),
    finished_at DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS teachers (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    user_login VARCHAR(20),
    birthday VARCHAR(20),
    gender gender,
    fullname VARCHAR(35),
    email VARCHAR(35),
    phone VARCHAR(20),
    user_password VARCHAR(20),
    salary DECIMAL,
    ielts_score DECIMAL,
    ielts_attempts_count INT,
    start_working DATE DEFAULT NOW(),
    end_working DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS support_teachers (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    user_login VARCHAR(20),
    birthday VARCHAR(20),
    gender gender,
    fullname VARCHAR(35),
    email VARCHAR(35),
    phone VARCHAR(20),
    user_password VARCHAR(20),
    salary DECIMAL,
    ielts_score DECIMAL,
    ielts_attempts_count INT,
    start_working DATE DEFAULT NOW(),
    end_working DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS managers (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    user_login VARCHAR(20),
    birthday VARCHAR(20),
    gender gender,
    fullname VARCHAR(35),
    email VARCHAR(35),
    phone VARCHAR(20),
    user_password VARCHAR(20),
    salary DECIMAL,
    start_working DATE DEFAULT NOW(),
    end_working DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS administrators (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    user_login VARCHAR(20),
    birthday VARCHAR(20),
    gender gender,
    fullname VARCHAR(35),
    email VARCHAR(35),
    phone VARCHAR(20),
    user_password VARCHAR(20),
    salary DECIMAL,
    start_working DATE DEFAULT NOW(),
    end_working DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS superadmins (
    id UUID PRIMARY KEY,
    user_login VARCHAR(20),
    birthday VARCHAR(20),
    gender gender,
    fullname VARCHAR(35),
    email VARCHAR(35),
    phone VARCHAR(20),
    user_password VARCHAR(20),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS groups (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    teacher_id UUID NOT NULL,
    support_teacher_id UUID NOT NULL,
    group_name VARCHAR(20),
    number_of_students INT,
    started_at DATE DEFAULT NOW(),
    finished_at DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS branches (
    id UUID PRIMARY KEY,
    branch_name VARCHAR(30),
    administration_id UUID,
    location POLYGON,
    phone VARCHAR(20),
    open_time TIMESTAMP,
    close_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS schedules (
    id UUID PRIMARY KEY,
    group_id UUID NOT NULL,
    lesson_id UUID NOT NULL,
    classroom VARCHAR(20),
    group_name VARCHAR(20),
    type_of_group level,
    task VARCHAR(35),
    deadline TIMESTAMP,
    score INT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS lessons (
    id UUID PRIMARY KEY,
    theme VARCHAR(35),
    links VARCHAR(255),
    type_of_group level,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS journals (
    id UUID PRIMARY KEY,
    schedule_id UUID NOT NULL,
    date_of_lesson DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    student_id UUID NOT NULL,
    topic VARCHAR(35),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS student_performance (
    id UUID PRIMARY KEY,
    student_id UUID NOT NULL,
    schedule_id UUID NOT NULL,
    attended BOOLEAN NOT NULL,
    task_score DECIMAL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
