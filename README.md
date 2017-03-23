# Definitely Not A Bot (DNAB)

DNAB is a pipeline-based, Docker-centric, lightweight task runner written in Go. DNAB aims to take the best parts of other tools to create a powerful, flexible, and easy-to-use task runner.

## Motivation

The inspiration for DNAB comes from a few places in particular; Concourse, Drone, Jenkins, and Travis. All of these tools are excellent, but none of them perfectly achieved what DNAB set's out to achieve as a CI / CD / task running tool.

* **Concourse**: Probably the favourite runner up, Concourse is an incredibly well-made tool, but it falls short in just a few key areas. Configuration, and resource usage. You have to configure _so much_ for even the simplest of things. A Travis build that was maybe 20 lines might be around 200 lines with Concourse because of the boilerplate. Other tools like Drone manage to get around this. As for resource usage, even though Concourse is written in Go, it still manages to use a startling amount of memory.

* **Drone**: An excellent self-hosted Travis alternative, but just too opinionated in some areas. It's very powerful, and might be all you ever need. Resource usage is excellent. Honestly, Drone is great, it's just the limitations in configuration that let it down. Those limitations can be quite specific, but after using Concourse and Jenkins which both have the ability to run artbirary tasks, Drone starts to feel a little too limited.

* **Jenkins**: A trusty old tool that's now showing it's age. It is going through some changes with the introduction of pipelines, but it still ends up being a special snowflake. It hoards configuration, and resources.

* **Travis**: Pretty much the same points as Drone, but also it's not self-hosted. Has a lot of great functionality built in, but it not being self-hosted made it a no-go.

## License

MIT
