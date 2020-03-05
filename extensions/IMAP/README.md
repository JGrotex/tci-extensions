# IMAP Extension
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This activity allows you to:
- can be used in a Timer Flogo Flow e.g. all 30sec.
- retrieve a Email from an IMAP Server

How it is working:
- Login to IMAP Server
- read first Email for specified mailbox e.g. default "INBOX"
- Data extract: 
  - Email Date, 
  - Email To, 
  - Email From, 
  - Email Subject, 
  - Email Body
- Flag indicate that there was a new email:
  - "none", no new Email
  - "error", a error occurred
  - "new", new Email to process in Flogo Flow!
- selected Email get deleted.

> WARNING :: Any Email get's deleted, after it is successfully read for the mailbox!

## Parameter Details
see Activity Tester      

## Third-party libraries used
- https://github.com/emersion/go-imap

## Background
This Implementation is inspired by TCI use cases where a Email could e.g. create a new case in TIBCO Cloud Live Apps.