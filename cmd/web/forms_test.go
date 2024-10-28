package main

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	t.Run("missing field", func(t *testing.T) {
		form := NewForm(nil)
		has := form.Has("test")
		if has {
			t.Errorf("Form.Has() = %v, want %v", has, false)
		}
	})

	t.Run("existing field", func(t *testing.T) {
		postedData := url.Values{}
		postedData.Add("test", "data")
		form := NewForm(postedData)
		has := form.Has("test")
		if !has {
			t.Errorf("Form.Has() = %v, want %v", has, true)
		}
	})
}

func TestForm_Required(t *testing.T) {
	t.Run("Missing required fields", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/test", nil)
		form := NewForm(req.PostForm)
		form.Required("a", "b", "c")

		if form.Valid() {
			t.Error("Form shows valid with missing required fields")
		}
	})

	t.Run("Required fields", func(t *testing.T) {
		postedData := url.Values{}
		postedData.Add("test", "data")
		postedData.Add("test2", "data")
		postedData.Add("test3", "data")

		req := httptest.NewRequest("POST", "/test", nil)
		req.PostForm = postedData

		form := NewForm(req.PostForm)
		form.Required("test", "test2", "test3")

		if !form.Valid() {
			t.Error("Form shows invalid with all required fields")
		}
	})
}

func TestForm_Check(t *testing.T) {
	t.Run("Check failed", func(t *testing.T) {
		form := NewForm(nil)
		form.Check(false, "test", "test failed")

		if form.Valid() {
			t.Error("Form shows valid with failed check")
		}
	})

	t.Run("Check passed", func(t *testing.T) {
		postedData := url.Values{}
		postedData.Add("test", "data")

		form := NewForm(postedData)
		form.Check(true, "test", "test failed")

		if !form.Valid() {
			t.Error("Form shows invalid with passed check")
		}
	})
}

func TestForm_ErrorGet(t *testing.T) {
	t.Run("Get error message", func(t *testing.T) {
		form := NewForm(nil)
		form.Check(false, "test", "test error")
		s := form.Errors.Get("test")

		if len(s) == 0 {
			t.Error("Should have an error returned from Get, but not found")
		}
	})

	t.Run("Get empty error message", func(t *testing.T) {
		form := NewForm(nil)
		s := form.Errors.Get("test")

		if len(s) != 0 {
			t.Error("Should not have an error returned from Get, but found")
		}
	})
}
