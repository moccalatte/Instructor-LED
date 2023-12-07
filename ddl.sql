CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    role VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_deleted BOOLEAN
);

CREATE TABLE student (
    student_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    fullname VARCHAR(255) NOT NULL,
    shortname VARCHAR(255) NOT NULL,
    birth_date DATE,
    birth_place VARCHAR(255),
    address VARCHAR(255),
    education VARCHAR(255),
    institution VARCHAR(255),
    job VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_deleted BOOLEAN
);

CREATE TABLE course (
    course_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_name VARCHAR(255) NOT NULL,
    course_detail_id uuid,
    is_deleted BOOLEAN
);

CREATE TABLE course_detail (
    course_detail_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    chapter VARCHAR(255) NOT NULL,
    course_id uuid,
    is_deleted BOOLEAN
);

CREATE TABLE question (
    question_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid,
    questionary VARCHAR(255),
    status BOOLEAN,
    is_deleted BOOLEAN
);

CREATE TABLE session (
    session_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    date DATE,
    trainer_admin_id UUID,
    is_deleted BOOLEAN
);

CREATE TABLE attendance (
    attendance_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    student_id uuid,
    course_id uuid,
    trainer_admin_id UUID,
    status BOOLEAN,
    is_deleted BOOLEAN
);

-- Add foreign key
ALTER TABLE course_detail
    ADD FOREIGN KEY (course_id) REFERENCES course(course_id);

ALTER TABLE course
    ADD FOREIGN KEY (course_detail_id) REFERENCES course_detail(course_detail_id);

ALTER TABLE question
    ADD FOREIGN KEY (user_id) REFERENCES users(user_id);

ALTER TABLE attendance
    ADD FOREIGN KEY (student_id) REFERENCES student(student_id);

ALTER TABLE session
    ADD FOREIGN KEY (trainer_admin_id) REFERENCES users(user_id);

ALTER TABLE attendance
    ADD FOREIGN KEY (course_id) REFERENCES course(course_id);
    
ALTER TABLE attendance
    ADD FOREIGN KEY (attendance_id) REFERENCES session(session_id);
