package name

func TestGenerate(t *testing.T) {
	generated := Generate()
	parts = strings.Split(generated, "-")
	if len(parts) != 3 {
		t.Error("wrong number of parts in generated name")
	}
}
