# Go Task Solution

I took a stab at this task. I've written a sort of half-baked/pseudocode implementation so **it does not run and there are probably typos** but hopefully you can see from this what I was trying to do.

My plan was to create separate functions for the handling of the database insert and messaging. To make these asynchronous I would call them in the LambdaHandler as goroutines and then communicate via 2 channels for the 2 different types of event (`NewScheduleCreatedEvent` and `ScheduledUserQuestionnairesCompletedEvent`).

I'm very interested to see what a "correct" implementation looks like (as someone who's still relatively new to Go). Perhaps this is something we could chat about.