db = db.getSiblingDB('admin');
 db.auth("root", "example");
 db = db.getSiblingDB('USERS');
 db.createUser({
'user': "dbUser",
'pwd': "dbPwd",
'roles': [{
    'role': 'dbOwner',
    'db': 'USERS'}]});
 db.createCollection('user_collection');
 db.createCollection('sequence')
 db.sequence.insert({"_id":"userid","sequence_value":1})