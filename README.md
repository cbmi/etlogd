# etlogd

etlogd is service for logging metadata during an ETL process. One of the challenges with complex ETL is keeping track of **where data comes from**, **how it was processed**, and **when it populated**. etlog messages contain information about the the _source_ and _target_ data as well as the _script_ that performed the ETL.

The intended outcomes of etlogd is to enable:

- Linking production data back to the source data
- Capturing the necessary metadata to inform data changes over time

The features implemented to support the above outcomes include:

- Flexible document-based data model for storing as much or as little necessary to provide utility
- Built-in support for common data stores such as files, relational databases, document types, etc.

Read about the [etlog client tools](https://github.com/cbmi/etlog/) for specifics on natively supported types.

## Install

```bash
% go get github.com/cbmi/etlogd
```

## Setup

Requires a mongodb instance.

## Usage

The default is an HTTP server which handles POST requests.

```bash
% etlogd
Listening on http://0.0.0.0:4000...
```

### Options

```bash
% etlogd -help
Usage of etlogd:
  -collection="": mongo collection name
  -config="": path to ini-style config file
  -database="": mongo database name
  -host="": server host
  -mongo="": mongo host
  -type="": server type [http (default), tcp, udp]
  -verbose=false: verbose output
```

#### INI Format

```dosini
[general]
verbose = true

[server]
type = http
host = 0.0.0.0:6201

[mongo]
host = 0.0.0.0:27017,0.0.0.0:27018,0.0.0.0:27019
```

Command line arguments take precedence over config file. For example, given the above config:

```bash
% etlogd -config etlogd.ini -database project1
```

the `project1` database will be used rather than the default (i.e. `etlog`).

## Issues

- The `http` server is the only server that currently works

## Notes

- The `http` server requires a `Content-Type` header of `application/json` and the body must be valid JSON

## TODOs

- Define and document message boundaries for TCP and UDP streams
- Implement the parsing for TCP and UDP streams
- Improve usage text for options
    - Show the defaults for each option

## Ideas / Questions

- Normalization of sources, targets, and scripts to normalize into separate collections
    - `_id` generation for sources, targets and scripts for normalizing into separate collections
    - This would reduce a lot redundancy for row/line/column-level logging
    - Flag for whether to keep the original document in place rather than replacing it with the `_id`

