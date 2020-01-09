# Aloesïa Management

![Github Actions Release status](https://github.com/SkYNewZ/aloesia/workflows/release/badge.svg)
![Github Actions Build status](https://github.com/SkYNewZ/aloesia/workflows/build/badge.svg)

## Needed environment variables

- `GOOGLE_CLOUD_PROJECT` : used to access Google Firestore (automatically provided by AppEngine)
- `GOOGLE_APPLICATION_CREDENTIALS` : used to log in to GCP (automatically provided by AppEngine)

## Docker

```bash
$ git clone --depth=1 --branch=master https://github.com/SkYNewZ/aloesia.git aloesia
$ cd $_
$ docker build --tag aloesia:latest .
$ docker run -it --rm -e GOOGLE_CLOUD_PROJECT -e GOOGLE_APPLICATION_CREDENTIALS -p 8080:8080 aloesia:latest
```
