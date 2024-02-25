SET  @api_user = 'library_api_user';

INSERT INTO library.author(first_name, middle_name, last_name, country, created_by, updated_by)
VALUES
  ('Stephen', NULL, 'King', 'United States', @api_user, @api_user),
  ('George', 'R.R.', 'Martin', 'United States', @api_user, @api_user),
  ('Sarah', 'J.', 'Maas', 'United States', @api_user, @api_user),
  ('Brandon', NULL, 'Sanderson', 'United States', @api_user, @api_user),
  ('Emily', NULL, 'Henry', 'United States', @api_user, @api_user);