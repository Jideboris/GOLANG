# PHP Task Solution

## Caveats
I wasn't entirely sure what was meant in the instructions by "a simple `auth` middleware on a route is enough for now", because the auth middleware redirects to /login and a user would need to authenticate to create a session, so I've just used auth.basic on the single route. I think this still satisfies the requirements. The password in UserFactory is `example`.

I assumed that the validation would need to check for an existing questionnaire and scheduled questionnaire so I've included a basic seeder as well.
