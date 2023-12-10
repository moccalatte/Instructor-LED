CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; 
 
CREATE TABLE users ( 
    user_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), 
    fullname VARCHAR(255) NOT NULL, 
    role VARCHAR(255), 
    email VARCHAR(255) UNIQUE NOT NULL, 
    password VARCHAR(255) NOT NULL, 
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
 updated_at TIMESTAMP, 
    is_deleted BOOLEAN 
); 
 
CREATE TABLE student ( 
    student_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), 
    fullname VARCHAR(255) NOT NULL, 
    birth_date VARCHAR(255), 
    birth_place VARCHAR(255), 
    address VARCHAR(255), 
    education VARCHAR(255), 
    institution VARCHAR(255), 
    job VARCHAR(255), 
    email VARCHAR(255) UNIQUE NOT NULL, 
    password VARCHAR(255) NOT NULL, 
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
updated_at TIMESTAMP, 
    is_deleted BOOLEAN 
); 
 
CREATE TABLE course ( 
    course_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), 
    course_name VARCHAR(255) NOT NULL, 
    description VARCHAR(255) NOT NULL, 
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
updated_at TIMESTAMP, 
    is_deleted BOOLEAN 
); 
 
CREATE TABLE course_detail ( 
    course_detail_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), 
    course_id uuid, 
    course_chapter VARCHAR(255) NOT NULL, 
course_content VARCHAR(255) NOT NULL, 
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
updated_at TIMESTAMP, 
    is_deleted BOOLEAN 
); 
 
CREATE TABLE question ( 
    question_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), 
 session_id uuid, 
    student_id uuid, 
 trainer_id uuid, 
    title VARCHAR(255) NOT NULL, 
    description VARCHAR(255) NOT NULL, 
 course_id uuid, 
 image VARCHAR(255) NOT NULL, 
 answer VARCHAR(255) NOT NULL, 
    status VARCHAR (10) NOT NULL, 
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
 updated_at TIMESTAMP, 
    is_deleted BOOLEAN 
); 
 
CREATE TABLE session ( 
    session_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), 
 title VARCHAR(255) NOT NULL, 
 description VARCHAR(255) NOT NULL, 
    session_date VARCHAR(20) NOT NULL, 
 session_time VARCHAR(20) NOT NULL, 
 session_link VARCHAR(255) NOT NULL, 
    trainer_id UUID, 
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
 updated_at TIMESTAMP, 
    is_deleted BOOLEAN 
); 
 
CREATE TABLE attendance ( 
    attendance_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), 
session_id uuid, 
    student_id uuid, 
attendance_student BOOLEAN, 
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
updated_at TIMESTAMP, 
    is_deleted BOOLEAN 
); 
 
-- Add foreign key 
 
ALTER TABLE course_detail 
    ADD FOREIGN KEY (course_id) REFERENCES course(course_id); 
 
ALTER TABLE question 
ADD FOREIGN KEY (student_id) REFERENCES student(student_id); 
 
ALTER TABLE question 
ADD FOREIGN KEY (trainer_id) REFERENCES users(user_id); 
 
ALTER TABLE question 
ADD FOREIGN KEY (session_id) REFERENCES session(session_id); 
 
ALTER TABLE question 
ADD FOREIGN KEY (course_id) REFERENCES course(course_id); 
 
ALTER TABLE attendance 
    ADD FOREIGN KEY (student_id) REFERENCES student(student_id); 
 
ALTER TABLE attendance 
    ADD FOREIGN KEY (session_id) REFERENCES session(session_id); 
 
ALTER TABLE session 
    ADD FOREIGN KEY (trainer_id) REFERENCES users(user_id);

-- Select all columns from the users table
SELECT * FROM users;

-- Select specific columns from users where is_deleted is true
SELECT user_id, fullname, role, email, password, is_deleted
FROM users
WHERE is_deleted = true;

-- Insert a new row into users with is_deleted set to true
INSERT INTO users(fullname, role, email, password, is_deleted)
VALUES ('gopan', 'admin', 'admin@gmail.com', '12345678', true)
RETURNING user_id, fullname, role, email, password, is_deleted;