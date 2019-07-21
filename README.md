# DraftCache

`RaftCache` aims to be the fastest & resilient cache system. It's mainly inspired [GroupCache](https://github.com/golang/groupcache) & [raft](https://raft.github.io/).

If you are familiar with GroupCache, you could consider DriftCache to be the GroupCache that supports write & linery consistency; but it's **NOT** using raft consensus algorithm.

## Design Goal

* `Key-value` cache store
* **Master-less** artchitecture, all nodes are equal
 * Any nodes could join/leave cluster at any time
* **CP** system in teams [CAP theorem](https://en.wikipedia.org/wiki/CAP_theorem)
  * **Linear consistency** must be achived for individual key
* System is always available to provide `Read Access`
  * `Write Access` is only affected during cluster re-grouping

## Network design

TBA

## Test

## Todo
