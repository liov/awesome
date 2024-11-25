# for
for file in $(ls /etc);do
  echo "$file"
done

for str in 'This is a string'
do
    echo $str
done

for((i=1;i<=5;i++));do
    echo "这是第 $i 次调用";
done;
#while
int=1
while(( $int<=5 ))
do
    echo $int
    let "int++"
done
#无限循环
while true;
do
  xxx
done

for (( ; ; ));
do
  xxx
done
#util
a=0

until [ ! $a -lt 10 ]
do
   echo $a
   a=`expr $a + 1`
done