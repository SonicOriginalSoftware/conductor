# A FIFO queue of pipeline jobs

Once a pipeline is verified it is added to the queue so the queue can triaging jobs, looking for applicable runners available for the job requirements.

The `queue` will also transmit any necessary `resources` along with the job to the runner that will execute that job.

Once a runner is ready to execute the job, it will assign an `id` to that job, hand it off along with all resources to the runner, and then await the runner's completion of that job.

It will then follow the pipeline's instructions for handling the success or failure of that job's commands using the `if` statements of the pipeline and determine if later jobs need to be run.

# Organization of Pipelines in the `queue`

A simple map of the pipeline with its associated jobs makes up the queue.
