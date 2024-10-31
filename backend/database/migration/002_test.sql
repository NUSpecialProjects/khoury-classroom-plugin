DROP TABLE IF EXISTS test;

CREATE TABLE test (
    id SERIAL PRIMARY KEY,
    content VARCHAR(255)
);


-- Insert data into the test table
INSERT INTO test (content) VALUES ('First entry');
INSERT INTO test (content) VALUES ('Second entry');
INSERT INTO test (content) VALUES ('Third entry');
INSERT INTO test (content) VALUES ('Fourth entry');
INSERT INTO test (content) VALUES ('Fifth entry');