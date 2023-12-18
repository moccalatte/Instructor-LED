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

-- dummy data 

INSERT INTO users (fullname, role, email, password, created_at, updated_at, is_deleted)
VALUES
    ('John Doe', 'admin', 'john.doe@example.com', 'password123', NOW(), NOW(), false),
    ('Jane Smith', 'trainer', 'jane.smith@example.com', 'studentpassword', NOW(), NOW(), false),
    ('Trainer One', 'trainer', 'trainer.one@example.com', 'trainerpassword', NOW(), NOW(), false);
