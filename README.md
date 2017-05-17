# usage:

./deadpbrewery $timestamp $snapshot $db-name <br />

# brew db from json dump:

curl "http://127.0.0.1:80/api/snapshots?timestamp=<unixtime>" > file <br />
mkdir slices; python slicer.py < file <br />
for s in `ls slices/`; do ./deadprewery $s slices/$s m30.db; done <br />

# push new snapshot to db:

curl [authoruized-reques] > new-snapshot <br />
./deadpdbbrewery $(date +%s) ./new-snapshot m30.db <br />
