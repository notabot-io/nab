# Workspace

## Problem

Initially I had the idea that a single volume would be shared between all parts of a job, and that resources would be pulled into that "workspace" volume. This sounds fine in principle, up to a point. Then you realise that you can't mount a specific resource into a specific location if you want to.

To elaborate; if you wanted to pull a git repository resource and then run a task that uses that, it may be the case that the task requires that the cloned git resource be in a specific location. You might not want to mount _all_ resources into that location, so you may end up having to do some weird (and slow) hacky stuff to either bind mount specific resources into specific places, or instead perhaps copy the resources around (bearing in mind that this would be a two-way thing, you'd have to copy the result back to the "workspace" volume too).

## Solution

DNAB initially set out to work more like Drone in how it handles workspaces. Upon realising the above problem (and simultaneously why Drone is so inflexible with it's workspace) it became apparent that an alternative solution was required.

Drone lets you configure where the workspace gets mounted to as a whole. The workspace is then re-used in all containers, and mounted to the same place. This may be fine for what you're doing, but in some cases simply doesn't offer enough flexibility and that's what DNAB is striving for. Drone is limited by it's aim to only need the Docker socket to be mounted. This means that the workspace is less flexible, and has knock-on effects on other areas too, for example how scripts are injected into the containers.

So, what if you just mounted a file-system based workspace onto DNAB alongside the Docker socket. Bear in mind that this is all only a problem if you're running DNAB workers in Docker. That's not to say that it doesn't make sense to do that (though, there should only need to be one worker per-node) - but it's to say that it shouldn't be DNAB's concern.

DNAB should be flexible enough to run bare-metal and containerised. Sound familiar? This is how Jenkins slaves have been handled in-docker before. You don't need to run those in containers, but they can be.

So, what's the caveat? The same as when you do this with Jenkins. The workspace directory for the worker has to be the exact same place on the host and in the container:

```
version: "2"

services:
  # ...
  dnab-worker:
    image: dnab/dnab
    environment:
      DNAB_WORKSPACE: /home/dnab
    volumes:
    - /home/dnab:/home/dnab
```

Of course there would be a default value for the DNAB_WORKSPACE environment variable so that you didn't need to specify it if you wanted to set it up with the least configuration possible.

This solves all of the problems I can see:

* Resources can be mounted to specific places, individually.
* Resources are still available in the automatically mounted workspace otherwise.
* The workspace as a whole can still be configured.
* It removes the need for tasks to specify output explicitly.
  * While this was "nice" in some ways also meant more bloat in the configuration.
* Changes that occur in the same volume that are mounted elsewhere are still reflected in both.
