The `conductor` component uses event-driven architecture to facilitate software automation.

It also maintains state/load-balancing of the runners and pipeline queue it is responible for.

**Can the conductor service be on the same machine as runner service(s)?**

>  Sure. But you are competing for resources.

# Events

Events are defined by you.

This allows events to also be defined through other 1st party services like the `issue-tracker` and `merge-manager` but also through 3rd party services like webhooks from GitHub or Slack.

Some events, like code automation, are ready to be wired-up out of the box, like `git hooks` that can be enabled through setting up new `git` code repositories (see [here](../code-repository/git)).

Other scaffold-ready 1st party events can be read about in the component documentation.

Events can be daisy-chained by calling the `queue` endpoint from your pipeline.

# [Pipelines](docs/pipeline)

# [Server](docs/server)

# [Queue](docs/queue)

# [Runner](docs/runner)
