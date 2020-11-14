# NIUTeam
Real Estate - Give Buildings a Personality

## Backend
To compile the backend run:

```
make
```

Upload the binaries as AWS Lambda functions. 


## Frontend

Note we use an Auth0 free tier as an OICD provider, the frontend expects a file: auth0.config.json in the root directy with the following fields:

```
{
  "domain": "",
  "clientId": "",
  "cacheLocation": "",
  "audience": ""
}
```
