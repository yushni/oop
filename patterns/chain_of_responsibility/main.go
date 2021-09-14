package main

import "fmt"

type section interface {
	execute(*task)
	setNext(section)
}

type material struct {
	next section
}

func (m *material) execute(t *task) {
	if t.materialCollected {
		fmt.Println("Вже всі матеріали на місці, працюйте халявщики")
		m.next.execute(t)
		return
	}
	fmt.Println("Починаємо збирати матеріали")
	t.materialCollected = true
	m.next.execute(t)
}

func (m *material) setNext(next section) {
	m.next = next
}

type assembly struct {
	next section
}

func (a *assembly) execute(t *task) {
	if t.assemblyExecuted {
		fmt.Println("Всьо вже позбирали, можете йти далі")
		a.next.execute(t)
		return
	}
	fmt.Println("Починаємо Збирати")
	t.assemblyExecuted = true
	a.next.execute(t)
}

func (a *assembly) setNext(next section) {
	a.next = next
}

type packaging struct {
	next section
}

func (p *packaging) execute(t *task) {
	if t.packagingExecuted {
		fmt.Println("Юж всьо поскладали")
		p.next.execute(t)
		return
	}
	fmt.Println("Секція пакування, починаємо пакувати!!!")
}

func (p *packaging) setNext(next section) {
	p.next = next
}

type task struct {
	name              string
	materialCollected bool
	assemblyExecuted  bool
	packagingExecuted bool
}

func main() {
	packaging := &packaging{}

	// set next for assembly section
	assembly := &assembly{}
	assembly.setNext(packaging)

	material := &material{}
	material.setNext(assembly)

	task := &task{name: "Іграшка"}
	material.execute(task)
}
