# socialloot with docker

Clone repository:
```
git clone https://github.com/KeKsBoTer/socialloot
```

Navigate into folder and build container:
```
cd socialloot
docker build -t socialloot .
```

Run docker container:
```
docker run -d -p 8080:8080 socialloot
```
Open http://localhost:8080 in browser