package cache

import "testing"

func TestBadgerCache_Has(t *testing.T) {
	err := testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("foo")
	if err == nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache, it should not be there")
	}

	err = testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	inCache, err = testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cache, it should be there")
	}

	err = testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}
}

func TestBadgerCache_Get(t *testing.T) {
	err := testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	val, err := testBadgerCache.Get("foo")
	if err != nil {
		t.Error(err)
	}

	if val != "bar" {
		t.Errorf(
			"did not get correct value in cache, got %s, instead of bar",
			val)
	}
}

func TestBadgerCache_Forget(t *testing.T) {
	err := testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}
	err = testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("foo")
	if err == nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache, it should not be there")
	}
}

func TestBadgerCache_Empty(t *testing.T) {
	err := testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Empty()
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("foo")
	if err == nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache, it should not be there")
	}

}

func TestBadgerCache_EmptyByMatch(t *testing.T) {
	err := testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Set("foo2", "bar2")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Set("alpha", "omega")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.EmptyByMatch("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("foo")
	if err == nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache, it should not be there")
	}

	inCache, err = testBadgerCache.Has("foo2")
	if err == nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo2 found in cache, it should not be there")
	}

	inCache, err = testBadgerCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("alpha not found in cache, it should   be there")
	}

}
