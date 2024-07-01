-- Create enum types
CREATE TYPE gender AS ENUM ('male', 'female', 'other');
CREATE TYPE course_level AS ENUM ('beginner', 'elementary', 'intermediate', 'ielts');

CREATE SEQUENCE student_external_id_seq START WITH 1;
CREATE SEQUENCE teacher_external_id_seq START WITH 1;
CREATE SEQUENCE support_teacher_external_id_seq START WITH 1;
CREATE SEQUENCE manager_external_id_seq START WITH 1;
CREATE SEQUENCE administration_external_id_seq START WITH 1;
CREATE SEQUENCE superadmin_external_id_seq START WITH 1;

CREATE TABLE IF NOT EXISTS students (
    id UUID PRIMARY KEY,
    group_id UUID NOT NULL,
    user_login VARCHAR(35),
    birthday DATE,
    gender gender,
    fullname VARCHAR(55),
    email VARCHAR(35) UNIQUE NOT NULL,
    phone VARCHAR(35),
    user_password VARCHAR,
    paid_sum NUMERIC,
    started_at DATE DEFAULT NOW(),
    finished_at DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS teachers (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    user_login VARCHAR,
    birthday DATE,
    gender gender,
    fullname VARCHAR(55),
    email VARCHAR(35),
    phone VARCHAR(35),
    user_password VARCHAR,
    salary NUMERIC,
    ielts_score NUMERIC,
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
    user_login VARCHAR,
    birthday DATE,
    gender gender,
    fullname VARCHAR(55),
    email VARCHAR(35),
    phone VARCHAR(35),
    user_password VARCHAR,
    salary NUMERIC,
    ielts_score NUMERIC,
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
    user_login VARCHAR,
    birthday DATE,
    gender gender,
    fullname VARCHAR(55),
    email VARCHAR(35),
    phone VARCHAR(35),
    user_password VARCHAR,
    salary NUMERIC,
    start_working DATE DEFAULT NOW(),
    end_working DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS administrators (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    user_login VARCHAR,
    birthday DATE,
    gender gender,
    fullname VARCHAR(55),
    email VARCHAR(35),
    phone VARCHAR(35),
    user_password VARCHAR,
    salary NUMERIC,
    start_working DATE DEFAULT NOW(),
    end_working DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS superadmins (
    id UUID PRIMARY KEY,
    user_login VARCHAR,
    birthday DATE,
    gender gender,
    fullname VARCHAR(55),
    email VARCHAR(35),
    phone VARCHAR(35),
    user_password VARCHAR,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS groups (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    teacher_id UUID NOT NULL,
    support_teacher_id UUID NOT NULL,
    group_name VARCHAR(40) UNIQUE,
    group_level course_level,
    started_at DATE DEFAULT NOW(),
    finished_at DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS branches (
    id UUID PRIMARY KEY,
    branch_name VARCHAR(40) UNIQUE,
    branch_location POLYGON,
    phone VARCHAR(35),
    open_time TIME,
    close_time TIME,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);


CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL,
    topic VARCHAR(35),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);



CREATE TABLE IF NOT EXISTS event_registrate (
    id UUID PRIMARY KEY,
    event_id UUID NOT NULL,
    student_id UUID NOT NULL 
);