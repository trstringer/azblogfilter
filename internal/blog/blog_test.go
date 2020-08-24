package blog

import "testing"

func TestCategoryEquals(t *testing.T) {
	var c1 Category = "hello"
	var c2 Category = "HeLlO"

	if !c1.Equals(c2) {
		t.Fatalf("Expected %s to equal %s", c1, c2)
	}
}

func TestCategoryNotEquals(t *testing.T) {
	var c1 Category = "hello"
	var c2 Category = "goodbye"

	if c1.Equals(c2) {
		t.Fatalf("Expected %s to not equal %s", c1, c2)
	}
}

func TestTitleHasKeyword(t *testing.T) {
	var title Title = "This is my title with an ElePhant in it"
	var keyword Keyword = "elephant"

	if !title.HasKeyword(keyword) {
		t.Fatalf("Expected title '%s' to have keyword '%s'", title, keyword)
	}
}

func TestTitleHasUppercaseKeyword(t *testing.T) {
	var title Title = "This is my title with an ElePhant in it"
	var keyword Keyword = "ELEPHANT"

	if !title.HasKeyword(keyword) {
		t.Fatalf("Expected title '%s' to have keyword '%s'", title, keyword)
	}
}

func TestTitleDoesNotHaveKeyword(t *testing.T) {
	var title Title = "This is my title with an ElePhant in it"
	var keyword Keyword = "giraffe"

	if title.HasKeyword(keyword) {
		t.Fatalf("Expected title '%s' to not have keyword '%s'", title, keyword)
	}
}

func TestCategoryContainsCategory(t *testing.T) {
	actualCategories := Categories{"category1", "category2", "category3"}
	post := Post{Categories: actualCategories}
	var desiredCategory Category = "category2"

	if !post.ContainsCategory(desiredCategory) {
		t.Fatalf(
			"Expected post categories %v to contain category %s",
			post.Categories,
			desiredCategory,
		)
	}
}

func TestPostDoesNotContainCategory(t *testing.T) {
	actualCategories := []Category{"category1", "category2", "category3"}
	post := Post{Categories: actualCategories}
	var desiredCategory Category = "category4"

	if post.ContainsCategory(desiredCategory) {
		t.Fatalf(
			"Expected post categories %v to not contain category %s",
			post.Categories,
			desiredCategory,
		)
	}
}
