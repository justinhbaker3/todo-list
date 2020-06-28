package list

import "fmt"

// List is a list of items
type List struct {
	Title string           `json:"title"`
	Items map[string]*item `json:"items"`
}

func NewList(title string) *List {
	return &List{
		Title: title,
		Items: map[string]*item{},
	}
}

// What happens if there are two items with the same title?
func (l *List) Append(i *item) error {
	if i.Title == "" {
		return fmt.Errorf("item must have a title")
	}
	l.Items[i.Title] = i
	return nil
}

func (l *List) AppendMultiple(items ...*item) error {
	for _, item := range items {
		err := l.Append(item)
		if err != nil {
			return fmt.Errorf("failed to add item %v, stopping: %s", item, err)
		}
	}
	return nil
}

func (l *List) Delete(title string) error {
	_, ok := l.Items[title]
	if !ok {
		return fmt.Errorf("item does not exist")
	}
	delete(l.Items, title)
	return nil
}
