# IMAP Extension
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> WARNING :: Draft Implementation not finalized yet.

This activity allows you to:
- retrieve a Email from an IMAP Server

Issues to be finalized:
- just retrieve 'unread' Emails
- mark read Emails as 'read'
- just get the oldest Email from Inbox
- Optional: handle Email attachments

First Test shows
- connection to IMAP is working
- Email Subject could be extracted

## Parameter Details
see Activity Tester      

## Third-party libraries used
- https://github.com/emersion/go-imap

## Background
This Implementation is inspired by TCI use cases where a Email could e.g. create a new case in TIBCO Cloud Live Apps.