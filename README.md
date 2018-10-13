[![Build Status](https://img.shields.io/travis/function61/staticwebsite.svg?style=for-the-badge)](https://travis-ci.org/function61/staticwebsite)
[![Download](https://img.shields.io/docker/pulls/fn61/staticwebsite.svg?style=for-the-badge)](https://hub.docker.com/r/fn61/staticwebsite/)

Just a simple static HTTP server container image.


Instructions
------------

When basing your image from this `FROM fn61/staticwebsite:...` in `Dockerfile`, add your
content to `/www`:

```
ADD index.html /www/index.html
```

Would become a valid "hello world" image.
