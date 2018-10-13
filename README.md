[![Build Status](https://img.shields.io/travis/function61/holepunch-client.svg?style=for-the-badge)](https://travis-ci.org/function61/holepunch-client)
[![Download](https://img.shields.io/bintray/v/function61/holepunch-client/main.svg?style=for-the-badge&label=Download)](https://bintray.com/function61/holepunch-client/main/_latestVersion#files)

Just a simple static HTTP server container image.


Instructions
------------

When basing your image from this `FROM fn61/staticwebsite:...` in `Dockerfile`, add your
content to `/www`:

```
ADD index.html /www/index.html
```

Would become a valid "hello world" image.
