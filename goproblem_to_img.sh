sgffile=$(find . -name $1.sgf)
outfile=$1
shift
params=$@
echo sgffile: $sgffile
echo outfile: $outfile
echo params: $params
vars=$(sgftopng -info ${sgffile}  | grep var | sed 's/var//' | sed 's/:.*//')
echo vars: $vars
sgftopng 0-0 $params $sgffile -o go_$outfile.png
for var in $vars
do
    echo var: $var
    sgftopng 1- $params -var $var -o go_${outfile}_var_$var.png $sgffile
done
