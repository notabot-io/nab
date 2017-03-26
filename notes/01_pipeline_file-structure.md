# Pipeline: File Structure

Similar to Concourse, DNAB will have a declarative syntax and will aim to keep all configuration for pipelines as code. This will mean that DNAB can be completely destroyed, and as long as the configuration is kept somewhere safe (i.e. Git) then it won't be the end of the world.

The structure will aim to stick to a few core concepts that are repeated throughout (again, like Concourse). The difference compared to Concourse will be that we'll provide a lot more defaults. Another aim will be to provide more things to help builds like Drone and Jenkins do in the form of variables (i.e. environment variables, and variables in the pipeline definition).

Around those core concepts DNAB will have some more opinionated building blocks. Essentially, this will make up the structure of the pipeline definition file.

## Resources

Almost exactly like Concourse's resources, DNAB's resources are very abstract building blocks for pipelines. The name "resource" implies that it will be something used by something else - and that's exactly what it is. A resource will have some state, it can be checked, retrieved, and pushed to. Any of those things can also be no-op.

Resources will be created as Docker images. That will mean that they are able to be written in pretty much any language, in any way a developer wants as long as they follow the correct interface.

## Jobs

Jobs pull together tasks, and form a part of a pipeline. Jobs run on their own. They take some input and produce some output, this can be fed into other jobs.
