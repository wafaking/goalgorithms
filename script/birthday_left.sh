#!/bin/bash
read -p "enter your birth year, please: " bir_year
read -p "enter your birth month, please: " bir_month
read -p "enter your birth day, please: " bir_day

bir=$bir_year$bir_month$bir_day
now=$(date +%Y%m%d)
echo "birthday:  ==> $bir"
echo "now date ==> $now"

if (( $bir = $now ))
then
    echo "wow, happy birthday !!!"
else
    bir_stamp=$(date -d $bir +%j) 
    now_stamp=$(date -d $now +%j) 
    diff=$[$bir_stamp-$now_stamp]
    echo "It's $diff days to left for your birthday!!!"
fi

echo "方法二: ---------------"

read -p "Pleas input your birthday (MMDD, ex> 0709): " bir
now=`date +%m%d`
if [ "$bir" == "$now" ]; then
    echo "Happy Birthday to you!!!"
elif [ "$bir" -gt "$now" ]; then
    year=`date +%Y`
    total_d=$(($((`date --date="$year$bir" +%s`-`date +%s`))/60/60/24))
    echo "Your birthday will be $total_d later"
else
    year=$((`date +%Y`+1))
    otal_d=$(($((`date --date="$year$bir" +%s`-`date +%s`))/60/60/24))