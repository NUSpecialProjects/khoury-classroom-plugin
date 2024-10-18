CREATE OR REPLACE FUNCTION gen_short_uuid() RETURNS VARCHAR(22) AS $$
BEGIN
  RETURN substr(rtrim(replace(replace(encode(uuid_send(gen_random_uuid()), 'base64'), '/', '~'), '+', '-'), '='), 1, 22);
END;
$$ LANGUAGE plpgsql;