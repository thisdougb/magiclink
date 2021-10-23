# magiclink

[![release](https://github.com/thisdougb/magiclink/actions/workflows/release.yaml/badge.svg)](https://github.com/thisdougb/magiclink/actions/workflows/release.yaml)

A Go package that implements magic-link login functionality.

### Login Request
To trigger a magic-link login request, call the /send/ url.
```
$ curl --data '{"email":"someuser@domain.com"}' -X POST http://localhost:8080/send/
OK
```
An expiring key is created, with the requesting email address as data.
This will be used when authenticating the login request, when the link in the login email is clicked.
```
redis> keys magiclink:auth:id*
1) "magiclink:auth:id:AlmmKroepZGnQ61RI8n2vwAZ1dUlhypji1ERGuhY1CwaKhi1fqyZUQuNSPjuavMJ"
2) "magiclink:auth:id:p2FJt1iUXKZU9OjIrzRjrbgr1Lj1momj7zKmm0wgSGPbRXnUcJo6IUuo4Wuxl2tW"
3) "magiclink:auth:id:sVm4ECyEaec1HYBI9yP8nqLPMP1f8PXSar2O1ZN5HzyNn1WCr5Zx7JuInMUB8o8t"

redis> get magiclink:auth:id:AlmmKroepZGnQ61RI8n2vwAZ1dUlhypji1ERGuhY1CwaKhi1fqyZUQuNSPjuavMJ
"someuser@domain.com"
redis> ttl magiclink:auth:id:AlmmKroepZGnQ61RI8n2vwAZ1dUlhypji1ERGuhY1CwaKhi1fqyZUQuNSPjuavMJ
(integer) 215
```
A job is added to the send queue, to be processed by some external smtp-sender process.
This package does not send smtp emails.
```
redis> lrange "magiclink:queue:send" 0 1
1) "{\"Email\":\"someuser@domain.com\",\"MagicLinkID\":\"AlmmKroepZGnQ61RI8n2vwAZ1dUlhypji1ERGuhY1CwaKhi1fqyZUQuNSPjuavMJ\",\"Timestamp\":1634976117}"
```
