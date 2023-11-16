FROM mysql

ADD init.sql /docker-entrypoint-initdb.d
ADD setdata.sql /docker-entrypoint-initdb.d


