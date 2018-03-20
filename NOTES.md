# Design notes

## Requirements

* Stable storage on S3 
* No replication
* Log on disk
* Atomic batch writes and snapshot reads (for ACID)
