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
go get github.com/cbmi/etlogd
go install github.com/cbmi/etlogd
```

## Setup

Requires a mongodb instance running on the default port (this will be configurable soon).

## Usage

The default is an HTTP server which handles POST reqeusts.

```bash
etlogd
Listening on http://:4000...
```

The `Content-Type` header is required and the body must be valid JSON, for example:

```
POST HTTP/1.1
Content-Type: application/json

{
    "timestamp": "2013-08-13T05:43:03.32344",
    "action": "update",
    "script": {
        "uri": "https://github.com/cbmi/project/blob/master/parse-users.py",
        "version": "a32f87cb"
    },
    "source": {
        "type": "delimited",
        "delimiter": ",",
        "uri": "148.29.12.100/path/to/users.csv",
        "name": "users.csv",
        "line": 5,
        "column": 4
    },
    "target": {
        "type": "relational",
        "uri": "148.29.12.101:5236",
        "database": "socialapp",
        "table": "users",
        "row": { "id": 38 },
        "column": "email"
    }
}
```

---

## Ideas / Questions

- Normalization of sources, targets, and scripts to normalize into separate collections
    - `_id` generation for sources, targets and scripts for normalizing into separate collections
    - This would reduce a lot redundancy for row/line/column-level logging
    - Flag for whether to keep the original document in place rather than replacing it with the `_id`
