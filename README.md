[![BCH compliance](https://bettercodehub.com/edge/badge/OdysseyMomentum/NIUTeam?branch=main&token=b56aa973546fcea86a3885958be946bfce4c025f)](https://bettercodehub.com/) For more info, click [here](https://odysseymomentum.github.io).

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
