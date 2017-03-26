# Pipelines... Or Not?

When looking for a more container-centric solution to a remote automation server, what was I looking
for really? Something like Drone that does one set of tasks in order? No. Something like Concourse
where jobs are tightly bound together (where "necessary")? Almost, but not quite. I was looking for
something that can just run arbitrary, uncomplicated, automated tasks in containers.

What I'm looking for, is Jenkins - but container centric.

When I say I'm looking for Jenkins, what I actually mean is Jenkins before pipelines took over. This
is really because pipelines are something more complex, and potentially dare I say... unnecessary.
We already had jobs that triggered other jobs by updating things after all, is that a pipeline? In
fact, what even is a pipeline? Does that even matter?

## What does this mean?

I think what I'm trying to say is, let's start off with something much simpler. Jobs don't need to
spit out specific pieces of output and be linked together and closely related. What problems (of
mine) would that solve? Well... none really.

Okay, so pipelines are a cool idea, but may not be "the answer" so to speak. And maybe jobs in the
way that I'm thinking are like "mini-pipelines". But what about problems that they solve then...
This seems more to do with automation than anything; for example, how do we have a job run based on
change propogated by some other job?

SHOCK HORROR, we can already do this. We already _did_ do this, with Jenkins. We set up jobs that
reacted to changes in code, and then did some stuff. It let you build jobs that knew about it's
environment, and things that triggered the job were passed in. It's so simple, yet it lets you
accomplish so much!

How about instead of "pipelines" with inputs and outputs, we instead have downstream jobs that we
can pass parameters to? This would let us reliably trigger jobs at the end of other jobs. Could this
be accomplished by pushing to resources (thinking of `passed` in Concourse as an example here).
Triggering jobs is simple, but updating resources is more scalable, because it could notify a LOT of
different things in different ways.

To push to resources in that manner, you'd need something that could apply metadata to a specific
version of resource (for example, that it passed testing).
