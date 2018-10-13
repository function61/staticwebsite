FROM scratch

CMD ["staticwebsite"]

ADD rel/staticwebsite_linux-amd64 /usr/local/bin/staticwebsite
