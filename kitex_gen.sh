kitex -module Douyin_Demo ./idl/"$1".proto

mkdir service/"$1"
cd service/"$1"
kitex -module Douyin_Demo -service "$1" -use Douyin_Demo/kitex_gen/ -I ../../idl/ ../../idl/"$1".proto

go mod tidy