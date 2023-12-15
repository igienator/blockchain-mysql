#challenge:

>“There is a table in MySQL, that needs to be converted to İPFS and get a chain of
>hashes, where a table cell is a hash, a row is a chain of cell hashes, and the whole
>table is a common hash bound to İPNS hash.”

Note:

1. The task is given only on decentralized networks and blockchain. We are not
   interested in cryptography in crypto.
2. Do not use frameworks.

How to use?
Create a mysql table and populate it with queries from database.sql file.

Mysql connection needs to be configured in controllers/mysql.go lines 25-31

I've decided to hardcode every setting.
IPFS connection is defined in controllers/ipfs.go line 12.
