# magiclink

[![release](https://github.com/thisdougb/magiclink/actions/workflows/release.yaml/badge.svg)](https://github.com/thisdougb/magiclink/actions/workflows/release.yaml)

A Go package that implements magic-link login functionality, using Redis as a backend.

### Kubernetes
Get the magiclink service and deployment spec:
```
$ curl -sSLO https://raw.githubusercontent.com/thisdougb/kubernetes/magiclink.yaml
$ kubectl apply -f magiclink.yaml    
service/magiclink created
deployment.apps/magiclink created
$
$ kubectl get po
NAME                         READY   STATUS    RESTARTS   AGE
magiclink-7c76857f8c-jghcr   1/1     Running   0          26s
redis-df87ffcd6-vqh2b        1/1     Running   0          14m
$
$ kubectl logs magiclink-7c76857f8c-jghcr
2021/11/07 12:45:55 webserver.Start(): listening on port 8080
$
```
For more details on running magiclink as a k8s service, see [these instructions](https://github.com/thisdougb/magiclink/tree/main/kubernetes)

# Configuration
Configuration is via env vars, which is easy for container environments.

All env vars are prefixed with *MAGICLINK_* to avoid clashes with other services.

Env Var Name| Default| Description
----|---|---
MAGICLINK_API_PORT| 8080 | The web server listens on this port.
MAGICLINK_REDIS_HOST | redis | Host name for the redis instance.
MAGICLINK_REDIS_PORT | 6379 | Port of the redis instance.
MAGICLINK_REDIS_KEY_PREFIX | magiclink | All redis database keys are prefixed with this string, to keep things isolated.
MAGICLINK_MAGICLINK_LENGTH | 64 | Length of the magiclink id string.
MAGICLINK_MAGICLINK_EXPIRES_MINS | 15 | Expiry time of magic link IDs, in minutes.
MAGICLINK_SESSION_NAME | MagicLinkSession | Cookie session ID name.
MAGICLINK_SESSION_ID_LENGTH | 64 | Length of cookie session ID string.
MAGICLINK_SESSION_EXPIRES_MINS | 10080 | Expire time of session ID, in minutes.
MAGICLINK_RATE_LIMIT_MAX_SEND_REQUESTS | 3 | Maximum number of send requests per email.
MAGICLINK_RATE_LIMIT_TIME_PERIOD_MINS | 15 | Time period over which max requests are limited, in minutes.

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
