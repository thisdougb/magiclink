# magiclink

[![release](https://github.com/thisdougb/magiclink/actions/workflows/release.yaml/badge.svg)](https://github.com/thisdougb/magiclink/actions/workflows/release.yaml)

A Go package that implements magic-link login functionality.

### Login Request
To trigger a magic-link login request, call the /send/ url.
```
$ curl --data '{"email":"someuser@domain.com"}' -X POST http://localhost:8080/send/
OK
```

### Authentication
Using the magic link creates a session ID linked to the email address.
This is returned in a cookie to the caller, as part of an http redirect.
```
$ curl -i http://localhost:8080/auth/AlmmKroepZGnQ61RI8n2vwAZ1dUlhypji1ERGuhY1CwaKhi1fqyZUQuNSPjuavMJ
HTTP/1.1 302 Found
Content-Type: text/html; charset=utf-8
Location: /
Set-Cookie: MagicLinkSession=HqnWnEwCGNqVQjXR24iQ5u0maK8VDSpqIk4uVH2TicotPdWfr2vfeEMLDaMvfX0o; Path=/; Expires=Sat, 30 Oct 2021 12:30:25 GMT; SameSite=Strict
Date: Sat, 23 Oct 2021 12:30:25 GMT
Content-Length: 24
```

### Kubernetes
To run magiclink as a k8s service, see [these instructions](https://github.com/thisdougb/magiclink/tree/main/kubernetes)
