# docker-build-nodot-git

Sourcehut refuses to work with clone URLs ending with `.git`. Docker can't handle git URLs not ending with `.git`.

```docker-compose
service:
  xyz:
    build: https://srhtgit.c7.ee/~foo/bar.git#main
```

^ gets transformed to `https://git.sr.ht/~foo/bar` (`#main` is removed by Docker before cloning)
