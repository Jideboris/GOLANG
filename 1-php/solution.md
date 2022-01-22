# PHP Task Solution

## Caveats & Instructions
- I wasn't entirely sure what was meant in the instructions by "a simple `auth` middleware on a route is enough for now", so I used auth.basic on the single route, which I think still satisfies the requirements. The password in UserFactory is `example`.

- I assumed that the validation would need to check for an existing questionnaire and scheduled questionnaire so I've included a basic seeder as well.

- I created a specific database connection for tests, with the environment variable `TEST_DB_DATABASE`. Just run `php artisan migrate --database=testing`

- I wasn't entirely sure what was meant for task four. Initially I wrote my own message logic with the AWS SDK (commits `1d2af21...30e2bba`) but later replaced this with a job, which was a bit simpler. Happy to explain both approaches.