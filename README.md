Package mold
============
![Project status](https://img.shields.io/badge/version-4.5.1-green.svg)
[![Build Status](https://travis-ci.org/go-playground/mold.svg?branch=v2)](https://travis-ci.org/go-playground/mold)
[![Coverage Status](https://coveralls.io/repos/github/go-playground/mold/badge.svg?branch=v2)](https://coveralls.io/github/go-playground/mold?branch=v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/mold)](https://goreportcard.com/report/github.com/go-playground/mold)
[![GoDoc](https://godoc.org/github.com/go-playground/mold?status.svg)](https://godoc.org/github.com/go-playground/mold)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Package mold is a general library to help modify or set data within data structures and other objects.

How can this help me you ask, please see the examples [here](_examples/full/main.go)

Requirements
------------
- Go 1.18+

Installation
------------

Use go get.
```shell
go get -u github.com/go-playground/mold/v4
```

Examples
----------
| Example                          | Description                                                        |
|----------------------------------|--------------------------------------------------------------------|
| [simple](_examples/mold/main.go) | A basic example with custom function.                              |
| [full](_examples/full/main.go)   | A more real life example combining the usage of multiple packages. |


Modifiers
----------
These functions modify the data in-place.

| Name                | Description                                                                               |
|---------------------|-------------------------------------------------------------------------------------------|
| camel               | Camel Cases the data.                                                                     |
| default             | Sets the provided default value only if the data is equal to it's default datatype value. |
| empty               | Sets the field equal to the datatype default value. e.g. 0 for int.                       |
| lcase               | lowercases the data.                                                                      |
| ltrim               | Trims spaces from the left of the data provided in the params.                            |
| rtrim               | Trims spaces from the right of the data provided in the params.                           |
| set                 | Set the provided value.                                                                   |
| slug                | Converts the field to a [slug](https://github.com/gosimple/slug)                          |
| snake               | Snake Cases the data.                                                                     |
| strip_alpha         | Strips all ascii characters from the data.                                                |
| strip_alpha_unicode | Strips all unicode characters from the data.                                              |
| strip_num           | Strips all ascii numeric characters from the data.                                        |
| strip_num_unicode   | Strips all unicode numeric characters from the data.                                      |
| strip_punctuation   | Strips all ascii punctuation from the data.                                               |
| title               | Title Cases the data.                                                                     |
| tprefix             | Trims a prefix from the value using the provided param value.                             |
| trim                | Trims space from the data.                                                                |
| tsuffix             | Trims a suffix from the value using the provided param value.                             |
| ucase               | Uppercases the data.                                                                      |
| ucfirst             | Upper cases the first character of the data.                                              |

**Special Notes:**
`default` and `set` modifiers are special in that they can be used to set the value of a field or underlying type information or attributes and both use the same underlying function to set the data.

Setting a Param will have the following special effects on data types where it's not just the value being set:
- Chan - param used to set the buffer size, default = 0.
- Slice - param used to set the capacity, default = 0.
- Map - param used to set the size, default = 0.
- time.Time - param used to set the time format OR value, default = time.Now(), `utc` = time.Now().UTC(), other tries to parse using RFC3339Nano and set a time value.

Scrubbers
----------
These functions obfuscate the specified types within the data for pii purposes.

| Name   | Description                                                       |
|--------|-------------------------------------------------------------------|
| emails | Scrubs multiple emails from data.                                 |
| email  | Scrubs the data from and specifies the sha name of the same name. |
| text   | Scrubs the data from and specifies the sha name of the same name. |
| name   | Scrubs the data from and specifies the sha name of the same name. |
| fname  | Scrubs the data from and specifies the sha name of the same name. |
| lname  | Scrubs the data from and specifies the sha name of the same name. |


Special Information
-------------------
- To use a comma(,) within your params replace use it's hex representation instead '0x2C' which will be replaced while caching.

Contributing
------------
I am definitely interested in the communities help in adding more scrubbers and modifiers.
Please send a PR with tests, and preferably no extra dependencies, at lease until a solid base
has been built.

Complimentary Software
----------------------

Here is a list of software that compliments using this library post decoding.

* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.
* [form](https://github.com/go-playground/form) - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.

License
------
Distributed under MIT License, please see license file in code for more details.
