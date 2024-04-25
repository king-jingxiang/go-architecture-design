package dag

import (
	"context"
	"fmt"
	"testing"
)

func TestDAGExecutor_executeNode(t *testing.T) {
	ctx := context.Background()

	// 创建DAG执行器
	dag := NewDAGExecutor()

	// 定义任务
	taskA := func(ctx context.Context) error {
		fmt.Println("Task A executed")
		return nil
	}

	taskB := func(ctx context.Context) error {
		fmt.Println("Task B executed")
		return nil
	}

	taskC := func(ctx context.Context) error {
		fmt.Println("Task C executed")
		return nil
	}

	// 添加节点
	dag.AddNode(&DAGNode{
		Name: "A",
		Task: taskA,
	})

	dag.AddNode(&DAGNode{
		Name:      "B",
		DependsOn: []string{"A"},
		Task:      taskB,
	})

	dag.AddNode(&DAGNode{
		Name:      "C",
		DependsOn: []string{"B"},
		Task:      taskC,
	})

	// 执行DAG
	if err := dag.Execute(ctx); err != nil {
		fmt.Println("Error executing DAG:", err)
	} else {
		fmt.Println("DAG executed successfully")
	}
}
