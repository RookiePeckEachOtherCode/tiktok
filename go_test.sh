#!/bin/bash

function clean_coverage() {
	if [ -f coverage.out ]; then
		rm coverage.out
	fi
}
# 设置默认参数
VERBOSE=""
COVER=false

# 解析输入参数
while [[ $# -gt 0 ]]; do
	case $1 in
	-v | -verbose)
		VERBOSE="-v"
		shift
		;;
	--cover)
		COVER=true
		shift
		;;
	*)
		echo "Unknown parameter: $1"
		exit 1
		;;
	esac
done

TEST_DIRS="./test/dao_test/"

# 构建测试命令
CMD="go test $TEST_DIRS"

if [ $VERBOSE ]; then
	CMD="$CMD $VERBOSE"
fi

# 仅在文件不存在时生成coverprofile
if [ $COVER = true ]; then
	if [ ! -f coverage.out ]; then
		CMD="$CMD -coverprofile=coverage.out"
	fi
fi

# 运行测试
$CMD

# 显示覆盖率报告
if [ -f coverage.out ]; then
	go tool cover -func coverage.out
	rm coverage.out
fi
