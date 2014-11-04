argo-test
=========

Test app for argo

### POST

    curl -d '{"name":"Second Company"}' localhost:3000/companies

    curl -d '{"id":1,"name":"First Company"}' localhost:3000/companies

    curl -d '{"hat":"Second Company"}' localhost:3000/companies

    curl -d 'gakoepf{fes' localhost:3000/companies

### PATCH

    curl -d '{"name":"Whatever Company"}' -X PATCH localhost:3000/companies/1

    curl -d 'gakoepf{fes' -X PATCH localhost:3000/companies/1

### DELETE

    curl -X DELETE localhost:3000/companies/1
