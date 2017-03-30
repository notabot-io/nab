# Job Parameters

Just like in Jenkins, jobs should accept parameters to populate some initial variables for the run. Also similarly to Jenkins, it should allow for pre-set values, for example, if a job is a downstream job, it's helpful to have a list of builds to choose to pull artifacts from. This kind of thing usually works by utlising the build number. In the case of Jenkins, artifacts are actually stored on Jenkins. For NAB this will be different however, but a build number for a job could still be useful to identify versions of resources to use.

In Jenkins, there is a plugin that can be used to handle this. How could we handle this with NAB?

* Some kind of plugin system of our own?
    * That's quite a complicated task, and opens many cans of worms.
* Use resources, and variables that are output from `pull`?
* Use resources, and versions that are output from `check`?
    * This seems like it could be the best approach, especially if NAB is going to keep track of allow versions of a resource somehow (which it probably should do?)
    * This approach would seem to indicate that a real database server somewhere would be necessary as an external dependency, e.g. Postgres, or MySQL. This sould allow versions of resources to be stored and retrieved efficiently, without complicating the code of NAB.
        * This is all unless there is an embedded "SQL-like" database solution we can use (that isn't SQLite!)

**Conclusion**: Resource versions can optionally be provided as parameters, with some kind of filtering (e.g. limit 10, order desc). Cache versions in the database when check runs (this is likely to happen anyway).

There could be other things that provide data to build parameters - but resource versions should cover a lot of bases, and in any case, you can still just manually input things.
