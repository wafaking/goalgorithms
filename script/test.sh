#！/bin/bash
sum=0
read -p "Enter the end num: " num
for ((i=0; i<=$num;i++))
do
    sum=$[$sum+$i]
done
echo "sum: $sum"