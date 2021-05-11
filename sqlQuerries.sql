CREATE TYPE ACCOUNTSTATUS AS ENUM (
    'verified',
    'upload_pending',
    'verification_pending',
    'failed'
);

--  Command to create user table
CREATE TABLE users 
(
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    fname TEXT NOT NULL,
    lname TEXT NOT NULL,
    sex TEXT NOT NULL,
    email TEXT NOT NULL,
    status ACCOUNTSTATUS DEFAULT 'upload_pending'
);

-- Command to add a default state to status
-- ALTER TABLE users ALTER COLUMN status SET DEFAULT 'upload_pending' ;

-- Command to describe table 
\d users

-- Command to check available tables in the DB
\dt+

-- Command to check the available types
\dT+
