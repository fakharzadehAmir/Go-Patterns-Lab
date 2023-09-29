package Memento

// Originator represents the originator object that holds a state to be saved.
type Originator struct {
	text string
}

// set sets the state of the originator.
func (o *Originator) set(text string) {
	o.text = text
}

// takeSnapshot creates a new memento with the current state.
func (o *Originator) takeSnapshot() *Memento {
	return &Memento{text: o.text}
}

// restore restores the originator's state from a memento.
func (o *Originator) restore(memento *Memento) {
	o.text = memento.getSavedText()
}

// Memento represents the memento object that stores a state.
type Memento struct {
	text string
}

// getSavedText retrieves the state stored in the memento.
func (m *Memento) getSavedText() string {
	return m.text
}

// Caretaker represents the caretaker object that manages mementos.
type Caretaker struct {
	originator   *Originator
	mementoArray []*Memento
}

// write saves the current state of the originator as a new memento.
func (c *Caretaker) write(text string) {
	c.originator.set(text)
	c.mementoArray = append(c.mementoArray, c.originator.takeSnapshot())
}

// undo restores the state of the originator from the last saved memento.
func (c *Caretaker) undo() {
	if len(c.mementoArray) > 1 {
		c.mementoArray = c.mementoArray[:len(c.mementoArray)-1]
		c.originator.restore(c.mementoArray[len(c.mementoArray)-1])
		return
	}
	c.originator.restore(&Memento{text: ""})
}
