package main

import (
  "fmt"
)

type XmlParser struct {
  Content string
  CreatedAt string
}

func (x XmlParser) Parse() string {
  return fmt.Sprintf("Parsing the content in XML: %s", x.Content)
}

func (x XmlParser) FormatDate() string {
  return fmt.Sprintf("Formatting date for XML: %s", x.CreatedAt)
}

type JsonParser struct {
  Content string
  CreatedAt string
}

func (j JsonParser) Parse() string {
  return fmt.Sprintf("Parsing the content in JSON: %s", j.Content)
}

func (j JsonParser) FormatDate() string {
  return fmt.Sprintf("Formatting date for JSON: %s", j.CreatedAt)
}

type CsvParser struct {
  Content string
  CreatedAt string
}

func (c CsvParser) Parse() string {
  return fmt.Sprintf("Parsing the content in CSV: %s", c.Content)
}

func (c CsvParser) FormatDate() string {
  return fmt.Sprintf("Formatting date for CSV: %s", c.CreatedAt)
}

type AvroParser struct {
  Content string
  CreatedAt string
}

func (a AvroParser) Parse() string {
  return fmt.Sprintf("Parsing the content in Avro: %s", a.Content)
}

func (a AvroParser) FormatDate() string {
  return fmt.Sprintf("Formatting date for Avro: %s", a.CreatedAt)
}

type Parser interface {
  Parse() string
  FormatDate() string
}

func Printer(format Parser) string {
  return format.Parse() + " - Date - " + format.FormatDate()
}

func main() {
  xmlParser := XmlParser{Content: "<title>My Title<title>", CreatedAt: "2018-01-20"}
  xmlPrinter := Printer(xmlParser)
  fmt.Println(xmlPrinter)

  jsonParser := JsonParser{Content: "{'title': 'My Title'}", CreatedAt: "2018-01-25"}
  jsonPrinter := Printer(jsonParser)
  fmt.Println(jsonPrinter)

  csvParser := CsvParser{Content: "My Content; Another Content", CreatedAt: "2018-01-25"}
  csvPrinter := Printer(csvParser)
  fmt.Println(csvPrinter)

  avroParser := AvroParser{Content: "My avro content", CreatedAt: "2018-01-25"}
  avroParser := Printer(avroParser)
  fmt.Println(avroParser)
}
