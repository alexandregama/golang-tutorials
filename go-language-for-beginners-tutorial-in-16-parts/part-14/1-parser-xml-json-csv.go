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

func XmlPrinter(xml XmlParser) string {
  return xml.Parse() + " - Date - " + xml.FormatDate()
}

func JsonPrinter(json JsonParser) string {
  return json.Parse() + " - Date - " + json.FormatDate()
}

func CsvPrinter(csv CsvParser) string {
  return csv.Parse() + " - Date - " + csv.FormatDate()
}

func main() {
  xmlParser := XmlParser{Content: "<title>My Title<title>", CreatedAt: "2018-01-20"}
  xmlPrinter := XmlPrinter(xmlParser)
  fmt.Println(xmlPrinter)

  jsonParser := JsonParser{Content: "{'title': 'My Title'}", CreatedAt: "2018-01-25"}
  jsonPrinter := JsonPrinter(jsonParser)
  fmt.Println(jsonPrinter)

  csvParser := CsvParser{Content: "My Content; Another Conent", CreatedAt: "2018-01-25"}
  csvPrinter := CsvPrinter(csvParser)
  fmt.Println(csvPrinter)
}
