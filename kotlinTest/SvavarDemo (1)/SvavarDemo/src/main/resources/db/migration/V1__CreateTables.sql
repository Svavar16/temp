CREATE TABLE testing (
    Id UUID primary key not null,
    Name varchar(255)
);

CREATE TABLE test_item(
    Id UUID primary key not null,
    Name varchar(255),
    Testing_id UUID
);

insert into testing (id, name) values ('e73fdcae-daec-484d-bbc6-05e362fade55', 'Svavar Páll');

insert into test_item (id, name, testing_id) values ('8b2f401c-3b0a-44a0-93a6-a087ebfd8e86','Svavars Test','e73fdcae-daec-484d-bbc6-05e362fade55');
insert into test_item (id, name, testing_id) values ('c6603f2f-edc5-489f-924b-ee119fe5110d','Svavars Test 2','e73fdcae-daec-484d-bbc6-05e362fade55');
