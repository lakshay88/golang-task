This is golang project Based on configuration.
you can updated configuration on config.yaml file, 
i have give support for multiple db, you can implemnet using database interface

Created multiple folders for each settings

Will be adding docker file in future.

currently it is  supporting 2 API only 

Assuming DB is already created 
DB looks like - 
1. Person -
  create table person(id int not null key auto_increment, name varchar(255), age int);
  insert into person(id, name, age) values (1, "mike", 31), (2, "John", 20), (3, "Joseph", 20);
  You can assume the following data in person:
  Name, age, id
  mike , 31, 1
  John, 45, 2
  Joseph, 20, 3

2. Phone -  
  create table phone(id int not null key auto_increment, number varchar(255), person_id int);
  insert into phone(id, person_id, number) values (1,1, "444-444-4444"), (8,2, "123-444-7777"),
  (3,3, "445-222-1234");
  You can assume the following data in phone:
  person_id, id, number
  1,1, 444-444-4444
  2,8, 123-444-7777
  3,3, 445-222-1234

3. Address 
  create table address(id int not null key auto_increment, city varchar(255), state varchar(255),
  street1 varchar(255), street2 varchar(255), zip_code varchar(255));
  insert into address(id , city , state , street1 , street2 , zip_code ) values (1,"Eugene", "OR", "111
  Main St", "", "98765"), (2, "Sacramento", "CA", "432 First St", "Apt 1", "22221"), (3, "Austin",
  "TX", "213 South 1st St", "", "78704");
  You can assume the following data in address:
  id, city, state, street1, street2, zip_code
  1,Eugene, OR, "111 Main St", "", 98765
  2, Sacramento, CA, "432 First St", "Apt 1", 22221
  3, Austin, TX, "213 South 1st St", "", 78704

4. Address_join
  create table address_join(id int not null key auto_increment, person_id int, address_id int);
  insert into address_join(id, person_id, address_id) values (1,1,3),(2,2,1),(3,3,2);
  You can assume the following data in address_join:
  id, person_id, address_id
  1,1,3
  2,2,1
  3,3,2