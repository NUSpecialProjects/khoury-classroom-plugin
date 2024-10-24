/* Automatically set the local id of an assignment before insertion */

DO $$
BEGIN
    -- Check if the trigger already exists
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger
        WHERE tgname = 'trigger_set_assignment_local_id'
        AND tgrelid = 'assignments'::regclass
    ) THEN
        -- Create the function
        CREATE OR REPLACE FUNCTION set_assignment_local_id()
        RETURNS TRIGGER AS $$
        BEGIN
            NEW.local_id := (
                SELECT COALESCE(MAX(local_id), 0) + 1
                FROM assignments
                WHERE classroom_id = NEW.classroom_id
            );
            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;

        -- Create the trigger
        CREATE TRIGGER trigger_set_assignment_local_id
        BEFORE INSERT ON assignments
        FOR EACH ROW EXECUTE FUNCTION set_assignment_local_id();
    END IF;
END $$;

/* Automatically set the local id of a student assignment before insertion */

DO $$
BEGIN
    -- Check if the trigger already exists
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger
        WHERE tgname = 'trigger_set_student_assignment_local_id'
        AND tgrelid = 'student_assignments'::regclass
    ) THEN
        -- Create the function
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

        -- Create the trigger
        CREATE TRIGGER trigger_set_student_assignment_local_id
        BEFORE INSERT ON student_assignments
        FOR EACH ROW EXECUTE FUNCTION set_student_assignment_local_id();
    END IF;
END $$;
