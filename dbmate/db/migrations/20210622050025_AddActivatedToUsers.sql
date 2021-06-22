-- migrate:up
ALTER TABLE users 
ADD COLUMN activated boolean DEFAULT false ;

-- migrate:down
ALTER TABLE users 
DROP COLUMN activated;

