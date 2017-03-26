# Config: File Types

Now that it seems to make more sense to use a more traditional "Jenkins-like" approach - with jobs, and not pipelines (at least in the "modern" sense) - it may also makes more sense tackle configuration in a bit of a different way compared to the very Concourse-like approach I've adopted thus far. Initially, I had the idea of using one big file with all of the configuration in it. This can result in things getting lost much easier when that file starts to get bigger.

What if we took a more Kubernetes-like approach to configuration and instead had the ability to also use a bunch of smaller files (but only if it was "necessary"). This may be accomplished by using YAML's multiple documents in a single file feature. This gives you complete flexibility over how you define your confiugration, and how you group things up.
