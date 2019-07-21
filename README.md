# RaftCache

`RaftCache` aims to be the fastest & resilient cache system. It's mainly inspired [GroupCache](https://github.com/golang/groupcache) & [raft](https://raft.github.io/).

If you are familiar with GroupCache, you could consider DriftCache to be the GroupCache that supports write & linery consistency; but it's **NOT** using raft consensus algorithm.

## Design Goal

* Designed for embeded cache, other than stand-alone cache
* `Key-value` cache store
* **Master-less** architecture, all nodes are equal
 * Any nodes could join/leave cluster at any time
* **CP** system in teams [CAP theorem](https://en.wikipedia.org/wiki/CAP_theorem)
  * **Linear consistency** must be achived for individual key
* System is always available to provide `Read Access`
  * `Write Access` is only affected during cluster re-grouping

# Architecture

## Master-less

All nodes in `RaftCache` are equal peer, there is no master, nor `master election` process.

Each node are `Owner` of certain keys through a consistent hash algorithm. The `Owner` node of given key is responsible for the key's writing process and replication to peer nodes. Thus, `linear consistency` is possible.

## Node Join process

### Node Statuses

Each node could be in the following status:

* Stand-alone
  * Initial status, node could remain in `stand-alone` status if no cluster is in used
* In-group
  * Status after joining with peer nodes to form a `group`
* Disconnected
  * Temporary status during network patition happened
* Hand-shaking
  * Temporary status of existing cluster peer when new node is joining
* Initiating
  * Temporary status for a node is joing a `group` cluster

> more status would be added for handling network partition


## Test

## Todo
