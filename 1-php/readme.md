PHP / Laravel
=============

In the `data-capture` directory, you will find a fresh Laravel project. Please edit this project to fill the brief.

**Read all of the tasks first**, then attempt them (as the later tasks may affect how you write the first two).

At the end of this, you should have one endpoint and one artisan command (and whatever other helpers, tests, etc. you create along the way).

## Task One

Make an authenticated API endpoint that handles the posting of questionnaire data.

```bash
POST api/questionnaire_result

{
	"questionnaire_id": '',
	"results": {}
}
```

The data should be validated, and inserted into the database.

## Task Two

Amend the endpoint so that it takes an additional optional parameter specifying the schedule it is related to.

```bash
POST api/questionnaire_result

{
	"questionnaire_id": '',
	"results": {},
	"questionnaire_schedule_id": ''|undefined
}
```

This, if present, should link the posted result to the schedule table entry, and should update the appropriate column on the schedules table.

## Task Three

Make an artisan console command that handles the posting of questionnaire data from the command line. The end goal should be the same as the above endpoint.

```bash
php artisan questionnaire_result:create [user id] [questionnaire id] [path to results json file] (optional --schedule=[schedule id])
```

## Task Four

When saving the results, if the questionnaire is part of a schedule, use the AWS SDK to push a message to a queue (the queue url is in config/services.php under the `queues.schedule` key), so that another application can respond to such an event.

## Assumptions / Notes

* **Do** write relevant tests!
* **Do not** worry about authentication / login etc. - a simple `auth` middleware on a route is enough for now, and you can get your user context from Laravel's auth helpers.
