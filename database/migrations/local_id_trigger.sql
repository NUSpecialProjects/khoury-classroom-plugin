/* Automatically set the local id of an assignment before insertion */

CREATE OR REPLACE FUNCTION set_assignment_local_id()
RETURNS TRIGGER AS $$
BEGIN
    NEW.local_id := (
        SELECT COALESCE(MAX(local_id), 0) + 1
        FROM assignments
        WHERE semester_id = NEW.semester_id
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_set_assignment_local_id
BEFORE INSERT ON assignments
FOR EACH ROW EXECUTE FUNCTION set_assignment_local_id();


/* Automatically set the local id of a student assignment before insertion */

CREATE OR REPLACE FUNCTION set_student_assignment_local_id()
RETURNS TRIGGER AS $$
BEGIN
    NEW.local_id := (
        SELECT COALESCE(MAX(local_id), 0) + 1
        FROM student_assignments
        WHERE assignment_id = NEW.assignment_id
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_set_student_assignment_local_id
BEFORE INSERT ON student_assignments
FOR EACH ROW EXECUTE FUNCTION set_student_assignment_local_id();