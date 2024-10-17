declare -a array=('a' '!' '\' '"')
arraylength=${#array[@]}

for (( i=0; i<${arraylength}; i++ ));
do
  touch ${array[$i]}
  echo creating ${array[$i]}
done

mkdir '`'

cp ! '`'/!

if [ "$MOVE_A" = "yes" ] ; then
    mv a '`'
elif [ "$MOVE_A" = "no" ] ; then
    rm a
else
    echo "Undefined"
fi
