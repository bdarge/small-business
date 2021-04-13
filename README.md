## Small business app

A sample small business app to track order, and quote. 

The technology stacks are:
- api in golang,
- database in mysql
- UI in angular

To run in development:
```
docker-compose up -d
```

To run in production:
```production
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build
```

Some angular UI features copied and repurpose from [here](https://https://github.com/tomastrajan/angular-ngrx-material-starter)


