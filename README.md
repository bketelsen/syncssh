# syncssh

Quickly add remote LXD instances to your ssh client configuration. Expecially useful for VS Code Remote/SSH development.

## Usage

`syncssh sync` will create a file at `$HOME/.ssh/config.lxd`. Add an import statment in your
`$HOME/.ssh/config` file like this:

```
... other stuff
Include ~/.ssh/config.lxd

```

Run `syncssh sync` manually, or add to cron/launchctl to synchronize your configs on a timer.
