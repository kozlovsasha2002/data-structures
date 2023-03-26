package graph

import (
	"testing"
)

func TestGraph_AddEdge(t *testing.T) {
	t.Run("adding one element", func(t *testing.T) {
		gr := New()
		gr.AddEdge("Казань", "Москва", 840)

		actualResult := gr.IsExist(&Edge{"Казань", "Москва", 840})
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})
}

func TestGraph_AddEdges(t *testing.T) {
	t.Run("adding zero elements", func(t *testing.T) {
		gr := New()
		var edges []*Edge

		gr.AddEdges(edges)
		actualResult := true
		expectedResult := true

		for _, edge := range edges {
			if !gr.IsExist(edge) {
				actualResult = false
			}
		}

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})

	t.Run("adding multiple elements", func(t *testing.T) {
		gr := New()
		edges := []*Edge{
			CreateEdge("Заинск", "Казань", 240),
			CreateEdge("Казань", "Нижний Новгород", 410),
			CreateEdge("Нижний Новгород", "Москва", 460),
			CreateEdge("Заинск", "Ульяновск", 470),
		}

		gr.AddEdges(edges)
		actualResult := true
		expectedResult := true

		for _, edge := range edges {
			if !gr.IsExist(edge) {
				actualResult = false
			}
		}

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})
}

func TestGraph_ChangeWeightInEdge(t *testing.T) {
	t.Run("graph consists target edge", func(t *testing.T) {
		gr := New()
		edges := []*Edge{
			CreateEdge("Заинск", "Казань", 240),
			CreateEdge("Казань", "Нижний Новгород", 410),
			CreateEdge("Нижний Новгород", "Москва", 460),
			CreateEdge("Заинск", "Ульяновск", 470),
		}
		gr.AddEdges(edges)

		gr.ChangeWeightInEdge("Казань", "Нижний Новгород", 500)
		actualResult := gr.IsExist(&Edge{"Казань", "Нижний Новгород", 500})
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})

	t.Run("graph doesn't consist target edge", func(t *testing.T) {
		gr := New()
		edges := []*Edge{
			CreateEdge("Заинск", "Казань", 240),
			CreateEdge("Казань", "Нижний Новгород", 410),
			CreateEdge("Нижний Новгород", "Москва", 460),
			CreateEdge("Заинск", "Ульяновск", 470),
		}
		gr.AddEdges(edges)

		err := gr.ChangeWeightInEdge("Казань", "Тамбов", 500)
		actualResult := ""
		if err != nil {
			actualResult = err.Error()
		}
		expectedResult := EdgeNotExists

		if actualResult != expectedResult {
			t.Errorf("expected result = %v, and actual result = %v", expectedResult, actualResult)
		}
	})
}

func TestGraph_Clear(t *testing.T) {
	t.Run("adding zero elements", func(t *testing.T) {
		gr := New()

		gr.Clear()
		actualResult := len(gr.listOfEdges) == 0
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})

	t.Run("adding multiple elements", func(t *testing.T) {
		gr := New()
		edges := []*Edge{
			CreateEdge("Заинск", "Казань", 240),
			CreateEdge("Казань", "Нижний Новгород", 410),
			CreateEdge("Нижний Новгород", "Москва", 460),
			CreateEdge("Заинск", "Ульяновск", 470),
		}

		gr.AddEdges(edges)
		gr.Clear()
		actualResult := len(gr.listOfEdges) == 0
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})
}

func TestGraph_FindShortestPathBFS(t *testing.T) {
	t.Run("graph doesn't consists loops", func(t *testing.T) {
		gr := New()

		edges := []*Edge{
			CreateEdge("Заинск", "Казань", 240),
			CreateEdge("Казань", "Нижний Новгород", 410),
			CreateEdge("Нижний Новгород", "Москва", 460),
			CreateEdge("Заинск", "Ульяновск", 470),
			CreateEdge("Ульяновск", "Саранск", 232),
			CreateEdge("Саранск", "Нижний Новгород", 287),
			CreateEdge("Ульяновск", "Саратов", 460),
			CreateEdge("Саратов", "Нижний Новгород", 549),
			CreateEdge("Нижний Новгород", "Заинск", 230),
			CreateEdge("Нижний Новгород", "Тамбов", 150),
			CreateEdge("Казань", "Тамбов", 150),
		}

		gr.AddEdges(edges)

		minPath, _ := gr.FindShortestPathBFS("Заинск", "Москва")
		expectedMinPath := 3

		if minPath != expectedMinPath {
			t.Errorf("expected min path = %v and actual min paht = %v", expectedMinPath, minPath)
		}
	})
}

func TestGraph_IsExist(t *testing.T) {
	t.Run("graph consists target edge", func(t *testing.T) {
		gr := New()
		edges := []*Edge{
			CreateEdge("Заинск", "Казань", 240),
			CreateEdge("Казань", "Нижний Новгород", 410),
			CreateEdge("Нижний Новгород", "Москва", 460),
			CreateEdge("Заинск", "Ульяновск", 470),
		}
		gr.AddEdges(edges)

		targetEdge := CreateEdge("Казань", "Нижний Новгород", 410)
		actualResult := gr.IsExist(targetEdge)
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})

	t.Run("graph doesn't consist target edge", func(t *testing.T) {
		gr := New()
		edges := []*Edge{
			CreateEdge("Заинск", "Казань", 240),
			CreateEdge("Казань", "Нижний Новгород", 410),
			CreateEdge("Нижний Новгород", "Москва", 460),
			CreateEdge("Заинск", "Ульяновск", 470),
		}
		gr.AddEdges(edges)

		targetEdge := CreateEdge("Казань", "Тамбов", 140)
		actualResult := gr.IsExist(targetEdge)
		expectedResult := false

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})
}

func TestGraph_RemoveEdgeBetweenNodes(t *testing.T) {
	t.Run("graph consists target edge", func(t *testing.T) {
		gr := New()
		edges := []*Edge{
			CreateEdge("Заинск", "Казань", 240),
			CreateEdge("Казань", "Нижний Новгород", 410),
			CreateEdge("Нижний Новгород", "Москва", 460),
			CreateEdge("Заинск", "Ульяновск", 470),
		}
		gr.AddEdges(edges)

		targetEdge := CreateEdge("Казань", "Нижний Новгород", 410)
		gr.RemoveEdgeBetweenNodes(targetEdge.start, targetEdge.end)
		actualResult := gr.IsExist(targetEdge)
		expectedResult := false

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})

	t.Run("graph doesn't consist target edge", func(t *testing.T) {
		gr := New()
		edges := []*Edge{
			CreateEdge("Заинск", "Казань", 240),
			CreateEdge("Казань", "Нижний Новгород", 410),
			CreateEdge("Нижний Новгород", "Москва", 460),
			CreateEdge("Заинск", "Ульяновск", 470),
		}
		gr.AddEdges(edges)

		targetEdge := CreateEdge("Серпухов", "Тамбов", 140)
		gr.RemoveEdgeBetweenNodes(targetEdge.start, targetEdge.end)
		actualResult := gr.IsExist(targetEdge)
		expectedResult := false

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})
}

func TestGraph_RemoveNode(t *testing.T) {
	t.Run("graph consists target node", func(t *testing.T) {
		gr := New()
		edges := []*Edge{
			CreateEdge("Заинск", "Казань", 240),
			CreateEdge("Казань", "Нижний Новгород", 410),
			CreateEdge("Нижний Новгород", "Москва", 460),
			CreateEdge("Заинск", "Ульяновск", 470),
		}
		gr.AddEdges(edges)

		targetNode := "Казань"
		gr.RemoveNode(targetNode)
		actualResult := true
		expectedResult := true

		for _, edges := range gr.listOfEdges {
			for _, edge := range edges {
				if edge.start == targetNode || edge.end == targetNode {
					actualResult = false
				}
			}
		}

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})

	t.Run("graph doesn't consist target node", func(t *testing.T) {
		gr := New()
		edges := []*Edge{
			CreateEdge("Заинск", "Казань", 240),
			CreateEdge("Казань", "Нижний Новгород", 410),
			CreateEdge("Нижний Новгород", "Москва", 460),
			CreateEdge("Заинск", "Ульяновск", 470),
		}
		gr.AddEdges(edges)

		targetNode := "Пермь"
		gr.RemoveNode(targetNode)
		actualResult := false
		expectedResult := false

		for _, edges := range gr.listOfEdges {
			for _, edge := range edges {
				if edge.start == targetNode || edge.end == targetNode {
					actualResult = true
				}
			}
		}

		if actualResult != expectedResult {
			t.Errorf("expected result = %t, and actual result = %t", expectedResult, actualResult)
		}
	})
}
