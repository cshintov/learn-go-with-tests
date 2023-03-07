package maps

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
    ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    def, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }
    return def, nil
}

func (d Dictionary) Add(word, def string) error {
    _, err := d.Search(word)
    switch err {
    case nil:
        return ErrWordExists
    case ErrNotFound:
        d[word] = def
        return nil
    default:
        return err
    }
}

func (d Dictionary) Update(word, def string) error {
    _, err := d.Search(word)
    switch err {
    case nil:
        d[word] = def
        return nil
    case ErrNotFound:
        return ErrWordDoesNotExist
    default:
        return err
    }
}

func (d Dictionary) Delete(word string) {
    delete(d, word)
}
