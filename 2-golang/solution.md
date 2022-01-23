# Go Task Solution

I took a stab at this task. I've written a sort of half-baked/pseudocode implementation so **it does not run and there are probably typos** but hopefully you can see from this what I was trying to do.

My original plan was to create separate functions for the handling of the database insert and messaging. To make these asynchronous I would call them in the LambdaHandler as goroutines and then communicate via 2 separate channels for each type of event.

I wasn't really satisfied with my implementation and (as I'm fairly inexperienced in Go) decided to use synchronous functions instead. I'm very interested to see what a correct asynchronous implementation looks like. Perhaps this is something we could chat about.