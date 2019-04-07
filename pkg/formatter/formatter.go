package formatter

type Formatter struct {
}

func (f *Formatter) Format(line map[string]interface{}) (string, error) {

	return line["message"].(string), nil
}
