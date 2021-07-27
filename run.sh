set -e
go build -o exec
./exec -start $1 -factor $2 -end $3 -folder $4
./tools/visualize_csv.py -folder $4
