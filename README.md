# Not A Bot (NAB)

NAB is a container-centric, lightweight, automation server written in Go. NAB aims to take the best parts of other tools to create a powerful, flexible, and easy-to-use tool.

## Motivation

The inspiration for NAB comes from a few places in particular; Concourse, Drone, Jenkins, and Travis. All of these tools are excellent, but none of them perfectly achieved what NAB sets out to achieve as an automation tool. A lot of the things that these tools do very well has been looked at and learnt from, and NAB sets out to solve as many of the problems with them as it can. This mainly boils down to:

* **Configuration**: NAB aims to  have a simple but flexible declarative configuration syntax that you keep as code in a repository that is applied to a server (think Kubernetes).
* **Resource Usage**: NAB aims to be easy and cheap to host. Resource usage needs to be kept to a minimum. A $5p/m VPS should be all you need to get started. It should grow as you need it to.
* **Portability**: NAB also aims to be simple to set up, and easy to deploy in a wide range of environments. It may not be perfect for everything, but it hopes to cover a lot of bases.

## Other Tools

### Concourse

While Concourse is an excellent tool, and is fantastic in it's own right, it falls short in some key areas: configuration, and resource usage.

**Configuration** in Concourse is very declarative. You don't have to worry about your CI server dying because all of it's configuration is just configuration files. Great! The problems come in when you want to express some things more concisely. There are no variables - by design, but this can be a hindrance and instead leads to much more verbose pipelines, or lots of custom resources being created.

Concourse's pipeline files can get pretty difficult to parse sometimes, and they can also grow to be very large at the best of times. This could be helped with the ability to split up configuration, and also if configuration was easier to reuse.

**Resource usage** is great, but not in the "cheap self-hosting" sort of range. It's a very complex tool thanks to how it handles running containers and it's configuration, but it does mean that self-hosting Concourse can become a little more expensive.

### Drone

An excellent self-hosted Travis clone. It's very powerful, and might be all you ever need. Resource usage is excellent too! The pipeline syntax is very simple and easy to parse. So where does Drone fall short?

**Configuration** again... Drone just isn't flexible enough. You can see some of the problems that the Drone developer(s) have tried to solve, things like how to have a shared workspace between jobs, and controlling how that workspace is mounted into jobs - there just might be a better solution to that problem than what they came up with. It's limitations in that area are very similar to that of Concourse.

You can't define arbitrary jobs like you can in Concourse or Jenkins too. This is probably fine for some people, but if you're here, maybe you're looking for an automation server in general - not just something to build a specific repository and go down a specific path in doing so.

### Jenkins

A trusty old tool that's now showing it's age. It is going through some changes with the introduction of pipelines, but it still ends up being a special snowflake.

**Configuration** is hoarded by Jenkins. Jenkins has your configuration, and it's a bit trickier to keep configuration as code - even with the new pipelines plugin. You still have to configure a lot in the jobs themselves - these can be kept as just the XML configuration files, but then what about all of the other things that come in to interact with the jobs.

**Resource usage**... well, it is built on Java, and as most people know, it's going to be fairly heavy straight out of the gate. Jenkins is also very complex which probably leads it to use more resources just to deal with all of the things that it can do (and all of the plugins that mount up that become _part of Jenkins_). This isn't a big issue, but again, self-hosting as an individual can become quite expensive.

### Travis

Pretty much the same points as Drone, but also it's not self-hosted. It has a lot of great functionality built in, and simple configuration.

## License

MIT
