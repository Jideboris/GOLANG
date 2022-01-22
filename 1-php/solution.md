# PHP Task Solution

## Caveats
- I wasn't entirely sure what was meant in the instructions by "a simple `auth` middleware on a route is enough for now", so I've just used auth.basic on the single route. I think this still satisfies the requirements. The password in UserFactory is `example`.

- I assumed that the validation would need to check for an existing questionnaire and scheduled questionnaire so I've included a basic seeder as well.

- I created a specific database connection for tests, with the environment variable `TEST_DB_DATABASE`. Just run `php artisan migrate --database=testing`