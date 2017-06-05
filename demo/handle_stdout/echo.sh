#测试 shift 命令(x_shift.sh)
until [ $# -eq 0 ]
do
echo "第一个参数为: $1 参数个数为: $#"
shift
done