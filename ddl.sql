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
role VARCHAR(10), 
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
    note VARCHAR, 
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
