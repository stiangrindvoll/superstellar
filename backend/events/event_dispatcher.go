package events

// GENERATED CODE! DO NOT EDIT THIS FILE!
// ADD YOUR EVENT AND RUN 'go generate' INSTEAD

import (
	"time"
)

const (
	buffersLength                         = 10000
	idleDispatcherSleepTime time.Duration = 5 * time.Millisecond
)

type TimeTickListener interface {
	HandleTimeTick(*TimeTick)
}

type ProjectileFiredListener interface {
	HandleProjectileFired(*ProjectileFired)
}

type UserInputListener interface {
	HandleUserInput(*UserInput)
}

type UserJoinedListener interface {
	HandleUserJoined(*UserJoined)
}

type UserLeftListener interface {
	HandleUserLeft(*UserLeft)
}

type EventDispatcher struct {

	// TimeTick
	timeTickQueue     chan *TimeTick
	timeTickListeners []TimeTickListener

	// ProjectileFired
	projectileFiredQueue     chan *ProjectileFired
	projectileFiredListeners []ProjectileFiredListener

	// UserInput
	userInputQueue     chan *UserInput
	userInputListeners []UserInputListener

	// UserJoined
	userJoinedQueue     chan *UserJoined
	userJoinedListeners []UserJoinedListener

	// UserLeft
	userLeftQueue     chan *UserLeft
	userLeftListeners []UserLeftListener
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{

		// TimeTick
		timeTickQueue:     make(chan *TimeTick, buffersLength),
		timeTickListeners: []TimeTickListener{},

		// ProjectileFired
		projectileFiredQueue:     make(chan *ProjectileFired, buffersLength),
		projectileFiredListeners: []ProjectileFiredListener{},

		// UserInput
		userInputQueue:     make(chan *UserInput, buffersLength),
		userInputListeners: []UserInputListener{},

		// UserJoined
		userJoinedQueue:     make(chan *UserJoined, buffersLength),
		userJoinedListeners: []UserJoinedListener{},

		// UserLeft
		userLeftQueue:     make(chan *UserLeft, buffersLength),
		userLeftListeners: []UserLeftListener{},
	}
}

func (d *EventDispatcher) RunEventLoop() {
	for {
		select {

		// TimeTick
		case event := <-d.timeTickQueue:
			for _, listener := range d.timeTickListeners {
				listener.HandleTimeTick(event)
			}

		// ProjectileFired
		case event := <-d.projectileFiredQueue:
			for _, listener := range d.projectileFiredListeners {
				listener.HandleProjectileFired(event)
			}

		// UserInput
		case event := <-d.userInputQueue:
			for _, listener := range d.userInputListeners {
				listener.HandleUserInput(event)
			}

		// UserJoined
		case event := <-d.userJoinedQueue:
			for _, listener := range d.userJoinedListeners {
				listener.HandleUserJoined(event)
			}

		// UserLeft
		case event := <-d.userLeftQueue:
			for _, listener := range d.userLeftListeners {
				listener.HandleUserLeft(event)
			}

		default:
			time.Sleep(idleDispatcherSleepTime)
		}
	}
}

// EVENT METHODS

// TimeTick

func (d *EventDispatcher) RegisterTimeTickListener(listener TimeTickListener) {
	d.timeTickListeners = append(d.timeTickListeners, listener)
}

func (d *EventDispatcher) FireTimeTick(e *TimeTick) {
	d.timeTickQueue <- e
}

// ProjectileFired

func (d *EventDispatcher) RegisterProjectileFiredListener(listener ProjectileFiredListener) {
	d.projectileFiredListeners = append(d.projectileFiredListeners, listener)
}

func (d *EventDispatcher) FireProjectileFired(e *ProjectileFired) {
	d.projectileFiredQueue <- e
}

// UserInput

func (d *EventDispatcher) RegisterUserInputListener(listener UserInputListener) {
	d.userInputListeners = append(d.userInputListeners, listener)
}

func (d *EventDispatcher) FireUserInput(e *UserInput) {
	d.userInputQueue <- e
}

// UserJoined

func (d *EventDispatcher) RegisterUserJoinedListener(listener UserJoinedListener) {
	d.userJoinedListeners = append(d.userJoinedListeners, listener)
}

func (d *EventDispatcher) FireUserJoined(e *UserJoined) {
	d.userJoinedQueue <- e
}

// UserLeft

func (d *EventDispatcher) RegisterUserLeftListener(listener UserLeftListener) {
	d.userLeftListeners = append(d.userLeftListeners, listener)
}

func (d *EventDispatcher) FireUserLeft(e *UserLeft) {
	d.userLeftQueue <- e
}