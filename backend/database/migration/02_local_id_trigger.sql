-- Check if the trigger for setting assignment local_id already exists before creating it
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'trigger_set_assignment_local_id'
    ) THEN
        -- Automatically set the local id of an assignment before insertion
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

        -- Create trigger using EXECUTE
        EXECUTE 'CREATE TRIGGER trigger_set_assignment_local_id
                 BEFORE INSERT ON assignments
                 FOR EACH ROW EXECUTE FUNCTION set_assignment_local_id()';
    END IF;
END $$;

-- Check if the trigger for setting student assignment local_id already exists before creating it
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'trigger_set_student_assignment_local_id'
    ) THEN
        -- Automatically set the local id of a student assignment before insertion
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

        -- Create trigger using EXECUTE
        EXECUTE 'CREATE TRIGGER trigger_set_student_assignment_local_id
                 BEFORE INSERT ON student_assignments
                 FOR EACH ROW EXECUTE FUNCTION set_student_assignment_local_id()';
    END IF;
END $$;
