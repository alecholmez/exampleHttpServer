FROM mongo

COPY users.json /users.json
CMD mongoimport --host mongo --db userStore --collection users --type json --file /users.json --jsonArray
