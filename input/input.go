package input

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

type Direction string

const (
	LEFT  Direction = "LEFT"
	RIGHT Direction = "RIGHT"
	UP    Direction = "UP"
	DOWN  Direction = "DOWN"
	EXIT            = "EXIT"
)

type InputBuffer struct {
	LastInputKeyCode keyboard.Key
	eventHandlers    []InputHandler
	InputChannel     chan string
	SyncChannel      chan int
	ExitChannel      chan int
}

func New(inputChan chan string, syncChan chan int, exitChan chan int) *InputBuffer {
	err := keyboard.Open()
	if err != nil {
		fmt.Println("Could not open channel for user input!")
	}

	return &InputBuffer{
		keyboard.KeySpace,
		[]InputHandler{},
		inputChan,
		syncChan,
		exitChan,
	}
}

func (ib *InputBuffer) Listen() {
	for {
		select {
		case x := <-ib.SyncChannel:
			if x == 0 {
				return
			}
			ib.Tick()
		default:
			_, k, err := keyboard.GetKey()

			if err != nil {
				if fmt.Sprintf("%v", err) != "operation canceled" {
					fmt.Printf("ERROR GETTING KEY PRESS: \n%v\n", err)
					panic("Error getting keyboard input, process will exit")
				}
			}
			ib.LastInputKeyCode = k
		}
	}
}

func (ib *InputBuffer) Close() {
	err := keyboard.Close()

	if err != nil {
		fmt.Printf("err: %v", err)
		panic("Could not close keyboard channel, crashing the process")
	}
}

func (ib *InputBuffer) Tick() {
	for _, handler := range ib.eventHandlers {
		if ib.LastInputKeyCode == handler.Key {
			handler.Handler()
		}
	}
}

type InputHandler struct {
	Key     keyboard.Key
	Handler func()
}

func (ib *InputBuffer) OnInput(handler InputHandler) {
	ib.eventHandlers = append(ib.eventHandlers, handler)
}
