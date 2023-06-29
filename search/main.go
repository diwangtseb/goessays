package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"
)

type Document struct {
	ID   string
	Text string
}

func main() {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		log.Fatal(err)
	}

	docs := []Document{
		{ID: "1", Text: "Hello world"},
		{ID: "2", Text: "Foo bar"},
		{ID: "3", Text: "world Hello"},
	}

	for _, doc := range docs {
		if err := index.Index(doc.ID, doc); err != nil {
			log.Fatal(err)
		}
	}

	query := bleve.NewMatchQuery("world")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		log.Fatal(err)
	}

	for _, hit := range searchResults.Hits {
		doc, err := index.Document(hit.ID)
		if err != nil {
			log.Fatal(err)
		}

		var document Document
		if err := mapDocument(doc, &document); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %s, Text: %s\n", document.ID, document.Text)
	}
}

func setField(target interface{}, fieldName string, value interface{}) error {
	targetValue := reflect.ValueOf(target).Elem()
	fieldValue := targetValue.FieldByName(fieldName)

	if !fieldValue.IsValid() {
		return fmt.Errorf("field %s doesn't exist", fieldName)
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("field %s is not settable", fieldName)
	}

	fieldType := fieldValue.Type()
	valueType := reflect.TypeOf(value)

	// Convert value to fieldType
	if valueType.ConvertibleTo(fieldType) {
		fieldValue.Set(reflect.ValueOf(value).Convert(fieldType))
	} else {
		return fmt.Errorf("cannot assign value of type %s to field %s with type %s", valueType, fieldName, fieldType)
	}

	return nil
}

func mapDocument(doc *document.Document, target interface{}) error {
	for _, field := range doc.Fields {
		if err := setField(target, field.Name(), field.Value()); err != nil {
			return err
		}
	}
	return nil
}
