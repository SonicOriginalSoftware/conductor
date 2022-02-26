Receives requests to queue pipelines, parses and verifies their payload, and then forwards the appropriate jobs for the pipeline event(s) on to the [queue](queue).

The `conductor` server only cares about whether it successfully adds a pipeline to a queue; the response should not dictate whether a runner has been assigned. It can report on where in the queue that pipeline was seated.

# Endpoints

- `queue` - Retrieves event payload and forwards pipeline to queue

  This will report back a failing status code if the pipeline could not be added to the queue.

  Otherwise responds with a `200` and the location/identifier of the pipeline in the queue.

# `portal` communication

Reports to the portal with
- the status of the conductor
  - the status of its runners
  - actions for its runners that the user wants to configure (see [Endpoints](#endpoints-1))

## Endpoints

- `add` - add a runner to be managed by the conductor
- `update` - update a runner to be managed by the conductor
- `remove` - remove a runner to be managed by the conductor
- `list` - list runners managed by the conductor

  This can be filtered by runner attributes
  Can be configured to retrieve status of runners as well
