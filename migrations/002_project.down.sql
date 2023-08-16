DROP SEQUENCE projects_id_seq CASCADE;
DROP TABLE projects;

ALTER TABLE services ADD COLUMN gh_repo text;
ALTER TABLE services ADD COLUMN gl_repo text;
ALTER TABLE services ADD COLUMN deploy_file text;
ALTER TABLE services ADD COLUMN namespace text;
ALTER TABLE services ADD COLUMN branch text;
