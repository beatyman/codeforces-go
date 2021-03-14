// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`"bank"`, `"kanb"`, 
			`true`,
		},
		{
			`"attack"`, `"defend"`, 
			`false`,
		},
		{
			`"kelb"`, `"kelb"`, 
			`true`,
		},
		{
			`"abcd"`, `"dcba"`, 
			`false`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, areAlmostEqual, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-232/problems/check-if-one-string-swap-can-make-strings-equal/
