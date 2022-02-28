<!-- runner software should be able to be run cross-platform and statically -->
<!-- i.e. should be able to spin-up a deamonized runner on macOS, linux, or Windows -->
<!-- linux runners could be containerized -->
<!-- should allow runners to be configured across networks -->
<!-- i.e. can use a self-hosted VM runner, self-hosted bare-metal runner, cloud-based runner container service, etc. -->

Configure runners attached to the `conductor` through the portal.

Configuring runners includes defining network location and attributes for each runner service.

# Purpose

`conductor` `runner`s run individual jobs passed on to it from the `conductor` `server`/`queue`.

# `gRPC`

Conductor `queue` sends jobs to a runner through `gRPC`.

## Endpoints

- `status` - reports back with the state of the runner (defined in [States](#states))
- `start` - receives job commands and begins execution

# States

- `available`
- `running`
- `undefined` - not a good state to be in; issue a restart of this runner's service

# Jobs

The runner executes job commands and when finished reports that the runner is available for new jobs.

# Attributes

## OS

### Windows

**TODO**

### Linux

**TODO**

### macOS

**TODO**

## Architecture

### x86_64

**TODO**

### arm64

**TODO**
