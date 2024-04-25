package dag

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// DAGNode 代表DAG中的一个节点，包含节点的名称、依赖关系和任务
type DAGNode struct {
	Name      string
	DependsOn []string
	Task      func(ctx context.Context) error
	Executed  bool
	ExecErr   error
}

// DAGExecutor 用于执行DAG
type DAGExecutor struct {
	nodes     map[string]*DAGNode
	nodeMutex sync.Mutex
}

// NewDAGExecutor 创建一个新的DAG执行器
func NewDAGExecutor() *DAGExecutor {
	return &DAGExecutor{
		nodes: make(map[string]*DAGNode),
	}
}

// AddNode 向DAG中添加一个节点
func (de *DAGExecutor) AddNode(node *DAGNode) error {
	de.nodeMutex.Lock()
	defer de.nodeMutex.Unlock()

	if _, exists := de.nodes[node.Name]; exists {
		return errors.New("node already exists")
	}

	de.nodes[node.Name] = node
	return nil
}

// Execute 执行DAG中的所有节点
func (de *DAGExecutor) Execute(ctx context.Context) error {
	var wg sync.WaitGroup
	var errChan = make(chan error, len(de.nodes))

	for _, node := range de.nodes {
		// 为无依赖的节点启动任务
		if len(node.DependsOn) == 0 {
			wg.Add(1)
			go de.executeNode(ctx, &wg, node, errChan)
		}
	}

	// 等待所有节点执行完毕
	wg.Wait()
	close(errChan)

	// 检查错误通道
	var err error
	for e := range errChan {
		if err == nil {
			err = e
		} else {
			err = fmt.Errorf("%w; %v", err, e)
		}
	}

	return err
}

// executeNode 执行一个节点及其后续节点
func (de *DAGExecutor) executeNode(ctx context.Context, wg *sync.WaitGroup, node *DAGNode, errChan chan<- error) {
	defer wg.Done()

	// 检查依赖关系
	for _, depName := range node.DependsOn {
		depNode, exists := de.nodes[depName]
		if !exists {
			errChan <- fmt.Errorf("dependency %s not found for node %s", depName, node.Name)
			return
		}

		// 确保依赖已执行
		if !depNode.Executed {
			errChan <- fmt.Errorf("dependency %s not executed for node %s", depName, node.Name)
			return
		}

		if depNode.ExecErr != nil {
			errChan <- depNode.ExecErr
			return
		}
	}

	// 执行任务
	node.ExecErr = node.Task(ctx)
	node.Executed = true

	// 如果没有错误，启动后续节点
	if node.ExecErr == nil {
		for _, nextNode := range de.nodes {
			if contains(nextNode.DependsOn, node.Name) && !nextNode.Executed {
				wg.Add(1)
				go de.executeNode(ctx, wg, nextNode, errChan)
			}
		}
	} else {
		errChan <- node.ExecErr
	}
}

// contains 检查一个字符串是否在切片中
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
